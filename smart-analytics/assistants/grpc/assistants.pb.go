// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: proto/assistants/assistants.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Запрос, включающий AssistantID
type AssistantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssistantId string `protobuf:"bytes,1,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
}

func (x *AssistantRequest) Reset() {
	*x = AssistantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_assistants_assistants_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssistantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssistantRequest) ProtoMessage() {}

func (x *AssistantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_assistants_assistants_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssistantRequest.ProtoReflect.Descriptor instead.
func (*AssistantRequest) Descriptor() ([]byte, []int) {
	return file_proto_assistants_assistants_proto_rawDescGZIP(), []int{0}
}

func (x *AssistantRequest) GetAssistantId() string {
	if x != nil {
		return x.AssistantId
	}
	return ""
}

type AssistantFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssistantId string `protobuf:"bytes,1,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
	FileId      string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *AssistantFileRequest) Reset() {
	*x = AssistantFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_assistants_assistants_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssistantFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssistantFileRequest) ProtoMessage() {}

func (x *AssistantFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_assistants_assistants_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssistantFileRequest.ProtoReflect.Descriptor instead.
func (*AssistantFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_assistants_assistants_proto_rawDescGZIP(), []int{1}
}

func (x *AssistantFileRequest) GetAssistantId() string {
	if x != nil {
		return x.AssistantId
	}
	return ""
}

func (x *AssistantFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

// Ответ с AssistantFiles
type ListAssistantFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssistantFiles *AssistantFiles `protobuf:"bytes,1,opt,name=assistant_files,json=assistantFiles,proto3" json:"assistant_files,omitempty"`
}

func (x *ListAssistantFilesResponse) Reset() {
	*x = ListAssistantFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_assistants_assistants_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssistantFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssistantFilesResponse) ProtoMessage() {}

func (x *ListAssistantFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_assistants_assistants_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssistantFilesResponse.ProtoReflect.Descriptor instead.
func (*ListAssistantFilesResponse) Descriptor() ([]byte, []int) {
	return file_proto_assistants_assistants_proto_rawDescGZIP(), []int{2}
}

func (x *ListAssistantFilesResponse) GetAssistantFiles() *AssistantFiles {
	if x != nil {
		return x.AssistantFiles
	}
	return nil
}

// сообщение для данных о файле
type AssistantFileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object      string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	CreatedAt   int32  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AssistantId string `protobuf:"bytes,4,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
}

func (x *AssistantFileData) Reset() {
	*x = AssistantFileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_assistants_assistants_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssistantFileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssistantFileData) ProtoMessage() {}

func (x *AssistantFileData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_assistants_assistants_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssistantFileData.ProtoReflect.Descriptor instead.
func (*AssistantFileData) Descriptor() ([]byte, []int) {
	return file_proto_assistants_assistants_proto_rawDescGZIP(), []int{3}
}

func (x *AssistantFileData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssistantFileData) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *AssistantFileData) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AssistantFileData) GetAssistantId() string {
	if x != nil {
		return x.AssistantId
	}
	return ""
}

// сообщение для списка файлов ассистента
type AssistantFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object  string               `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Data    []*AssistantFileData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	FirstId string               `protobuf:"bytes,3,opt,name=first_id,json=firstId,proto3" json:"first_id,omitempty"`
	LastId  string               `protobuf:"bytes,4,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	HasMore bool                 `protobuf:"varint,5,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *AssistantFiles) Reset() {
	*x = AssistantFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_assistants_assistants_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssistantFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssistantFiles) ProtoMessage() {}

func (x *AssistantFiles) ProtoReflect() protoreflect.Message {
	mi := &file_proto_assistants_assistants_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssistantFiles.ProtoReflect.Descriptor instead.
func (*AssistantFiles) Descriptor() ([]byte, []int) {
	return file_proto_assistants_assistants_proto_rawDescGZIP(), []int{4}
}

func (x *AssistantFiles) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *AssistantFiles) GetData() []*AssistantFileData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AssistantFiles) GetFirstId() string {
	if x != nil {
		return x.FirstId
	}
	return ""
}

func (x *AssistantFiles) GetLastId() string {
	if x != nil {
		return x.LastId
	}
	return ""
}

func (x *AssistantFiles) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type ThreadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadId string `protobuf:"bytes,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
}

func (x *ThreadRequest) Reset() {
	*x = ThreadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_assistants_assistants_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadRequest) ProtoMessage() {}

func (x *ThreadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_assistants_assistants_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadRequest.ProtoReflect.Descriptor instead.
func (*ThreadRequest) Descriptor() ([]byte, []int) {
	return file_proto_assistants_assistants_proto_rawDescGZIP(), []int{5}
}

func (x *ThreadRequest) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

// ответ DeleteAssistantFile
type DeletedObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object  string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Deleted bool   `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeletedObject) Reset() {
	*x = DeletedObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_assistants_assistants_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletedObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletedObject) ProtoMessage() {}

func (x *DeletedObject) ProtoReflect() protoreflect.Message {
	mi := &file_proto_assistants_assistants_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletedObject.ProtoReflect.Descriptor instead.
func (*DeletedObject) Descriptor() ([]byte, []int) {
	return file_proto_assistants_assistants_proto_rawDescGZIP(), []int{6}
}

func (x *DeletedObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeletedObject) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *DeletedObject) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

