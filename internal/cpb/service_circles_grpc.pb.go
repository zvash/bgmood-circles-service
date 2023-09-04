// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_circles.proto

package cpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Circles_CreateCircle_FullMethodName               = "/cpb.Circles/CreateCircle"
	Circles_EditCircle_FullMethodName                 = "/cpb.Circles/EditCircle"
	Circles_DestroyCircle_FullMethodName              = "/cpb.Circles/DestroyCircle"
	Circles_InviteUserToCircle_FullMethodName         = "/cpb.Circles/InviteUserToCircle"
	Circles_RollbackInviteUserToCircle_FullMethodName = "/cpb.Circles/RollbackInviteUserToCircle"
	Circles_KickFromCircle_FullMethodName             = "/cpb.Circles/KickFromCircle"
	Circles_SendJoinRequest_FullMethodName            = "/cpb.Circles/SendJoinRequest"
	Circles_AcceptJoinRequest_FullMethodName          = "/cpb.Circles/AcceptJoinRequest"
	Circles_RemoveJoinRequest_FullMethodName          = "/cpb.Circles/RemoveJoinRequest"
	Circles_LeaveCircle_FullMethodName                = "/cpb.Circles/LeaveCircle"
	Circles_GetJoinedCircles_FullMethodName           = "/cpb.Circles/GetJoinedCircles"
	Circles_GetRequestedCircles_FullMethodName        = "/cpb.Circles/GetRequestedCircles"
	Circles_ViewCircle_FullMethodName                 = "/cpb.Circles/ViewCircle"
	Circles_ExploreCircles_FullMethodName             = "/cpb.Circles/ExploreCircles"
	Circles_PromoteToPoster_FullMethodName            = "/cpb.Circles/PromoteToPoster"
	Circles_PromoteToAdmin_FullMethodName             = "/cpb.Circles/PromoteToAdmin"
	Circles_DemoteToViewer_FullMethodName             = "/cpb.Circles/DemoteToViewer"
	Circles_GetCircleWPChangeAccess_FullMethodName    = "/cpb.Circles/GetCircleWPChangeAccess"
	Circles_SetCircleWPChangeAccess_FullMethodName    = "/cpb.Circles/SetCircleWPChangeAccess"
	Circles_CreateMood_FullMethodName                 = "/cpb.Circles/CreateMood"
	Circles_React_FullMethodName                      = "/cpb.Circles/React"
	Circles_RemoveReaction_FullMethodName             = "/cpb.Circles/RemoveReaction"
	Circles_GetAvailableReactions_FullMethodName      = "/cpb.Circles/GetAvailableReactions"
)

