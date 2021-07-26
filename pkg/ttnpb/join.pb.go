// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/join.proto

package ttnpb

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_v3_pkg_types "go.thethings.network/lorawan-stack/v3/pkg/types"
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

type JoinRequest struct {
	RawPayload         []byte                                                  `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	Payload            *Message                                                `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	DevAddr            go_thethings_network_lorawan_stack_v3_pkg_types.DevAddr `protobuf:"bytes,3,opt,name=dev_addr,json=devAddr,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.DevAddr" json:"dev_addr"`
	SelectedMACVersion MACVersion                                              `protobuf:"varint,4,opt,name=selected_mac_version,json=selectedMacVersion,proto3,enum=ttn.lorawan.v3.MACVersion" json:"selected_mac_version,omitempty"`
	NetId              go_thethings_network_lorawan_stack_v3_pkg_types.NetID   `protobuf:"bytes,5,opt,name=net_id,json=netId,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.NetID" json:"net_id"`
	DownlinkSettings   DLSettings                                              `protobuf:"bytes,6,opt,name=downlink_settings,json=downlinkSettings,proto3" json:"downlink_settings"`
	RxDelay            RxDelay                                                 `protobuf:"varint,7,opt,name=rx_delay,json=rxDelay,proto3,enum=ttn.lorawan.v3.RxDelay" json:"rx_delay,omitempty"`
	// Optional CFList.
	CFList         *CFList  `protobuf:"bytes,8,opt,name=cf_list,json=cfList,proto3" json:"cf_list,omitempty"`
	CorrelationIDs []string `protobuf:"bytes,10,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	// Consumed airtime for the transmission of the join request. Calculated by Network Server using the RawPayload size and the transmission settings.
	ConsumedAirtime      *time.Duration `protobuf:"bytes,11,opt,name=consumed_airtime,json=consumedAirtime,proto3,stdduration" json:"consumed_airtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JoinRequest) Reset()      { *m = JoinRequest{} }
func (*JoinRequest) ProtoMessage() {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd69b88666e72e14, []int{0}
}
func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinRequest) GetPayload() *Message {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *JoinRequest) GetSelectedMACVersion() MACVersion {
	if m != nil {
		return m.SelectedMACVersion
	}
	return MAC_UNKNOWN
}

func (m *JoinRequest) GetDownlinkSettings() DLSettings {
	if m != nil {
		return m.DownlinkSettings
	}
	return DLSettings{}
}

func (m *JoinRequest) GetRxDelay() RxDelay {
	if m != nil {
		return m.RxDelay
	}
	return RX_DELAY_0
}

func (m *JoinRequest) GetCFList() *CFList {
	if m != nil {
		return m.CFList
	}
	return nil
}

func (m *JoinRequest) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

func (m *JoinRequest) GetConsumedAirtime() *time.Duration {
	if m != nil {
		return m.ConsumedAirtime
	}
	return nil
}

type JoinResponse struct {
	RawPayload           []byte `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	SessionKeys          `protobuf:"bytes,2,opt,name=session_keys,json=sessionKeys,proto3,embedded=session_keys" json:"session_keys"`
	Lifetime             time.Duration `protobuf:"bytes,3,opt,name=lifetime,proto3,stdduration" json:"lifetime"`
	CorrelationIDs       []string      `protobuf:"bytes,4,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *JoinResponse) Reset()      { *m = JoinResponse{} }
func (*JoinResponse) ProtoMessage() {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd69b88666e72e14, []int{1}
}
func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResponse.Unmarshal(m, b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return xxx_messageInfo_JoinResponse.Size(m)
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinResponse) GetLifetime() time.Duration {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *JoinResponse) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	golang_proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
	golang_proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
}

func init() { proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_dd69b88666e72e14) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_dd69b88666e72e14)
}

