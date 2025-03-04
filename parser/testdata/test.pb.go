// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: test.proto

package testdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Null           *NestedTest                       `protobuf:"bytes,1,opt,name=null,proto3" json:"null,omitempty"`
	BoolTrue       bool                              `protobuf:"varint,2,opt,name=bool_true,json=boolTrue,proto3" json:"bool_true,omitempty"`
	BoolFalse      bool                              `protobuf:"varint,3,opt,name=bool_false,json=boolFalse,proto3" json:"bool_false,omitempty"`
	NumInt32       int32                             `protobuf:"varint,4,opt,name=num_int32,json=numInt32,proto3" json:"num_int32,omitempty"`
	NumInt64       int64                             `protobuf:"varint,5,opt,name=num_int64,json=numInt64,proto3" json:"num_int64,omitempty"`
	NumFloat       float32                           `protobuf:"fixed32,6,opt,name=num_float,json=numFloat,proto3" json:"num_float,omitempty"`
	NumDouble      float64                           `protobuf:"fixed64,7,opt,name=num_double,json=numDouble,proto3" json:"num_double,omitempty"`
	NumUint32      uint32                            `protobuf:"varint,8,opt,name=num_uint32,json=numUint32,proto3" json:"num_uint32,omitempty"`
	NumUint64      uint64                            `protobuf:"varint,9,opt,name=num_uint64,json=numUint64,proto3" json:"num_uint64,omitempty"`
	Str            string                            `protobuf:"bytes,10,opt,name=str,proto3" json:"str,omitempty"`
	StrDate        string                            `protobuf:"bytes,11,opt,name=str_date,json=strDate,proto3" json:"str_date,omitempty"`
	Timestamp      *timestamppb.Timestamp            `protobuf:"bytes,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Obj            *NestedTest                       `protobuf:"bytes,13,opt,name=obj,proto3" json:"obj,omitempty"`
	ArrayEmpty     []int32                           `protobuf:"varint,14,rep,packed,name=array_empty,json=arrayEmpty,proto3" json:"array_empty,omitempty"`
	ArrayBool      []bool                            `protobuf:"varint,15,rep,packed,name=array_bool,json=arrayBool,proto3" json:"array_bool,omitempty"`
	ArrayNumInt32  []int32                           `protobuf:"varint,16,rep,packed,name=array_num_int32,json=arrayNumInt32,proto3" json:"array_num_int32,omitempty"`
	ArrayNumInt64  []int64                           `protobuf:"varint,17,rep,packed,name=array_num_int64,json=arrayNumInt64,proto3" json:"array_num_int64,omitempty"`
	ArrayNumFloat  []float32                         `protobuf:"fixed32,18,rep,packed,name=array_num_float,json=arrayNumFloat,proto3" json:"array_num_float,omitempty"`
	ArrayNumDouble []float64                         `protobuf:"fixed64,19,rep,packed,name=array_num_double,json=arrayNumDouble,proto3" json:"array_num_double,omitempty"`
	ArrayNumUint32 []uint32                          `protobuf:"varint,20,rep,packed,name=array_num_uint32,json=arrayNumUint32,proto3" json:"array_num_uint32,omitempty"`
	ArrayNumUint64 []uint64                          `protobuf:"varint,21,rep,packed,name=array_num_uint64,json=arrayNumUint64,proto3" json:"array_num_uint64,omitempty"`
	ArrayStr       []string                          `protobuf:"bytes,22,rep,name=array_str,json=arrayStr,proto3" json:"array_str,omitempty"`
	ArrayTimestamp []*timestamppb.Timestamp          `protobuf:"bytes,23,rep,name=array_timestamp,json=arrayTimestamp,proto3" json:"array_timestamp,omitempty"`
	Uuid           string                            `protobuf:"bytes,24,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ArrayUuid      []string                          `protobuf:"bytes,25,rep,name=array_uuid,json=arrayUuid,proto3" json:"array_uuid,omitempty"`
	Ipv4           string                            `protobuf:"bytes,26,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	ArrayIpv4      []string                          `protobuf:"bytes,27,rep,name=array_ipv4,json=arrayIpv4,proto3" json:"array_ipv4,omitempty"`
	Ipv6           string                            `protobuf:"bytes,28,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	ArrayIpv6      []string                          `protobuf:"bytes,29,rep,name=array_ipv6,json=arrayIpv6,proto3" json:"array_ipv6,omitempty"`
	StrTime        string                            `protobuf:"bytes,30,opt,name=str_time,json=strTime,proto3" json:"str_time,omitempty"`
	MapStrStr      map[string]string                 `protobuf:"bytes,31,rep,name=map_str_str,json=mapStrStr,proto3" json:"map_str_str,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStrUint32   map[string]uint32                 `protobuf:"bytes,32,rep,name=map_str_uint32,json=mapStrUint32,proto3" json:"map_str_uint32,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStrUint64   map[string]uint64                 `protobuf:"bytes,33,rep,name=map_str_uint64,json=mapStrUint64,proto3" json:"map_str_uint64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStrInt32    map[string]int32                  `protobuf:"bytes,34,rep,name=map_str_int32,json=mapStrInt32,proto3" json:"map_str_int32,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStrInt64    map[string]int64                  `protobuf:"bytes,35,rep,name=map_str_int64,json=mapStrInt64,proto3" json:"map_str_int64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStrFloat    map[string]float32                `protobuf:"bytes,36,rep,name=map_str_float,json=mapStrFloat,proto3" json:"map_str_float,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	MapStrDouble   map[string]float64                `protobuf:"bytes,37,rep,name=map_str_double,json=mapStrDouble,proto3" json:"map_str_double,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	MapInt64Str    map[int64]string                  `protobuf:"bytes,38,rep,name=map_int64_str,json=mapInt64Str,proto3" json:"map_int64_str,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStrBool     map[string]bool                   `protobuf:"bytes,39,rep,name=map_str_bool,json=mapStrBool,proto3" json:"map_str_bool,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStrDate     map[string]*timestamppb.Timestamp `protobuf:"bytes,40,rep,name=map_str_date,json=mapStrDate,proto3" json:"map_str_date,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapStrObj      map[string]*NestedTest            `protobuf:"bytes,41,rep,name=map_str_obj,json=mapStrObj,proto3" json:"map_str_obj,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetNull() *NestedTest {
	if x != nil {
		return x.Null
	}
	return nil
}

