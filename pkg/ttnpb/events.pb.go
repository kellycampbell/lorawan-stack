// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/events.proto

package ttnpb

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Event struct {
	// Name of the event. This can be used to find the (localized) event description.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Time at which the event was triggered.
	Time time.Time `protobuf:"bytes,2,opt,name=time,proto3,stdtime" json:"time"`
	// Identifiers of the entity (or entities) involved.
	Identifiers []*EntityIdentifiers `protobuf:"bytes,3,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	// Optional data attached to the event.
	Data *types.Any `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// Correlation IDs can be used to find related events and actions such as API calls.
	CorrelationIDs []string `protobuf:"bytes,5,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	// The origin of the event. Typically the hostname of the server that created it.
	Origin string `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	// Event context, internal use only.
	Context map[string][]byte `protobuf:"bytes,7,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The event will be visible to a caller that has any of these rights.
	Visibility *Rights `protobuf:"bytes,8,opt,name=visibility,proto3" json:"visibility,omitempty"`
	// Details on the authentication provided by the caller that triggered this event.
	Authentication *Event_Authentication `protobuf:"bytes,9,opt,name=authentication,proto3" json:"authentication,omitempty"`
	// The IP address of the caller that triggered this event.
	RemoteIP string `protobuf:"bytes,10,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	// The IP address of the caller that triggered this event.
	UserAgent string `protobuf:"bytes,11,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The unique identifier of the event, assigned on creation.
	UniqueID             string   `protobuf:"bytes,12,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()      { *m = Event{} }
func (*Event) ProtoMessage() {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Event) GetIdentifiers() []*EntityIdentifiers {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

func (m *Event) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

func (m *Event) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Event) GetContext() map[string][]byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Event) GetVisibility() *Rights {
	if m != nil {
		return m.Visibility
	}
	return nil
}

func (m *Event) GetAuthentication() *Event_Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Event) GetRemoteIP() string {
	if m != nil {
		return m.RemoteIP
	}
	return ""
}

func (m *Event) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *Event) GetUniqueID() string {
	if m != nil {
		return m.UniqueID
	}
	return ""
}

type Event_Authentication struct {
	// The type of authentication that was used. This is typically a bearer token.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The type of token that was used. Common types are APIKey, AccessToken and SessionToken.
	TokenType string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	// The ID of the token that was used.
	TokenID              string   `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_Authentication) Reset()      { *m = Event_Authentication{} }
func (*Event_Authentication) ProtoMessage() {}
func (*Event_Authentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{0, 1}
}
func (m *Event_Authentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Authentication.Unmarshal(m, b)
}
func (m *Event_Authentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Authentication.Marshal(b, m, deterministic)
}
func (m *Event_Authentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Authentication.Merge(m, src)
}
func (m *Event_Authentication) XXX_Size() int {
	return xxx_messageInfo_Event_Authentication.Size(m)
}
func (m *Event_Authentication) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Authentication.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Authentication proto.InternalMessageInfo

func (m *Event_Authentication) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event_Authentication) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *Event_Authentication) GetTokenID() string {
	if m != nil {
		return m.TokenID
	}
	return ""
}

type StreamEventsRequest struct {
	Identifiers []*EntityIdentifiers `protobuf:"bytes,1,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	// If greater than zero, this will return historical events, up to this maximum when the stream starts.
	// If used in combination with "after", the limit that is reached first, is used.
	// The availability of historical events depends on server support and retention policy.
	Tail uint32 `protobuf:"varint,2,opt,name=tail,proto3" json:"tail,omitempty"`
	// If not empty, this will return historical events after the given time when the stream starts.
	// If used in combination with "tail", the limit that is reached first, is used.
	// The availability of historical events depends on server support and retention policy.
	After                *time.Time `protobuf:"bytes,3,opt,name=after,proto3,stdtime" json:"after,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamEventsRequest) Reset()      { *m = StreamEventsRequest{} }
func (*StreamEventsRequest) ProtoMessage() {}
func (*StreamEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{1}
}
func (m *StreamEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventsRequest.Unmarshal(m, b)
}
func (m *StreamEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventsRequest.Marshal(b, m, deterministic)
}
func (m *StreamEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventsRequest.Merge(m, src)
}
func (m *StreamEventsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamEventsRequest.Size(m)
}
func (m *StreamEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventsRequest proto.InternalMessageInfo

func (m *StreamEventsRequest) GetIdentifiers() []*EntityIdentifiers {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

func (m *StreamEventsRequest) GetTail() uint32 {
	if m != nil {
		return m.Tail
	}
	return 0
}

func (m *StreamEventsRequest) GetAfter() *time.Time {
	if m != nil {
		return m.After
	}
	return nil
}

type FindRelatedEventsRequest struct {
	CorrelationID        string   `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRelatedEventsRequest) Reset()      { *m = FindRelatedEventsRequest{} }
