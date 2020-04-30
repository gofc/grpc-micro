// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.0
// source: proto/v1/pbinside/payment_service.proto

package pbinside

import (
	context "context"
	pb "github.com/gofc/grpc-micro/proto/v1/pb"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListPaymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPaymentsRequest) Reset() {
	*x = ListPaymentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_pbinside_payment_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentsRequest) ProtoMessage() {}

func (x *ListPaymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_pbinside_payment_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentsRequest.ProtoReflect.Descriptor instead.
func (*ListPaymentsRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_pbinside_payment_service_proto_rawDescGZIP(), []int{0}
}

type ListPaymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payments []*pb.Payment `protobuf:"bytes,1,rep,name=payments,proto3" json:"payments,omitempty"`
}

func (x *ListPaymentsResponse) Reset() {
	*x = ListPaymentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_pbinside_payment_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPaymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentsResponse) ProtoMessage() {}

func (x *ListPaymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_pbinside_payment_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentsResponse.ProtoReflect.Descriptor instead.
func (*ListPaymentsResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_pbinside_payment_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPaymentsResponse) GetPayments() []*pb.Payment {
	if x != nil {
		return x.Payments
	}
	return nil
}

type CreatePaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *pb.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *CreatePaymentRequest) Reset() {
	*x = CreatePaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_pbinside_payment_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentRequest) ProtoMessage() {}

func (x *CreatePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_pbinside_payment_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentRequest.ProtoReflect.Descriptor instead.
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_pbinside_payment_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePaymentRequest) GetPayment() *pb.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type CreatePaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *pb.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *CreatePaymentResponse) Reset() {
	*x = CreatePaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_pbinside_payment_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResponse) ProtoMessage() {}

func (x *CreatePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_pbinside_payment_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResponse.ProtoReflect.Descriptor instead.
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_pbinside_payment_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePaymentResponse) GetPayment() *pb.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

var File_proto_v1_pbinside_payment_service_proto protoreflect.FileDescriptor

var file_proto_v1_pbinside_payment_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x69, 0x6e, 0x73,
	0x69, 0x64, 0x65, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x62, 0x69, 0x6e, 0x73,
	0x69, 0x64, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xdd, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62,
	0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x69,
	0x6e, 0x73, 0x69, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x66, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x66, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x62, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_pbinside_payment_service_proto_rawDescOnce sync.Once
	file_proto_v1_pbinside_payment_service_proto_rawDescData = file_proto_v1_pbinside_payment_service_proto_rawDesc
)

func file_proto_v1_pbinside_payment_service_proto_rawDescGZIP() []byte {
	file_proto_v1_pbinside_payment_service_proto_rawDescOnce.Do(func() {
		file_proto_v1_pbinside_payment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_pbinside_payment_service_proto_rawDescData)
	})
	return file_proto_v1_pbinside_payment_service_proto_rawDescData
}

var file_proto_v1_pbinside_payment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1_pbinside_payment_service_proto_goTypes = []interface{}{
	(*ListPaymentsRequest)(nil),   // 0: pbinside.ListPaymentsRequest
	(*ListPaymentsResponse)(nil),  // 1: pbinside.ListPaymentsResponse
	(*CreatePaymentRequest)(nil),  // 2: pbinside.CreatePaymentRequest
	(*CreatePaymentResponse)(nil), // 3: pbinside.CreatePaymentResponse
	(*pb.Payment)(nil),            // 4: pb.Payment
}
var file_proto_v1_pbinside_payment_service_proto_depIdxs = []int32{
	4, // 0: pbinside.ListPaymentsResponse.payments:type_name -> pb.Payment
	4, // 1: pbinside.CreatePaymentRequest.payment:type_name -> pb.Payment
	4, // 2: pbinside.CreatePaymentResponse.payment:type_name -> pb.Payment
	0, // 3: pbinside.PaymentService.ListPayments:input_type -> pbinside.ListPaymentsRequest
	2, // 4: pbinside.PaymentService.CreatePayment:input_type -> pbinside.CreatePaymentRequest
	1, // 5: pbinside.PaymentService.ListPayments:output_type -> pbinside.ListPaymentsResponse
	3, // 6: pbinside.PaymentService.CreatePayment:output_type -> pbinside.CreatePaymentResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_v1_pbinside_payment_service_proto_init() }
func file_proto_v1_pbinside_payment_service_proto_init() {
	if File_proto_v1_pbinside_payment_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_pbinside_payment_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentsRequest); i {
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
		file_proto_v1_pbinside_payment_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPaymentsResponse); i {
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
		file_proto_v1_pbinside_payment_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentRequest); i {
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
		file_proto_v1_pbinside_payment_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentResponse); i {
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
			RawDescriptor: file_proto_v1_pbinside_payment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_pbinside_payment_service_proto_goTypes,
		DependencyIndexes: file_proto_v1_pbinside_payment_service_proto_depIdxs,
		MessageInfos:      file_proto_v1_pbinside_payment_service_proto_msgTypes,
	}.Build()
	File_proto_v1_pbinside_payment_service_proto = out.File
	file_proto_v1_pbinside_payment_service_proto_rawDesc = nil
	file_proto_v1_pbinside_payment_service_proto_goTypes = nil
	file_proto_v1_pbinside_payment_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentServiceClient interface {
	ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*ListPaymentsResponse, error)
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*ListPaymentsResponse, error) {
	out := new(ListPaymentsResponse)
	err := c.cc.Invoke(ctx, "/pbinside.PaymentService/ListPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, "/pbinside.PaymentService/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
type PaymentServiceServer interface {
	ListPayments(context.Context, *ListPaymentsRequest) (*ListPaymentsResponse, error)
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
}

// UnimplementedPaymentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (*UnimplementedPaymentServiceServer) ListPayments(context.Context, *ListPaymentsRequest) (*ListPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayments not implemented")
}
func (*UnimplementedPaymentServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&_PaymentService_serviceDesc, srv)
}

func _PaymentService_ListPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ListPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbinside.PaymentService/ListPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ListPayments(ctx, req.(*ListPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbinside.PaymentService/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbinside.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPayments",
			Handler:    _PaymentService_ListPayments_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentService_CreatePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/pbinside/payment_service.proto",
}