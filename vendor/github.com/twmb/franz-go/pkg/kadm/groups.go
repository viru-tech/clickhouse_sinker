package kadm

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kmsg"
)

// GroupMemberMetadata is the metadata that a client sent in a JoinGroup request.
// This can have one of three types:
//
//     *kmsg.ConsumerMemberMetadata, if the group's ProtocolType is "consumer"
//     *kmsg.ConnectMemberMetadata, if the group's ProtocolType is "connect"
//     []byte, if the group's ProtocolType is unknown
//
type GroupMemberMetadata struct{ i interface{} }

// AsConsumer returns the metadata as a ConsumerMemberMetadata if possible.
func (m GroupMemberMetadata) AsConsumer() (*kmsg.ConsumerMemberMetadata, bool) {
	c, ok := m.i.(*kmsg.ConsumerMemberMetadata)
	return c, ok
}

// AsConnect returns the metadata as ConnectMemberMetadata if possible.
func (m GroupMemberMetadata) AsConnect() (*kmsg.ConnectMemberMetadata, bool) {
	c, ok := m.i.(*kmsg.ConnectMemberMetadata)
	return c, ok
}

// Raw returns the metadata as a raw byte slice, if it is neither of consumer
// type nor connect type.
func (m GroupMemberMetadata) Raw() ([]byte, bool) {
	c, ok := m.i.([]byte)
	return c, ok
}

// GroupMemberAssignment is the assignment that a leader sent / a member
// received in a SyncGroup request.  This can have one of three types:
//
//     *kmsg.ConsumerMemberAssignment, if the group's ProtocolType is "consumer"
//     *kmsg.ConnectMemberAssignment, if the group's ProtocolType is "connect"
//     []byte, if the group's ProtocolType is unknown
//
type GroupMemberAssignment struct{ i interface{} }

// AsConsumer returns the assignment as a ConsumerMemberAssignment if possible.
func (m GroupMemberAssignment) AsConsumer() (*kmsg.ConsumerMemberAssignment, bool) {
	c, ok := m.i.(*kmsg.ConsumerMemberAssignment)
	return c, ok
}

// AsConnect returns the assignment as ConnectMemberAssignment if possible.
func (m GroupMemberAssignment) AsConnect() (*kmsg.ConnectMemberAssignment, bool) {
	c, ok := m.i.(*kmsg.ConnectMemberAssignment)
	return c, ok
}

// Raw returns the assignment as a raw byte slice, if it is neither of consumer
// type nor connect type.
func (m GroupMemberAssignment) Raw() ([]byte, bool) {
	c, ok := m.i.([]byte)
	return c, ok
}

// DescribedGroupMember is the detail of an individual group member as returned
// by a describe groups response.
type DescribedGroupMember struct {
	MemberID   string  // MemberID is the Kafka assigned member ID of this group member.
	InstanceID *string // InstanceID is a potential user assigned instance ID of this group member (KIP-345).
	ClientID   string  // ClientID is the Kafka client given ClientID of this group member.
	ClientHost string  // ClientHost is the host this member is running on.

	Join     GroupMemberMetadata   // Join is what this member sent in its join group request; what it wants to consume.
	Assigned GroupMemberAssignment // Assigned is what this member was assigned to consume by the leader.
}

// AssignedPartitions returns the set of unique topics and partitions that are
// assigned across all members in this group.
//
// This function is only relevant if the group is of type "consumer".
func (d *DescribedGroup) AssignedPartitions() TopicsSet {
	s := make(TopicsSet)
	for _, m := range d.Members {
		if c, ok := m.Assigned.AsConsumer(); ok {
			for _, t := range c.Topics {
				s.Add(t.Topic, t.Partitions...)
			}
		}
	}
	return s
}

// DescribedGroup contains data from a describe groups response for a single
// group.
type DescribedGroup struct {
	Group string // Group is the name of the described group.

	Coordinator  BrokerDetail           // Coordinator is the coordinator broker for this group.
	State        string                 // State is the state this group is in (Empty, Dead, Stable, etc.).
	ProtocolType string                 // ProtocolType is the type of protocol the group is using, "consumer" for normal consumers, "connect" for Kafka connect.
	Protocol     string                 // Protocol is the partition assignor strategy this group is using.
	Members      []DescribedGroupMember // Members contains the members of this group sorted first by InstanceID, or if nil, by MemberID.

	Err error // Err is non-nil if the group could not be described.
}

// DescribedGroups contains data for multiple groups from a describe groups
// response.
type DescribedGroups map[string]DescribedGroup

// AssignedPartitions returns the set of unique topics and partitions that are
// assigned across all members in all groups. This is the all-group analogue to
// DescribedGroup.AssignedPartitions.
//
// This function is only relevant for groups of type "consumer".
func (ds DescribedGroups) AssignedPartitions() TopicsSet {
	s := make(TopicsSet)
	for _, g := range ds {
		for _, m := range g.Members {
			if c, ok := m.Assigned.AsConsumer(); ok {
				for _, t := range c.Topics {
					s.Add(t.Topic, t.Partitions...)
				}
			}
		}
	}
	return s
}

