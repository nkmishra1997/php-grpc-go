// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/restaurant.proto

package restaurant

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Restaurant struct {
	RestId               int32    `protobuf:"varint,1,opt,name=rest_id,json=restId,proto3" json:"rest_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rating               float32  `protobuf:"fixed32,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Cuisines             string   `protobuf:"bytes,4,opt,name=cuisines,proto3" json:"cuisines,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Timing               string   `protobuf:"bytes,6,opt,name=timing,proto3" json:"timing,omitempty"`
	Cft                  int32    `protobuf:"varint,7,opt,name=cft,proto3" json:"cft,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Restaurant) Reset()         { *m = Restaurant{} }
func (m *Restaurant) String() string { return proto.CompactTextString(m) }
func (*Restaurant) ProtoMessage()    {}
func (*Restaurant) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d35531f06cd23d0, []int{0}
}

func (m *Restaurant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restaurant.Unmarshal(m, b)
}
func (m *Restaurant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restaurant.Marshal(b, m, deterministic)
}
func (m *Restaurant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restaurant.Merge(m, src)
}
func (m *Restaurant) XXX_Size() int {
	return xxx_messageInfo_Restaurant.Size(m)
}
func (m *Restaurant) XXX_DiscardUnknown() {
	xxx_messageInfo_Restaurant.DiscardUnknown(m)
}

var xxx_messageInfo_Restaurant proto.InternalMessageInfo

func (m *Restaurant) GetRestId() int32 {
	if m != nil {
		return m.RestId
	}
	return 0
}

func (m *Restaurant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Restaurant) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Restaurant) GetCuisines() string {
	if m != nil {
		return m.Cuisines
	}
	return ""
}

func (m *Restaurant) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Restaurant) GetTiming() string {
	if m != nil {
		return m.Timing
	}
	return ""
}

func (m *Restaurant) GetCft() int32 {
	if m != nil {
		return m.Cft
	}
	return 0
}

type RestaurantId struct {
	RestId               int32    `protobuf:"varint,1,opt,name=rest_id,json=restId,proto3" json:"rest_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestaurantId) Reset()         { *m = RestaurantId{} }
func (m *RestaurantId) String() string { return proto.CompactTextString(m) }
func (*RestaurantId) ProtoMessage()    {}
func (*RestaurantId) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d35531f06cd23d0, []int{1}
}

func (m *RestaurantId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestaurantId.Unmarshal(m, b)
}
func (m *RestaurantId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestaurantId.Marshal(b, m, deterministic)
}
func (m *RestaurantId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestaurantId.Merge(m, src)
}
func (m *RestaurantId) XXX_Size() int {
	return xxx_messageInfo_RestaurantId.Size(m)
}
func (m *RestaurantId) XXX_DiscardUnknown() {
	xxx_messageInfo_RestaurantId.DiscardUnknown(m)
}

var xxx_messageInfo_RestaurantId proto.InternalMessageInfo

func (m *RestaurantId) GetRestId() int32 {
	if m != nil {
		return m.RestId
	}
	return 0
}

type Status struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d35531f06cd23d0, []int{2}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d35531f06cd23d0, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Restaurant)(nil), "restaurant.Restaurant")
	proto.RegisterType((*RestaurantId)(nil), "restaurant.RestaurantId")
	proto.RegisterType((*Status)(nil), "restaurant.Status")
	proto.RegisterType((*Empty)(nil), "restaurant.Empty")
}

func init() { proto.RegisterFile("proto/restaurant.proto", fileDescriptor_9d35531f06cd23d0) }

