// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package cluster_manager_api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	HelloWorldMsg
	HelloWorldReply
	GetPodCountMsg
	GetPodCountReply
	CreateClusterMsg
	CreateClusterReply
	GetClusterMsg
	GetClusterReply
	DeleteClusterMsg
	DeleteClusterReply
*/
package cluster_manager_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The Hello World request
type HelloWorldMsg struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloWorldMsg) Reset()                    { *m = HelloWorldMsg{} }
func (m *HelloWorldMsg) String() string            { return proto.CompactTextString(m) }
func (*HelloWorldMsg) ProtoMessage()               {}
func (*HelloWorldMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloWorldMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response to Hello World
type HelloWorldReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloWorldReply) Reset()                    { *m = HelloWorldReply{} }
func (m *HelloWorldReply) String() string            { return proto.CompactTextString(m) }
func (*HelloWorldReply) ProtoMessage()               {}
func (*HelloWorldReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloWorldReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetPodCountMsg struct {
	// Namespace to fetch the pod count
	// Leave empty to query all namespaces
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *GetPodCountMsg) Reset()                    { *m = GetPodCountMsg{} }
func (m *GetPodCountMsg) String() string            { return proto.CompactTextString(m) }
func (*GetPodCountMsg) ProtoMessage()               {}
func (*GetPodCountMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetPodCountMsg) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type GetPodCountReply struct {
	// Number of pods in the namespace (or all namespaces)
	Pods int32 `protobuf:"varint,1,opt,name=pods" json:"pods,omitempty"`
}

func (m *GetPodCountReply) Reset()                    { *m = GetPodCountReply{} }
func (m *GetPodCountReply) String() string            { return proto.CompactTextString(m) }
func (*GetPodCountReply) ProtoMessage()               {}
func (*GetPodCountReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetPodCountReply) GetPods() int32 {
	if m != nil {
		return m.Pods
	}
	return 0
}

type CreateClusterMsg struct {
	// Name of the cluster to be provisioned
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CreateClusterMsg) Reset()                    { *m = CreateClusterMsg{} }
func (m *CreateClusterMsg) String() string            { return proto.CompactTextString(m) }
func (*CreateClusterMsg) ProtoMessage()               {}
func (*CreateClusterMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateClusterMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateClusterReply struct {
	// Whether or not the cluster was provisioned by this request
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	// The state of the cluster generation
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *CreateClusterReply) Reset()                    { *m = CreateClusterReply{} }
func (m *CreateClusterReply) String() string            { return proto.CompactTextString(m) }
func (*CreateClusterReply) ProtoMessage()               {}
func (*CreateClusterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateClusterReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *CreateClusterReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetClusterMsg struct {
	// Name of the cluster to be looked up
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetClusterMsg) Reset()                    { *m = GetClusterMsg{} }
func (m *GetClusterMsg) String() string            { return proto.CompactTextString(m) }
func (*GetClusterMsg) ProtoMessage()               {}
func (*GetClusterMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetClusterMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetClusterReply struct {
	// Is the cluster in the system
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	// What is the status of the cluster
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	// What is the kubeconfig to connect to the cluster
	Kubeconfig string `protobuf:"bytes,3,opt,name=kubeconfig" json:"kubeconfig,omitempty"`
}

func (m *GetClusterReply) Reset()                    { *m = GetClusterReply{} }
func (m *GetClusterReply) String() string            { return proto.CompactTextString(m) }
func (*GetClusterReply) ProtoMessage()               {}
func (*GetClusterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetClusterReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *GetClusterReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetClusterReply) GetKubeconfig() string {
	if m != nil {
		return m.Kubeconfig
	}
	return ""
}

type DeleteClusterMsg struct {
	// What is the cluster's name to destroy
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteClusterMsg) Reset()                    { *m = DeleteClusterMsg{} }
func (m *DeleteClusterMsg) String() string            { return proto.CompactTextString(m) }
func (*DeleteClusterMsg) ProtoMessage()               {}
func (*DeleteClusterMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteClusterMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteClusterReply struct {
	// Could the cluster be destroyed
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	// Status of the request
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *DeleteClusterReply) Reset()                    { *m = DeleteClusterReply{} }
func (m *DeleteClusterReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteClusterReply) ProtoMessage()               {}
func (*DeleteClusterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteClusterReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *DeleteClusterReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloWorldMsg)(nil), "cluster_manager_api.HelloWorldMsg")
	proto.RegisterType((*HelloWorldReply)(nil), "cluster_manager_api.HelloWorldReply")
	proto.RegisterType((*GetPodCountMsg)(nil), "cluster_manager_api.GetPodCountMsg")
	proto.RegisterType((*GetPodCountReply)(nil), "cluster_manager_api.GetPodCountReply")
	proto.RegisterType((*CreateClusterMsg)(nil), "cluster_manager_api.CreateClusterMsg")
	proto.RegisterType((*CreateClusterReply)(nil), "cluster_manager_api.CreateClusterReply")
	proto.RegisterType((*GetClusterMsg)(nil), "cluster_manager_api.GetClusterMsg")
	proto.RegisterType((*GetClusterReply)(nil), "cluster_manager_api.GetClusterReply")
	proto.RegisterType((*DeleteClusterMsg)(nil), "cluster_manager_api.DeleteClusterMsg")
	proto.RegisterType((*DeleteClusterReply)(nil), "cluster_manager_api.DeleteClusterReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cluster service

type ClusterClient interface {
	// Simple Hello World Service - will repeat a greeting to the name provided
	HelloWorld(ctx context.Context, in *HelloWorldMsg, opts ...grpc.CallOption) (*HelloWorldReply, error)
	// Simple pod count response.  Will respond with the number of pods in the namespace provided
	GetPodCount(ctx context.Context, in *GetPodCountMsg, opts ...grpc.CallOption) (*GetPodCountReply, error)
	// Will provision a cluster
	CreateCluster(ctx context.Context, in *CreateClusterMsg, opts ...grpc.CallOption) (*CreateClusterReply, error)
	// Will retrieve the status of a cluster and its kubeconfig for connectivity
	GetCluster(ctx context.Context, in *GetClusterMsg, opts ...grpc.CallOption) (*GetClusterReply, error)
	// Will delete a cluster
	DeleteCluster(ctx context.Context, in *DeleteClusterMsg, opts ...grpc.CallOption) (*DeleteClusterReply, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) HelloWorld(ctx context.Context, in *HelloWorldMsg, opts ...grpc.CallOption) (*HelloWorldReply, error) {
	out := new(HelloWorldReply)
	err := grpc.Invoke(ctx, "/cluster_manager_api.Cluster/HelloWorld", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) GetPodCount(ctx context.Context, in *GetPodCountMsg, opts ...grpc.CallOption) (*GetPodCountReply, error) {
	out := new(GetPodCountReply)
	err := grpc.Invoke(ctx, "/cluster_manager_api.Cluster/GetPodCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) CreateCluster(ctx context.Context, in *CreateClusterMsg, opts ...grpc.CallOption) (*CreateClusterReply, error) {
	out := new(CreateClusterReply)
	err := grpc.Invoke(ctx, "/cluster_manager_api.Cluster/CreateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) GetCluster(ctx context.Context, in *GetClusterMsg, opts ...grpc.CallOption) (*GetClusterReply, error) {
	out := new(GetClusterReply)
	err := grpc.Invoke(ctx, "/cluster_manager_api.Cluster/GetCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) DeleteCluster(ctx context.Context, in *DeleteClusterMsg, opts ...grpc.CallOption) (*DeleteClusterReply, error) {
	out := new(DeleteClusterReply)
	err := grpc.Invoke(ctx, "/cluster_manager_api.Cluster/DeleteCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cluster service

type ClusterServer interface {
	// Simple Hello World Service - will repeat a greeting to the name provided
	HelloWorld(context.Context, *HelloWorldMsg) (*HelloWorldReply, error)
	// Simple pod count response.  Will respond with the number of pods in the namespace provided
	GetPodCount(context.Context, *GetPodCountMsg) (*GetPodCountReply, error)
	// Will provision a cluster
	CreateCluster(context.Context, *CreateClusterMsg) (*CreateClusterReply, error)
	// Will retrieve the status of a cluster and its kubeconfig for connectivity
	GetCluster(context.Context, *GetClusterMsg) (*GetClusterReply, error)
	// Will delete a cluster
	DeleteCluster(context.Context, *DeleteClusterMsg) (*DeleteClusterReply, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster_manager_api.Cluster/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).HelloWorld(ctx, req.(*HelloWorldMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_GetPodCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPodCountMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).GetPodCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster_manager_api.Cluster/GetPodCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).GetPodCount(ctx, req.(*GetPodCountMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster_manager_api.Cluster/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).CreateCluster(ctx, req.(*CreateClusterMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster_manager_api.Cluster/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).GetCluster(ctx, req.(*GetClusterMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cluster_manager_api.Cluster/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).DeleteCluster(ctx, req.(*DeleteClusterMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cluster_manager_api.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _Cluster_HelloWorld_Handler,
		},
		{
			MethodName: "GetPodCount",
			Handler:    _Cluster_GetPodCount_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _Cluster_CreateCluster_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _Cluster_GetCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _Cluster_DeleteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x65, 0x53, 0x5a, 0x7a, 0x50, 0x9b, 0x32, 0x45, 0x34, 0xb2, 0x4a, 0x65, 0xb9, 0xdc,
	0x54, 0x88, 0xdd, 0x96, 0x5d, 0xc5, 0x82, 0x12, 0xa4, 0xd2, 0x45, 0xa4, 0x28, 0x0b, 0x10, 0xab,
	0x32, 0x71, 0x0e, 0x13, 0x13, 0xdb, 0x33, 0xf2, 0x8c, 0x13, 0x02, 0x3b, 0x1e, 0x01, 0x1e, 0x0d,
	0xb1, 0x60, 0x0d, 0x0f, 0x82, 0x3c, 0xb6, 0x95, 0x38, 0x17, 0x85, 0xae, 0x3c, 0x97, 0x6f, 0xfe,
	0xff, 0xcc, 0xf1, 0xaf, 0x81, 0x4d, 0x2a, 0x02, 0x57, 0x24, 0x5c, 0x71, 0xb2, 0xeb, 0x87, 0xa9,
	0x54, 0x98, 0x5c, 0x45, 0x34, 0xa6, 0x0c, 0x93, 0x2b, 0x2a, 0x02, 0x6b, 0x9f, 0x71, 0xce, 0x42,
	0xf4, 0xa8, 0x08, 0x3c, 0x1a, 0xc7, 0x5c, 0x51, 0x15, 0xf0, 0x58, 0xe6, 0x47, 0xac, 0x67, 0xfa,
	0xe3, 0x37, 0x18, 0xc6, 0x0d, 0x39, 0xa2, 0x8c, 0x61, 0xe2, 0x71, 0xa1, 0x89, 0x79, 0xda, 0x39,
	0x84, 0xad, 0x37, 0x18, 0x86, 0xfc, 0x1d, 0x4f, 0xc2, 0x5e, 0x4b, 0x32, 0x42, 0x60, 0x2d, 0xa6,
	0x11, 0xd6, 0x0d, 0xdb, 0x78, 0xb2, 0xd9, 0xd1, 0x63, 0xe7, 0x29, 0xd4, 0x26, 0x50, 0x07, 0x45,
	0x38, 0x26, 0x75, 0xd8, 0x88, 0x50, 0x4a, 0xca, 0x4a, 0xb2, 0x9c, 0x3a, 0x2e, 0x6c, 0x5f, 0xa0,
	0x6a, 0xf3, 0x5e, 0x93, 0xa7, 0xb1, 0xca, 0x24, 0xf7, 0x61, 0x33, 0x93, 0x91, 0x82, 0xfa, 0x25,
	0x3d, 0x59, 0x70, 0x1e, 0xc1, 0xce, 0x14, 0x9f, 0xab, 0x13, 0x58, 0x13, 0xbc, 0x27, 0x35, 0x7c,
	0xb3, 0xa3, 0xc7, 0x19, 0xd7, 0x4c, 0x90, 0x2a, 0x6c, 0xe6, 0x2d, 0x59, 0x56, 0xec, 0x0b, 0x20,
	0x15, 0x2e, 0x57, 0xdc, 0x06, 0x93, 0x0f, 0x34, 0x77, 0xab, 0x63, 0xf2, 0x01, 0xb9, 0x07, 0xeb,
	0x52, 0x51, 0x95, 0xca, 0xba, 0xa9, 0xcf, 0x16, 0xb3, 0xac, 0x1f, 0x17, 0xa8, 0x56, 0x58, 0xbc,
	0x87, 0xda, 0x04, 0xba, 0x96, 0x3e, 0x39, 0x00, 0x18, 0xa4, 0x5d, 0xf4, 0x79, 0xfc, 0x31, 0x60,
	0xf5, 0x1b, 0x7a, 0x6f, 0x6a, 0x25, 0xbb, 0xe5, 0x6b, 0x0c, 0xf1, 0x7f, 0x6e, 0x59, 0xe1, 0xae,
	0x55, 0xc5, 0xe9, 0x9f, 0x35, 0xd8, 0x28, 0x0e, 0x92, 0x14, 0x60, 0xf2, 0x73, 0x89, 0xe3, 0x2e,
	0x48, 0x9c, 0x5b, 0x89, 0x88, 0xf5, 0x60, 0x05, 0xa3, 0x6b, 0x71, 0xee, 0x7f, 0xfb, 0xf9, 0xf7,
	0x87, 0xb9, 0xe7, 0x10, 0x9d, 0xd3, 0xe1, 0x89, 0xd7, 0xcf, 0x80, 0x51, 0x06, 0x9c, 0x19, 0x47,
	0x64, 0x0c, 0xb7, 0xa7, 0x7e, 0x3b, 0x39, 0x5c, 0xa8, 0x59, 0x0d, 0x92, 0xf5, 0x70, 0x15, 0x94,
	0x3b, 0x1f, 0x68, 0xe7, 0xba, 0xb3, 0x5b, 0x3a, 0x33, 0x54, 0x82, 0xf7, 0xfc, 0x8c, 0xc8, 0xac,
	0xbf, 0xc2, 0x56, 0x25, 0x21, 0x64, 0xb1, 0xee, 0x6c, 0xda, 0xac, 0xc7, 0xab, 0xb1, 0xbc, 0x00,
	0x4b, 0x17, 0x70, 0xd7, 0xaa, 0x95, 0x05, 0x14, 0xe7, 0x32, 0x73, 0x0e, 0x30, 0xc9, 0xce, 0x92,
	0x76, 0x57, 0x12, 0xb8, 0xa4, 0xdd, 0x33, 0x01, 0x74, 0xf6, 0xb4, 0xe7, 0x1d, 0x32, 0xeb, 0x49,
	0x3e, 0xc3, 0x56, 0x25, 0x29, 0x4b, 0x6e, 0x3b, 0x9b, 0xba, 0x25, 0xb7, 0x9d, 0x0f, 0x5d, 0xe9,
	0x7c, 0x34, 0xeb, 0xfc, 0xea, 0xb7, 0xf1, 0xfd, 0xfc, 0x97, 0x41, 0x3e, 0xc0, 0x6e, 0xc1, 0xdb,
	0xad, 0x5c, 0xc9, 0x3e, 0x6f, 0x5f, 0x3a, 0xe7, 0x50, 0x6b, 0x05, 0x7e, 0x9f, 0x62, 0x68, 0xbf,
	0xc5, 0x18, 0xbf, 0x04, 0x94, 0x58, 0x7d, 0xa5, 0x84, 0x3c, 0xf3, 0x3c, 0x16, 0xa8, 0x7e, 0xda,
	0x75, 0x7d, 0x1e, 0x79, 0xc3, 0x7c, 0xcf, 0x22, 0x51, 0x31, 0x7a, 0xc9, 0x22, 0x1a, 0x84, 0xd9,
	0xde, 0xe9, 0xfa, 0xf0, 0xd8, 0x3d, 0x71, 0x8f, 0x8f, 0x4c, 0xc3, 0x3c, 0xdd, 0xa1, 0x42, 0x84,
	0x81, 0xaf, 0x1f, 0x37, 0xef, 0x93, 0xe4, 0xf1, 0xd9, 0xdc, 0x4a, 0x72, 0x09, 0x7b, 0x2d, 0x9e,
	0xa0, 0x4d, 0xbb, 0x3c, 0x55, 0x36, 0xeb, 0xb4, 0x9b, 0x8d, 0x0b, 0xaa, 0x70, 0x44, 0xc7, 0xc4,
	0x5d, 0x60, 0xcd, 0x12, 0xe1, 0x37, 0xd0, 0xe7, 0x72, 0x2c, 0x15, 0x16, 0x53, 0x96, 0xf3, 0xdd,
	0x75, 0xfd, 0x7c, 0x3e, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x99, 0xd0, 0xcf, 0x87, 0xac, 0x05,
	0x00, 0x00,
}
