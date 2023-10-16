// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/pb/healthProtocol.proto

package protocols_messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HealthRequest struct {
	MessageData          *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	Message              string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a24ea57b88489c19, []int{0}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

func (m *HealthRequest) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *HealthRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type HealthResponse struct {
	MessageData          *MessageData `protobuf:"bytes,1,opt,name=messageData,proto3" json:"messageData,omitempty"`
	Message              string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a24ea57b88489c19, []int{1}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *HealthResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthRequest)(nil), "protocols.messages.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "protocols.messages.HealthResponse")
}

func init() { proto.RegisterFile("server/pb/healthProtocol.proto", fileDescriptor_a24ea57b88489c19) }

var fileDescriptor_a24ea57b88489c19 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2f, 0x48, 0xd2, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0x08, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2b, 0x00, 0x31, 0x84, 0x84, 0x0a, 0xa0, 0xfc, 0x62, 0xbd, 0xdc,
	0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0xd4, 0x62, 0x29, 0x09, 0x84, 0x1e, 0x98, 0x18, 0x44, 0xb5, 0x52,
	0x0e, 0x17, 0xaf, 0x07, 0xd8, 0x94, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x47, 0x2e,
	0x6e, 0xa8, 0x12, 0x97, 0xc4, 0x92, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x79, 0x3d,
	0x4c, 0x43, 0xf5, 0x7c, 0x11, 0xca, 0x82, 0x90, 0xf5, 0x08, 0x49, 0x70, 0xb1, 0x43, 0xb9, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x52, 0x2e, 0x17, 0x1f, 0xcc, 0xb6, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x9a, 0x5a, 0x97, 0xc4, 0x06, 0x36, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xe4, 0x53, 0x46, 0xaf, 0x33, 0x01, 0x00, 0x00,
}