var fileDescriptor_dd69b88666e72e14 = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xbf, 0x6f, 0x23, 0x45,
	0x18, 0xf5, 0x3a, 0xfe, 0x39, 0xb6, 0x7c, 0xbe, 0x15, 0xba, 0x5b, 0x0c, 0xda, 0x35, 0xa9, 0x2c,
	0x24, 0xef, 0x8a, 0x8b, 0x10, 0x05, 0xa0, 0x93, 0x37, 0x16, 0xc8, 0x47, 0x12, 0x45, 0x1b, 0x42,
	0x91, 0x66, 0x35, 0xde, 0x19, 0xaf, 0x07, 0xaf, 0x67, 0x96, 0x99, 0xb1, 0x1d, 0xa7, 0x42, 0x94,
	0x54, 0x88, 0x02, 0x51, 0x52, 0xf2, 0x67, 0x50, 0xa6, 0xa4, 0x44, 0x14, 0x06, 0x36, 0x0d, 0x25,
	0x12, 0x9d, 0x2b, 0xe4, 0xfd, 0x81, 0x93, 0x38, 0x05, 0xa1, 0xda, 0xd9, 0xf9, 0xde, 0x7b, 0x7a,
	0xdf, 0x37, 0xef, 0x03, 0x6f, 0x06, 0x8c, 0xc3, 0x05, 0xa4, 0x5d, 0x21, 0xa1, 0x37, 0xb1, 0x60,
	0x48, 0xac, 0xcf, 0x19, 0xa1, 0x66, 0xc8, 0x99, 0x64, 0x6a, 0x43, 0x4a, 0x6a, 0xa6, 0x08, 0x73,
	0x7e, 0xd0, 0xea, 0xf9, 0x44, 0x8e, 0x67, 0x43, 0xd3, 0x63, 0x53, 0x0b, 0xd3, 0x39, 0x5b, 0x86,
	0x9c, 0x5d, 0x2e, 0xad, 0x18, 0xec, 0x75, 0x7d, 0x4c, 0xbb, 0x73, 0x18, 0x10, 0x04, 0x25, 0xb6,
	0x76, 0x0e, 0x89, 0x64, 0xab, 0x7b, 0x4b, 0xc2, 0x67, 0x3e, 0x4b, 0xc8, 0xc3, 0xd9, 0x28, 0xfe,
	0x8b, 0x7f, 0xe2, 0x53, 0x0a, 0xd7, 0x7d, 0xc6, 0xfc, 0x00, 0x6f, 0x51, 0x68, 0xc6, 0xa1, 0x24,
	0x2c, 0x75, 0xd8, 0x7a, 0xc0, 0xff, 0x04, 0x2f, 0x45, 0x5a, 0x35, 0x76, 0xab, 0x59, 0x37, 0x31,
	0x60, 0xff, 0xef, 0x22, 0xa8, 0xbd, 0x62, 0x84, 0x3a, 0xf8, 0x8b, 0x19, 0x16, 0x52, 0xed, 0x80,
	0x1a, 0x87, 0x0b, 0x37, 0x84, 0xcb, 0x80, 0x41, 0xa4, 0x29, 0x6d, 0xa5, 0x53, 0xb7, 0xcb, 0x6b,
	0xbb, 0x70, 0x95, 0x1f, 0x3f, 0x77, 0x00, 0x87, 0x8b, 0xd3, 0xa4, 0xa4, 0xbe, 0x03, 0xca, 0x19,
	0x2a, 0xdf, 0x56, 0x3a, 0xb5, 0x17, 0xcf, 0xcd, 0xbb, 0xc3, 0x32, 0x8f, 0xb1, 0x10, 0xd0, 0xc7,
	0x4e, 0x86, 0x53, 0x2f, 0x40, 0x05, 0xe1, 0xb9, 0x0b, 0x11, 0xe2, 0xda, 0x5e, 0xac, 0xfc, 0xf2,
	0x7a, 0x65, 0xe4, 0x7e, 0x5d, 0x19, 0xef, 0xf9, 0xcc, 0x94, 0x63, 0x2c, 0xc7, 0x84, 0xfa, 0xc2,
	0xa4, 0x58, 0x2e, 0x18, 0x9f, 0x58, 0x77, 0xcd, 0xcf, 0x0f, 0xac, 0x70, 0xe2, 0x5b, 0x72, 0x19,
	0x62, 0x61, 0xf6, 0xf1, 0xbc, 0x87, 0x10, 0x77, 0xca, 0x28, 0x39, 0xa8, 0x08, 0xbc, 0x26, 0x70,
	0x80, 0x3d, 0x89, 0x91, 0x3b, 0x85, 0x9e, 0x3b, 0xc7, 0x5c, 0x10, 0x46, 0xb5, 0x42, 0x5b, 0xe9,
	0x34, 0x5e, 0xb4, 0x76, 0xbc, 0xf5, 0x0e, 0x3f, 0x4b, 0x10, 0xf6, 0xb3, 0x68, 0x65, 0xa8, 0x67,
	0x29, 0x77, 0x7b, 0xef, 0xa8, 0x99, 0xde, 0x31, 0xf4, 0xd2, 0x3b, 0xf5, 0x53, 0x50, 0xa2, 0x58,
	0xba, 0x04, 0x69, 0xc5, 0xd8, 0xff, 0x87, 0xa9, 0xff, 0x77, 0x1f, 0xeb, 0xff, 0x04, 0xcb, 0x41,
	0xdf, 0x29, 0x52, 0x2c, 0x07, 0x48, 0x3d, 0x07, 0x4f, 0x11, 0x5b, 0xd0, 0x80, 0xd0, 0x89, 0x2b,
	0xb0, 0x94, 0x1b, 0x11, 0xad, 0x14, 0x0f, 0x75, 0xc7, 0x78, 0xff, 0xe8, 0x2c, 0x45, 0xd8, 0xf5,
	0xb5, 0x5d, 0xfc, 0x5a, 0xc9, 0x37, 0x95, 0x8d, 0x09, 0xa7, 0x99, 0x49, 0x64, 0x75, 0xf5, 0x03,
	0x50, 0xe1, 0x97, 0x2e, 0xc2, 0x01, 0x5c, 0x6a, 0xe5, 0x78, 0x0c, 0x3b, 0x4f, 0xe4, 0x5c, 0xf6,
	0x37, 0x65, 0xbb, 0xb2, 0xb6, 0x8b, 0x5f, 0x6d, 0xa4, 0x9c, 0x32, 0x4f, 0xae, 0xd4, 0xf7, 0x41,
	0xd9, 0x1b, 0xb9, 0x01, 0x11, 0x52, 0xab, 0xc4, 0x56, 0x9e, 0xdd, 0x27, 0x1f, 0x7e, 0x74, 0x44,
	0x84, 0xb4, 0x41, 0xb4, 0x32, 0x4a, 0xc9, 0xd9, 0x29, 0x79, 0xa3, 0xcd, 0x57, 0xfd, 0x18, 0x3c,
	0xf1, 0x18, 0xe7, 0x38, 0x88, 0xa3, 0xea, 0x12, 0x24, 0x34, 0xd0, 0xde, 0xeb, 0x54, 0x6d, 0x7d,
	0x6d, 0x57, 0xbf, 0x55, 0x4a, 0xfb, 0x05, 0x9e, 0xd7, 0x50, 0xb4, 0x32, 0x1a, 0x87, 0x5b, 0xd8,
	0xa0, 0x2f, 0x9c, 0xc6, 0x2d, 0xda, 0x00, 0x09, 0xf5, 0x04, 0x34, 0x3d, 0x46, 0xc5, 0x6c, 0x8a,
	0x91, 0x0b, 0x09, 0x97, 0x64, 0x8a, 0xb5, 0x5a, 0x6c, 0xe7, 0x75, 0x33, 0xd9, 0x0c, 0x33, 0xdb,
	0x0c, 0xb3, 0x9f, 0x6e, 0x86, 0x5d, 0xb9, 0x5e, 0x19, 0xca, 0xf7, 0xbf, 0x19, 0x8a, 0xf3, 0x24,
	0x23, 0xf7, 0x12, 0xee, 0xab, 0x42, 0xa5, 0xda, 0x04, 0xfb, 0xdf, 0xe5, 0x41, 0x3d, 0x49, 0xbd,
	0x08, 0x19, 0x15, 0x58, 0x7d, 0xfb, 0xa1, 0xd8, 0x57, 0xd7, 0x76, 0xe9, 0xaa, 0xd0, 0x7c, 0xaa,
	0xbd, 0x75, 0x27, 0xf8, 0xa7, 0xa0, 0x2e, 0xb0, 0xd8, 0xc4, 0xc1, 0xdd, 0x6c, 0x5a, 0x9a, 0xfe,
	0x37, 0xee, 0x4f, 0xe7, 0x2c, 0xc1, 0x7c, 0x82, 0x97, 0xc2, 0x6e, 0xde, 0x7e, 0xa9, 0x9f, 0x57,
	0x86, 0xe2, 0xd4, 0xc4, 0xb6, 0xac, 0xbe, 0x04, 0x95, 0x80, 0x8c, 0x70, 0xdc, 0xdc, 0xde, 0x7f,
	0x69, 0x2e, 0x17, 0x37, 0xf7, 0x2f, 0xe9, 0xa1, 0x71, 0x17, 0xfe, 0xcf, 0xb8, 0xed, 0xf3, 0x5f,
	0xfe, 0xd0, 0x73, 0x5f, 0x46, 0xba, 0xf2, 0x63, 0xa4, 0x2b, 0xbf, 0x47, 0xba, 0xf2, 0x67, 0xa4,
	0xe7, 0xfe, 0x8a, 0x74, 0xe5, 0x9b, 0x1b, 0x3d, 0xf7, 0xc3, 0x8d, 0x9e, 0xfb, 0xe9, 0x46, 0x57,
	0x2e, 0xac, 0x47, 0xa4, 0x5d, 0xd2, 0x70, 0x38, 0x2c, 0xc5, 0x6d, 0x1c, 0xfc, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x59, 0x5f, 0x7f, 0x6d, 0x05, 0x00, 0x00,
}

