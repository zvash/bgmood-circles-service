// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_mood_and_reactions.proto

package circlespb

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

type CreateMoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CircleId      string  `protobuf:"bytes,1,opt,name=circle_id,json=circleId,proto3" json:"circle_id,omitempty"`
	Image         string  `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	SetBackground bool    `protobuf:"varint,3,opt,name=set_background,json=setBackground,proto3" json:"set_background,omitempty"`
	Description   *string `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *CreateMoodRequest) Reset() {
	*x = CreateMoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_mood_and_reactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMoodRequest) ProtoMessage() {}

func (x *CreateMoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_mood_and_reactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMoodRequest.ProtoReflect.Descriptor instead.
func (*CreateMoodRequest) Descriptor() ([]byte, []int) {
	return file_rpc_mood_and_reactions_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMoodRequest) GetCircleId() string {
	if x != nil {
		return x.CircleId
	}
	return ""
}

func (x *CreateMoodRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateMoodRequest) GetSetBackground() bool {
	if x != nil {
		return x.SetBackground
	}
	return false
}

func (x *CreateMoodRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type ReactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoodId     string `protobuf:"bytes,1,opt,name=mood_id,json=moodId,proto3" json:"mood_id,omitempty"`
	ReactionId string `protobuf:"bytes,2,opt,name=reaction_id,json=reactionId,proto3" json:"reaction_id,omitempty"`
}

func (x *ReactRequest) Reset() {
	*x = ReactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_mood_and_reactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReactRequest) ProtoMessage() {}

func (x *ReactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_mood_and_reactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReactRequest.ProtoReflect.Descriptor instead.
func (*ReactRequest) Descriptor() ([]byte, []int) {
	return file_rpc_mood_and_reactions_proto_rawDescGZIP(), []int{1}
}

func (x *ReactRequest) GetMoodId() string {
	if x != nil {
		return x.MoodId
	}
	return ""
}

func (x *ReactRequest) GetReactionId() string {
	if x != nil {
		return x.ReactionId
	}
	return ""
}

type RemoveReactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoodId string `protobuf:"bytes,1,opt,name=mood_id,json=moodId,proto3" json:"mood_id,omitempty"`
}

func (x *RemoveReactionRequest) Reset() {
	*x = RemoveReactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_mood_and_reactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveReactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveReactionRequest) ProtoMessage() {}

func (x *RemoveReactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_mood_and_reactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveReactionRequest.ProtoReflect.Descriptor instead.
func (*RemoveReactionRequest) Descriptor() ([]byte, []int) {
	return file_rpc_mood_and_reactions_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveReactionRequest) GetMoodId() string {
	if x != nil {
		return x.MoodId
	}
	return ""
}

type MoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mood *Mood `protobuf:"bytes,1,opt,name=mood,proto3" json:"mood,omitempty"`
}

func (x *MoodResponse) Reset() {
	*x = MoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_mood_and_reactions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoodResponse) ProtoMessage() {}

func (x *MoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_mood_and_reactions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoodResponse.ProtoReflect.Descriptor instead.
func (*MoodResponse) Descriptor() ([]byte, []int) {
	return file_rpc_mood_and_reactions_proto_rawDescGZIP(), []int{3}
}

func (x *MoodResponse) GetMood() *Mood {
	if x != nil {
		return x.Mood
	}
	return nil
}

var File_rpc_mood_and_reactions_proto protoreflect.FileDescriptor

var file_rpc_mood_and_reactions_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x6f, 0x64, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x72,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x73, 0x70, 0x62, 0x1a, 0x0a, 0x6d, 0x6f, 0x6f, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6d, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0c, 0x4d, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x6d, 0x6f, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x73,
	0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x04, 0x6d, 0x6f, 0x6f, 0x64, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x76, 0x61, 0x73,
	0x68, 0x2f, 0x62, 0x67, 0x6d, 0x6f, 0x6f, 0x64, 0x2d, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x73,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x73,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_mood_and_reactions_proto_rawDescOnce sync.Once
	file_rpc_mood_and_reactions_proto_rawDescData = file_rpc_mood_and_reactions_proto_rawDesc
)

func file_rpc_mood_and_reactions_proto_rawDescGZIP() []byte {
	file_rpc_mood_and_reactions_proto_rawDescOnce.Do(func() {
		file_rpc_mood_and_reactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_mood_and_reactions_proto_rawDescData)
	})
	return file_rpc_mood_and_reactions_proto_rawDescData
}

var file_rpc_mood_and_reactions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_mood_and_reactions_proto_goTypes = []interface{}{
	(*CreateMoodRequest)(nil),     // 0: circlespb.CreateMoodRequest
	(*ReactRequest)(nil),          // 1: circlespb.ReactRequest
	(*RemoveReactionRequest)(nil), // 2: circlespb.RemoveReactionRequest
	(*MoodResponse)(nil),          // 3: circlespb.MoodResponse
	(*Mood)(nil),                  // 4: circlespb.Mood
}
var file_rpc_mood_and_reactions_proto_depIdxs = []int32{
	4, // 0: circlespb.MoodResponse.mood:type_name -> circlespb.Mood
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_mood_and_reactions_proto_init() }
func file_rpc_mood_and_reactions_proto_init() {
	if File_rpc_mood_and_reactions_proto != nil {
		return
	}
	file_mood_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_mood_and_reactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMoodRequest); i {
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
		file_rpc_mood_and_reactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReactRequest); i {
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
		file_rpc_mood_and_reactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveReactionRequest); i {
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
		file_rpc_mood_and_reactions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoodResponse); i {
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
	file_rpc_mood_and_reactions_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_mood_and_reactions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_mood_and_reactions_proto_goTypes,
		DependencyIndexes: file_rpc_mood_and_reactions_proto_depIdxs,
		MessageInfos:      file_rpc_mood_and_reactions_proto_msgTypes,
	}.Build()
	File_rpc_mood_and_reactions_proto = out.File
	file_rpc_mood_and_reactions_proto_rawDesc = nil
	file_rpc_mood_and_reactions_proto_goTypes = nil
	file_rpc_mood_and_reactions_proto_depIdxs = nil
}