// Sorted returns all groups sorted by group name.
func (ds DescribedGroups) Sorted() []DescribedGroup {
	s := make([]DescribedGroup, 0, len(ds))
	for _, d := range ds {
		s = append(s, d)
	}
	sort.Slice(s, func(i, j int) bool { return s[i].Group < s[j].Group })
	return s
}

// On calls fn for the group if it exists, returning the group and the error
// returned from fn. If fn is nil, this simply returns the group.
//
// The fn is given a shallow copy of the group. This function returns the copy
// as well; any modifications within fn are modifications on the returned copy.
// Modifications on a described group's inner fields are persisted to the
// original map (because slices are pointers).
//
// If the group does not exist, this returns kerr.GroupIDNotFound.
func (rs DescribedGroups) On(group string, fn func(*DescribedGroup) error) (DescribedGroup, error) {
	if len(rs) > 0 {
		r, ok := rs[group]
		if ok {
			if fn == nil {
				return r, nil
			}
			return r, fn(&r)
		}
	}
	return DescribedGroup{}, kerr.GroupIDNotFound
}

// Topics returns a sorted list of all group names.
func (ds DescribedGroups) Names() []string {
	all := make([]string, 0, len(ds))
	for g := range ds {
		all = append(all, g)
	}
	sort.Strings(all)
	return all
}

// ListedGroup contains data from a list groups response for a single group.
type ListedGroup struct {
	Coordinator  int32  // Coordinator is the node ID of the coordinator for this group.
	Group        string // Group is the name of this group.
	ProtocolType string // ProtocolType is the type of protocol the group is using, "consumer" for normal consumers, "connect" for Kafka connect.
	State        string // State is the state this group is in (Empty, Dead, Stable, etc.; only if talking to Kafka 2.6+).
}

// ListedGroups contains information from a list groups response.
type ListedGroups map[string]ListedGroup

// Sorted returns all groups sorted by group name.
func (ls ListedGroups) Sorted() []ListedGroup {
	s := make([]ListedGroup, 0, len(ls))
	for _, l := range ls {
		s = append(s, l)
	}
	sort.Slice(s, func(i, j int) bool { return s[i].Group < s[j].Group })
	return s
}

// Groups returns a sorted list of all group names.
func (ls ListedGroups) Groups() []string {
	all := make([]string, 0, len(ls))
	for g := range ls {
		all = append(all, g)
	}
	sort.Strings(all)
	return all
}

// ListGroups returns all groups in the cluster. If you are talking to Kafka
// 2.6+, filter states can be used to return groups only in the requested
// states. By default, this returns all groups. In almost all cases,
// DescribeGroups is more useful.
//
// This may return *ShardErrors.
func (cl *Client) ListGroups(ctx context.Context, filterStates ...string) (ListedGroups, error) {
	req := kmsg.NewPtrListGroupsRequest()
	req.StatesFilter = append(req.StatesFilter, filterStates...)
	shards := cl.cl.RequestSharded(ctx, req)
	list := make(ListedGroups)
	return list, shardErrEachBroker(req, shards, func(b BrokerDetail, kr kmsg.Response) error {
		resp := kr.(*kmsg.ListGroupsResponse)
		if err := maybeAuthErr(resp.ErrorCode); err != nil {
			return err
		}
		if err := kerr.ErrorForCode(resp.ErrorCode); err != nil {
			return err
		}
		for _, g := range resp.Groups {
			list[g.Group] = ListedGroup{
				Coordinator:  b.NodeID,
				Group:        g.Group,
				ProtocolType: g.ProtocolType,
				State:        g.GroupState,
			}
		}
		return nil
	})
}

