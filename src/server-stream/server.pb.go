// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: src/server-stream/server.proto

package server_stream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_server_stream_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_src_server_stream_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_src_server_stream_server_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Image) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frames   []*Image  `protobuf:"bytes,1,rep,name=frames,proto3" json:"frames,omitempty"`
	MetaData *MetaData `protobuf:"bytes,2,opt,name=metaData,proto3" json:"metaData,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_server_stream_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_src_server_stream_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_src_server_stream_server_proto_rawDescGZIP(), []int{1}
}

func (x *Video) GetFrames() []*Image {
	if x != nil {
		return x.Frames
	}
	return nil
}

func (x *Video) GetMetaData() *MetaData {
	if x != nil {
		return x.MetaData
	}
	return nil
}

type MetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId   string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Frequencies []byte `protobuf:"bytes,2,opt,name=frequencies,proto3" json:"frequencies,omitempty"`
}

func (x *MetaData) Reset() {
	*x = MetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_server_stream_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaData) ProtoMessage() {}

func (x *MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_src_server_stream_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaData.ProtoReflect.Descriptor instead.
func (*MetaData) Descriptor() ([]byte, []int) {
	return file_src_server_stream_server_proto_rawDescGZIP(), []int{2}
}

func (x *MetaData) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *MetaData) GetFrequencies() []byte {
	if x != nil {
		return x.Frequencies
	}
	return nil
}

type VideoConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId   string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	ImageFormat string `protobuf:"bytes,2,opt,name=imageFormat,proto3" json:"imageFormat,omitempty"`
}

func (x *VideoConfig) Reset() {
	*x = VideoConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_server_stream_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoConfig) ProtoMessage() {}

func (x *VideoConfig) ProtoReflect() protoreflect.Message {
	mi := &file_src_server_stream_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoConfig.ProtoReflect.Descriptor instead.
func (*VideoConfig) Descriptor() ([]byte, []int) {
	return file_src_server_stream_server_proto_rawDescGZIP(), []int{3}
}

func (x *VideoConfig) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *VideoConfig) GetImageFormat() string {
	if x != nil {
		return x.ImageFormat
	}
	return ""
}

var File_src_server_stream_server_proto protoreflect.FileDescriptor

var file_src_server_stream_server_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55,
	0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4e, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1e,
	0x0a, 0x06, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x25,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x22, 0x4d, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x32, 0x79, 0x0a, 0x0e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x06, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x12, 0x39, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x73,
	0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_server_stream_server_proto_rawDescOnce sync.Once
	file_src_server_stream_server_proto_rawDescData = file_src_server_stream_server_proto_rawDesc
)

func file_src_server_stream_server_proto_rawDescGZIP() []byte {
	file_src_server_stream_server_proto_rawDescOnce.Do(func() {
		file_src_server_stream_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_server_stream_server_proto_rawDescData)
	})
	return file_src_server_stream_server_proto_rawDescData
}

var file_src_server_stream_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_src_server_stream_server_proto_goTypes = []interface{}{
	(*Image)(nil),                 // 0: Image
	(*Video)(nil),                 // 1: Video
	(*MetaData)(nil),              // 2: MetaData
	(*VideoConfig)(nil),           // 3: VideoConfig
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_src_server_stream_server_proto_depIdxs = []int32{
	4, // 0: Image.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: Video.frames:type_name -> Image
	2, // 2: Video.metaData:type_name -> MetaData
	1, // 3: VideoStreaming.upload:input_type -> Video
	5, // 4: VideoStreaming.getUploadConfig:input_type -> google.protobuf.Empty
	5, // 5: VideoStreaming.upload:output_type -> google.protobuf.Empty
	3, // 6: VideoStreaming.getUploadConfig:output_type -> VideoConfig
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_server_stream_server_proto_init() }
func file_src_server_stream_server_proto_init() {
	if File_src_server_stream_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_server_stream_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_src_server_stream_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_src_server_stream_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaData); i {
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
		file_src_server_stream_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoConfig); i {
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
			RawDescriptor: file_src_server_stream_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_server_stream_server_proto_goTypes,
		DependencyIndexes: file_src_server_stream_server_proto_depIdxs,
		MessageInfos:      file_src_server_stream_server_proto_msgTypes,
	}.Build()
	File_src_server_stream_server_proto = out.File
	file_src_server_stream_server_proto_rawDesc = nil
	file_src_server_stream_server_proto_goTypes = nil
	file_src_server_stream_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VideoStreamingClient is the client API for VideoStreaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoStreamingClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (VideoStreaming_UploadClient, error)
	GetUploadConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VideoConfig, error)
}

type videoStreamingClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStreamingClient(cc grpc.ClientConnInterface) VideoStreamingClient {
	return &videoStreamingClient{cc}
}

func (c *videoStreamingClient) Upload(ctx context.Context, opts ...grpc.CallOption) (VideoStreaming_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VideoStreaming_serviceDesc.Streams[0], "/VideoStreaming/upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoStreamingUploadClient{stream}
	return x, nil
}

type VideoStreaming_UploadClient interface {
	Send(*Video) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type videoStreamingUploadClient struct {
	grpc.ClientStream
}

func (x *videoStreamingUploadClient) Send(m *Video) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoStreamingUploadClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoStreamingClient) GetUploadConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VideoConfig, error) {
	out := new(VideoConfig)
	err := c.cc.Invoke(ctx, "/VideoStreaming/getUploadConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoStreamingServer is the server API for VideoStreaming service.
type VideoStreamingServer interface {
	Upload(VideoStreaming_UploadServer) error
	GetUploadConfig(context.Context, *emptypb.Empty) (*VideoConfig, error)
}

// UnimplementedVideoStreamingServer can be embedded to have forward compatible implementations.
type UnimplementedVideoStreamingServer struct {
}

func (*UnimplementedVideoStreamingServer) Upload(VideoStreaming_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedVideoStreamingServer) GetUploadConfig(context.Context, *emptypb.Empty) (*VideoConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadConfig not implemented")
}

func RegisterVideoStreamingServer(s *grpc.Server, srv VideoStreamingServer) {
	s.RegisterService(&_VideoStreaming_serviceDesc, srv)
}

func _VideoStreaming_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoStreamingServer).Upload(&videoStreamingUploadServer{stream})
}

type VideoStreaming_UploadServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Video, error)
	grpc.ServerStream
}

type videoStreamingUploadServer struct {
	grpc.ServerStream
}

func (x *videoStreamingUploadServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoStreamingUploadServer) Recv() (*Video, error) {
	m := new(Video)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VideoStreaming_GetUploadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoStreamingServer).GetUploadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VideoStreaming/GetUploadConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoStreamingServer).GetUploadConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoStreaming_serviceDesc = grpc.ServiceDesc{
	ServiceName: "VideoStreaming",
	HandlerType: (*VideoStreamingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUploadConfig",
			Handler:    _VideoStreaming_GetUploadConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "upload",
			Handler:       _VideoStreaming_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "src/server-stream/server.proto",
}