func (x *Test) GetBoolTrue() bool {
	if x != nil {
		return x.BoolTrue
	}
	return false
}

func (x *Test) GetBoolFalse() bool {
	if x != nil {
		return x.BoolFalse
	}
	return false
}

func (x *Test) GetNumInt32() int32 {
	if x != nil {
		return x.NumInt32
	}
	return 0
}

func (x *Test) GetNumInt64() int64 {
	if x != nil {
		return x.NumInt64
	}
	return 0
}

func (x *Test) GetNumFloat() float32 {
	if x != nil {
		return x.NumFloat
	}
	return 0
}

func (x *Test) GetNumDouble() float64 {
	if x != nil {
		return x.NumDouble
	}
	return 0
}

func (x *Test) GetNumUint32() uint32 {
	if x != nil {
		return x.NumUint32
	}
	return 0
}

func (x *Test) GetNumUint64() uint64 {
	if x != nil {
		return x.NumUint64
	}
	return 0
}

func (x *Test) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *Test) GetStrDate() string {
	if x != nil {
		return x.StrDate
	}
	return ""
}

func (x *Test) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Test) GetObj() *NestedTest {
	if x != nil {
		return x.Obj
	}
	return nil
}

func (x *Test) GetArrayEmpty() []int32 {
	if x != nil {
		return x.ArrayEmpty
	}
	return nil
}

func (x *Test) GetArrayBool() []bool {
	if x != nil {
		return x.ArrayBool
	}
	return nil
}

func (x *Test) GetArrayNumInt32() []int32 {
	if x != nil {
		return x.ArrayNumInt32
	}
	return nil
}

func (x *Test) GetArrayNumInt64() []int64 {
	if x != nil {
		return x.ArrayNumInt64
	}
	return nil
}

func (x *Test) GetArrayNumFloat() []float32 {
	if x != nil {
		return x.ArrayNumFloat
	}
	return nil
}

func (x *Test) GetArrayNumDouble() []float64 {
	if x != nil {
		return x.ArrayNumDouble
	}
	return nil
}

func (x *Test) GetArrayNumUint32() []uint32 {
	if x != nil {
		return x.ArrayNumUint32
	}
	return nil
}

