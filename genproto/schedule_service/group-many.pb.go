// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: group-many.proto

package schedule_service

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_many_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_group_many_proto_msgTypes[0]
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
	return file_group_many_proto_rawDescGZIP(), []int{0}
}

type GMMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GMMessage) Reset() {
	*x = GMMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_many_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GMMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GMMessage) ProtoMessage() {}

func (x *GMMessage) ProtoReflect() protoreflect.Message {
	mi := &file_group_many_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GMMessage.ProtoReflect.Descriptor instead.
func (*GMMessage) Descriptor() ([]byte, []int) {
	return file_group_many_proto_rawDescGZIP(), []int{1}
}

func (x *GMMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateGroupMany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ScheduleId string `protobuf:"bytes,2,opt,name=scheduleId,proto3" json:"scheduleId,omitempty"`
	JournalId  string `protobuf:"bytes,3,opt,name=journalId,proto3" json:"journalId,omitempty"`
}

func (x *CreateGroupMany) Reset() {
	*x = CreateGroupMany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_many_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupMany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupMany) ProtoMessage() {}

func (x *CreateGroupMany) ProtoReflect() protoreflect.Message {
	mi := &file_group_many_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupMany.ProtoReflect.Descriptor instead.
func (*CreateGroupMany) Descriptor() ([]byte, []int) {
	return file_group_many_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGroupMany) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CreateGroupMany) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *CreateGroupMany) GetJournalId() string {
	if x != nil {
		return x.JournalId
	}
	return ""
}

type GroupMany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ScheduleId string `protobuf:"bytes,2,opt,name=scheduleId,proto3" json:"scheduleId,omitempty"`
	JournalId  string `protobuf:"bytes,3,opt,name=journalId,proto3" json:"journalId,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *GroupMany) Reset() {
	*x = GroupMany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_many_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMany) ProtoMessage() {}

func (x *GroupMany) ProtoReflect() protoreflect.Message {
	mi := &file_group_many_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMany.ProtoReflect.Descriptor instead.
func (*GroupMany) Descriptor() ([]byte, []int) {
	return file_group_many_proto_rawDescGZIP(), []int{3}
}

func (x *GroupMany) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupMany) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *GroupMany) GetJournalId() string {
	if x != nil {
		return x.JournalId
	}
	return ""
}

func (x *GroupMany) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GroupMany) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GroupMany) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type UpdateGroupManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ScheduleId string `protobuf:"bytes,2,opt,name=scheduleId,proto3" json:"scheduleId,omitempty"`
	JournalId  string `protobuf:"bytes,3,opt,name=journalId,proto3" json:"journalId,omitempty"`
}

func (x *UpdateGroupManyRequest) Reset() {
	*x = UpdateGroupManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_many_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupManyRequest) ProtoMessage() {}

func (x *UpdateGroupManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_many_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupManyRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupManyRequest) Descriptor() ([]byte, []int) {
	return file_group_many_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGroupManyRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *UpdateGroupManyRequest) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *UpdateGroupManyRequest) GetJournalId() string {
	if x != nil {
		return x.JournalId
	}
	return ""
}

type SchMonth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID            string `protobuf:"bytes,1,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	GroupType          string `protobuf:"bytes,2,opt,name=GroupType,proto3" json:"GroupType,omitempty"`
	StartTime          string `protobuf:"bytes,3,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime            string `protobuf:"bytes,4,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	BranchName         string `protobuf:"bytes,5,opt,name=BranchName,proto3" json:"BranchName,omitempty"`
	TeacherName        string `protobuf:"bytes,6,opt,name=TeacherName,proto3" json:"TeacherName,omitempty"`
	SupportTeacherName string `protobuf:"bytes,7,opt,name=SupportTeacherName,proto3" json:"SupportTeacherName,omitempty"`
	StudentCount       int32  `protobuf:"varint,8,opt,name=StudentCount,proto3" json:"StudentCount,omitempty"`
}

func (x *SchMonth) Reset() {
	*x = SchMonth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_many_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchMonth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchMonth) ProtoMessage() {}

func (x *SchMonth) ProtoReflect() protoreflect.Message {
	mi := &file_group_many_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchMonth.ProtoReflect.Descriptor instead.
func (*SchMonth) Descriptor() ([]byte, []int) {
	return file_group_many_proto_rawDescGZIP(), []int{5}
}

func (x *SchMonth) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *SchMonth) GetGroupType() string {
	if x != nil {
		return x.GroupType
	}
	return ""
}

