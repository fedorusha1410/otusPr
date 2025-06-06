// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: task/task.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Note          string                 `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	Priority      string                 `protobuf:"bytes,4,opt,name=priority,proto3" json:"priority,omitempty"`
	AuthorId      int32                  `protobuf:"varint,5,opt,name=authorId,proto3" json:"authorId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_task_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CreateTaskRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *CreateTaskRequest) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Note          string                 `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	Priority      string                 `protobuf:"bytes,5,opt,name=priority,proto3" json:"priority,omitempty"`
	AuthorId      int32                  `protobuf:"varint,6,opt,name=authorId,proto3" json:"authorId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	mi := &file_task_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateTaskResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateTaskResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskResponse) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *CreateTaskResponse) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *CreateTaskResponse) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Note          string                 `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	Priority      string                 `protobuf:"bytes,5,opt,name=priority,proto3" json:"priority,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	mi := &file_task_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTaskRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *UpdateTaskRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_task_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_task_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{4}
}

func (x *GetTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Note          string                 `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	Priority      string                 `protobuf:"bytes,4,opt,name=priority,proto3" json:"priority,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	AuthorId      int32                  `protobuf:"varint,6,opt,name=authorId,proto3" json:"authorId,omitempty"`
	CreatedTime   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	mi := &file_task_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskResponse) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *TaskResponse) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *TaskResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TaskResponse) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *TaskResponse) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *TaskResponse) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

type TaskListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*TaskResponse        `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	mi := &file_task_task_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{6}
}

func (x *TaskListResponse) GetTasks() []*TaskResponse {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_task_task_proto protoreflect.FileDescriptor

const file_task_task_proto_rawDesc = "" +
	"\n" +
	"\x0ftask/task.proto\x12\x05proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x85\x01\n" +
	"\x11CreateTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x12\n" +
	"\x04note\x18\x03 \x01(\tR\x04note\x12\x1a\n" +
	"\bpriority\x18\x04 \x01(\tR\bpriority\x12\x1a\n" +
	"\bauthorId\x18\x05 \x01(\x05R\bauthorId\"\x9e\x01\n" +
	"\x12CreateTaskResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x12\n" +
	"\x04note\x18\x04 \x01(\tR\x04note\x12\x1a\n" +
	"\bpriority\x18\x05 \x01(\tR\bpriority\x12\x1a\n" +
	"\bauthorId\x18\x06 \x01(\x05R\bauthorId\"\x81\x01\n" +
	"\x11UpdateTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x12\n" +
	"\x04note\x18\x04 \x01(\tR\x04note\x12\x1a\n" +
	"\bpriority\x18\x05 \x01(\tR\bpriority\"#\n" +
	"\x11DeleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\" \n" +
	"\x0eGetTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"\x94\x02\n" +
	"\fTaskResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x12\n" +
	"\x04note\x18\x03 \x01(\tR\x04note\x12\x1a\n" +
	"\bpriority\x18\x04 \x01(\tR\bpriority\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\x12\x1a\n" +
	"\bauthorId\x18\x06 \x01(\x05R\bauthorId\x12<\n" +
	"\vcreatedTime\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\vcreatedTime\x12<\n" +
	"\vupdatedTime\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\vupdatedTime\"=\n" +
	"\x10TaskListResponse\x12)\n" +
	"\x05tasks\x18\x01 \x03(\v2\x13.proto.TaskResponseR\x05tasks2\xcb\x02\n" +
	"\vTaskService\x12A\n" +
	"\n" +
	"CreateTask\x12\x18.proto.CreateTaskRequest\x1a\x19.proto.CreateTaskResponse\x12>\n" +
	"\n" +
	"UpdateTask\x12\x18.proto.UpdateTaskRequest\x1a\x16.google.protobuf.Empty\x12>\n" +
	"\n" +
	"DeleteTask\x12\x18.proto.DeleteTaskRequest\x1a\x16.google.protobuf.Empty\x129\n" +
	"\vGetTaskById\x12\x15.proto.GetTaskRequest\x1a\x13.proto.TaskResponse\x12>\n" +
	"\vGetAllTasks\x12\x16.google.protobuf.Empty\x1a\x17.proto.TaskListResponseB\x1fZ\x1dtask-manager/internal/grpc/pbb\x06proto3"

var (
	file_task_task_proto_rawDescOnce sync.Once
	file_task_task_proto_rawDescData []byte
)

func file_task_task_proto_rawDescGZIP() []byte {
	file_task_task_proto_rawDescOnce.Do(func() {
		file_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_task_task_proto_rawDesc), len(file_task_task_proto_rawDesc)))
	})
	return file_task_task_proto_rawDescData
}

var file_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_task_task_proto_goTypes = []any{
	(*CreateTaskRequest)(nil),     // 0: proto.CreateTaskRequest
	(*CreateTaskResponse)(nil),    // 1: proto.CreateTaskResponse
	(*UpdateTaskRequest)(nil),     // 2: proto.UpdateTaskRequest
	(*DeleteTaskRequest)(nil),     // 3: proto.DeleteTaskRequest
	(*GetTaskRequest)(nil),        // 4: proto.GetTaskRequest
	(*TaskResponse)(nil),          // 5: proto.TaskResponse
	(*TaskListResponse)(nil),      // 6: proto.TaskListResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_task_task_proto_depIdxs = []int32{
	7, // 0: proto.TaskResponse.createdTime:type_name -> google.protobuf.Timestamp
	7, // 1: proto.TaskResponse.updatedTime:type_name -> google.protobuf.Timestamp
	5, // 2: proto.TaskListResponse.tasks:type_name -> proto.TaskResponse
	0, // 3: proto.TaskService.CreateTask:input_type -> proto.CreateTaskRequest
	2, // 4: proto.TaskService.UpdateTask:input_type -> proto.UpdateTaskRequest
	3, // 5: proto.TaskService.DeleteTask:input_type -> proto.DeleteTaskRequest
	4, // 6: proto.TaskService.GetTaskById:input_type -> proto.GetTaskRequest
	8, // 7: proto.TaskService.GetAllTasks:input_type -> google.protobuf.Empty
	1, // 8: proto.TaskService.CreateTask:output_type -> proto.CreateTaskResponse
	8, // 9: proto.TaskService.UpdateTask:output_type -> google.protobuf.Empty
	8, // 10: proto.TaskService.DeleteTask:output_type -> google.protobuf.Empty
	5, // 11: proto.TaskService.GetTaskById:output_type -> proto.TaskResponse
	6, // 12: proto.TaskService.GetAllTasks:output_type -> proto.TaskListResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_task_task_proto_init() }
func file_task_task_proto_init() {
	if File_task_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_task_task_proto_rawDesc), len(file_task_task_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_task_proto_goTypes,
		DependencyIndexes: file_task_task_proto_depIdxs,
		MessageInfos:      file_task_task_proto_msgTypes,
	}.Build()
	File_task_task_proto = out.File
	file_task_task_proto_goTypes = nil
	file_task_task_proto_depIdxs = nil
}