// DescribeGroups describes either all groups specified, or all groups in the
// cluster if none are specified.
//
// This may return *ShardErrors.
//
// If no groups are specified and this method first lists groups, and list
// groups returns a *ShardErrors, this function describes all successfully
// listed groups and appends the list shard errors to any returned describe
// shard errors.
//
// If only one group is described, there will be at most one request issued,
// and there is no need to deeply inspect the error.
func (cl *Client) DescribeGroups(ctx context.Context, groups ...string) (DescribedGroups, error) {
	var seList *ShardErrors
	if len(groups) == 0 {
		listed, err := cl.ListGroups(ctx)
		switch {
		case err == nil:
		case errors.As(err, &seList):
		default:
			return nil, err
		}
		groups = listed.Groups()
		if len(groups) == 0 {
			return nil, err
		}
	}

	req := kmsg.NewPtrDescribeGroupsRequest()
	req.Groups = groups

	shards := cl.cl.RequestSharded(ctx, req)
	described := make(DescribedGroups)
	err := shardErrEachBroker(req, shards, func(b BrokerDetail, kr kmsg.Response) error {
		resp := kr.(*kmsg.DescribeGroupsResponse)
		for _, rg := range resp.Groups {
			if err := maybeAuthErr(rg.ErrorCode); err != nil {
				return err
			}
			g := DescribedGroup{
				Group:        rg.Group,
				Coordinator:  b,
				State:        rg.State,
				ProtocolType: rg.ProtocolType,
				Protocol:     rg.Protocol,
				Err:          kerr.ErrorForCode(rg.ErrorCode),
			}
			for _, rm := range rg.Members {
				gm := DescribedGroupMember{
					MemberID:   rm.MemberID,
					InstanceID: rm.InstanceID,
					ClientID:   rm.ClientID,
					ClientHost: rm.ClientHost,
				}

				var mi, ai interface{}
				switch g.ProtocolType {
				case "consumer":
					m := new(kmsg.ConsumerMemberMetadata)
					a := new(kmsg.ConsumerMemberAssignment)

					m.ReadFrom(rm.ProtocolMetadata)
					a.ReadFrom(rm.MemberAssignment)

					mi, ai = m, a
				case "connect":
					m := new(kmsg.ConnectMemberMetadata)
					a := new(kmsg.ConnectMemberAssignment)

					m.ReadFrom(rm.ProtocolMetadata)
					a.ReadFrom(rm.MemberAssignment)

					mi, ai = m, a
				default:
					mi, ai = rm.ProtocolMetadata, rm.MemberAssignment
				}

				gm.Join = GroupMemberMetadata{mi}
				gm.Assigned = GroupMemberAssignment{ai}
				g.Members = append(g.Members, gm)
			}
			sort.Slice(g.Members, func(i, j int) bool {
				if g.Members[i].InstanceID != nil {
					if g.Members[j].InstanceID == nil {
						return true
					}
					return *g.Members[i].InstanceID < *g.Members[j].InstanceID
				}
				if g.Members[j].InstanceID != nil {
					return false
				}
				return g.Members[i].MemberID < g.Members[j].MemberID
			})
			described[g.Group] = g
		}
		return nil
	})

	var seDesc *ShardErrors
	switch {
	case err == nil:
		return described, seList.into()
	case errors.As(err, &seDesc):
		if seList != nil {
			seDesc.Errs = append(seList.Errs, seDesc.Errs...)
		}
		return described, seDesc.into()
	default:
		return nil, err
	}
}

// DeleteGroupResponse contains the response for an individual deleted group.
type DeleteGroupResponse struct {
	Group string // Group is the group this response is for.
	Err   error  // Err is non-nil if the group failed to be deleted.
}

// DeleteGroupResponses contains per-group responses to deleted groups.
type DeleteGroupResponses map[string]DeleteGroupResponse

// Sorted returns all deleted group responses sorted by group name.
func (ds DeleteGroupResponses) Sorted() []DeleteGroupResponse {
	s := make([]DeleteGroupResponse, 0, len(ds))
	for _, d := range ds {
		s = append(s, d)
	}
	sort.Slice(s, func(i, j int) bool { return s[i].Group < s[j].Group })
	return s
}

// On calls fn for the response group if it exists, returning the response and
// the error returned from fn. If fn is nil, this simply returns the group.
//
// The fn is given a copy of the response. This function returns the copy as
// well; any modifications within fn are modifications on the returned copy.
//
// If the group does not exist, this returns kerr.GroupIDNotFound.
func (rs DeleteGroupResponses) On(group string, fn func(*DeleteGroupResponse) error) (DeleteGroupResponse, error) {
	if len(rs) > 0 {
		r, ok := rs[group]
		if ok {
			if fn == nil {
				return r, nil
			}
			return r, fn(&r)
		}
	}
	return DeleteGroupResponse{}, kerr.GroupIDNotFound
}

// DeleteGroups deletes all groups specified.
//
// The purpose of this request is to allow operators a way to delete groups
// after Kafka 1.1, which removed RetentionTimeMillis from offset commits. See
// KIP-229 for more details.
//
// This may return *ShardErrors. This does not return on authorization
// failures, instead, authorization failures are included in the responses.
func (cl *Client) DeleteGroups(ctx context.Context, groups ...string) (DeleteGroupResponses, error) {
	if len(groups) == 0 {
		return nil, nil
	}
	req := kmsg.NewPtrDeleteGroupsRequest()
	req.Groups = append(req.Groups, groups...)
	shards := cl.cl.RequestSharded(ctx, req)

	rs := make(map[string]DeleteGroupResponse)
	return rs, shardErrEach(req, shards, func(kr kmsg.Response) error {
		resp := kr.(*kmsg.DeleteGroupsResponse)
		for _, g := range resp.Groups {
			rs[g.Group] = DeleteGroupResponse{
				Group: g.Group,
				Err:   kerr.ErrorForCode(g.ErrorCode),
			}
		}
		return nil
	})
}

