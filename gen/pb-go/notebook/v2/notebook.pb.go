// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: notebook.proto

package v2

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Principal_Type int32

const (
	Principal_USER Principal_Type = 0
)

// Enum value maps for Principal_Type.
var (
	Principal_Type_name = map[int32]string{
		0: "USER",
	}
	Principal_Type_value = map[string]int32{
		"USER": 0,
	}
)

func (x Principal_Type) Enum() *Principal_Type {
	p := new(Principal_Type)
	*p = x
	return p
}

func (x Principal_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Principal_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_notebook_proto_enumTypes[0].Descriptor()
}

func (Principal_Type) Type() protoreflect.EnumType {
	return &file_notebook_proto_enumTypes[0]
}

func (x Principal_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Principal_Type.Descriptor instead.
func (Principal_Type) EnumDescriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{0, 0}
}

type Principal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Principal_Type `protobuf:"varint,1,opt,name=type,proto3,enum=notebook.v2.Principal_Type" json:"type,omitempty"`
	Id   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Principal) Reset() {
	*x = Principal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notebook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Principal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Principal) ProtoMessage() {}

func (x *Principal) ProtoReflect() protoreflect.Message {
	mi := &file_notebook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Principal.ProtoReflect.Descriptor instead.
func (*Principal) Descriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{0}
}

func (x *Principal) GetType() Principal_Type {
	if x != nil {
		return x.Type
	}
	return Principal_USER
}

func (x *Principal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEntryRequest) Reset() {
	*x = GetEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notebook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryRequest) ProtoMessage() {}

func (x *GetEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notebook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryRequest.ProtoReflect.Descriptor instead.
func (*GetEntryRequest) Descriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{1}
}

func (x *GetEntryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only.
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// Output only.
	CreatorId string `protobuf:"bytes,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	// Output only.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Output only.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notebook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_notebook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{2}
}

func (x *Entry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entry) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Entry) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Entry) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Entry) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *Entry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *CreateEntryRequest) Reset() {
	*x = CreateEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notebook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntryRequest) ProtoMessage() {}

func (x *CreateEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notebook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntryRequest.ProtoReflect.Descriptor instead.
func (*CreateEntryRequest) Descriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEntryRequest) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type UpdateEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Entry *Entry `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *UpdateEntryRequest) Reset() {
	*x = UpdateEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notebook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntryRequest) ProtoMessage() {}

func (x *UpdateEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notebook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntryRequest.ProtoReflect.Descriptor instead.
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEntryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEntryRequest) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type ListEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListEntryRequest) Reset() {
	*x = ListEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notebook_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntryRequest) ProtoMessage() {}

func (x *ListEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notebook_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntryRequest.ProtoReflect.Descriptor instead.
func (*ListEntryRequest) Descriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{5}
}

