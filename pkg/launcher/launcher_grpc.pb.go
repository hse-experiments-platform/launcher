// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: launcher/launcher.proto

package launcher

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
	LauncherService_GetLaunches_FullMethodName                     = "/github.hse_experiments_platform.launcher.LauncherService/GetLaunches"
	LauncherService_LaunchDatasetUpload_FullMethodName             = "/github.hse_experiments_platform.launcher.LauncherService/LaunchDatasetUpload"
	LauncherService_GetDatasetUploadLaunch_FullMethodName          = "/github.hse_experiments_platform.launcher.LauncherService/GetDatasetUploadLaunch"
	LauncherService_LaunchDatasetSetTypes_FullMethodName           = "/github.hse_experiments_platform.launcher.LauncherService/LaunchDatasetSetTypes"
	LauncherService_GetDatasetSetTypesLaunch_FullMethodName        = "/github.hse_experiments_platform.launcher.LauncherService/GetDatasetSetTypesLaunch"
	LauncherService_LaunchDatasetPreprocess_FullMethodName         = "/github.hse_experiments_platform.launcher.LauncherService/LaunchDatasetPreprocess"
	LauncherService_GetDatasetPreprocessLaunch_FullMethodName      = "/github.hse_experiments_platform.launcher.LauncherService/GetDatasetPreprocessLaunch"
	LauncherService_LaunchTrain_FullMethodName                     = "/github.hse_experiments_platform.launcher.LauncherService/LaunchTrain"
	LauncherService_GetTrainLaunch_FullMethodName                  = "/github.hse_experiments_platform.launcher.LauncherService/GetTrainLaunch"
	LauncherService_LaunchPredict_FullMethodName                   = "/github.hse_experiments_platform.launcher.LauncherService/LaunchPredict"
	LauncherService_GetPredictLaunch_FullMethodName                = "/github.hse_experiments_platform.launcher.LauncherService/GetPredictLaunch"
	LauncherService_GetPredictionResultDownloadLink_FullMethodName = "/github.hse_experiments_platform.launcher.LauncherService/GetPredictionResultDownloadLink"
	LauncherService_LaunchGenericConvert_FullMethodName            = "/github.hse_experiments_platform.launcher.LauncherService/LaunchGenericConvert"
	LauncherService_GetGenericConvertLaunch_FullMethodName         = "/github.hse_experiments_platform.launcher.LauncherService/GetGenericConvertLaunch"
	LauncherService_WaitForLaunchFinish_FullMethodName             = "/github.hse_experiments_platform.launcher.LauncherService/WaitForLaunchFinish"
)

// LauncherServiceClient is the client API for LauncherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LauncherServiceClient interface {
	GetLaunches(ctx context.Context, in *GetLaunchesRequest, opts ...grpc.CallOption) (*GetLaunchesResponse, error)
	LaunchDatasetUpload(ctx context.Context, in *LaunchDatasetUploadRequest, opts ...grpc.CallOption) (*LaunchDatasetUploadResponse, error)
	GetDatasetUploadLaunch(ctx context.Context, in *GetDatasetUploadLaunchRequest, opts ...grpc.CallOption) (*GetDatasetUploadLaunchResponse, error)
	LaunchDatasetSetTypes(ctx context.Context, in *LaunchDatasetSetTypesRequest, opts ...grpc.CallOption) (*LaunchDatasetSetTypesResponse, error)
	GetDatasetSetTypesLaunch(ctx context.Context, in *GetDatasetSetTypesLaunchRequest, opts ...grpc.CallOption) (*GetDatasetSetTypesLaunchResponse, error)
	LaunchDatasetPreprocess(ctx context.Context, in *LaunchDatasetPreprocessRequest, opts ...grpc.CallOption) (*LaunchDatasetPreprocessResponse, error)
	GetDatasetPreprocessLaunch(ctx context.Context, in *GetDatasetPreprocessLaunchRequest, opts ...grpc.CallOption) (*GetDatasetPreprocessLaunchResponse, error)
	LaunchTrain(ctx context.Context, in *LaunchTrainRequest, opts ...grpc.CallOption) (*LaunchTrainResponse, error)
	GetTrainLaunch(ctx context.Context, in *GetTrainLaunchRequest, opts ...grpc.CallOption) (*GetTrainLaunchResponse, error)
	LaunchPredict(ctx context.Context, in *LaunchPredictRequest, opts ...grpc.CallOption) (*LaunchPredictResponse, error)
	GetPredictLaunch(ctx context.Context, in *GetPredictLaunchRequest, opts ...grpc.CallOption) (*GetPredictLaunchResponse, error)
	GetPredictionResultDownloadLink(ctx context.Context, in *GetPredictionResultDownloadLinkRequest, opts ...grpc.CallOption) (*GetPredictionResultDownloadLinkResponse, error)
	LaunchGenericConvert(ctx context.Context, in *LaunchGenericConvertRequest, opts ...grpc.CallOption) (*LaunchGenericConvertResponse, error)
	GetGenericConvertLaunch(ctx context.Context, in *GetGenericConvertLaunchRequest, opts ...grpc.CallOption) (*GetGenericConvertLaunchResponse, error)
	WaitForLaunchFinish(ctx context.Context, in *WaitForLaunchFinishRequest, opts ...grpc.CallOption) (*WaitForLaunchFinishResponse, error)
}

type launcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLauncherServiceClient(cc grpc.ClientConnInterface) LauncherServiceClient {
	return &launcherServiceClient{cc}
}

func (c *launcherServiceClient) GetLaunches(ctx context.Context, in *GetLaunchesRequest, opts ...grpc.CallOption) (*GetLaunchesResponse, error) {
	out := new(GetLaunchesResponse)
	err := c.cc.Invoke(ctx, LauncherService_GetLaunches_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) LaunchDatasetUpload(ctx context.Context, in *LaunchDatasetUploadRequest, opts ...grpc.CallOption) (*LaunchDatasetUploadResponse, error) {
	out := new(LaunchDatasetUploadResponse)
	err := c.cc.Invoke(ctx, LauncherService_LaunchDatasetUpload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) GetDatasetUploadLaunch(ctx context.Context, in *GetDatasetUploadLaunchRequest, opts ...grpc.CallOption) (*GetDatasetUploadLaunchResponse, error) {
	out := new(GetDatasetUploadLaunchResponse)
	err := c.cc.Invoke(ctx, LauncherService_GetDatasetUploadLaunch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) LaunchDatasetSetTypes(ctx context.Context, in *LaunchDatasetSetTypesRequest, opts ...grpc.CallOption) (*LaunchDatasetSetTypesResponse, error) {
	out := new(LaunchDatasetSetTypesResponse)
	err := c.cc.Invoke(ctx, LauncherService_LaunchDatasetSetTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) GetDatasetSetTypesLaunch(ctx context.Context, in *GetDatasetSetTypesLaunchRequest, opts ...grpc.CallOption) (*GetDatasetSetTypesLaunchResponse, error) {
	out := new(GetDatasetSetTypesLaunchResponse)
	err := c.cc.Invoke(ctx, LauncherService_GetDatasetSetTypesLaunch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) LaunchDatasetPreprocess(ctx context.Context, in *LaunchDatasetPreprocessRequest, opts ...grpc.CallOption) (*LaunchDatasetPreprocessResponse, error) {
	out := new(LaunchDatasetPreprocessResponse)
	err := c.cc.Invoke(ctx, LauncherService_LaunchDatasetPreprocess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) GetDatasetPreprocessLaunch(ctx context.Context, in *GetDatasetPreprocessLaunchRequest, opts ...grpc.CallOption) (*GetDatasetPreprocessLaunchResponse, error) {
	out := new(GetDatasetPreprocessLaunchResponse)
	err := c.cc.Invoke(ctx, LauncherService_GetDatasetPreprocessLaunch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) LaunchTrain(ctx context.Context, in *LaunchTrainRequest, opts ...grpc.CallOption) (*LaunchTrainResponse, error) {
	out := new(LaunchTrainResponse)
	err := c.cc.Invoke(ctx, LauncherService_LaunchTrain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) GetTrainLaunch(ctx context.Context, in *GetTrainLaunchRequest, opts ...grpc.CallOption) (*GetTrainLaunchResponse, error) {
	out := new(GetTrainLaunchResponse)
	err := c.cc.Invoke(ctx, LauncherService_GetTrainLaunch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) LaunchPredict(ctx context.Context, in *LaunchPredictRequest, opts ...grpc.CallOption) (*LaunchPredictResponse, error) {
	out := new(LaunchPredictResponse)
	err := c.cc.Invoke(ctx, LauncherService_LaunchPredict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) GetPredictLaunch(ctx context.Context, in *GetPredictLaunchRequest, opts ...grpc.CallOption) (*GetPredictLaunchResponse, error) {
	out := new(GetPredictLaunchResponse)
	err := c.cc.Invoke(ctx, LauncherService_GetPredictLaunch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) GetPredictionResultDownloadLink(ctx context.Context, in *GetPredictionResultDownloadLinkRequest, opts ...grpc.CallOption) (*GetPredictionResultDownloadLinkResponse, error) {
	out := new(GetPredictionResultDownloadLinkResponse)
	err := c.cc.Invoke(ctx, LauncherService_GetPredictionResultDownloadLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) LaunchGenericConvert(ctx context.Context, in *LaunchGenericConvertRequest, opts ...grpc.CallOption) (*LaunchGenericConvertResponse, error) {
	out := new(LaunchGenericConvertResponse)
	err := c.cc.Invoke(ctx, LauncherService_LaunchGenericConvert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) GetGenericConvertLaunch(ctx context.Context, in *GetGenericConvertLaunchRequest, opts ...grpc.CallOption) (*GetGenericConvertLaunchResponse, error) {
	out := new(GetGenericConvertLaunchResponse)
	err := c.cc.Invoke(ctx, LauncherService_GetGenericConvertLaunch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherServiceClient) WaitForLaunchFinish(ctx context.Context, in *WaitForLaunchFinishRequest, opts ...grpc.CallOption) (*WaitForLaunchFinishResponse, error) {
	out := new(WaitForLaunchFinishResponse)
	err := c.cc.Invoke(ctx, LauncherService_WaitForLaunchFinish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LauncherServiceServer is the server API for LauncherService service.
// All implementations should embed UnimplementedLauncherServiceServer
// for forward compatibility
type LauncherServiceServer interface {
	GetLaunches(context.Context, *GetLaunchesRequest) (*GetLaunchesResponse, error)
	LaunchDatasetUpload(context.Context, *LaunchDatasetUploadRequest) (*LaunchDatasetUploadResponse, error)
	GetDatasetUploadLaunch(context.Context, *GetDatasetUploadLaunchRequest) (*GetDatasetUploadLaunchResponse, error)
	LaunchDatasetSetTypes(context.Context, *LaunchDatasetSetTypesRequest) (*LaunchDatasetSetTypesResponse, error)
	GetDatasetSetTypesLaunch(context.Context, *GetDatasetSetTypesLaunchRequest) (*GetDatasetSetTypesLaunchResponse, error)
	LaunchDatasetPreprocess(context.Context, *LaunchDatasetPreprocessRequest) (*LaunchDatasetPreprocessResponse, error)
	GetDatasetPreprocessLaunch(context.Context, *GetDatasetPreprocessLaunchRequest) (*GetDatasetPreprocessLaunchResponse, error)
	LaunchTrain(context.Context, *LaunchTrainRequest) (*LaunchTrainResponse, error)
	GetTrainLaunch(context.Context, *GetTrainLaunchRequest) (*GetTrainLaunchResponse, error)
	LaunchPredict(context.Context, *LaunchPredictRequest) (*LaunchPredictResponse, error)
	GetPredictLaunch(context.Context, *GetPredictLaunchRequest) (*GetPredictLaunchResponse, error)
	GetPredictionResultDownloadLink(context.Context, *GetPredictionResultDownloadLinkRequest) (*GetPredictionResultDownloadLinkResponse, error)
	LaunchGenericConvert(context.Context, *LaunchGenericConvertRequest) (*LaunchGenericConvertResponse, error)
	GetGenericConvertLaunch(context.Context, *GetGenericConvertLaunchRequest) (*GetGenericConvertLaunchResponse, error)
	WaitForLaunchFinish(context.Context, *WaitForLaunchFinishRequest) (*WaitForLaunchFinishResponse, error)
}

// UnimplementedLauncherServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLauncherServiceServer struct {
}

func (UnimplementedLauncherServiceServer) GetLaunches(context.Context, *GetLaunchesRequest) (*GetLaunchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLaunches not implemented")
}
func (UnimplementedLauncherServiceServer) LaunchDatasetUpload(context.Context, *LaunchDatasetUploadRequest) (*LaunchDatasetUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchDatasetUpload not implemented")
}
func (UnimplementedLauncherServiceServer) GetDatasetUploadLaunch(context.Context, *GetDatasetUploadLaunchRequest) (*GetDatasetUploadLaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetUploadLaunch not implemented")
}
func (UnimplementedLauncherServiceServer) LaunchDatasetSetTypes(context.Context, *LaunchDatasetSetTypesRequest) (*LaunchDatasetSetTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchDatasetSetTypes not implemented")
}
func (UnimplementedLauncherServiceServer) GetDatasetSetTypesLaunch(context.Context, *GetDatasetSetTypesLaunchRequest) (*GetDatasetSetTypesLaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetSetTypesLaunch not implemented")
}
func (UnimplementedLauncherServiceServer) LaunchDatasetPreprocess(context.Context, *LaunchDatasetPreprocessRequest) (*LaunchDatasetPreprocessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchDatasetPreprocess not implemented")
}
func (UnimplementedLauncherServiceServer) GetDatasetPreprocessLaunch(context.Context, *GetDatasetPreprocessLaunchRequest) (*GetDatasetPreprocessLaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetPreprocessLaunch not implemented")
}
func (UnimplementedLauncherServiceServer) LaunchTrain(context.Context, *LaunchTrainRequest) (*LaunchTrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchTrain not implemented")
}
func (UnimplementedLauncherServiceServer) GetTrainLaunch(context.Context, *GetTrainLaunchRequest) (*GetTrainLaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrainLaunch not implemented")
}
func (UnimplementedLauncherServiceServer) LaunchPredict(context.Context, *LaunchPredictRequest) (*LaunchPredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchPredict not implemented")
}
func (UnimplementedLauncherServiceServer) GetPredictLaunch(context.Context, *GetPredictLaunchRequest) (*GetPredictLaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredictLaunch not implemented")
}
func (UnimplementedLauncherServiceServer) GetPredictionResultDownloadLink(context.Context, *GetPredictionResultDownloadLinkRequest) (*GetPredictionResultDownloadLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredictionResultDownloadLink not implemented")
}
func (UnimplementedLauncherServiceServer) LaunchGenericConvert(context.Context, *LaunchGenericConvertRequest) (*LaunchGenericConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchGenericConvert not implemented")
}
func (UnimplementedLauncherServiceServer) GetGenericConvertLaunch(context.Context, *GetGenericConvertLaunchRequest) (*GetGenericConvertLaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenericConvertLaunch not implemented")
}
func (UnimplementedLauncherServiceServer) WaitForLaunchFinish(context.Context, *WaitForLaunchFinishRequest) (*WaitForLaunchFinishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitForLaunchFinish not implemented")
}

// UnsafeLauncherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LauncherServiceServer will
// result in compilation errors.
type UnsafeLauncherServiceServer interface {
	mustEmbedUnimplementedLauncherServiceServer()
}

func RegisterLauncherServiceServer(s grpc.ServiceRegistrar, srv LauncherServiceServer) {
	s.RegisterService(&LauncherService_ServiceDesc, srv)
}

func _LauncherService_GetLaunches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLaunchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).GetLaunches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_GetLaunches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).GetLaunches(ctx, req.(*GetLaunchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_LaunchDatasetUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchDatasetUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).LaunchDatasetUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_LaunchDatasetUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).LaunchDatasetUpload(ctx, req.(*LaunchDatasetUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_GetDatasetUploadLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetUploadLaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).GetDatasetUploadLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_GetDatasetUploadLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).GetDatasetUploadLaunch(ctx, req.(*GetDatasetUploadLaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_LaunchDatasetSetTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchDatasetSetTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).LaunchDatasetSetTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_LaunchDatasetSetTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).LaunchDatasetSetTypes(ctx, req.(*LaunchDatasetSetTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_GetDatasetSetTypesLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetSetTypesLaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).GetDatasetSetTypesLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_GetDatasetSetTypesLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).GetDatasetSetTypesLaunch(ctx, req.(*GetDatasetSetTypesLaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_LaunchDatasetPreprocess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchDatasetPreprocessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).LaunchDatasetPreprocess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_LaunchDatasetPreprocess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).LaunchDatasetPreprocess(ctx, req.(*LaunchDatasetPreprocessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_GetDatasetPreprocessLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetPreprocessLaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).GetDatasetPreprocessLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_GetDatasetPreprocessLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).GetDatasetPreprocessLaunch(ctx, req.(*GetDatasetPreprocessLaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_LaunchTrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchTrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).LaunchTrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_LaunchTrain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).LaunchTrain(ctx, req.(*LaunchTrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_GetTrainLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrainLaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).GetTrainLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_GetTrainLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).GetTrainLaunch(ctx, req.(*GetTrainLaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_LaunchPredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchPredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).LaunchPredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_LaunchPredict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).LaunchPredict(ctx, req.(*LaunchPredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_GetPredictLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPredictLaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).GetPredictLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_GetPredictLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).GetPredictLaunch(ctx, req.(*GetPredictLaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_GetPredictionResultDownloadLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPredictionResultDownloadLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).GetPredictionResultDownloadLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_GetPredictionResultDownloadLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).GetPredictionResultDownloadLink(ctx, req.(*GetPredictionResultDownloadLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_LaunchGenericConvert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchGenericConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).LaunchGenericConvert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_LaunchGenericConvert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).LaunchGenericConvert(ctx, req.(*LaunchGenericConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_GetGenericConvertLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGenericConvertLaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).GetGenericConvertLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_GetGenericConvertLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).GetGenericConvertLaunch(ctx, req.(*GetGenericConvertLaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LauncherService_WaitForLaunchFinish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitForLaunchFinishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServiceServer).WaitForLaunchFinish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LauncherService_WaitForLaunchFinish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServiceServer).WaitForLaunchFinish(ctx, req.(*WaitForLaunchFinishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LauncherService_ServiceDesc is the grpc.ServiceDesc for LauncherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LauncherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.hse_experiments_platform.launcher.LauncherService",
	HandlerType: (*LauncherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLaunches",
			Handler:    _LauncherService_GetLaunches_Handler,
		},
		{
			MethodName: "LaunchDatasetUpload",
			Handler:    _LauncherService_LaunchDatasetUpload_Handler,
		},
		{
			MethodName: "GetDatasetUploadLaunch",
			Handler:    _LauncherService_GetDatasetUploadLaunch_Handler,
		},
		{
			MethodName: "LaunchDatasetSetTypes",
			Handler:    _LauncherService_LaunchDatasetSetTypes_Handler,
		},
		{
			MethodName: "GetDatasetSetTypesLaunch",
			Handler:    _LauncherService_GetDatasetSetTypesLaunch_Handler,
		},
		{
			MethodName: "LaunchDatasetPreprocess",
			Handler:    _LauncherService_LaunchDatasetPreprocess_Handler,
		},
		{
			MethodName: "GetDatasetPreprocessLaunch",
			Handler:    _LauncherService_GetDatasetPreprocessLaunch_Handler,
		},
		{
			MethodName: "LaunchTrain",
			Handler:    _LauncherService_LaunchTrain_Handler,
		},
		{
			MethodName: "GetTrainLaunch",
			Handler:    _LauncherService_GetTrainLaunch_Handler,
		},
		{
			MethodName: "LaunchPredict",
			Handler:    _LauncherService_LaunchPredict_Handler,
		},
		{
			MethodName: "GetPredictLaunch",
			Handler:    _LauncherService_GetPredictLaunch_Handler,
		},
		{
			MethodName: "GetPredictionResultDownloadLink",
			Handler:    _LauncherService_GetPredictionResultDownloadLink_Handler,
		},
		{
			MethodName: "LaunchGenericConvert",
			Handler:    _LauncherService_LaunchGenericConvert_Handler,
		},
		{
			MethodName: "GetGenericConvertLaunch",
			Handler:    _LauncherService_GetGenericConvertLaunch_Handler,
		},
		{
			MethodName: "WaitForLaunchFinish",
			Handler:    _LauncherService_WaitForLaunchFinish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "launcher/launcher.proto",
}
