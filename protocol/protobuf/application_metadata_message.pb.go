// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application_metadata_message.proto

package protobuf

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

type ApplicationMetadataMessage_Type int32

const (
	ApplicationMetadataMessage_UNKNOWN                                 ApplicationMetadataMessage_Type = 0
	ApplicationMetadataMessage_CHAT_MESSAGE                            ApplicationMetadataMessage_Type = 1
	ApplicationMetadataMessage_CONTACT_UPDATE                          ApplicationMetadataMessage_Type = 2
	ApplicationMetadataMessage_MEMBERSHIP_UPDATE_MESSAGE               ApplicationMetadataMessage_Type = 3
	ApplicationMetadataMessage_PAIR_INSTALLATION                       ApplicationMetadataMessage_Type = 4
	ApplicationMetadataMessage_SYNC_INSTALLATION                       ApplicationMetadataMessage_Type = 5
	ApplicationMetadataMessage_REQUEST_ADDRESS_FOR_TRANSACTION         ApplicationMetadataMessage_Type = 6
	ApplicationMetadataMessage_ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION  ApplicationMetadataMessage_Type = 7
	ApplicationMetadataMessage_DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION ApplicationMetadataMessage_Type = 8
	ApplicationMetadataMessage_REQUEST_TRANSACTION                     ApplicationMetadataMessage_Type = 9
	ApplicationMetadataMessage_SEND_TRANSACTION                        ApplicationMetadataMessage_Type = 10
	ApplicationMetadataMessage_DECLINE_REQUEST_TRANSACTION             ApplicationMetadataMessage_Type = 11
	ApplicationMetadataMessage_SYNC_INSTALLATION_CONTACT               ApplicationMetadataMessage_Type = 12
	ApplicationMetadataMessage_SYNC_INSTALLATION_ACCOUNT               ApplicationMetadataMessage_Type = 13
	ApplicationMetadataMessage_SYNC_INSTALLATION_PUBLIC_CHAT           ApplicationMetadataMessage_Type = 14
	ApplicationMetadataMessage_CONTACT_CODE_ADVERTISEMENT              ApplicationMetadataMessage_Type = 15
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REGISTRATION          ApplicationMetadataMessage_Type = 16
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REGISTRATION_RESPONSE ApplicationMetadataMessage_Type = 17
	ApplicationMetadataMessage_PUSH_NOTIFICATION_QUERY                 ApplicationMetadataMessage_Type = 18
	ApplicationMetadataMessage_PUSH_NOTIFICATION_QUERY_RESPONSE        ApplicationMetadataMessage_Type = 19
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REQUEST               ApplicationMetadataMessage_Type = 20
	ApplicationMetadataMessage_PUSH_NOTIFICATION_RESPONSE              ApplicationMetadataMessage_Type = 21
	ApplicationMetadataMessage_EMOJI_REACTION                          ApplicationMetadataMessage_Type = 22
	ApplicationMetadataMessage_EMOJI_REACTION_RETRACTION               ApplicationMetadataMessage_Type = 23
)

var ApplicationMetadataMessage_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CHAT_MESSAGE",
	2:  "CONTACT_UPDATE",
	3:  "MEMBERSHIP_UPDATE_MESSAGE",
	4:  "PAIR_INSTALLATION",
	5:  "SYNC_INSTALLATION",
	6:  "REQUEST_ADDRESS_FOR_TRANSACTION",
	7:  "ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION",
	8:  "DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION",
	9:  "REQUEST_TRANSACTION",
	10: "SEND_TRANSACTION",
	11: "DECLINE_REQUEST_TRANSACTION",
	12: "SYNC_INSTALLATION_CONTACT",
	13: "SYNC_INSTALLATION_ACCOUNT",
	14: "SYNC_INSTALLATION_PUBLIC_CHAT",
	15: "CONTACT_CODE_ADVERTISEMENT",
	16: "PUSH_NOTIFICATION_REGISTRATION",
	17: "PUSH_NOTIFICATION_REGISTRATION_RESPONSE",
	18: "PUSH_NOTIFICATION_QUERY",
	19: "PUSH_NOTIFICATION_QUERY_RESPONSE",
	20: "PUSH_NOTIFICATION_REQUEST",
	21: "PUSH_NOTIFICATION_RESPONSE",
	22: "EMOJI_REACTION",
	23: "EMOJI_REACTION_RETRACTION",
}

