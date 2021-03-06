// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: proto/exam.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int32 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"` //¤ skal en user have lamport time, når vi skal broadcaste bids til andre replication managers
	Time   int32 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_exam_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserid() int32 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *User) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type Incval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Incval) Reset() {
	*x = Incval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Incval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Incval) ProtoMessage() {}

func (x *Incval) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Incval.ProtoReflect.Descriptor instead.
func (*Incval) Descriptor() ([]byte, []int) {
	return file_proto_exam_proto_rawDescGZIP(), []int{1}
}

func (x *Incval) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_exam_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type Grant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Grant) Reset() {
	*x = Grant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grant) ProtoMessage() {}

func (x *Grant) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grant.ProtoReflect.Descriptor instead.
func (*Grant) Descriptor() ([]byte, []int) {
	return file_proto_exam_proto_rawDescGZIP(), []int{3}
}

func (x *Grant) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_proto_exam_proto_rawDescGZIP(), []int{4}
}

func (x *Release) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exam_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exam_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_proto_exam_proto_rawDescGZIP(), []int{5}
}

func (x *State) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *State) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_exam_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_exam_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_exam_proto_rawDescGZIP(), []int{6}
}

var File_proto_exam_proto protoreflect.FileDescriptor

var file_proto_exam_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x20, 0x0a,
	0x06, 0x49, 0x6e, 0x63, 0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x2a, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x05, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x40, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x8a, 0x02, 0x0a,
	0x04, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x28, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x63, 0x76, 0x61, 0x6c, 0x12,
	0x2e, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x30, 0x01, 0x12,
	0x2c, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x29, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x0e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_exam_proto_rawDescOnce sync.Once
	file_proto_exam_proto_rawDescData = file_proto_exam_proto_rawDesc
)

func file_proto_exam_proto_rawDescGZIP() []byte {
	file_proto_exam_proto_rawDescOnce.Do(func() {
		file_proto_exam_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_exam_proto_rawDescData)
	})
	return file_proto_exam_proto_rawDescData
}