// OffsetResponse contains the response for an individual offset for offset
// methods.
type OffsetResponse struct {
	Offset
	Err error // Err is non-nil if the offset operation failed.
}

// OffsetResponses contains per-partition responses to offset methods.
type OffsetResponses map[string]map[int32]OffsetResponse

// Lookup returns the offset at t and p and whether it exists.
func (os OffsetResponses) Lookup(t string, p int32) (OffsetResponse, bool) {
	if len(os) == 0 {
		return OffsetResponse{}, false
	}
	ps := os[t]
	if len(ps) == 0 {
		return OffsetResponse{}, false
	}
	o, exists := ps[p]
	return o, exists
}

// Keep filters the responses to only keep the input offsets.
func (os OffsetResponses) Keep(o Offsets) {
	os.DeleteFunc(func(r OffsetResponse) bool {
		if len(o) == 0 {
			return true // keep nothing, delete
		}
		ot := o[r.Topic]
		if ot == nil {
			return true // topic missing, delete
		}
		_, ok := ot[r.Partition]
		return !ok // does not exist, delete
	})
}

// Offsets returns these offset responses as offsets.
func (os OffsetResponses) Offsets() Offsets {
	i := make(Offsets)
	os.Each(func(o OffsetResponse) {
		i.Add(o.Offset)
	})
	return i
}

// KOffsets returns these offset responses as a kgo offset map.
func (os OffsetResponses) KOffsets() map[string]map[int32]kgo.Offset {
	return os.Offsets().KOffsets()
}

// DeleteFunc keeps only the offsets for which fn returns true.
func (os OffsetResponses) KeepFunc(fn func(OffsetResponse) bool) {
	for t, ps := range os {
		for p, o := range ps {
			if !fn(o) {
				delete(ps, p)
			}
		}
		if len(ps) == 0 {
			delete(os, t)
		}
	}
}

// DeleteFunc deletes any offset for which fn returns true.
func (os OffsetResponses) DeleteFunc(fn func(OffsetResponse) bool) {
	os.KeepFunc(func(o OffsetResponse) bool { return !fn(o) })
}

// Add adds an offset for a given topic/partition to this OffsetResponses map
// (even if it exists).
func (os *OffsetResponses) Add(o OffsetResponse) {
	if *os == nil {
		*os = make(map[string]map[int32]OffsetResponse)
	}
	ot := (*os)[o.Topic]
	if ot == nil {
		ot = make(map[int32]OffsetResponse)
		(*os)[o.Topic] = ot
	}
	ot[o.Partition] = o
}

// EachError calls fn for every offset that as a non-nil error.
func (os OffsetResponses) EachError(fn func(o OffsetResponse)) {
	for _, ps := range os {
		for _, o := range ps {
			if o.Err != nil {
				fn(o)
			}
		}
	}
}

// Sorted returns the responses sorted by topic and partition.
func (os OffsetResponses) Sorted() []OffsetResponse {
	var s []OffsetResponse
	os.Each(func(o OffsetResponse) { s = append(s, o) })
	sort.Slice(s, func(i, j int) bool {
		return s[i].Topic < s[j].Topic ||
			s[i].Topic == s[j].Topic && s[i].Partition < s[j].Partition
	})
	return s
}

// Each calls fn for every offset.
func (os OffsetResponses) Each(fn func(OffsetResponse)) {
	for _, ps := range os {
		for _, o := range ps {
			fn(o)
		}
	}
}

// Partitions returns the set of unique topics and partitions in these offsets.
func (os OffsetResponses) Partitions() TopicsSet {
	s := make(TopicsSet)
	os.Each(func(o OffsetResponse) {
		s.Add(o.Topic, o.Partition)
	})
	return s
}

// Error iterates over all offsets and returns the first error encountered, if
// any. This can be used to check if an operation was entirely successful or
// not.
//
// Note that offset operations can be partially successful. For example, some
// offsets could succeed in an offset commit while others fail (maybe one topic
// does not exist for some reason, or you are not authorized for one topic). If
// this is something you need to worry about, you may need to check all offsets
// manually.
func (os OffsetResponses) Error() error {
	for _, ps := range os {
		for _, o := range ps {
			if o.Err != nil {
				return o.Err
			}
		}
	}
	return nil
}

// Ok returns true if there are no errors. This is a shortcut for os.Error() ==
// nil.
func (os OffsetResponses) Ok() bool {
	return os.Error() == nil
}

