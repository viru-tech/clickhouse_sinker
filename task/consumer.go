/*Copyright [2019] housepower

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package task

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/viru-tech/clickhouse_sinker/config"
	"github.com/viru-tech/clickhouse_sinker/input"
	"github.com/viru-tech/clickhouse_sinker/model"
	"github.com/viru-tech/clickhouse_sinker/util"
	"go.uber.org/zap"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

type Commit struct {
	group    string
	offsets  model.RecordMap
	wg       *sync.WaitGroup
	consumer *Consumer
}

type Consumer struct {
	sinker    *Sinker
	inputer   *input.KafkaFranz
	tasks     sync.Map
	grpConfig *config.GroupConfig
	fetchesCh chan input.Fetches
	processWg sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	state     atomic.Uint32
	errCommit bool

	numFlying  int32
	mux        sync.Mutex
	commitDone *sync.Cond
}

const (
	MaxCountInBuf  = 1 << 27
	MaxParallelism = 10
)

func newConsumer(s *Sinker, gCfg *config.GroupConfig) *Consumer {
	c := &Consumer{
		sinker:    s,
		numFlying: 0,
		errCommit: false,
		grpConfig: gCfg,
		fetchesCh: make(chan input.Fetches),
	}
	c.state.Store(util.StateStopped)
	c.commitDone = sync.NewCond(&c.mux)
	return c
}

func (c *Consumer) addTask(tsk *Service) {
	c.tasks.Store(tsk.taskCfg.Name, tsk)
}

func (c *Consumer) start() {
	if c.state.Load() == util.StateRunning {
		return
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.inputer = input.NewKafkaFranz()
	c.state.Store(util.StateRunning)
	if err := c.inputer.Init(c.sinker.curCfg, c.grpConfig, c.fetchesCh, c.cleanupFn); err == nil {
		go c.inputer.Run()
		go c.processFetch()
	} else {
		util.Logger.Fatal("failed to init consumer", zap.String("consumer", c.grpConfig.Name), zap.Error(err))
	}
}

func (c *Consumer) stop() {
	if c.state.Load() == util.StateStopped {
		return
	}
	c.state.Store(util.StateStopped)

	// stop the processFetch routine, make sure no more input to the commit chan & writing pool
	c.cancel()
	c.processWg.Wait()
	c.inputer.Stop()
}

func (c *Consumer) restart() {
	c.stop()
	c.start()
}

func (c *Consumer) cleanupFn() {
	// ensure the completion of writing to ck
	var wg sync.WaitGroup
	c.tasks.Range(func(key, value any) bool {
		wg.Add(1)
		go func(t *Service) {
			// drain ensure we have completeted persisting all received messages
			t.clickhouse.Drain()
			wg.Done()
		}(value.(*Service))
		return true
	})
	wg.Wait()

	// ensure the completion of offset submission
	c.mux.Lock()
	for c.numFlying != 0 {
		util.Logger.Debug("draining flying pending commits", zap.String("consumergroup", c.grpConfig.Name), zap.Int32("pending", c.numFlying))
		c.commitDone.Wait()
	}
	c.mux.Unlock()
}

func (c *Consumer) updateGroupConfig(g *config.GroupConfig) {
	if c.state.Load() == util.StateStopped {
		return
	}
	c.grpConfig = g
	// restart the processFetch routine because of potential BufferSize or FlushInterval change
	// make sure no more input to the commit chan & writing pool
	c.cancel()
	c.processWg.Wait()
	c.ctx, c.cancel = context.WithCancel(context.Background())
	go c.processFetch()
}

func (c *Consumer) processFetch() {
	c.processWg.Add(1)
	defer c.processWg.Done()

	flushers := make(map[string]*taskFlusher)
	topicsToFlushers := make(map[string][]*taskFlusher)
	c.tasks.Range(func(key, value any) bool {
		task := value.(*Service)
		bufSize := uint64(task.taskCfg.BufferSize * len(c.sinker.curCfg.Clickhouse.Hosts) * 4 / 5)
		flusher := &taskFlusher{
			threshold: bufSize,
			inputC:    make(chan messageWithTrace, bufSize),
			task:      task,
			recMap:    make(map[int32]*model.BatchRange),
		}
		if task.taskCfg.FlushInterval != 0 {
			flusher.duration = time.Duration(task.taskCfg.FlushInterval) * time.Second
			flusher.ticker = time.NewTicker(flusher.duration)
		}

		flushers[task.taskCfg.Name] = flusher
		topicsToFlushers[task.taskCfg.Topic] = append(topicsToFlushers[task.taskCfg.Topic], flusher)

		return true
	})

	wg := sync.WaitGroup{}
	wg.Add(len(flushers))
	defer wg.Wait()
	thresholdsCtx, cancel := context.WithCancel(c.ctx)
	defer cancel()
	for i := range flushers {
		task := i
		go func() {
			defer wg.Done()
			flusher := flushers[task]
			var traceID string
			for {
				select {
				case msg := <-flusher.inputC:
					partition := int32(msg.msg.Partition)
					if flusher.recMap[partition] == nil {
						flusher.recMap[partition] = &model.BatchRange{
							Begin: msg.msg.Offset,
						}
					}

					traceID = msg.traceID
					if msg.msg.Offset > flusher.recMap[partition].End {
						flusher.recMap[partition].End = msg.msg.Offset
					}
					flusher.current++
					err := flusher.task.Put(msg.msg, traceID, func(traceId, with string) {
						flusher.flushFn(c, traceId, with)
					})
					if flusher.current >= flusher.threshold {
						flusher.flushFn(c, traceID, "bufLength reached")
					}
					if err != nil {
						// decrease the error record
						util.Rs.Dec(1)
						util.Logger.Error("putting message in flusher failed", zap.Error(err))
					}
				case <-flusher.ticker.C:
					flusher.flushFn(c, traceID, "ticker.C triggered")
				case <-thresholdsCtx.Done():
					flusher.flushFn(c, traceID, "consumer is closing")
					return
				}
			}
		}()
	}

	for {
		select {
		case fetches := <-c.fetchesCh:
			if c.state.Load() == util.StateStopped {
				continue
			}
			fetch := fetches.Fetch.Records()
			traceId := fetches.TraceId
			util.LogTrace(traceId, util.TraceKindProcessStart, zap.Int("records", len(fetch)))

			for i := range fetch {
				if c.state.Load() == util.StateStopped {
					break
				}

				rec := fetch[i]
				msg := &model.InputMessage{
					Topic:     rec.Topic,
					Partition: int(rec.Partition),
					Key:       rec.Key,
					Value:     rec.Value,
					Offset:    rec.Offset,
					Timestamp: &rec.Timestamp,
				}

				workers, ok := topicsToFlushers[rec.Topic]
				if !ok {
					util.Logger.Warn("topic not found", zap.String("topic", rec.Topic))
					continue
				}

				for _, worker := range workers {
					select {
					case worker.inputC <- messageWithTrace{
						msg:     msg,
						traceID: traceId,
					}:
					case <-c.ctx.Done():
					}
				}
			}

		case <-c.ctx.Done():
			util.Logger.Info("stopped processing loop", zap.String("group", c.grpConfig.Name))
			cancel()
			return
		}
	}
}

type messageWithTrace struct {
	traceID string
	msg     *model.InputMessage
}

type taskFlusher struct {
	duration  time.Duration
	ticker    *time.Ticker
	threshold uint64
	current   uint64
	inputC    chan messageWithTrace
	task      *Service
	recMap    map[int32]*model.BatchRange
}

func (t *taskFlusher) flushFn(consumer *Consumer, traceId, with string) {
	if len(t.recMap) == 0 {
		return
	}

	bufLength := t.current
	if bufLength > 0 {
		util.LogTrace(traceId, util.TraceKindProcessEnd,
			zap.String("with", with),
			zap.Uint64("bufLength", bufLength),
		)
	}

	var wg sync.WaitGroup
	t.task.sharder.Flush(consumer.ctx, &wg, t.recMap, traceId)
	if consumer.ctx.Err() != nil {
		return
	}

	util.Logger.Warn("flushed",
		zap.Uint64("count", bufLength),
		zap.String("trace_id", traceId),
		zap.String("topic", t.task.taskCfg.Topic),
	)
	consumer.mux.Lock()
	consumer.numFlying++
	consumer.mux.Unlock()
	consumer.sinker.commitsCh <- &Commit{group: consumer.grpConfig.Name, offsets: model.RecordMap{
		t.task.taskCfg.Topic: t.recMap,
	}, wg: &wg, consumer: consumer}
	util.Logger.Info("commited offsets",
		zap.String("topic", t.task.taskCfg.Topic),
		zap.Any("offsets", t.recMap),
	)
	t.recMap = make(map[int32]*model.BatchRange)
	t.ticker.Reset(t.duration)
	t.current = 0
}
