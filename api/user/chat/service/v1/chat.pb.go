// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: chat.proto

package v1

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

type GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`               // 当前用户id
	ToUserId   int64 `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`       // 对方用户id
	PreMsgTime int64 `protobuf:"varint,3,opt,name=pre_msg_time,json=preMsgTime,proto3" json:"pre_msg_time,omitempty"` // 上次最新消息的时间
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) Reset() {
	*x = GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) ProtoMessage() {}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest.ProtoReflect.Descriptor instead.
func (*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest) GetPreMsgTime() int64 {
	if x != nil {
		return x.PreMsgTime
	}
	return 0
}

type GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   *string    `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	MessageList []*Message `protobuf:"bytes,3,rep,name=message_list,json=messageList,proto3" json:"message_list,omitempty"` // 消息列表
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse) Reset() {
	*x = GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse) ProtoMessage() {}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse.ProtoReflect.Descriptor instead.
func (*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse) GetMessageList() []*Message {
	if x != nil {
		return x.MessageList
	}
	return nil
}

type DouyinChatActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 当前用户id
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`     // 对方用户id
	ActionType int32  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-发送消息
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`                          // 消息内容
}

func (x *DouyinChatActionRequest) Reset() {
	*x = DouyinChatActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinChatActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinChatActionRequest) ProtoMessage() {}

func (x *DouyinChatActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinChatActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinChatActionRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *DouyinChatActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DouyinChatActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *DouyinChatActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *DouyinChatActionRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DouyinChatActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

func (x *DouyinChatActionResponse) Reset() {
	*x = DouyinChatActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinChatActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinChatActionResponse) ProtoMessage() {}

func (x *DouyinChatActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinChatActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinChatActionResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *DouyinChatActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinChatActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type GetLatestMsgByMyUserIdAndHisUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`         // 当前用户id
	ToUserId int64 `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"` // 对方用户id
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdRequest) Reset() {
	*x = GetLatestMsgByMyUserIdAndHisUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestMsgByMyUserIdAndHisUserIdRequest) ProtoMessage() {}

func (x *GetLatestMsgByMyUserIdAndHisUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestMsgByMyUserIdAndHisUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetLatestMsgByMyUserIdAndHisUserIdRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

type GetLatestMsgByMyUserIdAndHisUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	Message    *Message `protobuf:"bytes,3,opt,name=message,proto3,oneof" json:"message,omitempty"`                      // 最新消息
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdResponse) Reset() {
	*x = GetLatestMsgByMyUserIdAndHisUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestMsgByMyUserIdAndHisUserIdResponse) ProtoMessage() {}

func (x *GetLatestMsgByMyUserIdAndHisUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestMsgByMyUserIdAndHisUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetLatestMsgByMyUserIdAndHisUserIdResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *GetLatestMsgByMyUserIdAndHisUserIdResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                        // 消息id
	ToUserId   int64   `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`          // 该消息接收者的id
	FromUserId int64   `protobuf:"varint,3,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`    // 该消息发送者的id
	Content    string  `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`                               // 消息内容
	CreateTime *string `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"` // 消息创建时间
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *Message) GetFromUserId() int64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetCreateTime() string {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x93, 0x01,
	0x0a, 0x38, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x48, 0x69, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x4d, 0x73, 0x67, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x4d, 0x73, 0x67, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x39, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x41, 0x6e, 0x64, 0x48, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x50,
	0x72, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x73, 0x67, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x43, 0x68, 0x61,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x6e, 0x0a, 0x18, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x22, 0x62, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67,
	0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x48, 0x69, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0xc5, 0x01, 0x0a, 0x2a, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x4d, 0x73, 0x67, 0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41,
	0x6e, 0x64, 0x48, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa9, 0x01, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x72,
	0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xd6, 0x03, 0x0a, 0x04, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x61, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0xca, 0x01, 0x0a, 0x31, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x41, 0x6e, 0x64, 0x48, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64,
	0x50, 0x72, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4d, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x48, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x41, 0x6e, 0x64, 0x48, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x50,
	0x72, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x9d, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4d,
	0x73, 0x67, 0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x48,
	0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x41, 0x6e, 0x64, 0x48, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x42, 0x79, 0x4d, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x41, 0x6e, 0x64,
	0x48, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1b, 0x5a, 0x19, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chat_proto_goTypes = []interface{}{
	(*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest)(nil),  // 0: chat.service.v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest
	(*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse)(nil), // 1: chat.service.v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse
	(*DouyinChatActionRequest)(nil),                                   // 2: chat.service.v1.DouyinChatActionRequest
	(*DouyinChatActionResponse)(nil),                                  // 3: chat.service.v1.DouyinChatActionResponse
	(*GetLatestMsgByMyUserIdAndHisUserIdRequest)(nil),                 // 4: chat.service.v1.GetLatestMsgByMyUserIdAndHisUserIdRequest
	(*GetLatestMsgByMyUserIdAndHisUserIdResponse)(nil),                // 5: chat.service.v1.GetLatestMsgByMyUserIdAndHisUserIdResponse
	(*Message)(nil), // 6: chat.service.v1.Message
}
var file_chat_proto_depIdxs = []int32{
	6, // 0: chat.service.v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse.message_list:type_name -> chat.service.v1.Message
	6, // 1: chat.service.v1.GetLatestMsgByMyUserIdAndHisUserIdResponse.message:type_name -> chat.service.v1.Message
	2, // 2: chat.service.v1.Chat.ChatAction:input_type -> chat.service.v1.DouyinChatActionRequest
	0, // 3: chat.service.v1.Chat.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime:input_type -> chat.service.v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest
	4, // 4: chat.service.v1.Chat.GetLatestMsgByMyUserIdAndHisUserId:input_type -> chat.service.v1.GetLatestMsgByMyUserIdAndHisUserIdRequest
	3, // 5: chat.service.v1.Chat.ChatAction:output_type -> chat.service.v1.DouyinChatActionResponse
	1, // 6: chat.service.v1.Chat.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTime:output_type -> chat.service.v1.GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse
	5, // 7: chat.service.v1.Chat.GetLatestMsgByMyUserIdAndHisUserId:output_type -> chat.service.v1.GetLatestMsgByMyUserIdAndHisUserIdResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeRequest); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageListByMyUserIdAndHisUserIdAndPreMsgTimeResponse); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinChatActionRequest); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinChatActionResponse); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestMsgByMyUserIdAndHisUserIdRequest); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestMsgByMyUserIdAndHisUserIdResponse); i {
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
		file_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_chat_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_chat_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_chat_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_chat_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