var fileDescriptor_9d35531f06cd23d0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0xed, 0xf6, 0x63, 0xb7, 0x0e, 0x22, 0xdb, 0x39, 0xd4, 0xd0, 0x53, 0xc9, 0xc5, 0x9e, 0xaa,
	0xe8, 0x1f, 0xf0, 0xa3, 0x2a, 0xbd, 0xa6, 0x78, 0x96, 0xb8, 0x19, 0x25, 0xd0, 0x6e, 0x97, 0x24,
	0x2b, 0xf8, 0xb3, 0xfc, 0x5d, 0xfe, 0x09, 0x49, 0x6c, 0xbb, 0x2b, 0x18, 0xf0, 0x36, 0xef, 0xcd,
	0xe3, 0xcd, 0xcb, 0x23, 0x30, 0xae, 0xcc, 0xd6, 0x6d, 0xcf, 0x0d, 0x59, 0x27, 0x6b, 0x23, 0x4b,
	0x37, 0x0f, 0x04, 0x42, 0xc3, 0xf0, 0xcf, 0x04, 0x40, 0x1c, 0x20, 0x9e, 0x42, 0xe6, 0x97, 0xcf,
	0x5a, 0xb1, 0x64, 0x9a, 0xcc, 0x06, 0x22, 0xf5, 0x70, 0xa9, 0x10, 0xa1, 0x5f, 0xca, 0x0d, 0xb1,
	0xee, 0x34, 0x99, 0x1d, 0x89, 0x30, 0xe3, 0x18, 0x52, 0x23, 0x9d, 0x2e, 0xdf, 0x58, 0x6f, 0x9a,
	0xcc, 0xba, 0x62, 0x87, 0x70, 0x02, 0xc3, 0xa2, 0xd6, 0x56, 0x97, 0x64, 0x59, 0x3f, 0xe8, 0x0f,
	0x18, 0x19, 0x64, 0x52, 0x29, 0x43, 0xd6, 0xb2, 0x41, 0x58, 0xed, 0xa1, 0x77, 0x73, 0x7a, 0xe3,
	0xdd, 0xd2, 0xb0, 0xd8, 0x21, 0xcc, 0xa1, 0x57, 0xbc, 0x3a, 0x96, 0x85, 0x38, 0x7e, 0xe4, 0x67,
	0x70, 0xdc, 0x44, 0x5e, 0xaa, 0x68, 0x68, 0xce, 0x21, 0x5d, 0x39, 0xe9, 0xea, 0x70, 0xd6, 0xd6,
	0x45, 0xe1, 0xcf, 0x7a, 0xc9, 0x50, 0xec, 0x21, 0xcf, 0x60, 0x70, 0xbf, 0xa9, 0xdc, 0xc7, 0xe5,
	0x57, 0x17, 0x46, 0x8d, 0xed, 0x8a, 0xcc, 0xbb, 0x2e, 0x08, 0x1f, 0x20, 0xbf, 0x33, 0x24, 0x1d,
	0xb5, 0x4a, 0x1a, 0xcf, 0x5b, 0x95, 0x36, 0xfc, 0x84, 0xfd, 0xcd, 0x2f, 0x15, 0xef, 0xe0, 0x02,
	0x4e, 0x04, 0x49, 0xd5, 0x72, 0x89, 0xaa, 0x27, 0x11, 0x7f, 0xde, 0xc1, 0x6b, 0xc8, 0x9f, 0x2a,
	0xf5, 0xbf, 0x34, 0xd8, 0xe6, 0x7f, 0x6a, 0xe0, 0x1d, 0xbc, 0x85, 0x7c, 0x41, 0x6b, 0xfa, 0xe5,
	0x10, 0x4f, 0x12, 0xf3, 0x18, 0x3d, 0x92, 0xbb, 0x59, 0xaf, 0x1b, 0xad, 0xc5, 0x51, 0x5b, 0x1a,
	0x1a, 0x8d, 0xbf, 0xe3, 0x22, 0x79, 0x49, 0xc3, 0x57, 0xbc, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x93, 0x94, 0x9e, 0x4f, 0xa4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	// Create a new restaurant.
	CreateRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*RestaurantId, error)
	// Read a given restaurant by rest_id
	ReadRestaurant(ctx context.Context, in *RestaurantId, opts ...grpc.CallOption) (*Restaurant, error)
	// Update the details of given Restaurant
	UpdateRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*Status, error)
	// Delete Restaurant based on rest_id
	DeleteRestaurant(ctx context.Context, in *RestaurantId, opts ...grpc.CallOption) (*Status, error)
	// Get all restaurants
	GetAllRestaurants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RestaurantService_GetAllRestaurantsClient, error)
}

type restaurantServiceClient struct {
	cc *grpc.ClientConn
}