// CirclesClient is the client API for Circles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CirclesClient interface {
	CreateCircle(ctx context.Context, in *CreateCircleRequest, opts ...grpc.CallOption) (*CreateCircleResponse, error)
	EditCircle(ctx context.Context, in *EditCircleRequest, opts ...grpc.CallOption) (*EditCircleResponse, error)
	DestroyCircle(ctx context.Context, in *DestroyCircleRequest, opts ...grpc.CallOption) (*DestroyCircleResponse, error)
	InviteUserToCircle(ctx context.Context, in *InviteUserToCircleRequest, opts ...grpc.CallOption) (*InviteUserToCircleResponse, error)
	RollbackInviteUserToCircle(ctx context.Context, in *RollbackInviteUserToCircleRequest, opts ...grpc.CallOption) (*RollbackInviteUserToCircleResponse, error)
	KickFromCircle(ctx context.Context, in *KickFromCircleRequest, opts ...grpc.CallOption) (*KickFromCircleResponse, error)
	SendJoinRequest(ctx context.Context, in *SendJoinRequestRequest, opts ...grpc.CallOption) (*SendJoinRequestResponse, error)
	AcceptJoinRequest(ctx context.Context, in *AcceptJoinRequestRequest, opts ...grpc.CallOption) (*AcceptJoinRequestResponse, error)
	RemoveJoinRequest(ctx context.Context, in *RemoveJoinRequestRequest, opts ...grpc.CallOption) (*RemoveJoinRequestResponse, error)
	LeaveCircle(ctx context.Context, in *LeaveCircleRequest, opts ...grpc.CallOption) (*LeaveCircleResponse, error)
	GetJoinedCircles(ctx context.Context, in *GetJoinedCirclesRequest, opts ...grpc.CallOption) (*GetJoinedCirclesResponse, error)
	GetRequestedCircles(ctx context.Context, in *GetRequestedCirclesRequest, opts ...grpc.CallOption) (*GetRequestedCirclesResponse, error)
	ViewCircle(ctx context.Context, in *ViewCircleRequest, opts ...grpc.CallOption) (*ViewCircleResponse, error)
	ExploreCircles(ctx context.Context, in *ExploreCirclesRequest, opts ...grpc.CallOption) (*ExploreCirclesResponse, error)
	PromoteToPoster(ctx context.Context, in *PromoteToPosterRequest, opts ...grpc.CallOption) (*AccessModificationResponse, error)
	PromoteToAdmin(ctx context.Context, in *PromoteToAdminRequest, opts ...grpc.CallOption) (*AccessModificationResponse, error)
	DemoteToViewer(ctx context.Context, in *DemoteToViewerRequest, opts ...grpc.CallOption) (*AccessModificationResponse, error)
	GetCircleWPChangeAccess(ctx context.Context, in *GetCircleWPChangeAccessRequest, opts ...grpc.CallOption) (*CircleWPChangeAccessResponse, error)
	SetCircleWPChangeAccess(ctx context.Context, in *SetCircleWPChangeAccessRequest, opts ...grpc.CallOption) (*CircleWPChangeAccessResponse, error)
	CreateMood(ctx context.Context, in *CreateMoodRequest, opts ...grpc.CallOption) (*MoodResponse, error)
	React(ctx context.Context, in *ReactRequest, opts ...grpc.CallOption) (*MoodResponse, error)
	RemoveReaction(ctx context.Context, in *RemoveReactionRequest, opts ...grpc.CallOption) (*MoodResponse, error)
	GetAvailableReactions(ctx context.Context, in *GetAvailableReactionsRequest, opts ...grpc.CallOption) (*GetAvailableReactionsResponse, error)
}

type circlesClient struct {
	cc grpc.ClientConnInterface
}

func NewCirclesClient(cc grpc.ClientConnInterface) CirclesClient {
	return &circlesClient{cc}
}