func (x *Test) GetArrayNumUint64() []uint64 {
	if x != nil {
		return x.ArrayNumUint64
	}
	return nil
}

func (x *Test) GetArrayStr() []string {
	if x != nil {
		return x.ArrayStr
	}
	return nil
}

func (x *Test) GetArrayTimestamp() []*timestamppb.Timestamp {
	if x != nil {
		return x.ArrayTimestamp
	}
	return nil
}

func (x *Test) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Test) GetArrayUuid() []string {
	if x != nil {
		return x.ArrayUuid
	}
	return nil
}

func (x *Test) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *Test) GetArrayIpv4() []string {
	if x != nil {
		return x.ArrayIpv4
	}
	return nil
}

func (x *Test) GetIpv6() string {
	if x != nil {
		return x.Ipv6
	}
	return ""
}

func (x *Test) GetArrayIpv6() []string {
	if x != nil {
		return x.ArrayIpv6
	}
	return nil
}

func (x *Test) GetStrTime() string {
	if x != nil {
		return x.StrTime
	}
	return ""
}

func (x *Test) GetMapStrStr() map[string]string {
	if x != nil {
		return x.MapStrStr
	}
	return nil
}

func (x *Test) GetMapStrUint32() map[string]uint32 {
	if x != nil {
		return x.MapStrUint32
	}
	return nil
}

func (x *Test) GetMapStrUint64() map[string]uint64 {
	if x != nil {
		return x.MapStrUint64
	}
	return nil
}

func (x *Test) GetMapStrInt32() map[string]int32 {
	if x != nil {
		return x.MapStrInt32
	}
	return nil
}

func (x *Test) GetMapStrInt64() map[string]int64 {
	if x != nil {
		return x.MapStrInt64
	}
	return nil
}

func (x *Test) GetMapStrFloat() map[string]float32 {
	if x != nil {
		return x.MapStrFloat
	}
	return nil
}

func (x *Test) GetMapStrDouble() map[string]float64 {
	if x != nil {
		return x.MapStrDouble
	}
	return nil
}

func (x *Test) GetMapInt64Str() map[int64]string {
	if x != nil {
		return x.MapInt64Str
	}
	return nil
}

func (x *Test) GetMapStrBool() map[string]bool {
	if x != nil {
		return x.MapStrBool
	}
	return nil
}

func (x *Test) GetMapStrDate() map[string]*timestamppb.Timestamp {
	if x != nil {
		return x.MapStrDate
	}
	return nil
}

func (x *Test) GetMapStrObj() map[string]*NestedTest {
	if x != nil {
		return x.MapStrObj
	}
	return nil
}

type NestedTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *NestedTest) Reset() {
	*x = NestedTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedTest) ProtoMessage() {}

