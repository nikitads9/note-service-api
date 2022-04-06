// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package note_api

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

// NoteV1Client is the client API for NoteV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteV1Client interface {
	AddNote(ctx context.Context, in *AddNoteRequest, opts ...grpc.CallOption) (*AddNoteResponse, error)
	RemoveNote(ctx context.Context, in *RemoveNoteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type noteV1Client struct {
	cc grpc.ClientConnInterface
}

func NewNoteV1Client(cc grpc.ClientConnInterface) NoteV1Client {
	return &noteV1Client{cc}
}

func (c *noteV1Client) AddNote(ctx context.Context, in *AddNoteRequest, opts ...grpc.CallOption) (*AddNoteResponse, error) {
	out := new(AddNoteResponse)
	err := c.cc.Invoke(ctx, "/NoteV1/AddNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteV1Client) RemoveNote(ctx context.Context, in *RemoveNoteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/NoteV1/RemoveNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteV1Server is the server API for NoteV1 service.
// All implementations must embed UnimplementedNoteV1Server
// for forward compatibility
type NoteV1Server interface {
	AddNote(context.Context, *AddNoteRequest) (*AddNoteResponse, error)
	RemoveNote(context.Context, *RemoveNoteRequest) (*Empty, error)
	mustEmbedUnimplementedNoteV1Server()
}

// UnimplementedNoteV1Server must be embedded to have forward compatible implementations.
type UnimplementedNoteV1Server struct {
}

func (UnimplementedNoteV1Server) AddNote(context.Context, *AddNoteRequest) (*AddNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNote not implemented")
}
func (UnimplementedNoteV1Server) RemoveNote(context.Context, *RemoveNoteRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNote not implemented")
}
func (UnimplementedNoteV1Server) mustEmbedUnimplementedNoteV1Server() {}

// UnsafeNoteV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteV1Server will
// result in compilation errors.
type UnsafeNoteV1Server interface {
	mustEmbedUnimplementedNoteV1Server()
}

func RegisterNoteV1Server(s grpc.ServiceRegistrar, srv NoteV1Server) {
	s.RegisterService(&NoteV1_ServiceDesc, srv)
}

func _NoteV1_AddNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteV1Server).AddNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NoteV1/AddNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteV1Server).AddNote(ctx, req.(*AddNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteV1_RemoveNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteV1Server).RemoveNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NoteV1/RemoveNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteV1Server).RemoveNote(ctx, req.(*RemoveNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteV1_ServiceDesc is the grpc.ServiceDesc for NoteV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NoteV1",
	HandlerType: (*NoteV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNote",
			Handler:    _NoteV1_AddNote_Handler,
		},
		{
			MethodName: "RemoveNote",
			Handler:    _NoteV1_RemoveNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "note_v1.proto",
}