var ApplicationMetadataMessage_Type_value = map[string]int32{
	"UNKNOWN":                                 0,
	"CHAT_MESSAGE":                            1,
	"CONTACT_UPDATE":                          2,
	"MEMBERSHIP_UPDATE_MESSAGE":               3,
	"PAIR_INSTALLATION":                       4,
	"SYNC_INSTALLATION":                       5,
	"REQUEST_ADDRESS_FOR_TRANSACTION":         6,
	"ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION":  7,
	"DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION": 8,
	"REQUEST_TRANSACTION":                     9,
	"SEND_TRANSACTION":                        10,
	"DECLINE_REQUEST_TRANSACTION":             11,
	"SYNC_INSTALLATION_CONTACT":               12,
	"SYNC_INSTALLATION_ACCOUNT":               13,
	"SYNC_INSTALLATION_PUBLIC_CHAT":           14,
	"CONTACT_CODE_ADVERTISEMENT":              15,
	"PUSH_NOTIFICATION_REGISTRATION":          16,
	"PUSH_NOTIFICATION_REGISTRATION_RESPONSE": 17,
	"PUSH_NOTIFICATION_QUERY":                 18,
	"PUSH_NOTIFICATION_QUERY_RESPONSE":        19,
	"PUSH_NOTIFICATION_REQUEST":               20,
	"PUSH_NOTIFICATION_RESPONSE":              21,
	"EMOJI_REACTION":                          22,
	"EMOJI_REACTION_RETRACTION":               23,
}

func (x ApplicationMetadataMessage_Type) String() string {
	return proto.EnumName(ApplicationMetadataMessage_Type_name, int32(x))
}

func (ApplicationMetadataMessage_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad09a6406fcf24c7, []int{0, 0}
}

