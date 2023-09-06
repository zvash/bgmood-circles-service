// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: mood.proto

package cpb

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

type Mood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PosterId        string                 `protobuf:"bytes,2,opt,name=poster_id,json=posterId,proto3" json:"poster_id,omitempty"`
	CircleId        string                 `protobuf:"bytes,3,opt,name=circle_id,json=circleId,proto3" json:"circle_id,omitempty"`
	Image           string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	SetBackground   bool                   `protobuf:"varint,5,opt,name=set_background,json=setBackground,proto3" json:"set_background,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description     *string               `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
	UserReaction    *Reaction             `protobuf:"bytes,8,opt,name=user_reaction,json=userReaction,proto3,oneof" json:"user_reaction,omitempty"`
	OthersReactions []*ReactionAggregated `protobuf:"bytes,9,rep,name=others_reactions,json=othersReactions,proto3" json:"others_reactions,omitempty"`
}

func (x *Mood) Reset() {
	*x = Mood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mood_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mood) ProtoMessage() {}

func (x *Mood) ProtoReflect() protoreflect.Message {
	mi := &file_mood_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mood.ProtoReflect.Descriptor instead.
func (*Mood) Descriptor() ([]byte, []int) {
	return file_mood_proto_rawDescGZIP(), []int{0}
}

func (x *Mood) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Mood) GetPosterId() string {
	if x != nil {
		return x.PosterId
	}
	return ""
}

func (x *Mood) GetCircleId() string {
	if x != nil {
		return x.CircleId
	}
	return ""
}

func (x *Mood) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Mood) GetSetBackground() bool {
	if x != nil {
		return x.SetBackground
	}
	return false
}

func (x *Mood) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Mood) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Mood) GetUserReaction() *Reaction {
	if x != nil {
		return x.UserReaction
	}
	return nil
}

func (x *Mood) GetOthersReactions() []*ReactionAggregated {
	if x != nil {
		return x.OthersReactions
	}
	return nil
}

var File_mood_proto protoreflect.FileDescriptor

var file_mood_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x70,
	0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x04, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x72, 0x63,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x72,
	0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x65, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a,
	0x10, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x0f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x7a, 0x76, 0x61, 0x73, 0x68, 0x2f, 0x62, 0x67, 0x6d, 0x6f, 0x6f, 0x64, 0x2d, 0x63,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mood_proto_rawDescOnce sync.Once
	file_mood_proto_rawDescData = file_mood_proto_rawDesc
)

func file_mood_proto_rawDescGZIP() []byte {
	file_mood_proto_rawDescOnce.Do(func() {
		file_mood_proto_rawDescData = protoimpl.X.CompressGZIP(file_mood_proto_rawDescData)
	})
	return file_mood_proto_rawDescData
}

var file_mood_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mood_proto_goTypes = []interface{}{
	(*Mood)(nil),                  // 0: cpb.Mood
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Reaction)(nil),              // 2: cpb.Reaction
	(*ReactionAggregated)(nil),    // 3: cpb.ReactionAggregated
}
var file_mood_proto_depIdxs = []int32{
	1, // 0: cpb.Mood.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: cpb.Mood.user_reaction:type_name -> cpb.Reaction
	3, // 2: cpb.Mood.others_reactions:type_name -> cpb.ReactionAggregated
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mood_proto_init() }
func file_mood_proto_init() {
	if File_mood_proto != nil {
		return
	}
	file_reaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mood_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mood); i {
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
	file_mood_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mood_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mood_proto_goTypes,
		DependencyIndexes: file_mood_proto_depIdxs,
		MessageInfos:      file_mood_proto_msgTypes,
	}.Build()
	File_mood_proto = out.File
	file_mood_proto_rawDesc = nil
	file_mood_proto_goTypes = nil
	file_mood_proto_depIdxs = nil
}
