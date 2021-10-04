// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_identity.proto

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

// SourceType are the predefined types of image source allowed
type IdentityImage_SourceType int32

const (
	IdentityImage_UNKNOWN_SOURCE_TYPE IdentityImage_SourceType = 0
	// RAW_PAYLOAD image byte data
	IdentityImage_RAW_PAYLOAD IdentityImage_SourceType = 1
	// ENS_AVATAR uses the ENS record's resolver get-text-data.avatar data
	// The `payload` field will be ignored if ENS_AVATAR is selected
	// The application will read and parse the ENS avatar data as image payload data, URLs will be ignored
	// The parent `ChatMessageIdentity` must have a valid `ens_name` set
	IdentityImage_ENS_AVATAR IdentityImage_SourceType = 2
)

var IdentityImage_SourceType_name = map[int32]string{
	0: "UNKNOWN_SOURCE_TYPE",
	1: "RAW_PAYLOAD",
	2: "ENS_AVATAR",
}

var IdentityImage_SourceType_value = map[string]int32{
	"UNKNOWN_SOURCE_TYPE": 0,
	"RAW_PAYLOAD":         1,
	"ENS_AVATAR":          2,
}

func (x IdentityImage_SourceType) String() string {
	return proto.EnumName(IdentityImage_SourceType_name, int32(x))
}

func (IdentityImage_SourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{1, 0}
}