func (*FindRelatedEventsRequest) ProtoMessage() {}
func (*FindRelatedEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{2}
}
func (m *FindRelatedEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRelatedEventsRequest.Unmarshal(m, b)
}
func (m *FindRelatedEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRelatedEventsRequest.Marshal(b, m, deterministic)
}
func (m *FindRelatedEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRelatedEventsRequest.Merge(m, src)
}
func (m *FindRelatedEventsRequest) XXX_Size() int {
	return xxx_messageInfo_FindRelatedEventsRequest.Size(m)
}
func (m *FindRelatedEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRelatedEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRelatedEventsRequest proto.InternalMessageInfo

func (m *FindRelatedEventsRequest) GetCorrelationID() string {
	if m != nil {
		return m.CorrelationID
	}
	return ""
}

type FindRelatedEventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRelatedEventsResponse) Reset()      { *m = FindRelatedEventsResponse{} }
func (*FindRelatedEventsResponse) ProtoMessage() {}
func (*FindRelatedEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{3}
}
func (m *FindRelatedEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRelatedEventsResponse.Unmarshal(m, b)
}
func (m *FindRelatedEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRelatedEventsResponse.Marshal(b, m, deterministic)
}
func (m *FindRelatedEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRelatedEventsResponse.Merge(m, src)
}
func (m *FindRelatedEventsResponse) XXX_Size() int {
	return xxx_messageInfo_FindRelatedEventsResponse.Size(m)
}
func (m *FindRelatedEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRelatedEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindRelatedEventsResponse proto.InternalMessageInfo

func (m *FindRelatedEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "ttn.lorawan.v3.Event")
	golang_proto.RegisterType((*Event)(nil), "ttn.lorawan.v3.Event")
	proto.RegisterMapType((map[string][]byte)(nil), "ttn.lorawan.v3.Event.ContextEntry")
	golang_proto.RegisterMapType((map[string][]byte)(nil), "ttn.lorawan.v3.Event.ContextEntry")
	proto.RegisterType((*Event_Authentication)(nil), "ttn.lorawan.v3.Event.Authentication")
	golang_proto.RegisterType((*Event_Authentication)(nil), "ttn.lorawan.v3.Event.Authentication")
	proto.RegisterType((*StreamEventsRequest)(nil), "ttn.lorawan.v3.StreamEventsRequest")
	golang_proto.RegisterType((*StreamEventsRequest)(nil), "ttn.lorawan.v3.StreamEventsRequest")
	proto.RegisterType((*FindRelatedEventsRequest)(nil), "ttn.lorawan.v3.FindRelatedEventsRequest")
	golang_proto.RegisterType((*FindRelatedEventsRequest)(nil), "ttn.lorawan.v3.FindRelatedEventsRequest")
	proto.RegisterType((*FindRelatedEventsResponse)(nil), "ttn.lorawan.v3.FindRelatedEventsResponse")
	golang_proto.RegisterType((*FindRelatedEventsResponse)(nil), "ttn.lorawan.v3.FindRelatedEventsResponse")
}

func init() { proto.RegisterFile("lorawan-stack/api/events.proto", fileDescriptor_4fd8551d68f51e44) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/events.proto", fileDescriptor_4fd8551d68f51e44)
}

