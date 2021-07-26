// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/applicationserver_integrations_storage.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetStoredApplicationUpRequest struct {
	// Query upstream messages from all end devices of an application. Cannot be used in conjunction with end_device_ids.
	ApplicationIds *ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,proto3" json:"application_ids,omitempty"`
	// Query upstream messages from a single end device. Cannot be used in conjunction with application_ids.
	EndDeviceIds *EndDeviceIdentifiers `protobuf:"bytes,2,opt,name=end_device_ids,json=endDeviceIds,proto3" json:"end_device_ids,omitempty"`
	// Query upstream messages of a specific type. If not set, then all upstream messages are returned.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Limit number of results.
	Limit *types.UInt32Value `protobuf:"bytes,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Query upstream messages after this timestamp only. Cannot be used in conjunction with last.
	After *types.Timestamp `protobuf:"bytes,5,opt,name=after,proto3" json:"after,omitempty"`
	// Query upstream messages before this timestamp only. Cannot be used in conjunction with last.
	Before *types.Timestamp `protobuf:"bytes,6,opt,name=before,proto3" json:"before,omitempty"`
	// Query uplinks on a specific FPort only.
	FPort *types.UInt32Value `protobuf:"bytes,7,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// Order results.
	Order string `protobuf:"bytes,8,opt,name=order,proto3" json:"order,omitempty"`
	// The names of the upstream message fields that should be returned. See the API reference
	// for allowed field names for each type of upstream message.
	FieldMask *types.FieldMask `protobuf:"bytes,9,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Query upstream messages that have arrived in the last minutes or hours. Cannot be used in conjunction with after and before.
	Last                 *types.Duration `protobuf:"bytes,10,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetStoredApplicationUpRequest) Reset()      { *m = GetStoredApplicationUpRequest{} }
