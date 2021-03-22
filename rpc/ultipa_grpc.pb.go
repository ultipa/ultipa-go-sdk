// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ultipa

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

// UltipaRpcsClient is the client API for UltipaRpcs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UltipaRpcsClient interface {
	// 1.Sends a greeting
	SayHello(ctx context.Context, in *HelloUltipaRequest, opts ...grpc.CallOption) (*HelloUltipaReply, error)
	//2.uql
	Uql(ctx context.Context, in *UqlRequest, opts ...grpc.CallOption) (UltipaRpcs_UqlClient, error)
	//3.用户设置(用于存储用户配置信息,用户自主控制)
	UserSetting(ctx context.Context, in *UserSettingRequest, opts ...grpc.CallOption) (*UserSettingReply, error)
	//4.下载算法生成文件
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (UltipaRpcs_DownloadFileClient, error)
	//5.导出点,边数据
	Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (UltipaRpcs_ExportClient, error)
	//6. 获取raft的leader
	GetLeader(ctx context.Context, in *GetLeaderRequest, opts ...grpc.CallOption) (*GetLeaderReply, error)
	//7.插入点
	InsertNodes(ctx context.Context, in *InsertNodesRequest, opts ...grpc.CallOption) (*InsertNodesReply, error)
	//8.插入边
	InsertEdges(ctx context.Context, in *InsertEdgesRequest, opts ...grpc.CallOption) (*InsertEdgesReply, error)
	//9.uql扩展，以下命令在此接口执行执行 top, kill showTask, stopTask,clearTask show() stat listGraph
	// listAlgo getGraph createPolicy, deletePolicy, listPolicy, getPolicy,
	// grant, revoke, listPrivilege, getUser, getSelfInfo, createUser, updateUser, deleteUser, showIndex
	UqlEx(ctx context.Context, in *UqlRequest, opts ...grpc.CallOption) (UltipaRpcs_UqlExClient, error)
	//10.下载算法生成文件 v2 下载文件请求改为 算法名 + 任务号
	DownloadFileV2(ctx context.Context, in *DownloadFileRequestV2, opts ...grpc.CallOption) (UltipaRpcs_DownloadFileV2Client, error)
}

type ultipaRpcsClient struct {
	cc grpc.ClientConnInterface
}

func NewUltipaRpcsClient(cc grpc.ClientConnInterface) UltipaRpcsClient {
	return &ultipaRpcsClient{cc}
}