// CommitOffsets issues an offset commit request for the input offsets.
//
// This function can be used to manually commit offsets when directly consuming
// partitions outside of an actual consumer group. For example, if you assign
// partitions manually, but want still use Kafka to checkpoint what you have
// consumed, you can manually issue an offset commit request with this method.
//
// This does not return on authorization failures, instead, authorization
// failures are included in the responses.
func (cl *Client) CommitOffsets(ctx context.Context, group string, os Offsets) (OffsetResponses, error) {
	req := kmsg.NewPtrOffsetCommitRequest()
	req.Group = group
	for t, ps := range os {
		rt := kmsg.NewOffsetCommitRequestTopic()
		rt.Topic = t
		for p, o := range ps {
			rp := kmsg.NewOffsetCommitRequestTopicPartition()
			rp.Partition = p
			rp.Offset = o.At
			rp.LeaderEpoch = o.LeaderEpoch
			if len(o.Metadata) > 0 {
				rp.Metadata = kmsg.StringPtr(o.Metadata)
			}
			rt.Partitions = append(rt.Partitions, rp)
		}
		req.Topics = append(req.Topics, rt)
	}

	resp, err := req.RequestWith(ctx, cl.cl)
	if err != nil {
		return nil, err
	}

	rs := make(OffsetResponses)
	for _, t := range resp.Topics {
		rt := make(map[int32]OffsetResponse)
		rs[t.Topic] = rt
		for _, p := range t.Partitions {
			rt[p.Partition] = OffsetResponse{
				Offset: os[t.Topic][p.Partition],
				Err:    kerr.ErrorForCode(p.ErrorCode),
			}
		}
	}

	for t, ps := range os {
		respt := rs[t]
		if respt == nil {
			respt = make(map[int32]OffsetResponse)
			rs[t] = respt
		}
		for p, o := range ps {
			if _, exists := respt[p]; exists {
				continue
			}
			respt[p] = OffsetResponse{
				Offset: o,
				Err:    errOffsetCommitMissing,
			}
		}
	}

	return rs, nil
}

var errOffsetCommitMissing = errors.New("partition missing in commit response")

// CommitAllOffsets is identical to CommitOffsets, but returns an error if the
// offset commit was successful, but some offset within the commit failed to be
// committed.
//
// This is a shortcut function provided to avoid checking two errors, but you
// must be careful with this if partially successful commits can be a problem
// for you.
func (cl *Client) CommitAllOffsets(ctx context.Context, group string, os Offsets) error {
	commits, err := cl.CommitOffsets(ctx, group, os)
	if err != nil {
		return err
	}
	return commits.Error()
}

// FetchOffsets issues an offset fetch requests for all topics and partitions
// in the group. Because Kafka returns only partitions you are authorized to
// fetch, this only returns an auth error if you are not authorized to describe
// the group at all.
//
// This method requires talking to Kafka v0.11+.
func (cl *Client) FetchOffsets(ctx context.Context, group string) (OffsetResponses, error) {
	req := kmsg.NewPtrOffsetFetchRequest()
	req.Group = group
	resp, err := req.RequestWith(ctx, cl.cl)
	if err != nil {
		return nil, err
	}
	if err := maybeAuthErr(resp.ErrorCode); err != nil {
		return nil, err
	}
	if err := kerr.ErrorForCode(resp.ErrorCode); err != nil {
		return nil, err
	}
	rs := make(OffsetResponses)
	for _, t := range resp.Topics {
		rt := make(map[int32]OffsetResponse)
		rs[t.Topic] = rt
		for _, p := range t.Partitions {
			if err := maybeAuthErr(p.ErrorCode); err != nil {
				return nil, err
			}
			var meta string
			if p.Metadata != nil {
				meta = *p.Metadata
			}
			rt[p.Partition] = OffsetResponse{
				Offset: Offset{
					Topic:       t.Topic,
					Partition:   p.Partition,
					At:          p.Offset,
					LeaderEpoch: p.LeaderEpoch,
					Metadata:    meta,
				},
				Err: kerr.ErrorForCode(p.ErrorCode),
			}
		}
	}
	return rs, nil
}

// FetchOffsetsForTopics is a helper function that returns the currently
// committed offsets for the given group, as well as default -1 offsets for any
// topic/partition that does not yet have a commit.
//
// If any partition fetched or listed has an error, this function returns an
// error. The returned offset responses are ready to be used or converted
// directly to pure offsets with `Into`, and again into kgo offsets with
// another `Into`.
func (cl *Client) FetchOffsetsForTopics(ctx context.Context, group string, topics ...string) (OffsetResponses, error) {
	os := make(Offsets)

	if len(topics) > 0 {
		listed, err := cl.ListTopics(ctx, topics...)
		if err != nil {
			return nil, fmt.Errorf("unable to list topics: %w", err)
		}

		for _, topic := range topics {
			t := listed[topic]
			if t.Err != nil {
				return nil, fmt.Errorf("unable to describe topics, topic err: %w", t.Err)
			}
			for _, p := range t.Partitions {
				os.AddOffset(topic, p.Partition, -1, -1)
			}
		}
	}

	resps, err := cl.FetchOffsets(ctx, group)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch offsets: %w", err)
	}
	if err := resps.Error(); err != nil {
		return nil, fmt.Errorf("offset fetches had a load error, first error: %w", err)
	}
	os.Each(func(o Offset) {
		if _, ok := resps.Lookup(o.Topic, o.Partition); !ok {
			resps.Add(OffsetResponse{Offset: o})
		}
	})
	return resps, nil
}