func (c *circlesClient) CreateCircle(ctx context.Context, in *CreateCircleRequest, opts ...grpc.CallOption) (*CreateCircleResponse, error) {
	out := new(CreateCircleResponse)
	err := c.cc.Invoke(ctx, Circles_CreateCircle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) EditCircle(ctx context.Context, in *EditCircleRequest, opts ...grpc.CallOption) (*EditCircleResponse, error) {
	out := new(EditCircleResponse)
	err := c.cc.Invoke(ctx, Circles_EditCircle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) DestroyCircle(ctx context.Context, in *DestroyCircleRequest, opts ...grpc.CallOption) (*DestroyCircleResponse, error) {
	out := new(DestroyCircleResponse)
	err := c.cc.Invoke(ctx, Circles_DestroyCircle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) InviteUserToCircle(ctx context.Context, in *InviteUserToCircleRequest, opts ...grpc.CallOption) (*InviteUserToCircleResponse, error) {
	out := new(InviteUserToCircleResponse)
	err := c.cc.Invoke(ctx, Circles_InviteUserToCircle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) RollbackInviteUserToCircle(ctx context.Context, in *RollbackInviteUserToCircleRequest, opts ...grpc.CallOption) (*RollbackInviteUserToCircleResponse, error) {
	out := new(RollbackInviteUserToCircleResponse)
	err := c.cc.Invoke(ctx, Circles_RollbackInviteUserToCircle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) KickFromCircle(ctx context.Context, in *KickFromCircleRequest, opts ...grpc.CallOption) (*KickFromCircleResponse, error) {
	out := new(KickFromCircleResponse)
	err := c.cc.Invoke(ctx, Circles_KickFromCircle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) SendJoinRequest(ctx context.Context, in *SendJoinRequestRequest, opts ...grpc.CallOption) (*SendJoinRequestResponse, error) {
	out := new(SendJoinRequestResponse)
	err := c.cc.Invoke(ctx, Circles_SendJoinRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) AcceptJoinRequest(ctx context.Context, in *AcceptJoinRequestRequest, opts ...grpc.CallOption) (*AcceptJoinRequestResponse, error) {
	out := new(AcceptJoinRequestResponse)
	err := c.cc.Invoke(ctx, Circles_AcceptJoinRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) RemoveJoinRequest(ctx context.Context, in *RemoveJoinRequestRequest, opts ...grpc.CallOption) (*RemoveJoinRequestResponse, error) {
	out := new(RemoveJoinRequestResponse)
	err := c.cc.Invoke(ctx, Circles_RemoveJoinRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) LeaveCircle(ctx context.Context, in *LeaveCircleRequest, opts ...grpc.CallOption) (*LeaveCircleResponse, error) {
	out := new(LeaveCircleResponse)
	err := c.cc.Invoke(ctx, Circles_LeaveCircle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) GetJoinedCircles(ctx context.Context, in *GetJoinedCirclesRequest, opts ...grpc.CallOption) (*GetJoinedCirclesResponse, error) {
	out := new(GetJoinedCirclesResponse)
	err := c.cc.Invoke(ctx, Circles_GetJoinedCircles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) GetRequestedCircles(ctx context.Context, in *GetRequestedCirclesRequest, opts ...grpc.CallOption) (*GetRequestedCirclesResponse, error) {
	out := new(GetRequestedCirclesResponse)
	err := c.cc.Invoke(ctx, Circles_GetRequestedCircles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) ViewCircle(ctx context.Context, in *ViewCircleRequest, opts ...grpc.CallOption) (*ViewCircleResponse, error) {
	out := new(ViewCircleResponse)
	err := c.cc.Invoke(ctx, Circles_ViewCircle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) ExploreCircles(ctx context.Context, in *ExploreCirclesRequest, opts ...grpc.CallOption) (*ExploreCirclesResponse, error) {
	out := new(ExploreCirclesResponse)
	err := c.cc.Invoke(ctx, Circles_ExploreCircles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) PromoteToPoster(ctx context.Context, in *PromoteToPosterRequest, opts ...grpc.CallOption) (*AccessModificationResponse, error) {
	out := new(AccessModificationResponse)
	err := c.cc.Invoke(ctx, Circles_PromoteToPoster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) PromoteToAdmin(ctx context.Context, in *PromoteToAdminRequest, opts ...grpc.CallOption) (*AccessModificationResponse, error) {
	out := new(AccessModificationResponse)
	err := c.cc.Invoke(ctx, Circles_PromoteToAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) DemoteToViewer(ctx context.Context, in *DemoteToViewerRequest, opts ...grpc.CallOption) (*AccessModificationResponse, error) {
	out := new(AccessModificationResponse)
	err := c.cc.Invoke(ctx, Circles_DemoteToViewer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) GetCircleWPChangeAccess(ctx context.Context, in *GetCircleWPChangeAccessRequest, opts ...grpc.CallOption) (*CircleWPChangeAccessResponse, error) {
	out := new(CircleWPChangeAccessResponse)
	err := c.cc.Invoke(ctx, Circles_GetCircleWPChangeAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) SetCircleWPChangeAccess(ctx context.Context, in *SetCircleWPChangeAccessRequest, opts ...grpc.CallOption) (*CircleWPChangeAccessResponse, error) {
	out := new(CircleWPChangeAccessResponse)
	err := c.cc.Invoke(ctx, Circles_SetCircleWPChangeAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) CreateMood(ctx context.Context, in *CreateMoodRequest, opts ...grpc.CallOption) (*MoodResponse, error) {
	out := new(MoodResponse)
	err := c.cc.Invoke(ctx, Circles_CreateMood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) React(ctx context.Context, in *ReactRequest, opts ...grpc.CallOption) (*MoodResponse, error) {
	out := new(MoodResponse)
	err := c.cc.Invoke(ctx, Circles_React_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) RemoveReaction(ctx context.Context, in *RemoveReactionRequest, opts ...grpc.CallOption) (*MoodResponse, error) {
	out := new(MoodResponse)
	err := c.cc.Invoke(ctx, Circles_RemoveReaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circlesClient) GetAvailableReactions(ctx context.Context, in *GetAvailableReactionsRequest, opts ...grpc.CallOption) (*GetAvailableReactionsResponse, error) {
	out := new(GetAvailableReactionsResponse)
	err := c.cc.Invoke(ctx, Circles_GetAvailableReactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CirclesServer is the server API for Circles service.
// All implementations must embed UnimplementedCirclesServer
// for forward compatibility
type CirclesServer interface {
	CreateCircle(context.Context, *CreateCircleRequest) (*CreateCircleResponse, error)
	EditCircle(context.Context, *EditCircleRequest) (*EditCircleResponse, error)
	DestroyCircle(context.Context, *DestroyCircleRequest) (*DestroyCircleResponse, error)
	InviteUserToCircle(context.Context, *InviteUserToCircleRequest) (*InviteUserToCircleResponse, error)
	RollbackInviteUserToCircle(context.Context, *RollbackInviteUserToCircleRequest) (*RollbackInviteUserToCircleResponse, error)
	KickFromCircle(context.Context, *KickFromCircleRequest) (*KickFromCircleResponse, error)
	SendJoinRequest(context.Context, *SendJoinRequestRequest) (*SendJoinRequestResponse, error)
	AcceptJoinRequest(context.Context, *AcceptJoinRequestRequest) (*AcceptJoinRequestResponse, error)
	RemoveJoinRequest(context.Context, *RemoveJoinRequestRequest) (*RemoveJoinRequestResponse, error)
	LeaveCircle(context.Context, *LeaveCircleRequest) (*LeaveCircleResponse, error)
	GetJoinedCircles(context.Context, *GetJoinedCirclesRequest) (*GetJoinedCirclesResponse, error)
	GetRequestedCircles(context.Context, *GetRequestedCirclesRequest) (*GetRequestedCirclesResponse, error)
	ViewCircle(context.Context, *ViewCircleRequest) (*ViewCircleResponse, error)
	ExploreCircles(context.Context, *ExploreCirclesRequest) (*ExploreCirclesResponse, error)
	PromoteToPoster(context.Context, *PromoteToPosterRequest) (*AccessModificationResponse, error)
	PromoteToAdmin(context.Context, *PromoteToAdminRequest) (*AccessModificationResponse, error)
	DemoteToViewer(context.Context, *DemoteToViewerRequest) (*AccessModificationResponse, error)
	GetCircleWPChangeAccess(context.Context, *GetCircleWPChangeAccessRequest) (*CircleWPChangeAccessResponse, error)
	SetCircleWPChangeAccess(context.Context, *SetCircleWPChangeAccessRequest) (*CircleWPChangeAccessResponse, error)
	CreateMood(context.Context, *CreateMoodRequest) (*MoodResponse, error)
	React(context.Context, *ReactRequest) (*MoodResponse, error)
	RemoveReaction(context.Context, *RemoveReactionRequest) (*MoodResponse, error)
	GetAvailableReactions(context.Context, *GetAvailableReactionsRequest) (*GetAvailableReactionsResponse, error)
	mustEmbedUnimplementedCirclesServer()
}

// UnimplementedCirclesServer must be embedded to have forward compatible implementations.
type UnimplementedCirclesServer struct {
}

func (UnimplementedCirclesServer) CreateCircle(context.Context, *CreateCircleRequest) (*CreateCircleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCircle not implemented")
}
func (UnimplementedCirclesServer) EditCircle(context.Context, *EditCircleRequest) (*EditCircleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCircle not implemented")
}
func (UnimplementedCirclesServer) DestroyCircle(context.Context, *DestroyCircleRequest) (*DestroyCircleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyCircle not implemented")
}
func (UnimplementedCirclesServer) InviteUserToCircle(context.Context, *InviteUserToCircleRequest) (*InviteUserToCircleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteUserToCircle not implemented")
}
func (UnimplementedCirclesServer) RollbackInviteUserToCircle(context.Context, *RollbackInviteUserToCircleRequest) (*RollbackInviteUserToCircleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackInviteUserToCircle not implemented")
}
func (UnimplementedCirclesServer) KickFromCircle(context.Context, *KickFromCircleRequest) (*KickFromCircleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickFromCircle not implemented")
}
func (UnimplementedCirclesServer) SendJoinRequest(context.Context, *SendJoinRequestRequest) (*SendJoinRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendJoinRequest not implemented")
}
func (UnimplementedCirclesServer) AcceptJoinRequest(context.Context, *AcceptJoinRequestRequest) (*AcceptJoinRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptJoinRequest not implemented")
}
func (UnimplementedCirclesServer) RemoveJoinRequest(context.Context, *RemoveJoinRequestRequest) (*RemoveJoinRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveJoinRequest not implemented")
}
func (UnimplementedCirclesServer) LeaveCircle(context.Context, *LeaveCircleRequest) (*LeaveCircleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveCircle not implemented")
}
func (UnimplementedCirclesServer) GetJoinedCircles(context.Context, *GetJoinedCirclesRequest) (*GetJoinedCirclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJoinedCircles not implemented")
}
func (UnimplementedCirclesServer) GetRequestedCircles(context.Context, *GetRequestedCirclesRequest) (*GetRequestedCirclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestedCircles not implemented")
}
func (UnimplementedCirclesServer) ViewCircle(context.Context, *ViewCircleRequest) (*ViewCircleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewCircle not implemented")
}
func (UnimplementedCirclesServer) ExploreCircles(context.Context, *ExploreCirclesRequest) (*ExploreCirclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExploreCircles not implemented")
}
func (UnimplementedCirclesServer) PromoteToPoster(context.Context, *PromoteToPosterRequest) (*AccessModificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteToPoster not implemented")
}
func (UnimplementedCirclesServer) PromoteToAdmin(context.Context, *PromoteToAdminRequest) (*AccessModificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteToAdmin not implemented")
}
func (UnimplementedCirclesServer) DemoteToViewer(context.Context, *DemoteToViewerRequest) (*AccessModificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemoteToViewer not implemented")
}
func (UnimplementedCirclesServer) GetCircleWPChangeAccess(context.Context, *GetCircleWPChangeAccessRequest) (*CircleWPChangeAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCircleWPChangeAccess not implemented")
}
func (UnimplementedCirclesServer) SetCircleWPChangeAccess(context.Context, *SetCircleWPChangeAccessRequest) (*CircleWPChangeAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCircleWPChangeAccess not implemented")
}
func (UnimplementedCirclesServer) CreateMood(context.Context, *CreateMoodRequest) (*MoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMood not implemented")
}
func (UnimplementedCirclesServer) React(context.Context, *ReactRequest) (*MoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method React not implemented")
}
func (UnimplementedCirclesServer) RemoveReaction(context.Context, *RemoveReactionRequest) (*MoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReaction not implemented")
}
func (UnimplementedCirclesServer) GetAvailableReactions(context.Context, *GetAvailableReactionsRequest) (*GetAvailableReactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableReactions not implemented")
}
func (UnimplementedCirclesServer) mustEmbedUnimplementedCirclesServer() {}

// UnsafeCirclesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CirclesServer will
// result in compilation errors.
type UnsafeCirclesServer interface {
	mustEmbedUnimplementedCirclesServer()
}

func RegisterCirclesServer(s grpc.ServiceRegistrar, srv CirclesServer) {
	s.RegisterService(&Circles_ServiceDesc, srv)
}

func _Circles_CreateCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCircleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).CreateCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_CreateCircle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).CreateCircle(ctx, req.(*CreateCircleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_EditCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCircleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).EditCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_EditCircle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).EditCircle(ctx, req.(*EditCircleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_DestroyCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyCircleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).DestroyCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_DestroyCircle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).DestroyCircle(ctx, req.(*DestroyCircleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_InviteUserToCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUserToCircleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).InviteUserToCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_InviteUserToCircle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).InviteUserToCircle(ctx, req.(*InviteUserToCircleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_RollbackInviteUserToCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackInviteUserToCircleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).RollbackInviteUserToCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_RollbackInviteUserToCircle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).RollbackInviteUserToCircle(ctx, req.(*RollbackInviteUserToCircleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_KickFromCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickFromCircleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).KickFromCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_KickFromCircle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).KickFromCircle(ctx, req.(*KickFromCircleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_SendJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendJoinRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).SendJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_SendJoinRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).SendJoinRequest(ctx, req.(*SendJoinRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_AcceptJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptJoinRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).AcceptJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_AcceptJoinRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).AcceptJoinRequest(ctx, req.(*AcceptJoinRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_RemoveJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveJoinRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).RemoveJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_RemoveJoinRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).RemoveJoinRequest(ctx, req.(*RemoveJoinRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_LeaveCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveCircleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).LeaveCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_LeaveCircle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).LeaveCircle(ctx, req.(*LeaveCircleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_GetJoinedCircles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJoinedCirclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).GetJoinedCircles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_GetJoinedCircles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).GetJoinedCircles(ctx, req.(*GetJoinedCirclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_GetRequestedCircles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestedCirclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).GetRequestedCircles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_GetRequestedCircles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).GetRequestedCircles(ctx, req.(*GetRequestedCirclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_ViewCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewCircleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).ViewCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_ViewCircle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).ViewCircle(ctx, req.(*ViewCircleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_ExploreCircles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExploreCirclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).ExploreCircles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_ExploreCircles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).ExploreCircles(ctx, req.(*ExploreCirclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_PromoteToPoster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteToPosterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).PromoteToPoster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_PromoteToPoster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).PromoteToPoster(ctx, req.(*PromoteToPosterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_PromoteToAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteToAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).PromoteToAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_PromoteToAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).PromoteToAdmin(ctx, req.(*PromoteToAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_DemoteToViewer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoteToViewerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).DemoteToViewer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_DemoteToViewer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).DemoteToViewer(ctx, req.(*DemoteToViewerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_GetCircleWPChangeAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCircleWPChangeAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).GetCircleWPChangeAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_GetCircleWPChangeAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).GetCircleWPChangeAccess(ctx, req.(*GetCircleWPChangeAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_SetCircleWPChangeAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCircleWPChangeAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).SetCircleWPChangeAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_SetCircleWPChangeAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).SetCircleWPChangeAccess(ctx, req.(*SetCircleWPChangeAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_CreateMood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).CreateMood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_CreateMood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).CreateMood(ctx, req.(*CreateMoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_React_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).React(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_React_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).React(ctx, req.(*ReactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_RemoveReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).RemoveReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_RemoveReaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).RemoveReaction(ctx, req.(*RemoveReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Circles_GetAvailableReactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableReactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CirclesServer).GetAvailableReactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Circles_GetAvailableReactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CirclesServer).GetAvailableReactions(ctx, req.(*GetAvailableReactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Circles_ServiceDesc is the grpc.ServiceDesc for Circles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Circles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cpb.Circles",
	HandlerType: (*CirclesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCircle",
			Handler:    _Circles_CreateCircle_Handler,
		},
		{
			MethodName: "EditCircle",
			Handler:    _Circles_EditCircle_Handler,
		},
		{
			MethodName: "DestroyCircle",
			Handler:    _Circles_DestroyCircle_Handler,
		},
		{
			MethodName: "InviteUserToCircle",
			Handler:    _Circles_InviteUserToCircle_Handler,
		},
		{
			MethodName: "RollbackInviteUserToCircle",
			Handler:    _Circles_RollbackInviteUserToCircle_Handler,
		},
		{
			MethodName: "KickFromCircle",
			Handler:    _Circles_KickFromCircle_Handler,
		},
		{
			MethodName: "SendJoinRequest",
			Handler:    _Circles_SendJoinRequest_Handler,
		},
		{
			MethodName: "AcceptJoinRequest",
			Handler:    _Circles_AcceptJoinRequest_Handler,
		},
		{
			MethodName: "RemoveJoinRequest",
			Handler:    _Circles_RemoveJoinRequest_Handler,
		},
		{
			MethodName: "LeaveCircle",
			Handler:    _Circles_LeaveCircle_Handler,
		},
		{
			MethodName: "GetJoinedCircles",
			Handler:    _Circles_GetJoinedCircles_Handler,
		},
		{
			MethodName: "GetRequestedCircles",
			Handler:    _Circles_GetRequestedCircles_Handler,
		},
		{
			MethodName: "ViewCircle",
			Handler:    _Circles_ViewCircle_Handler,
		},
		{
			MethodName: "ExploreCircles",
			Handler:    _Circles_ExploreCircles_Handler,
		},
		{
			MethodName: "PromoteToPoster",
			Handler:    _Circles_PromoteToPoster_Handler,
		},
		{
			MethodName: "PromoteToAdmin",
			Handler:    _Circles_PromoteToAdmin_Handler,
		},
		{
			MethodName: "DemoteToViewer",
			Handler:    _Circles_DemoteToViewer_Handler,
		},
		{
			MethodName: "GetCircleWPChangeAccess",
			Handler:    _Circles_GetCircleWPChangeAccess_Handler,
		},
		{
			MethodName: "SetCircleWPChangeAccess",
			Handler:    _Circles_SetCircleWPChangeAccess_Handler,
		},
		{
			MethodName: "CreateMood",
			Handler:    _Circles_CreateMood_Handler,
		},
		{
			MethodName: "React",
			Handler:    _Circles_React_Handler,
		},
		{
			MethodName: "RemoveReaction",
			Handler:    _Circles_RemoveReaction_Handler,
		},
		{
			MethodName: "GetAvailableReactions",
			Handler:    _Circles_GetAvailableReactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_circles.proto",
}