var fileDescriptor_4fd8551d68f51e44 = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0xef, 0xe4, 0x8f, 0x93, 0x4c, 0xd2, 0x6c, 0x19, 0x96, 0xc5, 0x6b, 0x81, 0x53, 0xbc, 0x08,
	0x65, 0x91, 0x62, 0xa3, 0x56, 0x5a, 0xa1, 0x8a, 0x03, 0x49, 0x5b, 0x90, 0x11, 0x07, 0x34, 0x74,
	0x2f, 0x2b, 0xa1, 0x68, 0x12, 0x4f, 0x9d, 0x51, 0x92, 0x19, 0xaf, 0x3d, 0xc9, 0xae, 0xc5, 0x05,
	0xf1, 0x05, 0x58, 0xc1, 0x17, 0xe0, 0x84, 0x38, 0xf3, 0x09, 0x38, 0x72, 0xe7, 0xc2, 0xa9, 0xb0,
	0x29, 0x07, 0x8e, 0x9c, 0x7b, 0x42, 0x9e, 0x71, 0xb6, 0x49, 0x1a, 0x10, 0xda, 0xdb, 0x9b, 0xf7,
	0x7e, 0xef, 0xbd, 0xdf, 0xfb, 0xf9, 0x3d, 0x43, 0x7b, 0x22, 0x62, 0xf2, 0x84, 0xf0, 0x4e, 0x22,
	0xc9, 0x70, 0xec, 0x91, 0x88, 0x79, 0x74, 0x4e, 0xb9, 0x4c, 0xdc, 0x28, 0x16, 0x52, 0xa0, 0xa6,
	0x94, 0xdc, 0xcd, 0x31, 0xee, 0xfc, 0xd0, 0xea, 0x86, 0x4c, 0x8e, 0x66, 0x03, 0x77, 0x28, 0xa6,
	0x1e, 0xe5, 0x73, 0x91, 0x46, 0xb1, 0x78, 0x9a, 0x7a, 0x0a, 0x3c, 0xec, 0x84, 0x94, 0x77, 0xe6,
	0x64, 0xc2, 0x02, 0x22, 0xa9, 0x77, 0xc3, 0xd0, 0x25, 0xad, 0xce, 0x4a, 0x89, 0x50, 0x84, 0x42,
	0x27, 0x0f, 0x66, 0xe7, 0xea, 0xa5, 0x1e, 0xca, 0xca, 0xe1, 0x6f, 0x84, 0x42, 0x84, 0x13, 0xaa,
	0xa8, 0x11, 0xce, 0x85, 0x24, 0x92, 0x09, 0x9e, 0xf3, 0xb3, 0xee, 0xe6, 0xd1, 0x17, 0x35, 0x08,
	0x4f, 0xf3, 0x50, 0x6b, 0x33, 0x24, 0xd9, 0x94, 0x26, 0x92, 0x4c, 0xa3, 0x1c, 0x70, 0xef, 0xe6,
	0xec, 0x2c, 0xa0, 0x5c, 0xb2, 0x73, 0x46, 0xe3, 0x65, 0x83, 0x2d, 0x02, 0xc5, 0x2c, 0x1c, 0x2d,
	0x05, 0x72, 0xbe, 0x31, 0x60, 0xf9, 0x34, 0x53, 0x0c, 0x21, 0x58, 0xe2, 0x64, 0x4a, 0x4d, 0xb0,
	0x0f, 0xda, 0x35, 0xac, 0x6c, 0xf4, 0x21, 0x2c, 0x65, 0x5d, 0xcd, 0xc2, 0x3e, 0x68, 0xd7, 0x0f,
	0x2c, 0x57, 0x53, 0x72, 0x97, 0x94, 0xdc, 0xb3, 0x25, 0xa5, 0xde, 0xde, 0x55, 0xaf, 0xfc, 0x13,
	0x28, 0x54, 0xc1, 0x2f, 0x17, 0xad, 0x9d, 0x67, 0xbf, 0xb7, 0x00, 0x56, 0x99, 0xe8, 0x18, 0xd6,
	0x57, 0x48, 0x99, 0xc5, 0xfd, 0x62, 0xbb, 0x7e, 0xf0, 0x96, 0xbb, 0xfe, 0x59, 0xdc, 0x53, 0x2e,
	0x99, 0x4c, 0xfd, 0x6b, 0x20, 0x5e, 0xcd, 0x42, 0x6d, 0x58, 0x0a, 0x88, 0x24, 0x66, 0x49, 0xd1,
	0xb8, 0x7d, 0x83, 0x46, 0x97, 0xa7, 0x58, 0x21, 0xd0, 0xc7, 0xf0, 0xd6, 0x50, 0xc4, 0x31, 0x9d,
	0x28, 0x95, 0xfb, 0x2c, 0x48, 0xcc, 0xf2, 0x7e, 0xb1, 0x5d, 0xeb, 0xd9, 0x57, 0xbd, 0xda, 0xb7,
	0xc0, 0x70, 0x4a, 0x71, 0xc1, 0x0c, 0x16, 0x17, 0xad, 0xe6, 0xf1, 0x35, 0xcc, 0x3f, 0x49, 0x70,
	0x73, 0x25, 0xcd, 0x0f, 0x12, 0x74, 0x07, 0x1a, 0x22, 0x66, 0x21, 0xe3, 0xa6, 0xa1, 0xf4, 0xc8,
	0x5f, 0xe8, 0x03, 0x58, 0x19, 0x0a, 0x2e, 0xe9, 0x53, 0x69, 0x56, 0xd4, 0x2c, 0xce, 0x8d, 0x59,
	0x32, 0x35, 0xdd, 0x63, 0x0d, 0x3a, 0xe5, 0x32, 0x4e, 0xf1, 0x32, 0x05, 0x3d, 0x80, 0x70, 0xce,
	0x12, 0x36, 0x60, 0x13, 0x26, 0x53, 0xb3, 0xaa, 0xc6, 0xb9, 0xb3, 0x59, 0x00, 0xab, 0xef, 0x83,
	0x57, 0x90, 0xe8, 0x53, 0xd8, 0x24, 0x33, 0x39, 0xca, 0x14, 0x19, 0x2a, 0x8a, 0x66, 0x4d, 0xe5,
	0xbe, 0xbd, 0xbd, 0x79, 0x77, 0x0d, 0x8b, 0x37, 0x72, 0xd1, 0x7d, 0x58, 0x8b, 0xe9, 0x54, 0x48,
	0xda, 0x67, 0x91, 0x09, 0xb3, 0xf1, 0x7a, 0x8d, 0xc5, 0x45, 0xab, 0x8a, 0x95, 0xd3, 0xff, 0x0c,
	0x57, 0x75, 0xd8, 0x8f, 0xd0, 0x9b, 0x10, 0xce, 0x12, 0x1a, 0xf7, 0x49, 0x48, 0xb9, 0x34, 0xeb,
	0x4a, 0x8a, 0x5a, 0xe6, 0xe9, 0x66, 0x8e, 0xac, 0xd2, 0x8c, 0xb3, 0xc7, 0x33, 0xda, 0x67, 0x81,
	0xd9, 0xb8, 0xae, 0xf4, 0x50, 0x39, 0xfd, 0x13, 0x5c, 0xd5, 0x61, 0x3f, 0xb0, 0x8e, 0x60, 0x63,
	0x55, 0x13, 0xb4, 0x07, 0x8b, 0x63, 0x9a, 0xe6, 0xdb, 0x96, 0x99, 0xe8, 0x36, 0x2c, 0xcf, 0xc9,
	0x64, 0xa6, 0xb7, 0xad, 0x81, 0xf5, 0xe3, 0xa8, 0xf0, 0x3e, 0xb0, 0xc6, 0xb0, 0xb9, 0x3e, 0x52,
	0xb6, 0xac, 0x32, 0x8d, 0x5e, 0x2c, 0x6b, 0x66, 0x67, 0x5c, 0xa5, 0x18, 0x53, 0xde, 0x57, 0x91,
	0x82, 0xe6, 0xaa, 0x3c, 0x67, 0x59, 0xf8, 0x1d, 0x58, 0xd5, 0x61, 0x16, 0x98, 0x45, 0x45, 0xb5,
	0xbe, 0xb8, 0x68, 0x55, 0xce, 0x32, 0x9f, 0x7f, 0x82, 0x2b, 0x2a, 0xe8, 0x07, 0xce, 0x0f, 0x00,
	0xbe, 0xfa, 0xb9, 0x8c, 0x29, 0x99, 0x2a, 0x31, 0x13, 0x4c, 0x1f, 0xcf, 0x68, 0x22, 0x37, 0x37,
	0x19, 0xbc, 0xd4, 0x26, 0x67, 0xbc, 0x09, 0x9b, 0x28, 0x76, 0xbb, 0x58, 0xd9, 0xe8, 0x01, 0x2c,
	0x93, 0x73, 0x49, 0x63, 0xc5, 0xea, 0xbf, 0xaf, 0xac, 0xa4, 0x2e, 0x4b, 0xc3, 0x9d, 0x2f, 0xa0,
	0xf9, 0x11, 0xe3, 0x01, 0xce, 0xb6, 0x96, 0x06, 0xeb, 0x64, 0xbb, 0xb0, 0xb9, 0x7e, 0x07, 0x5a,
	0xa9, 0x9e, 0x75, 0xd5, 0x33, 0xe2, 0xd2, 0x1e, 0x50, 0x37, 0xb0, 0xbb, 0x76, 0x03, 0x78, 0x77,
	0xed, 0x04, 0x9c, 0x4f, 0xe0, 0xdd, 0x2d, 0xe5, 0x93, 0x48, 0xf0, 0x84, 0xa2, 0x0e, 0x34, 0xf4,
	0x7f, 0x36, 0xd7, 0xe1, 0xb5, 0xad, 0x8b, 0x88, 0x73, 0xd0, 0xc1, 0x73, 0x00, 0x0d, 0x5d, 0x01,
	0x3d, 0x82, 0x86, 0x56, 0x17, 0xdd, 0xdb, 0xcc, 0xd9, 0xa2, 0xba, 0xb5, 0xbd, 0xb0, 0x83, 0xbe,
	0xfe, 0xf5, 0xcf, 0xef, 0x0a, 0x0d, 0xa7, 0x92, 0xff, 0xee, 0x8f, 0xc0, 0xbb, 0xef, 0x01, 0xf4,
	0x25, 0xac, 0xaf, 0x50, 0x46, 0xed, 0xcd, 0xdc, 0x7f, 0x93, 0xcb, 0xba, 0xff, 0x3f, 0x90, 0x7a,
	0x72, 0xe7, 0x75, 0xd5, 0xf9, 0x15, 0x74, 0x2b, 0xef, 0xec, 0xc5, 0x1a, 0xd6, 0x7b, 0xf8, 0xdb,
	0x73, 0x7b, 0xe7, 0xab, 0x85, 0x0d, 0x7e, 0x5c, 0xd8, 0xe0, 0x8f, 0x85, 0x0d, 0xfe, 0x5a, 0xd8,
	0x3b, 0x7f, 0x2f, 0x6c, 0xf0, 0xec, 0xd2, 0xde, 0xf9, 0xfe, 0xd2, 0xde, 0xf9, 0xf9, 0xd2, 0x06,
	0x8f, 0xbc, 0x50, 0xb8, 0x72, 0x44, 0xe5, 0x88, 0xf1, 0x30, 0x71, 0x39, 0x95, 0x4f, 0x44, 0x3c,
	0xf6, 0xd6, 0x7f, 0xd2, 0xf3, 0x43, 0x2f, 0x1a, 0x87, 0x9e, 0x94, 0x3c, 0x1a, 0x0c, 0x0c, 0xb5,
	0x06, 0x87, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x47, 0x1a, 0x13, 0xea, 0x06, 0x00, 0x00,
}