// FetchOffsetsResponse contains a fetch offsets response for a single group.
type FetchOffsetsResponse struct {
	Group   string          // Group is the offsets these fetches correspond to.
	Fetched OffsetResponses // Fetched contains offsets fetched for this group, if any.
	Err     error           // Err contains any error preventing offsets from being fetched.
}

// CommittedPartitions returns the set of unique topics and partitions that
// have been committed to in this group.
func (r FetchOffsetsResponse) CommittedPartitions() TopicsSet {
	return r.Fetched.Partitions()
}

// FetchOFfsetsResponses contains responses for many fetch offsets requests.
type FetchOffsetsResponses map[string]FetchOffsetsResponse

// EachError calls fn for every response that as a non-nil error.
func (rs FetchOffsetsResponses) EachError(fn func(FetchOffsetsResponse)) {
	for _, r := range rs {
		if r.Err != nil {
			fn(r)
		}
	}
}

// AllFailed returns whether all fetch offsets requests failed.
func (rs FetchOffsetsResponses) AllFailed() bool {
	var n int
	rs.EachError(func(FetchOffsetsResponse) { n++ })
	return n == len(rs)
}

// CommittedPartitions returns the set of unique topics and partitions that
// have been committed to across all members in all responses. This is the
// all-group analogue to FetchOffsetsResponse.CommittedPartitions.
func (rs FetchOffsetsResponses) CommittedPartitions() TopicsSet {
	s := make(TopicsSet)
	for _, r := range rs {
		s.Merge(r.CommittedPartitions())
	}
	return s
}

// On calls fn for the response group if it exists, returning the response and
// the error returned from fn. If fn is nil, this simply returns the group.
//
// The fn is given a copy of the response. This function returns the copy as
// well; any modifications within fn are modifications on the returned copy.
//
// If the group does not exist, this returns kerr.GroupIDNotFound.
func (rs FetchOffsetsResponses) On(group string, fn func(*FetchOffsetsResponse) error) (FetchOffsetsResponse, error) {
	if len(rs) > 0 {
		r, ok := rs[group]
		if ok {
			if fn == nil {
				return r, nil
			}
			return r, fn(&r)
		}
	}
	return FetchOffsetsResponse{}, kerr.GroupIDNotFound
}