func (*GetStoredApplicationUpRequest) ProtoMessage() {}
func (*GetStoredApplicationUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0e9f52f73d254, []int{0}
}
func (m *GetStoredApplicationUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Unmarshal(m, b)
}
func (m *GetStoredApplicationUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Marshal(b, m, deterministic)
}
func (m *GetStoredApplicationUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoredApplicationUpRequest.Merge(m, src)
}
func (m *GetStoredApplicationUpRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Size(m)
}
func (m *GetStoredApplicationUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoredApplicationUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoredApplicationUpRequest proto.InternalMessageInfo

func (m *GetStoredApplicationUpRequest) GetApplicationIds() *ApplicationIdentifiers {
	if m != nil {
		return m.ApplicationIds
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetEndDeviceIds() *EndDeviceIdentifiers {
	if m != nil {
		return m.EndDeviceIds
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetStoredApplicationUpRequest) GetLimit() *types.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetAfter() *types.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetBefore() *types.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetFPort() *types.UInt32Value {
	if m != nil {
		return m.FPort
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *GetStoredApplicationUpRequest) GetFieldMask() *types.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetLast() *types.Duration {
	if m != nil {
		return m.Last
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStoredApplicationUpRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpRequest")
	golang_proto.RegisterType((*GetStoredApplicationUpRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpRequest")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/applicationserver_integrations_storage.proto", fileDescriptor_6ff0e9f52f73d254)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/applicationserver_integrations_storage.proto", fileDescriptor_6ff0e9f52f73d254)
}

var fileDescriptor_6ff0e9f52f73d254 = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xbf, 0x6f, 0xe4, 0x44,
	0x14, 0xde, 0x09, 0xd9, 0x40, 0x26, 0x21, 0x91, 0x2c, 0x84, 0xcc, 0xea, 0xe2, 0x44, 0x01, 0xa1,
	0x6b, 0xd6, 0x3e, 0xed, 0x36, 0x40, 0x81, 0x74, 0xd1, 0x01, 0x0a, 0x08, 0x81, 0xe6, 0x08, 0xc5,
	0x35, 0xd6, 0xac, 0xe7, 0xd9, 0x19, 0xd6, 0x3b, 0x33, 0x37, 0xf3, 0xec, 0x10, 0x45, 0x91, 0x10,
	0x7f, 0x01, 0x12, 0xff, 0xc0, 0x95, 0x27, 0x51, 0xd0, 0x52, 0x50, 0x50, 0x50, 0xd1, 0x21, 0x84,
	0x44, 0x09, 0x1b, 0x0a, 0x4a, 0xea, 0x13, 0x05, 0x5a, 0xdb, 0x9b, 0xfd, 0x45, 0xee, 0xe8, 0xe6,
	0xbd, 0xf9, 0xbe, 0xcf, 0xcf, 0xf3, 0xde, 0xf7, 0xe8, 0xdb, 0xb9, 0xb6, 0xfc, 0x8c, 0xab, 0xae,
	0x43, 0x9e, 0x0c, 0x23, 0x6e, 0x64, 0xc4, 0x8d, 0xc9, 0x65, 0xc2, 0x51, 0x6a, 0xe5, 0xc0, 0x96,
	0x60, 0x63, 0xa9, 0x10, 0x32, 0x5b, 0x67, 0x62, 0x87, 0xda, 0xf2, 0x0c, 0x42, 0x63, 0x35, 0x6a,
	0x6f, 0x07, 0x51, 0x85, 0x8d, 0x46, 0x58, 0xf6, 0x3b, 0x77, 0x33, 0x89, 0xa7, 0xc5, 0x20, 0x4c,
	0xf4, 0x28, 0x02, 0x55, 0xea, 0x73, 0x63, 0xf5, 0xe7, 0xe7, 0x51, 0x05, 0x4e, 0xba, 0x19, 0xa8,
	0x6e, 0xc9, 0x73, 0x29, 0x38, 0x42, 0xb4, 0x72, 0xa8, 0x25, 0x3b, 0xdd, 0x39, 0x89, 0x4c, 0x67,
	0xba, 0x26, 0x0f, 0x8a, 0xb4, 0x8a, 0xaa, 0xa0, 0x3a, 0x35, 0xf0, 0x5b, 0x99, 0xd6, 0x59, 0x0e,
	0x75, 0xe9, 0x4a, 0x69, 0xac, 0xeb, 0x6c, 0x6e, 0x0f, 0x9a, 0xdb, 0x6b, 0x8d, 0x54, 0x42, 0x2e,
	0xe2, 0x11, 0x77, 0xc3, 0x06, 0xb1, 0xbf, 0x8c, 0x40, 0x39, 0x02, 0x87, 0x7c, 0x64, 0x1a, 0x40,
	0xb0, 0x0c, 0x10, 0x45, 0xfd, 0x16, 0x37, 0xdd, 0x9f, 0x59, 0x6e, 0x0c, 0xd8, 0x69, 0x09, 0xaf,
	0xae, 0x3e, 0xb1, 0x14, 0xa0, 0x50, 0xa6, 0x72, 0x06, 0x3a, 0x58, 0x05, 0x8d, 0xc0, 0x39, 0x9e,
	0x41, 0x83, 0x38, 0xfc, 0xa7, 0x4d, 0xf7, 0xde, 0x03, 0xbc, 0x8f, 0xda, 0x82, 0xb8, 0x3b, 0xeb,
	0xd1, 0x89, 0x61, 0xf0, 0xb0, 0x00, 0x87, 0xde, 0x47, 0x74, 0x77, 0xae, 0x77, 0xb1, 0x14, 0xce,
	0x27, 0x07, 0xe4, 0xf6, 0x56, 0xef, 0xf5, 0x70, 0xb1, 0x4b, 0xe1, 0x1c, 0xfd, 0x78, 0x56, 0x0a,
	0xdb, 0xe1, 0xf3, 0x79, 0xe7, 0xbd, 0x4f, 0x77, 0x40, 0x89, 0x58, 0x40, 0x29, 0x13, 0xa8, 0xf4,
	0xd6, 0x2a, 0xbd, 0xd7, 0x96, 0xf5, 0xde, 0x51, 0xe2, 0x5e, 0x05, 0x9a, 0x57, 0xdb, 0x86, 0x59,
	0xd6, 0x79, 0x3f, 0x12, 0xba, 0x8e, 0xe7, 0x06, 0xfc, 0xe7, 0x0e, 0xc8, 0xed, 0xcd, 0xa3, 0x6f,
	0xc9, 0x93, 0xa3, 0x6f, 0x88, 0x7d, 0x4c, 0x58, 0x8b, 0xed, 0x14, 0x26, 0x97, 0x6a, 0x18, 0x37,
	0x3f, 0xcc, 0xb6, 0x3e, 0xd3, 0x52, 0xc5, 0x3c, 0x49, 0xc0, 0x20, 0xdb, 0x16, 0xfa, 0x4c, 0x55,
	0xd7, 0x3c, 0x19, 0xb2, 0x17, 0xaf, 0x23, 0xb5, 0x18, 0x3a, 0x50, 0xc8, 0x76, 0xaf, 0xc3, 0x94,
	0xcb, 0x1c, 0xc4, 0x5c, 0xe2, 0x61, 0x01, 0x05, 0x08, 0xd6, 0x59, 0x4c, 0xc4, 0x52, 0x4d, 0x87,
	0x4f, 0xb0, 0xdd, 0x5c, 0x37, 0x2f, 0xe7, 0x74, 0x5e, 0x82, 0x60, 0xdb, 0x93, 0xf1, 0x9f, 0xfc,
	0xb9, 0xe0, 0xc8, 0x59, 0x55, 0xbd, 0xd7, 0xa3, 0xed, 0x5c, 0x8e, 0x24, 0xfa, 0xeb, 0xd5, 0x4b,
	0xdc, 0x0a, 0xeb, 0xe6, 0x87, 0xd3, 0xe6, 0x87, 0x27, 0xc7, 0x0a, 0xfb, 0xbd, 0x4f, 0x79, 0x5e,
	0x00, 0xab, 0xa1, 0xde, 0x1d, 0xda, 0xe6, 0x29, 0x82, 0xf5, 0xdb, 0x15, 0xa7, 0xb3, 0xc2, 0xf9,
	0x64, 0x3a, 0x71, 0xac, 0x06, 0x7a, 0x3d, 0xba, 0x31, 0x80, 0x54, 0x5b, 0xf0, 0x37, 0x9e, 0x49,
	0x69, 0x90, 0x5e, 0x9f, 0x6e, 0xa4, 0xb1, 0xd1, 0x16, 0xfd, 0xe7, 0xff, 0x4f, 0x69, 0xe9, 0xc7,
	0xda, 0xa2, 0xf7, 0x06, 0x6d, 0x6b, 0x2b, 0xc0, 0xfa, 0x2f, 0x54, 0x5d, 0x39, 0x7c, 0x72, 0xb4,
	0x6f, 0xf7, 0x58, 0x8b, 0x6d, 0x77, 0x2d, 0x24, 0x20, 0x4b, 0x10, 0x31, 0x47, 0xb6, 0x35, 0x1f,
	0xd4, 0x04, 0xef, 0x4d, 0x4a, 0x67, 0x56, 0xf2, 0x37, 0x6f, 0x28, 0xf3, 0xdd, 0x09, 0xe4, 0x43,
	0xee, 0x86, 0x6c, 0x33, 0x9d, 0x1e, 0xbd, 0x2e, 0x5d, 0xcf, 0xb9, 0x43, 0x9f, 0x56, 0xa4, 0x57,
	0x56, 0x48, 0xf7, 0x1a, 0x7f, 0xb1, 0x0a, 0xf6, 0xd6, 0xfa, 0x77, 0x8f, 0xf6, 0x49, 0xef, 0xd7,
	0x35, 0xfa, 0xd2, 0xc2, 0xd4, 0xdf, 0xaf, 0xf7, 0x90, 0xf7, 0xfd, 0x1a, 0x7d, 0xf9, 0xbf, 0x7d,
	0xe1, 0x75, 0x97, 0xe7, 0xf4, 0xa9, 0xfe, 0xe9, 0xec, 0x3d, 0xc5, 0x26, 0x27, 0xe6, 0xf0, 0x67,
	0xf2, 0xe5, 0x2f, 0x7f, 0x7e, 0xbd, 0xf6, 0x13, 0xf1, 0x2e, 0x22, 0xee, 0x16, 0xd6, 0x64, 0x74,
	0xb1, 0xe8, 0x93, 0x70, 0xc9, 0x87, 0x4b, 0xf1, 0x65, 0x54, 0x43, 0x57, 0x79, 0xd7, 0xc7, 0xcb,
	0xc8, 0xf0, 0x64, 0x38, 0xb1, 0x7f, 0xd4, 0x2c, 0xdc, 0xe8, 0x62, 0x32, 0x80, 0x97, 0x0f, 0x3e,
	0xf0, 0x8e, 0x57, 0x3f, 0xff, 0xac, 0xef, 0xdd, 0x20, 0x76, 0x87, 0x1c, 0x9d, 0xfc, 0xf6, 0x47,
	0xd0, 0xfa, 0x62, 0x1c, 0x90, 0xc7, 0xe3, 0x80, 0xfc, 0x3e, 0x0e, 0xc8, 0x5f, 0xe3, 0xa0, 0xf5,
	0xf7, 0x38, 0x20, 0x5f, 0x5d, 0x05, 0xad, 0x47, 0x57, 0x41, 0xeb, 0x87, 0xab, 0x80, 0x3c, 0x88,
	0x32, 0x1d, 0xe2, 0x29, 0xe0, 0xa9, 0x54, 0x99, 0x0b, 0x15, 0xe0, 0x99, 0xb6, 0xc3, 0x68, 0x71,
	0x67, 0x95, 0xfd, 0xc8, 0x0c, 0xb3, 0x08, 0x51, 0x99, 0xc1, 0x60, 0xa3, 0xea, 0x66, 0xff, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xf8, 0x84, 0x53, 0x60, 0x06, 0x00, 0x00,
}

func (this *GetStoredApplicationUpRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetStoredApplicationUpRequest)
	if !ok {
		that2, ok := that.(GetStoredApplicationUpRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ApplicationIds.Equal(that1.ApplicationIds) {
		return false
	}
	if !this.EndDeviceIds.Equal(that1.EndDeviceIds) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !this.Limit.Equal(that1.Limit) {
		return false
	}
	if !this.After.Equal(that1.After) {
		return false
	}
	if !this.Before.Equal(that1.Before) {
		return false
	}
	if !this.FPort.Equal(that1.FPort) {
		return false
	}
	if this.Order != that1.Order {
		return false
	}
	if !this.FieldMask.Equal(that1.FieldMask) {
		return false
	}
	if !this.Last.Equal(that1.Last) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationUpStorageClient is the client API for ApplicationUpStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationUpStorageClient interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error)
}

type applicationUpStorageClient struct {
	cc *grpc.ClientConn
}

func NewApplicationUpStorageClient(cc *grpc.ClientConn) ApplicationUpStorageClient {
	return &applicationUpStorageClient{cc}
}

func (c *applicationUpStorageClient) GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApplicationUpStorage_serviceDesc.Streams[0], "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUp", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationUpStorageGetStoredApplicationUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationUpStorage_GetStoredApplicationUpClient interface {
	Recv() (*ApplicationUp, error)
	grpc.ClientStream
}

type applicationUpStorageGetStoredApplicationUpClient struct {
	grpc.ClientStream
}

func (x *applicationUpStorageGetStoredApplicationUpClient) Recv() (*ApplicationUp, error) {
	m := new(ApplicationUp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApplicationUpStorageServer is the server API for ApplicationUpStorage service.
type ApplicationUpStorageServer interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(*GetStoredApplicationUpRequest, ApplicationUpStorage_GetStoredApplicationUpServer) error
}

// UnimplementedApplicationUpStorageServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationUpStorageServer struct {
}

func (*UnimplementedApplicationUpStorageServer) GetStoredApplicationUp(req *GetStoredApplicationUpRequest, srv ApplicationUpStorage_GetStoredApplicationUpServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStoredApplicationUp not implemented")
}

func RegisterApplicationUpStorageServer(s *grpc.Server, srv ApplicationUpStorageServer) {
	s.RegisterService(&_ApplicationUpStorage_serviceDesc, srv)
}

func _ApplicationUpStorage_GetStoredApplicationUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStoredApplicationUpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationUpStorageServer).GetStoredApplicationUp(m, &applicationUpStorageGetStoredApplicationUpServer{stream})
}

type ApplicationUpStorage_GetStoredApplicationUpServer interface {
	Send(*ApplicationUp) error
	grpc.ServerStream
}

type applicationUpStorageGetStoredApplicationUpServer struct {
	grpc.ServerStream
}

func (x *applicationUpStorageGetStoredApplicationUpServer) Send(m *ApplicationUp) error {
	return x.ServerStream.SendMsg(m)
}

var _ApplicationUpStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationUpStorage",
	HandlerType: (*ApplicationUpStorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStoredApplicationUp",
			Handler:       _ApplicationUpStorage_GetStoredApplicationUp_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/applicationserver_integrations_storage.proto",
}

func NewPopulatedGetStoredApplicationUpRequest(r randyApplicationserverIntegrationsStorage, easy bool) *GetStoredApplicationUpRequest {
	this := &GetStoredApplicationUpRequest{}
	if r.Intn(5) != 0 {
		this.ApplicationIds = NewPopulatedApplicationIdentifiers(r, easy)
	}
	if r.Intn(5) != 0 {
		this.EndDeviceIds = NewPopulatedEndDeviceIdentifiers(r, easy)
	}
	this.Type = string(randStringApplicationserverIntegrationsStorage(r))
	if r.Intn(5) != 0 {
		this.Limit = types.NewPopulatedUInt32Value(r, easy)
	}
	if r.Intn(5) != 0 {
		this.After = types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Before = types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(5) != 0 {
		this.FPort = types.NewPopulatedUInt32Value(r, easy)
	}
	this.Order = string(randStringApplicationserverIntegrationsStorage(r))
	if r.Intn(5) != 0 {
		this.FieldMask = types.NewPopulatedFieldMask(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Last = types.NewPopulatedDuration(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyApplicationserverIntegrationsStorage interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneApplicationserverIntegrationsStorage(r randyApplicationserverIntegrationsStorage) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringApplicationserverIntegrationsStorage(r randyApplicationserverIntegrationsStorage) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneApplicationserverIntegrationsStorage(r)
	}
	return string(tmps)
}
func randUnrecognizedApplicationserverIntegrationsStorage(r randyApplicationserverIntegrationsStorage, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldApplicationserverIntegrationsStorage(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldApplicationserverIntegrationsStorage(dAtA []byte, r randyApplicationserverIntegrationsStorage, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GetStoredApplicationUpRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApplicationIds != nil {
		l = m.ApplicationIds.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.EndDeviceIds != nil {
		l = m.EndDeviceIds.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.Limit != nil {
		l = m.Limit.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.After != nil {
		l = m.After.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.Before != nil {
		l = m.Before.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.FPort != nil {
		l = m.FPort.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	l = len(m.Order)
	if l > 0 {
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.FieldMask != nil {
		l = m.FieldMask.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.Last != nil {
		l = m.Last.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	return n
}

func sovApplicationserverIntegrationsStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApplicationserverIntegrationsStorage(x uint64) (n int) {
	return sovApplicationserverIntegrationsStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetStoredApplicationUpRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetStoredApplicationUpRequest{`,
		`ApplicationIds:` + strings.Replace(fmt.Sprintf("%v", this.ApplicationIds), "ApplicationIdentifiers", "ApplicationIdentifiers", 1) + `,`,
		`EndDeviceIds:` + strings.Replace(fmt.Sprintf("%v", this.EndDeviceIds), "EndDeviceIdentifiers", "EndDeviceIdentifiers", 1) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Limit:` + strings.Replace(fmt.Sprintf("%v", this.Limit), "UInt32Value", "types.UInt32Value", 1) + `,`,
		`After:` + strings.Replace(fmt.Sprintf("%v", this.After), "Timestamp", "types.Timestamp", 1) + `,`,
		`Before:` + strings.Replace(fmt.Sprintf("%v", this.Before), "Timestamp", "types.Timestamp", 1) + `,`,
		`FPort:` + strings.Replace(fmt.Sprintf("%v", this.FPort), "UInt32Value", "types.UInt32Value", 1) + `,`,
		`Order:` + fmt.Sprintf("%v", this.Order) + `,`,
		`FieldMask:` + strings.Replace(fmt.Sprintf("%v", this.FieldMask), "FieldMask", "types.FieldMask", 1) + `,`,
		`Last:` + strings.Replace(fmt.Sprintf("%v", this.Last), "Duration", "types.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApplicationserverIntegrationsStorage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
