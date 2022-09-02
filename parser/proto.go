package parser

import (
	"encoding/binary"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/golang/protobuf/jsonpb"
	"github.com/housepower/clickhouse_sinker/model"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/thanos-io/thanos/pkg/errors"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"math"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

// ProtoParser knows how to get data from proto format.
type ProtoParser struct {
	pp           *Pool
	deserializer *ProtoDeserializer
}

func (p *ProtoParser) Parse(bs []byte) (metric model.Metric, err error) {
	// use json
	//json, err := p.deserializer.ToJSON(bs)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("parsed json: %s", json)
	//
	//var value *fastjson.Value
	//if value, err = p.fjp.ParseBytes(json); err != nil {
	//	err = errors.Wrapf(err, "")
	//	return
	//}
	//return &FastjsonMetric{pp: p.pp, value: value}

	dynamicMsg, err := p.deserializer.ToDynamicMessage(bs)
	if err != nil {
		err = errors.Wrapf(err, "")
		return
	}
	return &ProtoMetric{pp: p.pp, msg: dynamicMsg}, nil
}

type ProtoMetric struct {
	pp  *Pool
	msg *dynamic.Message
}

func (m *ProtoMetric) GetBool(key string, nullable bool) (val interface{}) {
	fieldVal, _ := m.msg.TryGetFieldByName(key)
	if fieldVal == nil {
		return getDefaultBool(nullable)
	}

	rv := reflect.ValueOf(fieldVal)
	if rv.Kind() == reflect.Bool {
		return rv.Bool()
	}

	return getDefaultBool(nullable)
}

func (m *ProtoMetric) GetInt8(key string, nullable bool) (val interface{}) {
	return getIntFromProto[int8](m, key, nullable, math.MinInt8, math.MaxInt8)
}

func (m *ProtoMetric) GetInt16(key string, nullable bool) (val interface{}) {
	return getIntFromProto[int16](m, key, nullable, math.MinInt16, math.MaxInt16)
}

func (m *ProtoMetric) GetInt32(key string, nullable bool) (val interface{}) {
	return getIntFromProto[int32](m, key, nullable, math.MinInt32, math.MaxInt32)
}

func (m *ProtoMetric) GetInt64(key string, nullable bool) (val interface{}) {
	return getIntFromProto[int64](m, key, nullable, math.MinInt64, math.MaxInt64)
}

func (m *ProtoMetric) GetUint8(key string, nullable bool) (val interface{}) {
	return getUIntFromProto[uint8](m, key, nullable, math.MaxUint8)
}

func (m *ProtoMetric) GetUint16(key string, nullable bool) (val interface{}) {
	return getUIntFromProto[uint16](m, key, nullable, math.MaxUint16)
}

func (m *ProtoMetric) GetUint32(key string, nullable bool) (val interface{}) {
	return getUIntFromProto[uint32](m, key, nullable, math.MaxUint32)
}

func (m *ProtoMetric) GetUint64(key string, nullable bool) (val interface{}) {
	return getUIntFromProto[uint64](m, key, nullable, math.MaxUint64)
}

func (m *ProtoMetric) GetFloat32(key string, nullable bool) (val interface{}) {
	// TODO: support float32
	panic("not supported")
}
func (m *ProtoMetric) GetFloat64(key string, nullable bool) (val interface{}) {
	// TODO: support float64
	panic("not supported")
}
func (m *ProtoMetric) GetDecimal(key string, nullable bool) (val interface{}) {
	// TODO: support decimal
	panic("not supported")
}

func (m *ProtoMetric) GetDateTime(key string, nullable bool) (val interface{}) {
	fieldVal, _ := m.msg.TryGetFieldByName(key)
	if fieldVal == nil {
		return getDefaultDateTime(nullable)
	}

	rv := reflect.ValueOf(fieldVal)
	if ts, ok := rv.Interface().(*timestamppb.Timestamp); ok {
		return ts.AsTime()
	}

	return getDefaultDateTime(nullable)
}

func (m *ProtoMetric) GetString(key string, nullable bool) (val interface{}) {
	fieldVal, _ := m.msg.TryGetFieldByName(key)
	if fieldVal == nil {
		if nullable {
			return
		}
		return ""
	}

	rv := reflect.ValueOf(fieldVal)
	if rv.Kind() != reflect.String {
		return ""
	}
	return rv.String()
}

func (m *ProtoMetric) GetArray(key string, t int) (val interface{}) {
	// TODO: support array
	panic("not supported")
}
func (m *ProtoMetric) GetNewKeys(knownKeys, newKeys, warnKeys *sync.Map, white, black *regexp.Regexp, partition int, offset int64) bool {
	return false
}

func getIntFromProto[T constraints.Signed](m *ProtoMetric, key string, nullable bool, min, max int64) (val interface{}) {
	fieldVal, _ := m.msg.TryGetFieldByName(key)
	if fieldVal == nil {
		return getDefaultInt[T](nullable)
	}

	rv := reflect.ValueOf(fieldVal)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		int64Val := rv.Int()
		if int64Val < min {
			return T(min)
		}
		if int64Val > max {
			return T(max)
		}
		return T(int64Val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uint64Val := rv.Uint()
		if uint64Val > uint64(max) {
			return T(max)
		}
		return T(uint64Val)
	case reflect.Bool:
		if rv.Bool() {
			return T(1)
		}
		return T(0)
	default:
		return getDefaultInt[T](nullable)
	}
}

func getUIntFromProto[T constraints.Unsigned](m *ProtoMetric, key string, nullable bool, max uint64) (val interface{}) {
	fieldVal, _ := m.msg.TryGetFieldByName(key)
	if fieldVal == nil {
		return getDefaultInt[T](nullable)
	}

	rv := reflect.ValueOf(fieldVal)
	switch rv.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uint64Val := rv.Uint()
		if uint64Val > max {
			return T(max)
		}
		return T(uint64Val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intV := rv.Int()
		if intV < 0 {
			return getDefaultInt[T](nullable)
		}
		if uint64(intV) > max {
			return T(max)
		}
		return T(intV)
	case reflect.Bool:
		if rv.Bool() {
			return T(1)
		}
		return T(0)
	default:
		return getDefaultInt[T](nullable)
	}
}

// ProtoDeserializer represents a Protobuf deserializer
type ProtoDeserializer struct {
	srClient         schemaregistry.Client
	baseDeserializer *serde.BaseDeserializer
	topic            string
}

func (p *ProtoDeserializer) ToDynamicMessage(bytes []byte) (*dynamic.Message, error) {
	if len(bytes) == 0 {
		return nil, nil
	}
	schemaInfo, err := p.baseDeserializer.GetSchema(p.topic, bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema info: %w", err)
	}
	fileDesc, err := p.toFileDescriptor(schemaInfo)
	if err != nil {
		return nil, err
	}
	bytesRead, msgIndexes, err := readMessageIndexes(bytes[5:])
	if err != nil {
		return nil, err
	}
	messageDesc, err := toMessageDescriptor(fileDesc, msgIndexes)
	if err != nil {
		return nil, err
	}

	msg := dynamic.NewMessage(messageDesc)
	if err := msg.Unmarshal(bytes[5+bytesRead:]); err != nil {
		return nil, fmt.Errorf("failed to unmarshal value into protobuf message: %w", err)
	}

	return msg, nil
}

func (p *ProtoDeserializer) ToJSON(bytes []byte) ([]byte, error) {
	if len(bytes) == 0 {
		return nil, nil
	}
	schemaInfo, err := p.baseDeserializer.GetSchema(p.topic, bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema info")
	}
	fd, err := p.toFileDescriptor(schemaInfo)
	if err != nil {
		return nil, err
	}
	bytesRead, msgIndexes, err := readMessageIndexes(bytes[5:])
	if err != nil {
		return nil, err
	}
	messageDesc, err := toMessageDescriptor(fd, msgIndexes)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := p.protobufToJSON(bytes[5+bytesRead:], messageDesc)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func (p *ProtoDeserializer) toFileDescriptor(info schemaregistry.SchemaInfo) (*desc.FileDescriptor, error) {
	deps := make(map[string]string)
	err := serde.ResolveReferences(p.srClient, info, deps)
	if err != nil {
		return nil, err
	}
	parser := protoparse.Parser{
		Accessor: func(filename string) (io.ReadCloser, error) {
			var schema string
			if filename == "." {
				schema = info.Schema
			} else {
				schema = deps[filename]
			}
			if schema == "" {
				// these may be google types
				return nil, errors.Newf("unknown reference")
			}
			return io.NopCloser(strings.NewReader(schema)), nil
		},
	}

	fileDescriptors, err := parser.ParseFiles(".")
	if err != nil {
		return nil, err
	}

	if len(fileDescriptors) != 1 {
		return nil, fmt.Errorf("could not resolve schema")
	}
	return fileDescriptors[0], nil
}

func (p *ProtoDeserializer) protobufToJSON(bytes []byte, md *desc.MessageDescriptor) ([]byte, error) {
	msg := dynamic.NewMessage(md)
	err := msg.Unmarshal(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal value into protobuf message: %w", err)
	}

	marshaler := &jsonpb.Marshaler{
		OrigName:    true,
		EnumsAsInts: true,
	}
	jsonBytes, err := msg.MarshalJSONPB(marshaler)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal protobuf message to JSON: %w", err)
	}

	return jsonBytes, nil
}

func readMessageIndexes(bytes []byte) (int, []int, error) {
	arrayLen, bytesRead := binary.Varint(bytes)
	if bytesRead <= 0 {
		return bytesRead, nil, fmt.Errorf("unable to read message indexes")
	}
	if arrayLen == 0 {
		// handle the optimization for the first message in the schema
		return bytesRead, []int{0}, nil
	}
	msgIndexes := make([]int, arrayLen)
	for i := 0; i < int(arrayLen); i++ {
		idx, read := binary.Varint(bytes[bytesRead:])
		if read <= 0 {
			return bytesRead, nil, fmt.Errorf("unable to read message indexes")
		}
		bytesRead += read
		msgIndexes[i] = int(idx)
	}
	return bytesRead, msgIndexes, nil
}

func toMessageDescriptor(descriptor desc.Descriptor, msgIndexes []int) (*desc.MessageDescriptor, error) {
	index := msgIndexes[0]

	switch v := descriptor.(type) {
	case *desc.FileDescriptor:
		if len(msgIndexes) == 1 {
			return v.GetMessageTypes()[index], nil
		}
		return toMessageDescriptor(v.GetMessageTypes()[index], msgIndexes[1:])
	case *desc.MessageDescriptor:
		if len(msgIndexes) == 1 {
			return v.GetNestedMessageTypes()[index], nil
		}
		return toMessageDescriptor(v.GetNestedMessageTypes()[index], msgIndexes[1:])
	default:
		return nil, fmt.Errorf("unexpected type")
	}
}