func (this *JoinRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JoinRequest)
	if !ok {
		that2, ok := that.(JoinRequest)
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
	if !bytes.Equal(this.RawPayload, that1.RawPayload) {
		return false
	}
	if !this.Payload.Equal(that1.Payload) {
		return false
	}
	if !this.DevAddr.Equal(that1.DevAddr) {
		return false
	}
	if this.SelectedMACVersion != that1.SelectedMACVersion {
		return false
	}
	if !this.NetId.Equal(that1.NetId) {
		return false
	}
	if !this.DownlinkSettings.Equal(&that1.DownlinkSettings) {
		return false
	}
	if this.RxDelay != that1.RxDelay {
		return false
	}
	if !this.CFList.Equal(that1.CFList) {
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
	if this.ConsumedAirtime != nil && that1.ConsumedAirtime != nil {
		if *this.ConsumedAirtime != *that1.ConsumedAirtime {
			return false
		}
	} else if this.ConsumedAirtime != nil {
		return false
	} else if that1.ConsumedAirtime != nil {
		return false
	}
	return true
}
func (this *JoinResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JoinResponse)
	if !ok {
		that2, ok := that.(JoinResponse)
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
	if !bytes.Equal(this.RawPayload, that1.RawPayload) {
		return false
	}
	if !this.SessionKeys.Equal(&that1.SessionKeys) {
		return false
	}
	if this.Lifetime != that1.Lifetime {
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
	return true
}
func (m *JoinRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.DevAddr.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.SelectedMACVersion != 0 {
		n += 1 + sovJoin(uint64(m.SelectedMACVersion))
	}
	l = m.NetId.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = m.DownlinkSettings.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.RxDelay != 0 {
		n += 1 + sovJoin(uint64(m.RxDelay))
	}
	if m.CFList != nil {
		l = m.CFList.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovJoin(uint64(l))
		}
	}
	if m.ConsumedAirtime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ConsumedAirtime)
		n += 1 + l + sovJoin(uint64(l))
	}
	return n
}