type ApplicationMetadataMessage struct {
	// Signature of the payload field
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// This is the encoded protobuf of the application level message, i.e ChatMessage
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// The type of protobuf message sent
	Type                 ApplicationMetadataMessage_Type `protobuf:"varint,3,opt,name=type,proto3,enum=protobuf.ApplicationMetadataMessage_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ApplicationMetadataMessage) Reset()         { *m = ApplicationMetadataMessage{} }
func (m *ApplicationMetadataMessage) String() string { return proto.CompactTextString(m) }
func (*ApplicationMetadataMessage) ProtoMessage()    {}
func (*ApplicationMetadataMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad09a6406fcf24c7, []int{0}
}

func (m *ApplicationMetadataMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMetadataMessage.Unmarshal(m, b)
}
func (m *ApplicationMetadataMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMetadataMessage.Marshal(b, m, deterministic)
}
func (m *ApplicationMetadataMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMetadataMessage.Merge(m, src)
}
func (m *ApplicationMetadataMessage) XXX_Size() int {
	return xxx_messageInfo_ApplicationMetadataMessage.Size(m)
}
func (m *ApplicationMetadataMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMetadataMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMetadataMessage proto.InternalMessageInfo

func (m *ApplicationMetadataMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *ApplicationMetadataMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ApplicationMetadataMessage) GetType() ApplicationMetadataMessage_Type {
	if m != nil {
		return m.Type
	}
	return ApplicationMetadataMessage_UNKNOWN
}

func init() {
	proto.RegisterEnum("protobuf.ApplicationMetadataMessage_Type", ApplicationMetadataMessage_Type_name, ApplicationMetadataMessage_Type_value)
	proto.RegisterType((*ApplicationMetadataMessage)(nil), "protobuf.ApplicationMetadataMessage")
}

func init() { proto.RegisterFile("application_metadata_message.proto", fileDescriptor_ad09a6406fcf24c7) }

var fileDescriptor_ad09a6406fcf24c7 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0x52, 0xd3, 0x4e,
	0x14, 0xc6, 0xff, 0x85, 0xd2, 0xc2, 0xa1, 0xff, 0xba, 0x1c, 0xc0, 0x56, 0x10, 0xa8, 0xd5, 0x51,
	0xd4, 0x99, 0x5e, 0xe8, 0xb5, 0x17, 0xcb, 0xe6, 0x40, 0xa3, 0xcd, 0x26, 0xec, 0x6e, 0x74, 0xb8,
	0xda, 0x09, 0x12, 0x99, 0xce, 0x00, 0xcd, 0xd0, 0x70, 0xd1, 0x67, 0xf5, 0x29, 0x7c, 0x03, 0x27,
	0x69, 0x6a, 0xa9, 0x2d, 0x72, 0x95, 0xd9, 0xef, 0xfb, 0x9d, 0x73, 0xe6, 0x7c, 0x9b, 0x85, 0x76,
	0x94, 0x24, 0x57, 0xfd, 0xef, 0x51, 0xda, 0x1f, 0xdc, 0xd8, 0xeb, 0x38, 0x8d, 0x2e, 0xa2, 0x34,
	0xb2, 0xd7, 0xf1, 0x70, 0x18, 0x5d, 0xc6, 0x9d, 0xe4, 0x76, 0x90, 0x0e, 0x70, 0x35, 0xff, 0x9c,
	0xdf, 0xfd, 0x68, 0xff, 0xaa, 0xc0, 0x0e, 0x9f, 0x16, 0x78, 0x05, 0xef, 0x8d, 0x71, 0x7c, 0x0e,
	0x6b, 0xc3, 0xfe, 0xe5, 0x4d, 0x94, 0xde, 0xdd, 0xc6, 0xcd, 0x52, 0xab, 0x74, 0x58, 0x53, 0x53,
	0x01, 0x9b, 0x50, 0x4d, 0xa2, 0xd1, 0xd5, 0x20, 0xba, 0x68, 0x2e, 0xe5, 0xde, 0xe4, 0x88, 0x9f,
	0xa0, 0x9c, 0x8e, 0x92, 0xb8, 0xb9, 0xdc, 0x2a, 0x1d, 0xd6, 0x3f, 0xbc, 0xed, 0x4c, 0xe6, 0x75,
	0x1e, 0x9e, 0xd5, 0x31, 0xa3, 0x24, 0x56, 0x79, 0x59, 0xfb, 0xe7, 0x0a, 0x94, 0xb3, 0x23, 0xae,
	0x43, 0x35, 0x94, 0x5f, 0xa4, 0xff, 0x4d, 0xb2, 0xff, 0x90, 0x41, 0x4d, 0x74, 0xb9, 0xb1, 0x1e,
	0x69, 0xcd, 0x4f, 0x88, 0x95, 0x10, 0xa1, 0x2e, 0x7c, 0x69, 0xb8, 0x30, 0x36, 0x0c, 0x1c, 0x6e,
	0x88, 0x2d, 0xe1, 0x1e, 0x3c, 0xf3, 0xc8, 0x3b, 0x22, 0xa5, 0xbb, 0x6e, 0x50, 0xc8, 0x7f, 0x4a,
	0x96, 0x71, 0x1b, 0x36, 0x02, 0xee, 0x2a, 0xeb, 0x4a, 0x6d, 0x78, 0xaf, 0xc7, 0x8d, 0xeb, 0x4b,
	0x56, 0xce, 0x64, 0x7d, 0x26, 0xc5, 0xac, 0xbc, 0x82, 0x2f, 0xe1, 0x40, 0xd1, 0x69, 0x48, 0xda,
	0x58, 0xee, 0x38, 0x8a, 0xb4, 0xb6, 0xc7, 0xbe, 0xb2, 0x46, 0x71, 0xa9, 0xb9, 0xc8, 0xa1, 0x0a,
	0xbe, 0x83, 0xd7, 0x5c, 0x08, 0x0a, 0x8c, 0x7d, 0x8c, 0xad, 0xe2, 0x7b, 0x78, 0xe3, 0x90, 0xe8,
	0xb9, 0x92, 0x1e, 0x85, 0x57, 0xb1, 0x01, 0x9b, 0x13, 0xe8, 0xbe, 0xb1, 0x86, 0x5b, 0xc0, 0x34,
	0x49, 0x67, 0x46, 0x05, 0x3c, 0x80, 0xdd, 0xbf, 0x7b, 0xdf, 0x07, 0xd6, 0xb3, 0x68, 0xe6, 0x96,
	0xb4, 0x45, 0x80, 0xac, 0xb6, 0xd8, 0xe6, 0x42, 0xf8, 0xa1, 0x34, 0xec, 0x7f, 0x7c, 0x01, 0x7b,
	0xf3, 0x76, 0x10, 0x1e, 0xf5, 0x5c, 0x61, 0xb3, 0x7b, 0x61, 0x75, 0xdc, 0x87, 0x9d, 0xc9, 0x7d,
	0x08, 0xdf, 0x21, 0xcb, 0x9d, 0xaf, 0xa4, 0x8c, 0xab, 0xc9, 0x23, 0x69, 0xd8, 0x13, 0x6c, 0xc3,
	0x7e, 0x10, 0xea, 0xae, 0x95, 0xbe, 0x71, 0x8f, 0x5d, 0x31, 0x6e, 0xa1, 0xe8, 0xc4, 0xd5, 0x46,
	0x8d, 0x23, 0x67, 0x59, 0x42, 0xff, 0x66, 0xac, 0x22, 0x1d, 0xf8, 0x52, 0x13, 0xdb, 0xc0, 0x5d,
	0x68, 0xcc, 0xc3, 0xa7, 0x21, 0xa9, 0x33, 0x86, 0xf8, 0x0a, 0x5a, 0x0f, 0x98, 0xd3, 0x16, 0x9b,
	0xd9, 0xd6, 0x8b, 0xe6, 0xe5, 0xf9, 0xb1, 0xad, 0x6c, 0xa5, 0x45, 0x76, 0x51, 0xbe, 0x9d, 0xfd,
	0x82, 0xe4, 0xf9, 0x9f, 0x5d, 0xab, 0xa8, 0xc8, 0xf9, 0x69, 0xd6, 0x72, 0x56, 0xb3, 0x8a, 0x8c,
	0x2a, 0xec, 0xc6, 0x79, 0x25, 0x7f, 0x0d, 0x1f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x79, 0x47,
	0x6f, 0x0d, 0xaa, 0x03, 0x00, 0x00,
}