func (c *ultipaRpcsClient) SayHello(ctx context.Context, in *HelloUltipaRequest, opts ...grpc.CallOption) (*HelloUltipaReply, error) {
	out := new(HelloUltipaReply)
	err := c.cc.Invoke(ctx, "/ultipa.UltipaRpcs/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultipaRpcsClient) Uql(ctx context.Context, in *UqlRequest, opts ...grpc.CallOption) (UltipaRpcs_UqlClient, error) {
	stream, err := c.cc.NewStream(ctx, &UltipaRpcs_ServiceDesc.Streams[0], "/ultipa.UltipaRpcs/Uql", opts...)
	if err != nil {
		return nil, err
	}
	x := &ultipaRpcsUqlClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UltipaRpcs_UqlClient interface {
	Recv() (*UqlReply, error)
	grpc.ClientStream
}

type ultipaRpcsUqlClient struct {
	grpc.ClientStream
}

func (x *ultipaRpcsUqlClient) Recv() (*UqlReply, error) {
	m := new(UqlReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ultipaRpcsClient) UserSetting(ctx context.Context, in *UserSettingRequest, opts ...grpc.CallOption) (*UserSettingReply, error) {
	out := new(UserSettingReply)
	err := c.cc.Invoke(ctx, "/ultipa.UltipaRpcs/UserSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultipaRpcsClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (UltipaRpcs_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &UltipaRpcs_ServiceDesc.Streams[1], "/ultipa.UltipaRpcs/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &ultipaRpcsDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UltipaRpcs_DownloadFileClient interface {
	Recv() (*DownloadFileReply, error)
	grpc.ClientStream
}

type ultipaRpcsDownloadFileClient struct {
	grpc.ClientStream
}

func (x *ultipaRpcsDownloadFileClient) Recv() (*DownloadFileReply, error) {
	m := new(DownloadFileReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ultipaRpcsClient) Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (UltipaRpcs_ExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &UltipaRpcs_ServiceDesc.Streams[2], "/ultipa.UltipaRpcs/Export", opts...)
	if err != nil {
		return nil, err
	}
	x := &ultipaRpcsExportClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UltipaRpcs_ExportClient interface {
	Recv() (*ExportReply, error)
	grpc.ClientStream
}

type ultipaRpcsExportClient struct {
	grpc.ClientStream
}

func (x *ultipaRpcsExportClient) Recv() (*ExportReply, error) {
	m := new(ExportReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ultipaRpcsClient) GetLeader(ctx context.Context, in *GetLeaderRequest, opts ...grpc.CallOption) (*GetLeaderReply, error) {
	out := new(GetLeaderReply)
	err := c.cc.Invoke(ctx, "/ultipa.UltipaRpcs/GetLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultipaRpcsClient) InsertNodes(ctx context.Context, in *InsertNodesRequest, opts ...grpc.CallOption) (*InsertNodesReply, error) {
	out := new(InsertNodesReply)
	err := c.cc.Invoke(ctx, "/ultipa.UltipaRpcs/InsertNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultipaRpcsClient) InsertEdges(ctx context.Context, in *InsertEdgesRequest, opts ...grpc.CallOption) (*InsertEdgesReply, error) {
	out := new(InsertEdgesReply)
	err := c.cc.Invoke(ctx, "/ultipa.UltipaRpcs/InsertEdges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ultipaRpcsClient) UqlEx(ctx context.Context, in *UqlRequest, opts ...grpc.CallOption) (UltipaRpcs_UqlExClient, error) {
	stream, err := c.cc.NewStream(ctx, &UltipaRpcs_ServiceDesc.Streams[3], "/ultipa.UltipaRpcs/UqlEx", opts...)
	if err != nil {
		return nil, err
	}
	x := &ultipaRpcsUqlExClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UltipaRpcs_UqlExClient interface {
	Recv() (*UqlReply, error)
	grpc.ClientStream
}

type ultipaRpcsUqlExClient struct {
	grpc.ClientStream
}

func (x *ultipaRpcsUqlExClient) Recv() (*UqlReply, error) {
	m := new(UqlReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ultipaRpcsClient) DownloadFileV2(ctx context.Context, in *DownloadFileRequestV2, opts ...grpc.CallOption) (UltipaRpcs_DownloadFileV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &UltipaRpcs_ServiceDesc.Streams[4], "/ultipa.UltipaRpcs/DownloadFileV2", opts...)
	if err != nil {
		return nil, err
	}
	x := &ultipaRpcsDownloadFileV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UltipaRpcs_DownloadFileV2Client interface {
	Recv() (*DownloadFileReply, error)
	grpc.ClientStream
}

type ultipaRpcsDownloadFileV2Client struct {
	grpc.ClientStream
}

func (x *ultipaRpcsDownloadFileV2Client) Recv() (*DownloadFileReply, error) {
	m := new(DownloadFileReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UltipaRpcsServer is the server API for UltipaRpcs service.
// All implementations must embed UnimplementedUltipaRpcsServer
// for forward compatibility
type UltipaRpcsServer interface {
	// 1.Sends a greeting
	SayHello(context.Context, *HelloUltipaRequest) (*HelloUltipaReply, error)
	//2.uql
	Uql(*UqlRequest, UltipaRpcs_UqlServer) error
	//3.用户设置(用于存储用户配置信息,用户自主控制)
	UserSetting(context.Context, *UserSettingRequest) (*UserSettingReply, error)
	//4.下载算法生成文件
	DownloadFile(*DownloadFileRequest, UltipaRpcs_DownloadFileServer) error
	//5.导出点,边数据
	Export(*ExportRequest, UltipaRpcs_ExportServer) error
	//6. 获取raft的leader
	GetLeader(context.Context, *GetLeaderRequest) (*GetLeaderReply, error)
	//7.插入点
	InsertNodes(context.Context, *InsertNodesRequest) (*InsertNodesReply, error)
	//8.插入边
	InsertEdges(context.Context, *InsertEdgesRequest) (*InsertEdgesReply, error)
	//9.uql扩展，以下命令在此接口执行执行 top, kill showTask, stopTask,clearTask show() stat listGraph
	// listAlgo getGraph createPolicy, deletePolicy, listPolicy, getPolicy,
	// grant, revoke, listPrivilege, getUser, getSelfInfo, createUser, updateUser, deleteUser, showIndex
	UqlEx(*UqlRequest, UltipaRpcs_UqlExServer) error
	//10.下载算法生成文件 v2 下载文件请求改为 算法名 + 任务号
	DownloadFileV2(*DownloadFileRequestV2, UltipaRpcs_DownloadFileV2Server) error
	mustEmbedUnimplementedUltipaRpcsServer()
}

// UnimplementedUltipaRpcsServer must be embedded to have forward compatible implementations.
type UnimplementedUltipaRpcsServer struct {
}

func (UnimplementedUltipaRpcsServer) SayHello(context.Context, *HelloUltipaRequest) (*HelloUltipaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedUltipaRpcsServer) Uql(*UqlRequest, UltipaRpcs_UqlServer) error {
	return status.Errorf(codes.Unimplemented, "method Uql not implemented")
}
func (UnimplementedUltipaRpcsServer) UserSetting(context.Context, *UserSettingRequest) (*UserSettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSetting not implemented")
}
func (UnimplementedUltipaRpcsServer) DownloadFile(*DownloadFileRequest, UltipaRpcs_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedUltipaRpcsServer) Export(*ExportRequest, UltipaRpcs_ExportServer) error {
	return status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedUltipaRpcsServer) GetLeader(context.Context, *GetLeaderRequest) (*GetLeaderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeader not implemented")
}
func (UnimplementedUltipaRpcsServer) InsertNodes(context.Context, *InsertNodesRequest) (*InsertNodesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNodes not implemented")
}
func (UnimplementedUltipaRpcsServer) InsertEdges(context.Context, *InsertEdgesRequest) (*InsertEdgesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEdges not implemented")
}
func (UnimplementedUltipaRpcsServer) UqlEx(*UqlRequest, UltipaRpcs_UqlExServer) error {
	return status.Errorf(codes.Unimplemented, "method UqlEx not implemented")
}
func (UnimplementedUltipaRpcsServer) DownloadFileV2(*DownloadFileRequestV2, UltipaRpcs_DownloadFileV2Server) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFileV2 not implemented")
}
func (UnimplementedUltipaRpcsServer) mustEmbedUnimplementedUltipaRpcsServer() {}

// UnsafeUltipaRpcsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UltipaRpcsServer will
// result in compilation errors.
type UnsafeUltipaRpcsServer interface {
	mustEmbedUnimplementedUltipaRpcsServer()
}

func RegisterUltipaRpcsServer(s grpc.ServiceRegistrar, srv UltipaRpcsServer) {
	s.RegisterService(&UltipaRpcs_ServiceDesc, srv)
}

func _UltipaRpcs_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloUltipaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltipaRpcsServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultipa.UltipaRpcs/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltipaRpcsServer).SayHello(ctx, req.(*HelloUltipaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltipaRpcs_Uql_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UqlRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UltipaRpcsServer).Uql(m, &ultipaRpcsUqlServer{stream})
}

type UltipaRpcs_UqlServer interface {
	Send(*UqlReply) error
	grpc.ServerStream
}

type ultipaRpcsUqlServer struct {
	grpc.ServerStream
}

func (x *ultipaRpcsUqlServer) Send(m *UqlReply) error {
	return x.ServerStream.SendMsg(m)
}

func _UltipaRpcs_UserSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltipaRpcsServer).UserSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultipa.UltipaRpcs/UserSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltipaRpcsServer).UserSetting(ctx, req.(*UserSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltipaRpcs_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UltipaRpcsServer).DownloadFile(m, &ultipaRpcsDownloadFileServer{stream})
}

type UltipaRpcs_DownloadFileServer interface {
	Send(*DownloadFileReply) error
	grpc.ServerStream
}

type ultipaRpcsDownloadFileServer struct {
	grpc.ServerStream
}

func (x *ultipaRpcsDownloadFileServer) Send(m *DownloadFileReply) error {
	return x.ServerStream.SendMsg(m)
}

func _UltipaRpcs_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExportRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UltipaRpcsServer).Export(m, &ultipaRpcsExportServer{stream})
}

type UltipaRpcs_ExportServer interface {
	Send(*ExportReply) error
	grpc.ServerStream
}

type ultipaRpcsExportServer struct {
	grpc.ServerStream
}

func (x *ultipaRpcsExportServer) Send(m *ExportReply) error {
	return x.ServerStream.SendMsg(m)
}

func _UltipaRpcs_GetLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltipaRpcsServer).GetLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultipa.UltipaRpcs/GetLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltipaRpcsServer).GetLeader(ctx, req.(*GetLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltipaRpcs_InsertNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltipaRpcsServer).InsertNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultipa.UltipaRpcs/InsertNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltipaRpcsServer).InsertNodes(ctx, req.(*InsertNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltipaRpcs_InsertEdges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertEdgesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltipaRpcsServer).InsertEdges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ultipa.UltipaRpcs/InsertEdges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltipaRpcsServer).InsertEdges(ctx, req.(*InsertEdgesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UltipaRpcs_UqlEx_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UqlRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UltipaRpcsServer).UqlEx(m, &ultipaRpcsUqlExServer{stream})
}

type UltipaRpcs_UqlExServer interface {
	Send(*UqlReply) error
	grpc.ServerStream
}

type ultipaRpcsUqlExServer struct {
	grpc.ServerStream
}

func (x *ultipaRpcsUqlExServer) Send(m *UqlReply) error {
	return x.ServerStream.SendMsg(m)
}

func _UltipaRpcs_DownloadFileV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequestV2)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UltipaRpcsServer).DownloadFileV2(m, &ultipaRpcsDownloadFileV2Server{stream})
}

type UltipaRpcs_DownloadFileV2Server interface {
	Send(*DownloadFileReply) error
	grpc.ServerStream
}

type ultipaRpcsDownloadFileV2Server struct {
	grpc.ServerStream
}

func (x *ultipaRpcsDownloadFileV2Server) Send(m *DownloadFileReply) error {
	return x.ServerStream.SendMsg(m)
}

// UltipaRpcs_ServiceDesc is the grpc.ServiceDesc for UltipaRpcs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UltipaRpcs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ultipa.UltipaRpcs",
	HandlerType: (*UltipaRpcsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _UltipaRpcs_SayHello_Handler,
		},
		{
			MethodName: "UserSetting",
			Handler:    _UltipaRpcs_UserSetting_Handler,
		},
		{
			MethodName: "GetLeader",
			Handler:    _UltipaRpcs_GetLeader_Handler,
		},
		{
			MethodName: "InsertNodes",
			Handler:    _UltipaRpcs_InsertNodes_Handler,
		},
		{
			MethodName: "InsertEdges",
			Handler:    _UltipaRpcs_InsertEdges_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Uql",
			Handler:       _UltipaRpcs_Uql_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _UltipaRpcs_DownloadFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Export",
			Handler:       _UltipaRpcs_Export_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UqlEx",
			Handler:       _UltipaRpcs_UqlEx_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DownloadFileV2",
			Handler:       _UltipaRpcs_DownloadFileV2_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ultipa.proto",
}
