// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: student.proto

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

type StudentPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StudentPrimaryKey) Reset() {
	*x = StudentPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentPrimaryKey) ProtoMessage() {}

func (x *StudentPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentPrimaryKey.ProtoReflect.Descriptor instead.
func (*StudentPrimaryKey) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{0}
}

func (x *StudentPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StudentGmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
}

func (x *StudentGmail) Reset() {
	*x = StudentGmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentGmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentGmail) ProtoMessage() {}

func (x *StudentGmail) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentGmail.ProtoReflect.Descriptor instead.
func (*StudentGmail) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{1}
}

func (x *StudentGmail) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

type STMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *STMessage) Reset() {
	*x = STMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *STMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*STMessage) ProtoMessage() {}

func (x *STMessage) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use STMessage.ProtoReflect.Descriptor instead.
func (*STMessage) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{2}
}

func (x *STMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateStudent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone    string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Branchid string `protobuf:"bytes,4,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	GroupId  string `protobuf:"bytes,6,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *CreateStudent) Reset() {
	*x = CreateStudent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudent) ProtoMessage() {}

func (x *CreateStudent) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudent.ProtoReflect.Descriptor instead.
func (*CreateStudent) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStudent) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreateStudent) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateStudent) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateStudent) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *CreateStudent) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateStudent) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname    string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone       string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Branchid    string `protobuf:"bytes,3,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PaidSum     int32  `protobuf:"varint,5,opt,name=paidSum,proto3" json:"paidSum,omitempty"`
	CourseCount int32  `protobuf:"varint,6,opt,name=courseCount,proto3" json:"courseCount,omitempty"`
	TotalSum    int32  `protobuf:"varint,7,opt,name=totalSum,proto3" json:"totalSum,omitempty"`
	Id          string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	GroupId     string `protobuf:"bytes,12,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{4}
}

func (x *Student) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *Student) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Student) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *Student) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Student) GetPaidSum() int32 {
	if x != nil {
		return x.PaidSum
	}
	return 0
}

func (x *Student) GetCourseCount() int32 {
	if x != nil {
		return x.CourseCount
	}
	return 0
}

func (x *Student) GetTotalSum() int32 {
	if x != nil {
		return x.TotalSum
	}
	return 0
}

func (x *Student) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Student) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Student) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Student) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Student) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type UpdateStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname    string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Phone       string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Branchid    string `protobuf:"bytes,3,opt,name=branchid,proto3" json:"branchid,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PaidSum     int32  `protobuf:"varint,5,opt,name=paidSum,proto3" json:"paidSum,omitempty"`
	CourseCount int32  `protobuf:"varint,6,opt,name=courseCount,proto3" json:"courseCount,omitempty"`
	TotalSum    int32  `protobuf:"varint,7,opt,name=totalSum,proto3" json:"totalSum,omitempty"`
	GroupId     string `protobuf:"bytes,8,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Id          string `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateStudentRequest) Reset() {
	*x = UpdateStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentRequest) ProtoMessage() {}