// FetchManyOffsets issues a fetch offsets requests for each group specified.
//
// This API is slightly different from others on the admin client: the
// underlying FetchOFfsets request only supports one group at a time. Unlike
// all other methods, which build and issue a single request, this method
// issues many requests and captures all responses into the return map
// (disregarding sharded functions, which actually have the input request
// split).
//
// More importantly, FetchOffsets and CommitOffsets are important to provide as
// simple APIs for users that manage group offsets outside of a consumer group.
// This function complements FetchOffsets by supporting the metric gathering
// use case of fetching offsets for many groups at once.
func (cl *Client) FetchManyOffsets(ctx context.Context, groups ...string) FetchOffsetsResponses {
	if len(groups) == 0 {
		return nil
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	fetched := make(FetchOffsetsResponses)
	for i := range groups {
		group := groups[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			offsets, err := cl.FetchOffsets(ctx, group)
			mu.Lock()
			defer mu.Unlock()
			fetched[group] = FetchOffsetsResponse{
				Group:   group,
				Fetched: offsets,
				Err:     err,
			}
		}()
	}
	wg.Wait()
	return fetched
}

// DeleteOffsetsResponses contains the per topic, per partition errors. If an
// offset deletion for a partition was successful, the error will be nil.
type DeleteOffsetsResponses map[string]map[int32]error

// Lookup returns the response at t and p and whether it exists.
func (ds DeleteOffsetsResponses) Lookup(t string, p int32) (error, bool) {
	if len(ds) == 0 {
		return nil, false
	}
	ps := ds[t]
	if len(ps) == 0 {
		return nil, false
	}
	r, exists := ps[p]
	return r, exists
}

// EachError calls fn for every partition that as a non-nil deletion error.
func (ds DeleteOffsetsResponses) EachError(fn func(string, int32, error)) {
	for t, ps := range ds {
		for p, err := range ps {
			if err != nil {
				fn(t, p, err)
			}
		}
	}
}

// DeleteOffsets deletes offsets for the given group.
//
// Originally, offset commits were persisted in Kafka for some retention time.
// This posed problematic for infrequently committing consumers, so the
// retention time concept was removed in Kafka v2.1 in favor of deleting
// offsets for a group only when the group became empty. However, if a group
// stops consuming from a topic, then the offsets will persist and lag
// monitoring for the group will notice an ever increasing amount of lag for
// these no-longer-consumed topics. Thus, Kafka v2.4 introduced an OffsetDelete
// request to allow admins to manually delete offsets for no longer consumed
// topics.
//
// This method requires talking to Kafka v2.4+. This returns an *AuthErr if the
// user is not authorized to delete offsets in the group at all. This does not
// return on per-topic authorization failures, instead, per-topic authorization
// failures are included in the responses.
func (cl *Client) DeleteOffsets(ctx context.Context, group string, s TopicsSet) (DeleteOffsetsResponses, error) {
	if len(s) == 0 {
		return nil, nil
	}

	req := kmsg.NewPtrOffsetDeleteRequest()
	req.Group = group
	for t, ps := range s {
		rt := kmsg.NewOffsetDeleteRequestTopic()
		rt.Topic = t
		for p := range ps {
			rp := kmsg.NewOffsetDeleteRequestTopicPartition()
			rp.Partition = p
			rt.Partitions = append(rt.Partitions, rp)
		}
		req.Topics = append(req.Topics, rt)
	}

	resp, err := req.RequestWith(ctx, cl.cl)
	if err != nil {
		return nil, err
	}
	if err := maybeAuthErr(resp.ErrorCode); err != nil {
		return nil, err
	}
	if err := kerr.ErrorForCode(resp.ErrorCode); err != nil {
		return nil, err
	}

	r := make(DeleteOffsetsResponses)
	for _, t := range resp.Topics {
		rt := make(map[int32]error)
		r[t.Topic] = rt
		for _, p := range t.Partitions {
			rt[p.Partition] = kerr.ErrorForCode(p.ErrorCode)
		}
	}
	return r, nil
}

// GroupMemberLag is the lag between a group member's current offset commit and
// the current end offset.
//
// If either the offset commits have load errors, or the listed end offsets
// have load errors, the Lag field will be -1 and the Err field will be set (to
// the first of either the commit error, or else the list error).
//
// If the group is in the Empty state, lag is calculated for all partitions in
// a topic, but the member is nil. The calculate function assumes that any
// assigned topic is meant to be entirely consumed. If the group is Empty and
// topics could not be listed, some partitions may be missing.
type GroupMemberLag struct {
	// Member is a reference to the group member consuming this partition.
	// If the group is in state Empty, the member will be nil.
	Member *DescribedGroupMember

	Commit Offset       // Commit is this member's current offset commit.
	End    ListedOffset // EndOffset is a reference to the end offset of this partition.
	Lag    int64        // Lag is how far behind this member is, or -1 if there is a commit error or list offset error.

	Err error // Err is either the commit error, or the list end offsets error, or nil.
}

// IsEmpty returns if the this lag is for a group in the Empty state.
func (g *GroupMemberLag) IsEmpty() bool { return g.Member == nil }

// GroupLag is the per-topic, per-partition lag of members in a group.
type GroupLag map[string]map[int32]GroupMemberLag

// Lookup returns the lag at t and p and whether it exists.
func (l GroupLag) Lookup(t string, p int32) (GroupMemberLag, bool) {
	if len(l) == 0 {
		return GroupMemberLag{}, false
	}
	ps := l[t]
	if len(ps) == 0 {
		return GroupMemberLag{}, false
	}
	m, exists := ps[p]
	return m, exists
}

// Sorted returns the per-topic, per-partition lag by member sorted in order by
// topic then partition.
func (l GroupLag) Sorted() []GroupMemberLag {
	var all []GroupMemberLag
	for _, ps := range l {
		for _, l := range ps {
			all = append(all, l)
		}
	}
	sort.Slice(all, func(i, j int) bool {
		l, r := all[i], all[j]
		if l.End.Topic < r.End.Topic {
			return true
		}
		if l.End.Topic > r.End.Topic {
			return false
		}
		return l.End.Partition < r.End.Partition
	})
	return all
}

// IsEmpty returns if the group is empty.
func (l GroupLag) IsEmpty() bool {
	for _, ps := range l {
		for _, m := range ps {
			return m.IsEmpty()
		}
	}
	return false
}

// Total returns the total lag across all topics.
func (l GroupLag) Total() int64 {
	var tot int64
	for _, tl := range l.TotalByTopic() {
		tot += tl.Lag
	}
	return tot
}

// TotalByTopic returns the total lag for each topic.
func (l GroupLag) TotalByTopic() GroupTopicsLag {
	m := make(map[string]TopicLag)
	for t, ps := range l {
		mt := TopicLag{
			Topic: t,
		}
		for _, l := range ps {
			if l.Lag > 0 {
				mt.Lag += l.Lag
			}
		}
		m[t] = mt
	}
	return m
}

// GroupTopicsLag is the total lag per topic within a group.
type GroupTopicsLag map[string]TopicLag

// TopicLag is the lag for an individual topic within a group.
type TopicLag struct {
	Topic string
	Lag   int64
}

// Sorted returns the per-topic lag, sorted by topic.
func (l GroupTopicsLag) Sorted() []TopicLag {
	var all []TopicLag
	for _, tl := range l {
		all = append(all, tl)
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].Topic < all[j].Topic
	})
	return all
}