func (x *ListEntryRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListEntryRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries       []*Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalSize     int32    `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	HasNextPage   bool     `protobuf:"varint,4,opt,name=has_next_page,json=hasNextPage,proto3" json:"has_next_page,omitempty"`
}

func (x *ListEntriesResponse) Reset() {
	*x = ListEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notebook_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntriesResponse) ProtoMessage() {}

func (x *ListEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notebook_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntriesResponse.ProtoReflect.Descriptor instead.
func (*ListEntriesResponse) Descriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{6}
}

func (x *ListEntriesResponse) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *ListEntriesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListEntriesResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *ListEntriesResponse) GetHasNextPage() bool {
	if x != nil {
		return x.HasNextPage
	}
	return false
}

type DeleteEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEntryRequest) Reset() {
	*x = DeleteEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notebook_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntryRequest) ProtoMessage() {}

func (x *DeleteEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notebook_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntryRequest.ProtoReflect.Descriptor instead.
func (*DeleteEntryRequest) Descriptor() ([]byte, []int) {
	return file_notebook_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEntryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_notebook_proto protoreflect.FileDescriptor

var file_notebook_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x09, 0x50, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc0, 0x01, 0x0a,
	0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x3e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x4e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x4e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xae, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc3, 0x04, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x12, 0x67, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x1c, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x0e, 0x62, 0x0c,
	0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x12, 0x6f, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0b, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x92, 0x41, 0x0e, 0x62,
	0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x12, 0x74, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x32, 0x10, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x00, 0x12, 0x74, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a,
	0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x12, 0x71, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x0e, 0x62, 0x0c,
	0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x42, 0x3c, 0x5a, 0x0b,
	0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x32, 0x92, 0x41, 0x2c, 0x22, 0x09,
	0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5a, 0x1f, 0x0a, 0x1d, 0x0a, 0x06, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_notebook_proto_rawDescOnce sync.Once
	file_notebook_proto_rawDescData = file_notebook_proto_rawDesc
)

func file_notebook_proto_rawDescGZIP() []byte {
	file_notebook_proto_rawDescOnce.Do(func() {
		file_notebook_proto_rawDescData = protoimpl.X.CompressGZIP(file_notebook_proto_rawDescData)
	})
	return file_notebook_proto_rawDescData
}

var file_notebook_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notebook_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_notebook_proto_goTypes = []interface{}{
	(Principal_Type)(0),         // 0: notebook.v2.Principal.Type
	(*Principal)(nil),           // 1: notebook.v2.Principal
	(*GetEntryRequest)(nil),     // 2: notebook.v2.GetEntryRequest
	(*Entry)(nil),               // 3: notebook.v2.Entry
	(*CreateEntryRequest)(nil),  // 4: notebook.v2.CreateEntryRequest
	(*UpdateEntryRequest)(nil),  // 5: notebook.v2.UpdateEntryRequest
	(*ListEntryRequest)(nil),    // 6: notebook.v2.ListEntryRequest
	(*ListEntriesResponse)(nil), // 7: notebook.v2.ListEntriesResponse
	(*DeleteEntryRequest)(nil),  // 8: notebook.v2.DeleteEntryRequest
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_notebook_proto_depIdxs = []int32{
	0,  // 0: notebook.v2.Principal.type:type_name -> notebook.v2.Principal.Type
	9,  // 1: notebook.v2.Entry.created_at:type_name -> google.protobuf.Timestamp
	9,  // 2: notebook.v2.Entry.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 3: notebook.v2.CreateEntryRequest.entry:type_name -> notebook.v2.Entry
	3,  // 4: notebook.v2.UpdateEntryRequest.entry:type_name -> notebook.v2.Entry
	3,  // 5: notebook.v2.ListEntriesResponse.entries:type_name -> notebook.v2.Entry
	2,  // 6: notebook.v2.Notebook.GetEntry:input_type -> notebook.v2.GetEntryRequest
	4,  // 7: notebook.v2.Notebook.CreateEntry:input_type -> notebook.v2.CreateEntryRequest
	5,  // 8: notebook.v2.Notebook.UpdateEntry:input_type -> notebook.v2.UpdateEntryRequest
	6,  // 9: notebook.v2.Notebook.ListEntries:input_type -> notebook.v2.ListEntryRequest
	8,  // 10: notebook.v2.Notebook.DeleteEntry:input_type -> notebook.v2.DeleteEntryRequest
	3,  // 11: notebook.v2.Notebook.GetEntry:output_type -> notebook.v2.Entry
	3,  // 12: notebook.v2.Notebook.CreateEntry:output_type -> notebook.v2.Entry
	3,  // 13: notebook.v2.Notebook.UpdateEntry:output_type -> notebook.v2.Entry
	7,  // 14: notebook.v2.Notebook.ListEntries:output_type -> notebook.v2.ListEntriesResponse
	10, // 15: notebook.v2.Notebook.DeleteEntry:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_notebook_proto_init() }
func file_notebook_proto_init() {
	if File_notebook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notebook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Principal); i {
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
		file_notebook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntryRequest); i {
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
		file_notebook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_notebook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntryRequest); i {
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
		file_notebook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntryRequest); i {
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
		file_notebook_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntryRequest); i {
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
		file_notebook_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntriesResponse); i {
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
		file_notebook_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntryRequest); i {
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
			RawDescriptor: file_notebook_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notebook_proto_goTypes,
		DependencyIndexes: file_notebook_proto_depIdxs,
		EnumInfos:         file_notebook_proto_enumTypes,
		MessageInfos:      file_notebook_proto_msgTypes,
	}.Build()
	File_notebook_proto = out.File
	file_notebook_proto_rawDesc = nil
	file_notebook_proto_goTypes = nil
	file_notebook_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotebookClient is the client API for Notebook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotebookClient interface {
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*Entry, error)
	ListEntries(ctx context.Context, in *ListEntryRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type notebookClient struct {
	cc grpc.ClientConnInterface
}

func NewNotebookClient(cc grpc.ClientConnInterface) NotebookClient {
	return &notebookClient{cc}
}

func (c *notebookClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/notebook.v2.Notebook/GetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/notebook.v2.Notebook/CreateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/notebook.v2.Notebook/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookClient) ListEntries(ctx context.Context, in *ListEntryRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error) {
	out := new(ListEntriesResponse)
	err := c.cc.Invoke(ctx, "/notebook.v2.Notebook/ListEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/notebook.v2.Notebook/DeleteEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotebookServer is the server API for Notebook service.
type NotebookServer interface {
	GetEntry(context.Context, *GetEntryRequest) (*Entry, error)
	CreateEntry(context.Context, *CreateEntryRequest) (*Entry, error)
	UpdateEntry(context.Context, *UpdateEntryRequest) (*Entry, error)
	ListEntries(context.Context, *ListEntryRequest) (*ListEntriesResponse, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*empty.Empty, error)
}

// UnimplementedNotebookServer can be embedded to have forward compatible implementations.
type UnimplementedNotebookServer struct {
}

func (*UnimplementedNotebookServer) GetEntry(context.Context, *GetEntryRequest) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (*UnimplementedNotebookServer) CreateEntry(context.Context, *CreateEntryRequest) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntry not implemented")
}
func (*UnimplementedNotebookServer) UpdateEntry(context.Context, *UpdateEntryRequest) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntry not implemented")
}
func (*UnimplementedNotebookServer) ListEntries(context.Context, *ListEntryRequest) (*ListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntries not implemented")
}
func (*UnimplementedNotebookServer) DeleteEntry(context.Context, *DeleteEntryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntry not implemented")
}

func RegisterNotebookServer(s *grpc.Server, srv NotebookServer) {
	s.RegisterService(&_Notebook_serviceDesc, srv)
}

func _Notebook_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notebook.v2.Notebook/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notebook_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notebook.v2.Notebook/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notebook_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notebook.v2.Notebook/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notebook_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notebook.v2.Notebook/ListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServer).ListEntries(ctx, req.(*ListEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notebook_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notebook.v2.Notebook/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notebook_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notebook.v2.Notebook",
	HandlerType: (*NotebookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _Notebook_GetEntry_Handler,
		},
		{
			MethodName: "CreateEntry",
			Handler:    _Notebook_CreateEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _Notebook_UpdateEntry_Handler,
		},
		{
			MethodName: "ListEntries",
			Handler:    _Notebook_ListEntries_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _Notebook_DeleteEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notebook.proto",
}