func (x *NestedTest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedTest.ProtoReflect.Descriptor instead.
func (*NestedTest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *NestedTest) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x76, 0x69,
	0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x17,
	0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68,
	0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x74,
	0x72, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x54,
	0x72, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x61, 0x6c, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x61, 0x6c,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6e, 0x75, 0x6d, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d,
	0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6e,
	0x75, 0x6d, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x75,
	0x6d, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x75, 0x6d,
	0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4c, 0x0a,
	0x03, 0x6f, 0x62, 0x6a, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x76, 0x69, 0x72,
	0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x03, 0x6f, 0x62, 0x6a, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x08,
	0x52, 0x09, 0x61, 0x72, 0x72, 0x61, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x10,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4e, 0x75, 0x6d, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x6e, 0x75, 0x6d,
	0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x11, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x26, 0x0a, 0x0f, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x12,
	0x20, 0x03, 0x28, 0x02, 0x52, 0x0d, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4e, 0x75, 0x6d, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x6e, 0x75, 0x6d,
	0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x13, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0e, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x4e, 0x75, 0x6d, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4e, 0x75,
	0x6d, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x15, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x0e, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4e, 0x75, 0x6d, 0x55, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x16,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x72, 0x61, 0x79, 0x53, 0x74, 0x72, 0x12, 0x43,
	0x0a, 0x0f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x61, 0x72, 0x72, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x19, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x49, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76,
	0x36, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x18, 0x1d, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x72, 0x72, 0x61, 0x79, 0x49, 0x70, 0x76, 0x36, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x63, 0x0a, 0x0b, 0x6d, 0x61, 0x70, 0x5f, 0x73,
	0x74, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x76,
	0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x53, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x09, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x53, 0x74, 0x72, 0x12, 0x6c, 0x0a, 0x0e,
	0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x20,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68,
	0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x61,
	0x70, 0x53, 0x74, 0x72, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x6c, 0x0a, 0x0e, 0x6d, 0x61,
	0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x21, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x46, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63,
	0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72,
	0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x55,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x61, 0x70, 0x53,
	0x74, 0x72, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x69, 0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x5f,
	0x73, 0x74, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x22, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x45, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x69, 0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x76, 0x69, 0x72,
	0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x69,
	0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18,
	0x24, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63,
	0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53,
	0x74, 0x72, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x61,
	0x70, 0x53, 0x74, 0x72, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x6c, 0x0a, 0x0e, 0x6d, 0x61, 0x70,
	0x5f, 0x73, 0x74, 0x72, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x25, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x46, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x69, 0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x5f, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x26, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45,
	0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x53, 0x74, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x53,
	0x74, 0x72, 0x12, 0x66, 0x0a, 0x0c, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x62, 0x6f,
	0x6f, 0x6c, 0x18, 0x27, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f,
	0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f,
	0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d,
	0x61, 0x70, 0x53, 0x74, 0x72, 0x42, 0x6f, 0x6f, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x66, 0x0a, 0x0c, 0x6d, 0x61,
	0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x44, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x63, 0x0a, 0x0b, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x6f, 0x62,
	0x6a, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74,
	0x65, 0x63, 0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73,
	0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61,
	0x70, 0x53, 0x74, 0x72, 0x4f, 0x62, 0x6a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6d, 0x61,
	0x70, 0x53, 0x74, 0x72, 0x4f, 0x62, 0x6a, 0x1a, 0x3c, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x53, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72,
	0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x4d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x53, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x53,
	0x74, 0x72, 0x42, 0x6f, 0x6f, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x59, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x53, 0x74,
	0x72, 0x44, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x78, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x4f, 0x62, 0x6a, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x50, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x76, 0x69, 0x72, 0x75, 0x5f, 0x74, 0x65, 0x63,
	0x68, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1e, 0x0a, 0x0a,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x42, 0x2d, 0x5a, 0x2b,
	0x76, 0x69, 0x72, 0x75, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_test_proto_goTypes = []any{
	(*Test)(nil),                  // 0: viru_tech.clickhouse_sinker.parser.testdata.v1.Test
	(*NestedTest)(nil),            // 1: viru_tech.clickhouse_sinker.parser.testdata.v1.NestedTest
	nil,                           // 2: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrStrEntry
	nil,                           // 3: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrUint32Entry
	nil,                           // 4: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrUint64Entry
	nil,                           // 5: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrInt32Entry
	nil,                           // 6: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrInt64Entry
	nil,                           // 7: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrFloatEntry
	nil,                           // 8: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrDoubleEntry
	nil,                           // 9: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapInt64StrEntry
	nil,                           // 10: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrBoolEntry
	nil,                           // 11: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrDateEntry
	nil,                           // 12: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrObjEntry
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_test_proto_depIdxs = []int32{
	1,  // 0: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.null:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.NestedTest
	13, // 1: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 2: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.obj:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.NestedTest
	13, // 3: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.array_timestamp:type_name -> google.protobuf.Timestamp
	2,  // 4: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_str:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrStrEntry
	3,  // 5: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_uint32:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrUint32Entry
	4,  // 6: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_uint64:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrUint64Entry
	5,  // 7: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_int32:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrInt32Entry
	6,  // 8: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_int64:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrInt64Entry
	7,  // 9: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_float:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrFloatEntry
	8,  // 10: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_double:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrDoubleEntry
	9,  // 11: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_int64_str:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapInt64StrEntry
	10, // 12: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_bool:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrBoolEntry
	11, // 13: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_date:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrDateEntry
	12, // 14: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.map_str_obj:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrObjEntry
	13, // 15: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrDateEntry.value:type_name -> google.protobuf.Timestamp
	1,  // 16: viru_tech.clickhouse_sinker.parser.testdata.v1.Test.MapStrObjEntry.value:type_name -> viru_tech.clickhouse_sinker.parser.testdata.v1.NestedTest
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Test); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NestedTest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