func (this *Event) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Event)
	if !ok {
		that2, ok := that.(Event)
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
	if this.Name != that1.Name {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if len(this.Identifiers) != len(that1.Identifiers) {
		return false
	}
	for i := range this.Identifiers {
		if !this.Identifiers[i].Equal(that1.Identifiers[i]) {
			return false
		}
	}
	if !this.Data.Equal(that1.Data) {
		return false
	}
	if len(this.CorrelationIDs) != len(that1.CorrelationIDs) {
		return false
	}
	for i := range this.CorrelationIDs {
		if this.CorrelationIDs[i] != that1.CorrelationIDs[i] {
			return false
		}
	}
	if this.Origin != that1.Origin {
		return false
	}
	if len(this.Context) != len(that1.Context) {
		return false
	}
	for i := range this.Context {
		if !bytes.Equal(this.Context[i], that1.Context[i]) {
			return false
		}
	}
	if !this.Visibility.Equal(that1.Visibility) {
		return false
	}
	if !this.Authentication.Equal(that1.Authentication) {
		return false
	}
	if this.RemoteIP != that1.RemoteIP {
		return false
	}
	if this.UserAgent != that1.UserAgent {
		return false
	}
	if this.UniqueID != that1.UniqueID {
		return false
	}
	return true
}
func (this *Event_Authentication) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Event_Authentication)
	if !ok {
		that2, ok := that.(Event_Authentication)
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
	if this.Type != that1.Type {
		return false
	}
	if this.TokenType != that1.TokenType {
		return false
	}
	if this.TokenID != that1.TokenID {
		return false
	}
	return true
}
func (this *StreamEventsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StreamEventsRequest)
	if !ok {
		that2, ok := that.(StreamEventsRequest)
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
	if len(this.Identifiers) != len(that1.Identifiers) {
		return false
	}
	for i := range this.Identifiers {
		if !this.Identifiers[i].Equal(that1.Identifiers[i]) {
			return false
		}
	}
	if this.Tail != that1.Tail {
		return false
	}
	if that1.After == nil {
		if this.After != nil {
			return false
		}
	} else if !this.After.Equal(*that1.After) {
		return false
	}
	return true
}
func (this *FindRelatedEventsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FindRelatedEventsRequest)
	if !ok {
		that2, ok := that.(FindRelatedEventsRequest)
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
	if this.CorrelationID != that1.CorrelationID {
		return false
	}
	return true
}
func (this *FindRelatedEventsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FindRelatedEventsResponse)
	if !ok {
		that2, ok := that.(FindRelatedEventsResponse)
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
	if len(this.Events) != len(that1.Events) {
		return false
	}
	for i := range this.Events {
		if !this.Events[i].Equal(that1.Events[i]) {
			return false
		}
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	// Stream live events, optionally with a tail of historical events (depending on server support and retention policy).
	// Events may arrive out-of-order.
	Stream(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (Events_StreamClient, error)
	FindRelated(ctx context.Context, in *FindRelatedEventsRequest, opts ...grpc.CallOption) (*FindRelatedEventsResponse, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Stream(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (Events_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Events_serviceDesc.Streams[0], "/ttn.lorawan.v3.Events/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Events_StreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventsStreamClient struct {
	grpc.ClientStream
}

func (x *eventsStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventsClient) FindRelated(ctx context.Context, in *FindRelatedEventsRequest, opts ...grpc.CallOption) (*FindRelatedEventsResponse, error) {
	out := new(FindRelatedEventsResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Events/FindRelated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	// Stream live events, optionally with a tail of historical events (depending on server support and retention policy).
	// Events may arrive out-of-order.
	Stream(*StreamEventsRequest, Events_StreamServer) error
	FindRelated(context.Context, *FindRelatedEventsRequest) (*FindRelatedEventsResponse, error)
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) Stream(req *StreamEventsRequest, srv Events_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedEventsServer) FindRelated(ctx context.Context, req *FindRelatedEventsRequest) (*FindRelatedEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRelated not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServer).Stream(m, &eventsStreamServer{stream})
}

type Events_StreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventsStreamServer struct {
	grpc.ServerStream
}

func (x *eventsStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Events_FindRelated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRelatedEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).FindRelated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Events/FindRelated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).FindRelated(ctx, req.(*FindRelatedEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindRelated",
			Handler:    _Events_FindRelated_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Events_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/events.proto",
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovEvents(uint64(l))
	if len(m.Identifiers) > 0 {
		for _, e := range m.Identifiers {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Context) > 0 {
		for k, v := range m.Context {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovEvents(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovEvents(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovEvents(uint64(mapEntrySize))
		}
	}
	if m.Visibility != nil {
		l = m.Visibility.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Authentication != nil {
		l = m.Authentication.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RemoteIP)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.UserAgent)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.UniqueID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *Event_Authentication) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.TokenType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.TokenID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *StreamEventsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Identifiers) > 0 {
		for _, e := range m.Identifiers {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if m.Tail != 0 {
		n += 1 + sovEvents(uint64(m.Tail))
	}
	if m.After != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.After)
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *FindRelatedEventsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CorrelationID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *FindRelatedEventsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForIdentifiers := "[]*EntityIdentifiers{"
	for _, f := range this.Identifiers {
		repeatedStringForIdentifiers += strings.Replace(fmt.Sprintf("%v", f), "EntityIdentifiers", "EntityIdentifiers", 1) + ","
	}
	repeatedStringForIdentifiers += "}"
	keysForContext := make([]string, 0, len(this.Context))
	for k := range this.Context {
		keysForContext = append(keysForContext, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForContext)
	mapStringForContext := "map[string][]byte{"
	for _, k := range keysForContext {
		mapStringForContext += fmt.Sprintf("%v: %v,", k, this.Context[k])
	}
	mapStringForContext += "}"
	s := strings.Join([]string{`&Event{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Time:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Time), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`Identifiers:` + repeatedStringForIdentifiers + `,`,
		`Data:` + strings.Replace(fmt.Sprintf("%v", this.Data), "Any", "types.Any", 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`Origin:` + fmt.Sprintf("%v", this.Origin) + `,`,
		`Context:` + mapStringForContext + `,`,
		`Visibility:` + strings.Replace(fmt.Sprintf("%v", this.Visibility), "Rights", "Rights", 1) + `,`,
		`Authentication:` + strings.Replace(fmt.Sprintf("%v", this.Authentication), "Event_Authentication", "Event_Authentication", 1) + `,`,
		`RemoteIP:` + fmt.Sprintf("%v", this.RemoteIP) + `,`,
		`UserAgent:` + fmt.Sprintf("%v", this.UserAgent) + `,`,
		`UniqueID:` + fmt.Sprintf("%v", this.UniqueID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Event_Authentication) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event_Authentication{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`TokenType:` + fmt.Sprintf("%v", this.TokenType) + `,`,
		`TokenID:` + fmt.Sprintf("%v", this.TokenID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StreamEventsRequest) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForIdentifiers := "[]*EntityIdentifiers{"
	for _, f := range this.Identifiers {
		repeatedStringForIdentifiers += strings.Replace(fmt.Sprintf("%v", f), "EntityIdentifiers", "EntityIdentifiers", 1) + ","
	}
	repeatedStringForIdentifiers += "}"
	s := strings.Join([]string{`&StreamEventsRequest{`,
		`Identifiers:` + repeatedStringForIdentifiers + `,`,
		`Tail:` + fmt.Sprintf("%v", this.Tail) + `,`,
		`After:` + strings.Replace(fmt.Sprintf("%v", this.After), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FindRelatedEventsRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FindRelatedEventsRequest{`,
		`CorrelationID:` + fmt.Sprintf("%v", this.CorrelationID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FindRelatedEventsResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForEvents := "[]*Event{"
	for _, f := range this.Events {
		repeatedStringForEvents += strings.Replace(f.String(), "Event", "Event", 1) + ","
	}
	repeatedStringForEvents += "}"
	s := strings.Join([]string{`&FindRelatedEventsResponse{`,
		`Events:` + repeatedStringForEvents + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEvents(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