// ответ AssistantObject
type AssistantObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object    string     `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	CreatedAt int32      `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Metadata  *anypb.Any `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *AssistantObject) Reset() {
	*x = AssistantObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_assistants_assistants_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssistantObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssistantObject) ProtoMessage() {}

func (x *AssistantObject) ProtoReflect() protoreflect.Message {
	mi := &file_proto_assistants_assistants_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssistantObject.ProtoReflect.Descriptor instead.
func (*AssistantObject) Descriptor() ([]byte, []int) {
	return file_proto_assistants_assistants_proto_rawDescGZIP(), []int{7}
}

func (x *AssistantObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssistantObject) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *AssistantObject) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AssistantObject) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_proto_assistants_assistants_proto protoreflect.FileDescriptor

var file_proto_assistants_assistants_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x52,
	0x0a, 0x14, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x22, 0x61, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x7d, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x72, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72,
	0x65, 0x22, 0x2c, 0x0a, 0x0d, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x22,
	0x51, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0f, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xea,
	0x03, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x56, 0x0a, 0x13, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x6f, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73,
	0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x43, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x19, 0x2e,
	0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x44, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x73, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x21, 0x5a, 0x1f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_assistants_assistants_proto_rawDescOnce sync.Once
	file_proto_assistants_assistants_proto_rawDescData = file_proto_assistants_assistants_proto_rawDesc
)

func file_proto_assistants_assistants_proto_rawDescGZIP() []byte {
	file_proto_assistants_assistants_proto_rawDescOnce.Do(func() {
		file_proto_assistants_assistants_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_assistants_assistants_proto_rawDescData)
	})
	return file_proto_assistants_assistants_proto_rawDescData
}

var file_proto_assistants_assistants_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_assistants_assistants_proto_goTypes = []interface{}{
	(*AssistantRequest)(nil),           // 0: assistants.AssistantRequest
	(*AssistantFileRequest)(nil),       // 1: assistants.AssistantFileRequest
	(*ListAssistantFilesResponse)(nil), // 2: assistants.ListAssistantFilesResponse
	(*AssistantFileData)(nil),          // 3: assistants.AssistantFileData
	(*AssistantFiles)(nil),             // 4: assistants.AssistantFiles
	(*ThreadRequest)(nil),              // 5: assistants.ThreadRequest
	(*DeletedObject)(nil),              // 6: assistants.DeletedObject
	(*AssistantObject)(nil),            // 7: assistants.AssistantObject
	(*anypb.Any)(nil),                  // 8: google.protobuf.Any
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_proto_assistants_assistants_proto_depIdxs = []int32{
	4, // 0: assistants.ListAssistantFilesResponse.assistant_files:type_name -> assistants.AssistantFiles
	3, // 1: assistants.AssistantFiles.data:type_name -> assistants.AssistantFileData
	8, // 2: assistants.AssistantObject.metadata:type_name -> google.protobuf.Any
	0, // 3: assistants.AssistantService.ListAssistantFiles:input_type -> assistants.AssistantRequest
	1, // 4: assistants.AssistantService.DeleteAssistantFile:input_type -> assistants.AssistantFileRequest
	1, // 5: assistants.AssistantService.LinkFileToAssistant:input_type -> assistants.AssistantFileRequest
	9, // 6: assistants.AssistantService.CreateThread:input_type -> google.protobuf.Empty
	5, // 7: assistants.AssistantService.GetThread:input_type -> assistants.ThreadRequest
	5, // 8: assistants.AssistantService.DeleteThread:input_type -> assistants.ThreadRequest
	2, // 9: assistants.AssistantService.ListAssistantFiles:output_type -> assistants.ListAssistantFilesResponse
	6, // 10: assistants.AssistantService.DeleteAssistantFile:output_type -> assistants.DeletedObject
	3, // 11: assistants.AssistantService.LinkFileToAssistant:output_type -> assistants.AssistantFileData
	7, // 12: assistants.AssistantService.CreateThread:output_type -> assistants.AssistantObject
	7, // 13: assistants.AssistantService.GetThread:output_type -> assistants.AssistantObject
	6, // 14: assistants.AssistantService.DeleteThread:output_type -> assistants.DeletedObject
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_assistants_assistants_proto_init() }
func file_proto_assistants_assistants_proto_init() {
	if File_proto_assistants_assistants_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_assistants_assistants_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssistantRequest); i {
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
		file_proto_assistants_assistants_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssistantFileRequest); i {
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
		file_proto_assistants_assistants_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssistantFilesResponse); i {
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
		file_proto_assistants_assistants_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssistantFileData); i {
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
		file_proto_assistants_assistants_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssistantFiles); i {
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
		file_proto_assistants_assistants_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadRequest); i {
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
		file_proto_assistants_assistants_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletedObject); i {
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
		file_proto_assistants_assistants_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssistantObject); i {
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
			RawDescriptor: file_proto_assistants_assistants_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_assistants_assistants_proto_goTypes,
		DependencyIndexes: file_proto_assistants_assistants_proto_depIdxs,
		MessageInfos:      file_proto_assistants_assistants_proto_msgTypes,
	}.Build()
	File_proto_assistants_assistants_proto = out.File
	file_proto_assistants_assistants_proto_rawDesc = nil
	file_proto_assistants_assistants_proto_goTypes = nil
	file_proto_assistants_assistants_proto_depIdxs = nil
}