func (x *SchMonth) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *SchMonth) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *SchMonth) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *SchMonth) GetTeacherName() string {
	if x != nil {
		return x.TeacherName
	}
	return ""
}

func (x *SchMonth) GetSupportTeacherName() string {
	if x != nil {
		return x.SupportTeacherName
	}
	return ""
}

func (x *SchMonth) GetStudentCount() int32 {
	if x != nil {
		return x.StudentCount
	}
	return 0
}

type ScheduleMonth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int32       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	SMonth []*SchMonth `protobuf:"bytes,2,rep,name=sMonth,proto3" json:"sMonth,omitempty"`
}

func (x *ScheduleMonth) Reset() {
	*x = ScheduleMonth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_many_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleMonth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleMonth) ProtoMessage() {}

func (x *ScheduleMonth) ProtoReflect() protoreflect.Message {
	mi := &file_group_many_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleMonth.ProtoReflect.Descriptor instead.
func (*ScheduleMonth) Descriptor() ([]byte, []int) {
	return file_group_many_proto_rawDescGZIP(), []int{6}
}

func (x *ScheduleMonth) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ScheduleMonth) GetSMonth() []*SchMonth {
	if x != nil {
		return x.SMonth
	}
	return nil
}

var File_group_many_proto protoreflect.FileDescriptor

var file_group_many_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x6d, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x25, 0x0a,
	0x09, 0x47, 0x4d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22,
	0xc0, 0x01, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x70, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x22, 0x90, 0x02, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x4d, 0x6f, 0x6e, 0x74,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32,
	0x0a, 0x06, 0x73, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x63, 0x68, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x06, 0x73, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x32, 0xa7, 0x01, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x61, 0x6e, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x61, 0x6e, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x4d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4d,
	0x12, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_group_many_proto_rawDescOnce sync.Once
	file_group_many_proto_rawDescData = file_group_many_proto_rawDesc
)

func file_group_many_proto_rawDescGZIP() []byte {
	file_group_many_proto_rawDescOnce.Do(func() {
		file_group_many_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_many_proto_rawDescData)
	})
	return file_group_many_proto_rawDescData
}

var file_group_many_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_group_many_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: schedule_service.Empty
	(*GMMessage)(nil),              // 1: schedule_service.GMMessage
	(*CreateGroupMany)(nil),        // 2: schedule_service.CreateGroupMany
	(*GroupMany)(nil),              // 3: schedule_service.GroupMany
	(*UpdateGroupManyRequest)(nil), // 4: schedule_service.UpdateGroupManyRequest
	(*SchMonth)(nil),               // 5: schedule_service.SchMonth
	(*ScheduleMonth)(nil),          // 6: schedule_service.ScheduleMonth
}
var file_group_many_proto_depIdxs = []int32{
	5, // 0: schedule_service.ScheduleMonth.sMonth:type_name -> schedule_service.SchMonth
	2, // 1: schedule_service.GroupManyService.Create:input_type -> schedule_service.CreateGroupMany
	0, // 2: schedule_service.GroupManyService.ScheduleM:input_type -> schedule_service.Empty
	1, // 3: schedule_service.GroupManyService.Create:output_type -> schedule_service.GMMessage
	6, // 4: schedule_service.GroupManyService.ScheduleM:output_type -> schedule_service.ScheduleMonth
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_group_many_proto_init() }
func file_group_many_proto_init() {
	if File_group_many_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group_many_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_group_many_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GMMessage); i {
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
		file_group_many_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupMany); i {
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
		file_group_many_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMany); i {
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
		file_group_many_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupManyRequest); i {
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
		file_group_many_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchMonth); i {
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
		file_group_many_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleMonth); i {
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
			RawDescriptor: file_group_many_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_many_proto_goTypes,
		DependencyIndexes: file_group_many_proto_depIdxs,
		MessageInfos:      file_group_many_proto_msgTypes,
	}.Build()
	File_group_many_proto = out.File
	file_group_many_proto_rawDesc = nil
	file_group_many_proto_goTypes = nil
	file_group_many_proto_depIdxs = nil
}