func (m *JoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.SessionKeys.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Lifetime)
	n += 1 + l + sovJoin(uint64(l))
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovJoin(uint64(l))
		}
	}
	return n
}

func sovJoin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJoin(x uint64) (n int) {
	return sovJoin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *JoinRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinRequest{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`Payload:` + strings.Replace(fmt.Sprintf("%v", this.Payload), "Message", "Message", 1) + `,`,
		`DevAddr:` + fmt.Sprintf("%v", this.DevAddr) + `,`,
		`SelectedMACVersion:` + fmt.Sprintf("%v", this.SelectedMACVersion) + `,`,
		`NetId:` + fmt.Sprintf("%v", this.NetId) + `,`,
		`DownlinkSettings:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.DownlinkSettings), "DLSettings", "DLSettings", 1), `&`, ``, 1) + `,`,
		`RxDelay:` + fmt.Sprintf("%v", this.RxDelay) + `,`,
		`CFList:` + strings.Replace(fmt.Sprintf("%v", this.CFList), "CFList", "CFList", 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`ConsumedAirtime:` + strings.Replace(fmt.Sprintf("%v", this.ConsumedAirtime), "Duration", "types.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JoinResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinResponse{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`SessionKeys:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.SessionKeys), "SessionKeys", "SessionKeys", 1), `&`, ``, 1) + `,`,
		`Lifetime:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Lifetime), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringJoin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