var file_proto_exam_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_exam_proto_goTypes = []interface{}{
	(*User)(nil),    // 0: proto.User
	(*Incval)(nil),  // 1: proto.Incval
	(*Request)(nil), // 2: proto.Request
	(*Grant)(nil),   // 3: proto.Grant
	(*Release)(nil), // 4: proto.Release
	(*State)(nil),   // 5: proto.State
	(*Empty)(nil),   // 6: proto.Empty
}
var file_proto_exam_proto_depIdxs = []int32{
	0,  // 0: proto.Request.user:type_name -> proto.User
	0,  // 1: proto.Grant.user:type_name -> proto.User
	0,  // 2: proto.Release.user:type_name -> proto.User
	0,  // 3: proto.State.user:type_name -> proto.User
	6,  // 4: proto.Exam.Increment:input_type -> proto.Empty
	2,  // 5: proto.Exam.RequestToken:input_type -> proto.Request
	4,  // 6: proto.Exam.ReleaseToken:input_type -> proto.Release
	6,  // 7: proto.Exam.Ping:input_type -> proto.Empty
	5,  // 8: proto.Exam.Coordinator:input_type -> proto.State
	6,  // 9: proto.Exam.RequestLamport:input_type -> proto.Empty
	1,  // 10: proto.Exam.Increment:output_type -> proto.Incval
	3,  // 11: proto.Exam.RequestToken:output_type -> proto.Grant
	6,  // 12: proto.Exam.ReleaseToken:output_type -> proto.Empty
	6,  // 13: proto.Exam.Ping:output_type -> proto.Empty
	6,  // 14: proto.Exam.Coordinator:output_type -> proto.Empty
	0,  // 15: proto.Exam.RequestLamport:output_type -> proto.User
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_exam_proto_init() }
func file_proto_exam_proto_init() {
	if File_proto_exam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_exam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_exam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Incval); i {
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
		file_proto_exam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_exam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grant); i {
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
		file_proto_exam_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
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
		file_proto_exam_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_proto_exam_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_exam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_exam_proto_goTypes,
		DependencyIndexes: file_proto_exam_proto_depIdxs,
		MessageInfos:      file_proto_exam_proto_msgTypes,
	}.Build()
	File_proto_exam_proto = out.File
	file_proto_exam_proto_rawDesc = nil
	file_proto_exam_proto_goTypes = nil
	file_proto_exam_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExamClient is the client API for Exam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExamClient interface {
	Increment(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Incval, error)
	RequestToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (Exam_RequestTokenClient, error)
	ReleaseToken(ctx context.Context, in *Release, opts ...grpc.CallOption) (*Empty, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Coordinator(ctx context.Context, in *State, opts ...grpc.CallOption) (*Empty, error)
	RequestLamport(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error)
}

type examClient struct {
	cc grpc.ClientConnInterface
}

func NewExamClient(cc grpc.ClientConnInterface) ExamClient {
	return &examClient{cc}
}

func (c *examClient) Increment(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Incval, error) {
	out := new(Incval)
	err := c.cc.Invoke(ctx, "/proto.Exam/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examClient) RequestToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (Exam_RequestTokenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Exam_serviceDesc.Streams[0], "/proto.Exam/RequestToken", opts...)
	if err != nil {
		return nil, err
	}
	x := &examRequestTokenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Exam_RequestTokenClient interface {
	Recv() (*Grant, error)
	grpc.ClientStream
}

type examRequestTokenClient struct {
	grpc.ClientStream
}

func (x *examRequestTokenClient) Recv() (*Grant, error) {
	m := new(Grant)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *examClient) ReleaseToken(ctx context.Context, in *Release, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Exam/ReleaseToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Exam/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examClient) Coordinator(ctx context.Context, in *State, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Exam/Coordinator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examClient) RequestLamport(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.Exam/RequestLamport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExamServer is the server API for Exam service.
type ExamServer interface {
	Increment(context.Context, *Empty) (*Incval, error)
	RequestToken(*Request, Exam_RequestTokenServer) error
	ReleaseToken(context.Context, *Release) (*Empty, error)
	Ping(context.Context, *Empty) (*Empty, error)
	Coordinator(context.Context, *State) (*Empty, error)
	RequestLamport(context.Context, *Empty) (*User, error)
}

// UnimplementedExamServer can be embedded to have forward compatible implementations.
type UnimplementedExamServer struct {
}

func (*UnimplementedExamServer) Increment(context.Context, *Empty) (*Incval, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (*UnimplementedExamServer) RequestToken(*Request, Exam_RequestTokenServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestToken not implemented")
}
func (*UnimplementedExamServer) ReleaseToken(context.Context, *Release) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseToken not implemented")
}
func (*UnimplementedExamServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedExamServer) Coordinator(context.Context, *State) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coordinator not implemented")
}
func (*UnimplementedExamServer) RequestLamport(context.Context, *Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLamport not implemented")
}

func RegisterExamServer(s *grpc.Server, srv ExamServer) {
	s.RegisterService(&_Exam_serviceDesc, srv)
}

func _Exam_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Exam/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServer).Increment(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exam_RequestToken_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExamServer).RequestToken(m, &examRequestTokenServer{stream})
}

type Exam_RequestTokenServer interface {
	Send(*Grant) error
	grpc.ServerStream
}

type examRequestTokenServer struct {
	grpc.ServerStream
}

func (x *examRequestTokenServer) Send(m *Grant) error {
	return x.ServerStream.SendMsg(m)
}

func _Exam_ReleaseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Release)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServer).ReleaseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Exam/ReleaseToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServer).ReleaseToken(ctx, req.(*Release))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exam_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Exam/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exam_Coordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServer).Coordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Exam/Coordinator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServer).Coordinator(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exam_RequestLamport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServer).RequestLamport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Exam/RequestLamport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServer).RequestLamport(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exam_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Exam",
	HandlerType: (*ExamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _Exam_Increment_Handler,
		},
		{
			MethodName: "ReleaseToken",
			Handler:    _Exam_ReleaseToken_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Exam_Ping_Handler,
		},
		{
			MethodName: "Coordinator",
			Handler:    _Exam_Coordinator_Handler,
		},
		{
			MethodName: "RequestLamport",
			Handler:    _Exam_RequestLamport_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestToken",
			Handler:       _Exam_RequestToken_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/exam.proto",
}
