// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: website.proto

// proto 包名

package websiterpc

import (
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

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{0}
}

type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResp.ProtoReflect.Descriptor instead.
func (*EmptyResp) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{1}
}

type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{2}
}

func (x *IdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsReq) Reset() {
	*x = IdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsReq) ProtoMessage() {}

func (x *IdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsReq.ProtoReflect.Descriptor instead.
func (*IdsReq) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{3}
}

func (x *IdsReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserIdReq) Reset() {
	*x = UserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdReq) ProtoMessage() {}

func (x *UserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdReq.ProtoReflect.Descriptor instead.
func (*UserIdReq) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{4}
}

func (x *UserIdReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BatchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessCount int64 `protobuf:"varint,1,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
}

func (x *BatchResp) Reset() {
	*x = BatchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResp) ProtoMessage() {}

func (x *BatchResp) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResp.ProtoReflect.Descriptor instead.
func (*BatchResp) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{5}
}

func (x *BatchResp) GetSuccessCount() int64 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

type CountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountResp) Reset() {
	*x = CountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResp) ProtoMessage() {}

func (x *CountResp) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountResp.ProtoReflect.Descriptor instead.
func (*CountResp) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{6}
}

func (x *CountResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UserVisit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date      string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	ViewCount int64  `protobuf:"varint,2,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
}

func (x *UserVisit) Reset() {
	*x = UserVisit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVisit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVisit) ProtoMessage() {}

func (x *UserVisit) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVisit.ProtoReflect.Descriptor instead.
func (*UserVisit) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{7}
}

func (x *UserVisit) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *UserVisit) GetViewCount() int64 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

type UserDailyVisitRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*UserVisit `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserDailyVisitRsp) Reset() {
	*x = UserDailyVisitRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_website_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDailyVisitRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDailyVisitRsp) ProtoMessage() {}

func (x *UserDailyVisitRsp) ProtoReflect() protoreflect.Message {
	mi := &file_website_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDailyVisitRsp.ProtoReflect.Descriptor instead.
func (*UserDailyVisitRsp) Descriptor() ([]byte, []int) {
	return file_website_proto_rawDescGZIP(), []int{8}
}

func (x *UserDailyVisitRsp) GetList() []*UserVisit {
	if x != nil {
		return x.List
	}
	return nil
}

var File_website_proto protoreflect.FileDescriptor

var file_website_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x72, 0x70, 0x63, 0x22, 0x0a, 0x0a, 0x08, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x22, 0x0b, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a,
	0x06, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x24, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x30, 0x0a, 0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x21, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x73, 0x69,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x32, 0x56, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52,
	0x70, 0x63, 0x12, 0x48, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x73, 0x70, 0x42, 0x0e, 0x5a, 0x0c,
	0x2e, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_website_proto_rawDescOnce sync.Once
	file_website_proto_rawDescData = file_website_proto_rawDesc
)

func file_website_proto_rawDescGZIP() []byte {
	file_website_proto_rawDescOnce.Do(func() {
		file_website_proto_rawDescData = protoimpl.X.CompressGZIP(file_website_proto_rawDescData)
	})
	return file_website_proto_rawDescData
}

var file_website_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_website_proto_goTypes = []any{
	(*EmptyReq)(nil),          // 0: websiterpc.EmptyReq
	(*EmptyResp)(nil),         // 1: websiterpc.EmptyResp
	(*IdReq)(nil),             // 2: websiterpc.IdReq
	(*IdsReq)(nil),            // 3: websiterpc.IdsReq
	(*UserIdReq)(nil),         // 4: websiterpc.UserIdReq
	(*BatchResp)(nil),         // 5: websiterpc.BatchResp
	(*CountResp)(nil),         // 6: websiterpc.CountResp
	(*UserVisit)(nil),         // 7: websiterpc.UserVisit
	(*UserDailyVisitRsp)(nil), // 8: websiterpc.UserDailyVisitRsp
}
var file_website_proto_depIdxs = []int32{
	7, // 0: websiterpc.UserDailyVisitRsp.list:type_name -> websiterpc.UserVisit
	0, // 1: websiterpc.WebsiteRpc.GetUserDailyVisit:input_type -> websiterpc.EmptyReq
	8, // 2: websiterpc.WebsiteRpc.GetUserDailyVisit:output_type -> websiterpc.UserDailyVisitRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_website_proto_init() }
func file_website_proto_init() {
	if File_website_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_website_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyReq); i {
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
		file_website_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyResp); i {
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
		file_website_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IdReq); i {
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
		file_website_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IdsReq); i {
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
		file_website_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserIdReq); i {
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
		file_website_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BatchResp); i {
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
		file_website_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CountResp); i {
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
		file_website_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UserVisit); i {
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
		file_website_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UserDailyVisitRsp); i {
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
			RawDescriptor: file_website_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_website_proto_goTypes,
		DependencyIndexes: file_website_proto_depIdxs,
		MessageInfos:      file_website_proto_msgTypes,
	}.Build()
	File_website_proto = out.File
	file_website_proto_rawDesc = nil
	file_website_proto_goTypes = nil
	file_website_proto_depIdxs = nil
}