func (x *UpdateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentRequest.ProtoReflect.Descriptor instead.
func (*UpdateStudentRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateStudentRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UpdateStudentRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateStudentRequest) GetBranchid() string {
	if x != nil {
		return x.Branchid
	}
	return ""
}

func (x *UpdateStudentRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateStudentRequest) GetPaidSum() int32 {
	if x != nil {
		return x.PaidSum
	}
	return 0
}

func (x *UpdateStudentRequest) GetCourseCount() int32 {
	if x != nil {
		return x.CourseCount
	}
	return 0
}

func (x *UpdateStudentRequest) GetTotalSum() int32 {
	if x != nil {
		return x.TotalSum
	}
	return 0
}

func (x *UpdateStudentRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *UpdateStudentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListStudentRequest) Reset() {
	*x = GetListStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListStudentRequest) ProtoMessage() {}

func (x *GetListStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListStudentRequest.ProtoReflect.Descriptor instead.
func (*GetListStudentRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{6}
}

func (x *GetListStudentRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListStudentRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListStudentRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int64      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Student []*Student `protobuf:"bytes,2,rep,name=student,proto3" json:"student,omitempty"`
}

func (x *GetListStudentResponse) Reset() {
	*x = GetListStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListStudentResponse) ProtoMessage() {}

func (x *GetListStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListStudentResponse.ProtoReflect.Descriptor instead.
func (*GetListStudentResponse) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{7}
}

func (x *GetListStudentResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListStudentResponse) GetStudent() []*Student {
	if x != nil {
		return x.Student
	}
	return nil
}

var File_student_proto protoreflect.FileDescriptor

var file_student_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x25, 0x0a, 0x09,
	0x53, 0x54, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0xcc, 0x02, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x69, 0x64, 0x53, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x61, 0x69, 0x64, 0x53, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xfc,
	0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x69, 0x64, 0x53, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61,
	0x69, 0x64, 0x53, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x63, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x07,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x32, 0x83, 0x04, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a,
	0x23, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x23, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x54, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x54, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x47, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x47, 0x6d, 0x61, 0x69, 0x6c,
	0x1a, 0x23, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_proto_rawDescOnce sync.Once
	file_student_proto_rawDescData = file_student_proto_rawDesc
)

func file_student_proto_rawDescGZIP() []byte {
	file_student_proto_rawDescOnce.Do(func() {
		file_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_proto_rawDescData)
	})
	return file_student_proto_rawDescData
}

var file_student_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_student_proto_goTypes = []interface{}{
	(*StudentPrimaryKey)(nil),      // 0: schedule_service.StudentPrimaryKey
	(*StudentGmail)(nil),           // 1: schedule_service.StudentGmail
	(*STMessage)(nil),              // 2: schedule_service.STMessage
	(*CreateStudent)(nil),          // 3: schedule_service.CreateStudent
	(*Student)(nil),                // 4: schedule_service.Student
	(*UpdateStudentRequest)(nil),   // 5: schedule_service.UpdateStudentRequest
	(*GetListStudentRequest)(nil),  // 6: schedule_service.GetListStudentRequest
	(*GetListStudentResponse)(nil), // 7: schedule_service.GetListStudentResponse
}
var file_student_proto_depIdxs = []int32{
	4, // 0: schedule_service.GetListStudentResponse.student:type_name -> schedule_service.Student
	3, // 1: schedule_service.StudentService.Create:input_type -> schedule_service.CreateStudent
	0, // 2: schedule_service.StudentService.GetByID:input_type -> schedule_service.StudentPrimaryKey
	6, // 3: schedule_service.StudentService.GetList:input_type -> schedule_service.GetListStudentRequest
	5, // 4: schedule_service.StudentService.Update:input_type -> schedule_service.UpdateStudentRequest
	0, // 5: schedule_service.StudentService.Delete:input_type -> schedule_service.StudentPrimaryKey
	1, // 6: schedule_service.StudentService.GetByGmail:input_type -> schedule_service.StudentGmail
	0, // 7: schedule_service.StudentService.Create:output_type -> schedule_service.StudentPrimaryKey
	4, // 8: schedule_service.StudentService.GetByID:output_type -> schedule_service.Student
	7, // 9: schedule_service.StudentService.GetList:output_type -> schedule_service.GetListStudentResponse
	2, // 10: schedule_service.StudentService.Update:output_type -> schedule_service.STMessage
	2, // 11: schedule_service.StudentService.Delete:output_type -> schedule_service.STMessage
	0, // 12: schedule_service.StudentService.GetByGmail:output_type -> schedule_service.StudentPrimaryKey
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_student_proto_init() }
func file_student_proto_init() {
	if File_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentPrimaryKey); i {
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
		file_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentGmail); i {
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
		file_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*STMessage); i {
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
		file_student_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudent); i {
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
		file_student_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_student_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudentRequest); i {
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
		file_student_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListStudentRequest); i {
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
		file_student_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListStudentResponse); i {
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
			RawDescriptor: file_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_proto_goTypes,
		DependencyIndexes: file_student_proto_depIdxs,
		MessageInfos:      file_student_proto_msgTypes,
	}.Build()
	File_student_proto = out.File
	file_student_proto_rawDesc = nil
	file_student_proto_goTypes = nil
	file_student_proto_depIdxs = nil
}