// CalculateGroupLag returns the per-partition lag of all members in a group.
// The input to this method is the returns from the three following methods,
//
//     described := DescribeGroups(ctx, group)
//     fetched := FetchOffsets(ctx, group)
//     toList := described.AssignedPartitions()
//     toList.Merge(fetched.CommittedPartitions()
//     ListEndOffsets(ctx, toList.Topics())
//
// If assigned partitions are missing in the listed end offsets listed end
// offsets, the partition will have an error indicating it is missing. A
// missing topic or partition in the commits is assumed to be nothing
// committing yet.
func CalculateGroupLag(
	group DescribedGroup,
	commit OffsetResponses,
	offsets ListedOffsets,
) GroupLag {
	if group.State == "Empty" {
		return calculateEmptyLag(commit, offsets)
	}
	if commit == nil { // avoid panics below
		commit = make(OffsetResponses)
	}
	if offsets == nil {
		offsets = make(ListedOffsets)
	}

	l := make(map[string]map[int32]GroupMemberLag)
	for mi, m := range group.Members {
		c, ok := m.Assigned.AsConsumer()
		if !ok {
			continue
		}
		for _, t := range c.Topics {
			lt := l[t.Topic]
			if lt == nil {
				lt = make(map[int32]GroupMemberLag)
				l[t.Topic] = lt
			}

			tcommit := commit[t.Topic]
			tend := offsets[t.Topic]
			for _, p := range t.Partitions {
				var (
					pcommit OffsetResponse
					pend    ListedOffset
					perr    error
					ok      bool
				)

				if tcommit != nil {
					if pcommit, ok = tcommit[p]; !ok {
						pcommit = OffsetResponse{Offset: Offset{
							Topic:     t.Topic,
							Partition: p,
							At:        -1,
						}}
					}
				}
				if tend == nil {
					perr = errListMissing
				} else {
					if pend, ok = tend[p]; !ok {
						perr = errListMissing
					}
				}

				if perr == nil {
					if perr = pcommit.Err; perr == nil {
						perr = pend.Err
					}
				}

				lag := int64(-1)
				if perr == nil {
					lag = pend.Offset
					if pcommit.At >= 0 {
						lag = pend.Offset - pcommit.At
					}
				}

				lt[p] = GroupMemberLag{
					Member: &group.Members[mi],
					Commit: pcommit.Offset,
					End:    pend,
					Lag:    lag,
					Err:    perr,
				}

			}
		}
	}

	return l
}

func calculateEmptyLag(commit OffsetResponses, offsets ListedOffsets) GroupLag {
	l := make(map[string]map[int32]GroupMemberLag)
	for t, ps := range commit {
		lt := l[t]
		if lt == nil {
			lt = make(map[int32]GroupMemberLag)
			l[t] = lt
		}
		tend := offsets[t]
		for p, pcommit := range ps {
			var (
				pend ListedOffset
				perr error
				ok   bool
			)
			if tend == nil {
				perr = errListMissing
			} else {
				if pend, ok = tend[p]; !ok {
					perr = errListMissing
				}
			}
			if perr == nil {
				if perr = pcommit.Err; perr == nil {
					perr = pend.Err
				}
			}

			lag := int64(-1)
			if perr == nil {
				lag = pend.Offset
				if pcommit.At >= 0 {
					lag = pend.Offset - pcommit.At
				}
			}

			lt[p] = GroupMemberLag{
				Commit: pcommit.Offset,
				End:    pend,
				Lag:    lag,
				Err:    perr,
			}
		}
	}

	// Now we look at all topics that we calculated lag for, and check out
	// the partitions we listed. If those partitions are missing from the
	// lag calculations above, the partitions were not committed to and we
	// count that as entirely lagging.
	for t, lt := range l {
		tend := offsets[t]
		for p, pend := range tend {
			if _, ok := lt[p]; ok {
				continue
			}
			pcommit := Offset{
				Topic:       t,
				Partition:   p,
				At:          -1,
				LeaderEpoch: -1,
			}
			perr := pend.Err
			lag := int64(-1)
			if perr == nil {
				lag = pend.Offset
			}
			lt[p] = GroupMemberLag{
				Commit: pcommit,
				End:    pend,
				Lag:    lag,
				Err:    perr,
			}
		}
	}

	return l
}

var errListMissing = errors.New("missing from list offsets")