func NewRestaurantServiceClient(cc *grpc.ClientConn) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) CreateRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*RestaurantId, error) {
	out := new(RestaurantId)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantService/CreateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) ReadRestaurant(ctx context.Context, in *RestaurantId, opts ...grpc.CallOption) (*Restaurant, error) {
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantService/ReadRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) UpdateRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantService/UpdateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) DeleteRestaurant(ctx context.Context, in *RestaurantId, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantService/DeleteRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetAllRestaurants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RestaurantService_GetAllRestaurantsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RestaurantService_serviceDesc.Streams[0], "/restaurant.RestaurantService/GetAllRestaurants", opts...)
	if err != nil {
		return nil, err
	}
	x := &restaurantServiceGetAllRestaurantsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RestaurantService_GetAllRestaurantsClient interface {
	Recv() (*Restaurant, error)
	grpc.ClientStream
}

type restaurantServiceGetAllRestaurantsClient struct {
	grpc.ClientStream
}

func (x *restaurantServiceGetAllRestaurantsClient) Recv() (*Restaurant, error) {
	m := new(Restaurant)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RestaurantServiceServer is the server API for RestaurantService service.
type RestaurantServiceServer interface {
	// Create a new restaurant.
	CreateRestaurant(context.Context, *Restaurant) (*RestaurantId, error)
	// Read a given restaurant by rest_id
	ReadRestaurant(context.Context, *RestaurantId) (*Restaurant, error)
	// Update the details of given Restaurant
	UpdateRestaurant(context.Context, *Restaurant) (*Status, error)
	// Delete Restaurant based on rest_id
	DeleteRestaurant(context.Context, *RestaurantId) (*Status, error)
	// Get all restaurants
	GetAllRestaurants(*Empty, RestaurantService_GetAllRestaurantsServer) error
}

// UnimplementedRestaurantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRestaurantServiceServer struct {
}

func (*UnimplementedRestaurantServiceServer) CreateRestaurant(ctx context.Context, req *Restaurant) (*RestaurantId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRestaurant not implemented")
}
func (*UnimplementedRestaurantServiceServer) ReadRestaurant(ctx context.Context, req *RestaurantId) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRestaurant not implemented")
}
func (*UnimplementedRestaurantServiceServer) UpdateRestaurant(ctx context.Context, req *Restaurant) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestaurant not implemented")
}
func (*UnimplementedRestaurantServiceServer) DeleteRestaurant(ctx context.Context, req *RestaurantId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRestaurant not implemented")
}
func (*UnimplementedRestaurantServiceServer) GetAllRestaurants(req *Empty, srv RestaurantService_GetAllRestaurantsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllRestaurants not implemented")
}

func RegisterRestaurantServiceServer(s *grpc.Server, srv RestaurantServiceServer) {
	s.RegisterService(&_RestaurantService_serviceDesc, srv)
}

func _RestaurantService_CreateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Restaurant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).CreateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantService/CreateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).CreateRestaurant(ctx, req.(*Restaurant))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_ReadRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).ReadRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantService/ReadRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).ReadRestaurant(ctx, req.(*RestaurantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_UpdateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Restaurant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).UpdateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantService/UpdateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).UpdateRestaurant(ctx, req.(*Restaurant))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_DeleteRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).DeleteRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantService/DeleteRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).DeleteRestaurant(ctx, req.(*RestaurantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetAllRestaurants_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RestaurantServiceServer).GetAllRestaurants(m, &restaurantServiceGetAllRestaurantsServer{stream})
}

type RestaurantService_GetAllRestaurantsServer interface {
	Send(*Restaurant) error
	grpc.ServerStream
}

type restaurantServiceGetAllRestaurantsServer struct {
	grpc.ServerStream
}

func (x *restaurantServiceGetAllRestaurantsServer) Send(m *Restaurant) error {
	return x.ServerStream.SendMsg(m)
}

var _RestaurantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant.RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRestaurant",
			Handler:    _RestaurantService_CreateRestaurant_Handler,
		},
		{
			MethodName: "ReadRestaurant",
			Handler:    _RestaurantService_ReadRestaurant_Handler,
		},
		{
			MethodName: "UpdateRestaurant",
			Handler:    _RestaurantService_UpdateRestaurant_Handler,
		},
		{
			MethodName: "DeleteRestaurant",
			Handler:    _RestaurantService_DeleteRestaurant_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllRestaurants",
			Handler:       _RestaurantService_GetAllRestaurants_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/restaurant.proto",
}