// ChatIdentity represents the user defined identity associated with their public chat key
type ChatIdentity struct {
	// Lamport timestamp of the message
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// ens_name is the valid ENS name associated with the chat key
	EnsName string `protobuf:"bytes,2,opt,name=ens_name,json=ensName,proto3" json:"ens_name,omitempty"`
	// images is a string indexed mapping of images associated with an identity
	Images map[string]*IdentityImage `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// display name is the user set identity, valid only for organisations
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// description is the user set description, valid only for organisations
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Color                string   `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Emoji                string   `protobuf:"bytes,7,opt,name=emoji,proto3" json:"emoji,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatIdentity) Reset()         { *m = ChatIdentity{} }
func (m *ChatIdentity) String() string { return proto.CompactTextString(m) }
func (*ChatIdentity) ProtoMessage()    {}
func (*ChatIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{0}
}

func (m *ChatIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatIdentity.Unmarshal(m, b)
}
func (m *ChatIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatIdentity.Marshal(b, m, deterministic)
}
func (m *ChatIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatIdentity.Merge(m, src)
}
func (m *ChatIdentity) XXX_Size() int {
	return xxx_messageInfo_ChatIdentity.Size(m)
}
func (m *ChatIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_ChatIdentity proto.InternalMessageInfo

func (m *ChatIdentity) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *ChatIdentity) GetEnsName() string {
	if m != nil {
		return m.EnsName
	}
	return ""
}

func (m *ChatIdentity) GetImages() map[string]*IdentityImage {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *ChatIdentity) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ChatIdentity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ChatIdentity) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *ChatIdentity) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

// ProfileImage represents data associated with a user's profile image
type IdentityImage struct {
	// payload is a context based payload for the profile image data,
	// context is determined by the `source_type`
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// source_type signals the image payload source
	SourceType IdentityImage_SourceType `protobuf:"varint,2,opt,name=source_type,json=sourceType,proto3,enum=protobuf.IdentityImage_SourceType" json:"source_type,omitempty"`
	// image_type signals the image type and method of parsing the payload
	ImageType ImageType `protobuf:"varint,3,opt,name=image_type,json=imageType,proto3,enum=protobuf.ImageType" json:"image_type,omitempty"`
	// encryption_keys is a list of encrypted keys that can be used to decrypted an encrypted payload
	EncryptionKeys [][]byte `protobuf:"bytes,4,rep,name=encryption_keys,json=encryptionKeys,proto3" json:"encryption_keys,omitempty"`
	// encrypted signals the encryption state of the payload, default is false.
	Encrypted            bool     `protobuf:"varint,5,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityImage) Reset()         { *m = IdentityImage{} }
func (m *IdentityImage) String() string { return proto.CompactTextString(m) }
func (*IdentityImage) ProtoMessage()    {}
func (*IdentityImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a652489000a5879, []int{1}
}

func (m *IdentityImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityImage.Unmarshal(m, b)
}
func (m *IdentityImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityImage.Marshal(b, m, deterministic)
}
func (m *IdentityImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityImage.Merge(m, src)
}
func (m *IdentityImage) XXX_Size() int {
	return xxx_messageInfo_IdentityImage.Size(m)
}
func (m *IdentityImage) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityImage.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityImage proto.InternalMessageInfo

func (m *IdentityImage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *IdentityImage) GetSourceType() IdentityImage_SourceType {
	if m != nil {
		return m.SourceType
	}
	return IdentityImage_UNKNOWN_SOURCE_TYPE
}

func (m *IdentityImage) GetImageType() ImageType {
	if m != nil {
		return m.ImageType
	}
	return ImageType_UNKNOWN_IMAGE_TYPE
}

func (m *IdentityImage) GetEncryptionKeys() [][]byte {
	if m != nil {
		return m.EncryptionKeys
	}
	return nil
}

func (m *IdentityImage) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

func init() {
	proto.RegisterEnum("protobuf.IdentityImage_SourceType", IdentityImage_SourceType_name, IdentityImage_SourceType_value)
	proto.RegisterType((*ChatIdentity)(nil), "protobuf.ChatIdentity")
	proto.RegisterMapType((map[string]*IdentityImage)(nil), "protobuf.ChatIdentity.ImagesEntry")
	proto.RegisterType((*IdentityImage)(nil), "protobuf.IdentityImage")
}

func init() {
	proto.RegisterFile("chat_identity.proto", fileDescriptor_7a652489000a5879)
}

var fileDescriptor_7a652489000a5879 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0x9d, 0xbf, 0xe3, 0x34, 0x8d, 0x36, 0x48, 0x35, 0x15, 0x07, 0x93, 0x0b, 0xb9,
	0x60, 0xa4, 0x70, 0x41, 0xe5, 0x64, 0x42, 0x90, 0xaa, 0x22, 0xa7, 0xda, 0xa4, 0x54, 0xe5, 0x62,
	0x6d, 0xed, 0x85, 0x9a, 0xd8, 0x5e, 0xcb, 0xeb, 0x20, 0xed, 0x63, 0xf1, 0x06, 0x3c, 0x1a, 0xca,
	0xac, 0x5d, 0x27, 0xa7, 0x9d, 0xf9, 0xe6, 0x9b, 0xdf, 0xce, 0x0c, 0x4c, 0xa3, 0x27, 0x56, 0x85,
	0x49, 0xcc, 0xf3, 0x2a, 0xa9, 0x94, 0x57, 0x94, 0xa2, 0x12, 0x64, 0x80, 0xcf, 0xe3, 0xfe, 0xe7,
	0xa5, 0xcd, 0xf3, 0x7d, 0x26, 0xb5, 0x3c, 0xfb, 0x67, 0xc2, 0x68, 0xf9, 0xc4, 0xaa, 0xeb, 0xda,
	0x4d, 0x5e, 0x42, 0x37, 0x4a, 0x45, 0xb4, 0x73, 0x0c, 0xd7, 0x98, 0x77, 0xa8, 0x4e, 0xc8, 0x2b,
	0x18, 0xf0, 0x5c, 0x86, 0x39, 0xcb, 0xb8, 0x63, 0xba, 0xc6, 0x7c, 0x48, 0xfb, 0x3c, 0x97, 0x01,
	0xcb, 0x38, 0xb9, 0x82, 0x5e, 0x92, 0xb1, 0x5f, 0x5c, 0x3a, 0x96, 0x6b, 0xcd, 0xed, 0xc5, 0xcc,
	0x6b, 0x7e, 0xf2, 0x8e, 0xc1, 0xde, 0x35, 0x9a, 0x56, 0x79, 0x55, 0x2a, 0x5a, 0x77, 0x90, 0x37,
	0x30, 0x8a, 0x13, 0x59, 0xa4, 0x4c, 0x69, 0x74, 0x07, 0xd1, 0x76, 0xad, 0x21, 0xde, 0x05, 0x3b,
	0xe6, 0x32, 0x2a, 0x93, 0xa2, 0x4a, 0x44, 0xee, 0x74, 0x6b, 0x47, 0x2b, 0xe1, 0xc4, 0x22, 0x15,
	0xa5, 0xd3, 0xc3, 0x9a, 0x4e, 0x0e, 0x2a, 0xcf, 0xc4, 0xef, 0xc4, 0xe9, 0x6b, 0x15, 0x93, 0x4b,
	0x0a, 0xf6, 0xd1, 0x1c, 0x64, 0x02, 0xd6, 0x8e, 0x2b, 0x5c, 0x75, 0x48, 0x0f, 0x21, 0x79, 0x07,
	0xdd, 0x3f, 0x2c, 0xdd, 0xeb, 0x2d, 0xed, 0xc5, 0x45, 0xbb, 0x4c, 0xb3, 0x08, 0xf6, 0x53, 0xed,
	0xba, 0x32, 0x3f, 0x1a, 0xb3, 0xbf, 0x26, 0x9c, 0x9d, 0x14, 0x89, 0x03, 0xfd, 0x82, 0xa9, 0x54,
	0xb0, 0x18, 0xd1, 0x23, 0xda, 0xa4, 0x64, 0x09, 0xb6, 0x14, 0xfb, 0x32, 0xe2, 0x61, 0xa5, 0x0a,
	0xfd, 0xc9, 0xf8, 0xf8, 0x62, 0x27, 0x1c, 0x6f, 0x83, 0xd6, 0xad, 0x2a, 0x38, 0x05, 0xf9, 0x1c,
	0x93, 0x05, 0x00, 0xde, 0x4f, 0x33, 0x2c, 0x64, 0x4c, 0x8f, 0x18, 0x87, 0x1a, 0x36, 0x0d, 0x93,
	0x26, 0x24, 0x6f, 0xe1, 0x9c, 0xe7, 0x51, 0xa9, 0xf0, 0x64, 0xe1, 0x8e, 0x2b, 0xe9, 0x74, 0x5c,
	0x6b, 0x3e, 0xa2, 0xe3, 0x56, 0xbe, 0xe1, 0x4a, 0x92, 0xd7, 0x30, 0xac, 0x15, 0x1e, 0xe3, 0xb5,
	0x07, 0xb4, 0x15, 0x66, 0x5f, 0x01, 0xda, 0xa1, 0xc8, 0x05, 0x4c, 0xef, 0x82, 0x9b, 0x60, 0x7d,
	0x1f, 0x84, 0x9b, 0xf5, 0x1d, 0x5d, 0xae, 0xc2, 0xed, 0xc3, 0xed, 0x6a, 0xf2, 0x82, 0x9c, 0x83,
	0x4d, 0xfd, 0xfb, 0xf0, 0xd6, 0x7f, 0xf8, 0xb6, 0xf6, 0xbf, 0x4c, 0x0c, 0x32, 0x06, 0x58, 0x05,
	0x9b, 0xd0, 0xff, 0xee, 0x6f, 0x7d, 0x3a, 0x31, 0x3f, 0x9f, 0xfd, 0xb0, 0xbd, 0xf7, 0x9f, 0x9a,
	0x91, 0x1f, 0x7b, 0x18, 0x7d, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x76, 0x41, 0x72, 0x35, 0xba,
	0x02, 0x00, 0x00,
}
