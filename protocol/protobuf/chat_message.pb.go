// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_message.proto

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

type AudioMessage_AudioType int32

const (
	AudioMessage_UNKNOWN_AUDIO_TYPE AudioMessage_AudioType = 0
	AudioMessage_AAC                AudioMessage_AudioType = 1
	AudioMessage_AMR                AudioMessage_AudioType = 2
)

var AudioMessage_AudioType_name = map[int32]string{
	0: "UNKNOWN_AUDIO_TYPE",
	1: "AAC",
	2: "AMR",
}

var AudioMessage_AudioType_value = map[string]int32{
	"UNKNOWN_AUDIO_TYPE": 0,
	"AAC":                1,
	"AMR":                2,
}

func (x AudioMessage_AudioType) String() string {
	return proto.EnumName(AudioMessage_AudioType_name, int32(x))
}

func (AudioMessage_AudioType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{2, 0}
}

type ChatMessage_ContentType int32

const (
	ChatMessage_UNKNOWN_CONTENT_TYPE ChatMessage_ContentType = 0
	ChatMessage_TEXT_PLAIN           ChatMessage_ContentType = 1
	ChatMessage_STICKER              ChatMessage_ContentType = 2
	ChatMessage_STATUS               ChatMessage_ContentType = 3
	ChatMessage_EMOJI                ChatMessage_ContentType = 4
	ChatMessage_TRANSACTION_COMMAND  ChatMessage_ContentType = 5
	// Only local
	ChatMessage_SYSTEM_MESSAGE_CONTENT_PRIVATE_GROUP ChatMessage_ContentType = 6
	ChatMessage_IMAGE                                ChatMessage_ContentType = 7
	ChatMessage_AUDIO                                ChatMessage_ContentType = 8
	ChatMessage_COMMUNITY                            ChatMessage_ContentType = 9
	// Only local
	ChatMessage_SYSTEM_MESSAGE_GAP ChatMessage_ContentType = 10
)

var ChatMessage_ContentType_name = map[int32]string{
	0:  "UNKNOWN_CONTENT_TYPE",
	1:  "TEXT_PLAIN",
	2:  "STICKER",
	3:  "STATUS",
	4:  "EMOJI",
	5:  "TRANSACTION_COMMAND",
	6:  "SYSTEM_MESSAGE_CONTENT_PRIVATE_GROUP",
	7:  "IMAGE",
	8:  "AUDIO",
	9:  "COMMUNITY",
	10: "SYSTEM_MESSAGE_GAP",
}

var ChatMessage_ContentType_value = map[string]int32{
	"UNKNOWN_CONTENT_TYPE":                 0,
	"TEXT_PLAIN":                           1,
	"STICKER":                              2,
	"STATUS":                               3,
	"EMOJI":                                4,
	"TRANSACTION_COMMAND":                  5,
	"SYSTEM_MESSAGE_CONTENT_PRIVATE_GROUP": 6,
	"IMAGE":                                7,
	"AUDIO":                                8,
	"COMMUNITY":                            9,
	"SYSTEM_MESSAGE_GAP":                   10,
}

func (x ChatMessage_ContentType) String() string {
	return proto.EnumName(ChatMessage_ContentType_name, int32(x))
}

func (ChatMessage_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{4, 0}
}

type StickerMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Pack                 int32    `protobuf:"varint,2,opt,name=pack,proto3" json:"pack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StickerMessage) Reset()         { *m = StickerMessage{} }
func (m *StickerMessage) String() string { return proto.CompactTextString(m) }
func (*StickerMessage) ProtoMessage()    {}
func (*StickerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{0}
}

func (m *StickerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StickerMessage.Unmarshal(m, b)
}
func (m *StickerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StickerMessage.Marshal(b, m, deterministic)
}
func (m *StickerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StickerMessage.Merge(m, src)
}
func (m *StickerMessage) XXX_Size() int {
	return xxx_messageInfo_StickerMessage.Size(m)
}
func (m *StickerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StickerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StickerMessage proto.InternalMessageInfo

func (m *StickerMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *StickerMessage) GetPack() int32 {
	if m != nil {
		return m.Pack
	}
	return 0
}

type ImageMessage struct {
	Payload              []byte    `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Type                 ImageType `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.ImageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ImageMessage) Reset()         { *m = ImageMessage{} }
func (m *ImageMessage) String() string { return proto.CompactTextString(m) }
func (*ImageMessage) ProtoMessage()    {}
func (*ImageMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{1}
}

func (m *ImageMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageMessage.Unmarshal(m, b)
}
func (m *ImageMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageMessage.Marshal(b, m, deterministic)
}
func (m *ImageMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageMessage.Merge(m, src)
}
func (m *ImageMessage) XXX_Size() int {
	return xxx_messageInfo_ImageMessage.Size(m)
}
func (m *ImageMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ImageMessage proto.InternalMessageInfo

func (m *ImageMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ImageMessage) GetType() ImageType {
	if m != nil {
		return m.Type
	}
	return ImageType_UNKNOWN_IMAGE_TYPE
}

type AudioMessage struct {
	Payload              []byte                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Type                 AudioMessage_AudioType `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.AudioMessage_AudioType" json:"type,omitempty"`
	DurationMs           uint64                 `protobuf:"varint,3,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AudioMessage) Reset()         { *m = AudioMessage{} }
func (m *AudioMessage) String() string { return proto.CompactTextString(m) }
func (*AudioMessage) ProtoMessage()    {}
func (*AudioMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{2}
}

func (m *AudioMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioMessage.Unmarshal(m, b)
}
func (m *AudioMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioMessage.Marshal(b, m, deterministic)
}
func (m *AudioMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioMessage.Merge(m, src)
}
func (m *AudioMessage) XXX_Size() int {
	return xxx_messageInfo_AudioMessage.Size(m)
}
func (m *AudioMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AudioMessage proto.InternalMessageInfo

func (m *AudioMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *AudioMessage) GetType() AudioMessage_AudioType {
	if m != nil {
		return m.Type
	}
	return AudioMessage_UNKNOWN_AUDIO_TYPE
}

func (m *AudioMessage) GetDurationMs() uint64 {
	if m != nil {
		return m.DurationMs
	}
	return 0
}

type EditMessage struct {
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// Text of the message
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	ChatId    string `protobuf:"bytes,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	MessageId string `protobuf:"bytes,4,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// Grant for community edit messages
	Grant []byte `protobuf:"bytes,5,opt,name=grant,proto3" json:"grant,omitempty"`
	// The type of message (public/one-to-one/private-group-chat)
	MessageType          MessageType `protobuf:"varint,6,opt,name=message_type,json=messageType,proto3,enum=protobuf.MessageType" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EditMessage) Reset()         { *m = EditMessage{} }
func (m *EditMessage) String() string { return proto.CompactTextString(m) }
func (*EditMessage) ProtoMessage()    {}
func (*EditMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{3}
}

func (m *EditMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditMessage.Unmarshal(m, b)
}
func (m *EditMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditMessage.Marshal(b, m, deterministic)
}
func (m *EditMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditMessage.Merge(m, src)
}
func (m *EditMessage) XXX_Size() int {
	return xxx_messageInfo_EditMessage.Size(m)
}
func (m *EditMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EditMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EditMessage proto.InternalMessageInfo

func (m *EditMessage) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *EditMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *EditMessage) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *EditMessage) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *EditMessage) GetGrant() []byte {
	if m != nil {
		return m.Grant
	}
	return nil
}

func (m *EditMessage) GetMessageType() MessageType {
	if m != nil {
		return m.MessageType
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

type ChatMessage struct {
	// Lamport timestamp of the chat message
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// Unix timestamps in milliseconds, currently not used as we use whisper as more reliable, but here
	// so that we don't rely on it
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Text of the message
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// Id of the message that we are replying to
	ResponseTo string `protobuf:"bytes,4,opt,name=response_to,json=responseTo,proto3" json:"response_to,omitempty"`
	// Ens name of the sender
	EnsName string `protobuf:"bytes,5,opt,name=ens_name,json=ensName,proto3" json:"ens_name,omitempty"`
	// Chat id, this field is symmetric for public-chats and private group chats,
	// but asymmetric in case of one-to-ones, as the sender will use the chat-id
	// of the received, while the receiver will use the chat-id of the sender.
	// Probably should be the concatenation of sender-pk & receiver-pk in alphabetical order
	ChatId string `protobuf:"bytes,6,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// The type of message (public/one-to-one/private-group-chat)
	MessageType MessageType `protobuf:"varint,7,opt,name=message_type,json=messageType,proto3,enum=protobuf.MessageType" json:"message_type,omitempty"`
	// The type of the content of the message
	ContentType ChatMessage_ContentType `protobuf:"varint,8,opt,name=content_type,json=contentType,proto3,enum=protobuf.ChatMessage_ContentType" json:"content_type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ChatMessage_Sticker
	//	*ChatMessage_Image
	//	*ChatMessage_Audio
	//	*ChatMessage_Community
	Payload isChatMessage_Payload `protobuf_oneof:"payload"`
	// Grant for community chat messages
	Grant                []byte   `protobuf:"bytes,13,opt,name=grant,proto3" json:"grant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_263952f55fd35689, []int{4}
}

func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (m *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(m, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *ChatMessage) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ChatMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *ChatMessage) GetResponseTo() string {
	if m != nil {
		return m.ResponseTo
	}
	return ""
}

func (m *ChatMessage) GetEnsName() string {
	if m != nil {
		return m.EnsName
	}
	return ""
}

func (m *ChatMessage) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *ChatMessage) GetMessageType() MessageType {
	if m != nil {
		return m.MessageType
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *ChatMessage) GetContentType() ChatMessage_ContentType {
	if m != nil {
		return m.ContentType
	}
	return ChatMessage_UNKNOWN_CONTENT_TYPE
}

type isChatMessage_Payload interface {
	isChatMessage_Payload()
}

type ChatMessage_Sticker struct {
	Sticker *StickerMessage `protobuf:"bytes,9,opt,name=sticker,proto3,oneof"`
}

type ChatMessage_Image struct {
	Image *ImageMessage `protobuf:"bytes,10,opt,name=image,proto3,oneof"`
}

type ChatMessage_Audio struct {
	Audio *AudioMessage `protobuf:"bytes,11,opt,name=audio,proto3,oneof"`
}

type ChatMessage_Community struct {
	Community []byte `protobuf:"bytes,12,opt,name=community,proto3,oneof"`
}

func (*ChatMessage_Sticker) isChatMessage_Payload() {}

func (*ChatMessage_Image) isChatMessage_Payload() {}

func (*ChatMessage_Audio) isChatMessage_Payload() {}

func (*ChatMessage_Community) isChatMessage_Payload() {}

func (m *ChatMessage) GetPayload() isChatMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChatMessage) GetSticker() *StickerMessage {
	if x, ok := m.GetPayload().(*ChatMessage_Sticker); ok {
		return x.Sticker
	}
	return nil
}

func (m *ChatMessage) GetImage() *ImageMessage {
	if x, ok := m.GetPayload().(*ChatMessage_Image); ok {
		return x.Image
	}
	return nil
}

func (m *ChatMessage) GetAudio() *AudioMessage {
	if x, ok := m.GetPayload().(*ChatMessage_Audio); ok {
		return x.Audio
	}
	return nil
}

func (m *ChatMessage) GetCommunity() []byte {
	if x, ok := m.GetPayload().(*ChatMessage_Community); ok {
		return x.Community
	}
	return nil
}

func (m *ChatMessage) GetGrant() []byte {
	if m != nil {
		return m.Grant
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ChatMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ChatMessage_Sticker)(nil),
		(*ChatMessage_Image)(nil),
		(*ChatMessage_Audio)(nil),
		(*ChatMessage_Community)(nil),
	}
}

func init() {
	proto.RegisterEnum("protobuf.AudioMessage_AudioType", AudioMessage_AudioType_name, AudioMessage_AudioType_value)
	proto.RegisterEnum("protobuf.ChatMessage_ContentType", ChatMessage_ContentType_name, ChatMessage_ContentType_value)
	proto.RegisterType((*StickerMessage)(nil), "protobuf.StickerMessage")
	proto.RegisterType((*ImageMessage)(nil), "protobuf.ImageMessage")
	proto.RegisterType((*AudioMessage)(nil), "protobuf.AudioMessage")
	proto.RegisterType((*EditMessage)(nil), "protobuf.EditMessage")
	proto.RegisterType((*ChatMessage)(nil), "protobuf.ChatMessage")
}

func init() {
	proto.RegisterFile("chat_message.proto", fileDescriptor_263952f55fd35689)
}

var fileDescriptor_263952f55fd35689 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0x9b, 0x4c,
	0x10, 0x0e, 0xf1, 0x07, 0x66, 0x70, 0x22, 0xb4, 0xc9, 0x9b, 0xf0, 0xbe, 0x7a, 0xdb, 0xb8, 0x56,
	0xa5, 0xfa, 0xe4, 0x43, 0x9a, 0x4a, 0xb9, 0x52, 0x07, 0x39, 0x34, 0x05, 0xbb, 0xcb, 0xba, 0x6d,
	0x4e, 0x68, 0x83, 0xb7, 0x31, 0x4a, 0xf8, 0x90, 0x59, 0x4b, 0xf5, 0x1f, 0xeb, 0xb5, 0x3f, 0xa4,
	0x87, 0xfe, 0x95, 0x6a, 0x17, 0x63, 0x88, 0x2f, 0xe9, 0xc9, 0x33, 0xc3, 0x3c, 0x0f, 0xf3, 0x3c,
	0xcc, 0x18, 0x50, 0xb8, 0xa0, 0x3c, 0x88, 0x59, 0x9e, 0xd3, 0x7b, 0x36, 0xcc, 0x96, 0x29, 0x4f,
	0x51, 0x47, 0xfe, 0xdc, 0xad, 0xbe, 0xfd, 0xa7, 0xb3, 0x64, 0x15, 0xe7, 0x45, 0xb9, 0x7f, 0x09,
	0x87, 0x3e, 0x8f, 0xc2, 0x07, 0xb6, 0x74, 0x8b, 0x76, 0x84, 0xa0, 0xb9, 0xa0, 0xf9, 0xc2, 0x54,
	0x7a, 0xca, 0x40, 0xc3, 0x32, 0x16, 0xb5, 0x8c, 0x86, 0x0f, 0xe6, 0x7e, 0x4f, 0x19, 0xb4, 0xb0,
	0x8c, 0xfb, 0x9f, 0xa0, 0xeb, 0xc4, 0xf4, 0x9e, 0x95, 0x38, 0x13, 0xd4, 0x8c, 0xae, 0x1f, 0x53,
	0x3a, 0x97, 0xd0, 0x2e, 0x2e, 0x53, 0xf4, 0x06, 0x9a, 0x7c, 0x9d, 0x31, 0x89, 0x3e, 0x3c, 0x3f,
	0x1a, 0x96, 0x93, 0x0c, 0x25, 0x9e, 0xac, 0x33, 0x86, 0x65, 0x43, 0xff, 0x87, 0x02, 0x5d, 0x6b,
	0x35, 0x8f, 0xd2, 0xe7, 0x39, 0x2f, 0x9e, 0x70, 0xf6, 0x2a, 0xce, 0x3a, 0xbe, 0x48, 0xaa, 0x17,
	0xa0, 0x33, 0xd0, 0xe7, 0xab, 0x25, 0xe5, 0x51, 0x9a, 0x04, 0x71, 0x6e, 0x36, 0x7a, 0xca, 0xa0,
	0x89, 0xa1, 0x2c, 0xb9, 0x79, 0xff, 0x1d, 0x68, 0x5b, 0x0c, 0x3a, 0x01, 0x34, 0xf3, 0x6e, 0xbc,
	0xc9, 0x17, 0x2f, 0xb0, 0x66, 0x57, 0xce, 0x24, 0x20, 0xb7, 0x53, 0xdb, 0xd8, 0x43, 0x2a, 0x34,
	0x2c, 0x6b, 0x64, 0x28, 0x32, 0x70, 0xb1, 0xb1, 0xdf, 0xff, 0xa9, 0x80, 0x6e, 0xcf, 0x23, 0x5e,
	0xce, 0x7d, 0x0c, 0xad, 0xf0, 0x31, 0x0d, 0x1f, 0xe4, 0xd4, 0x4d, 0x5c, 0x24, 0xc2, 0x45, 0xce,
	0xbe, 0x73, 0x39, 0xb3, 0x86, 0x65, 0x8c, 0x4e, 0x41, 0x95, 0x1f, 0x2b, 0x9a, 0xcb, 0x69, 0x34,
	0xdc, 0x16, 0xa9, 0x33, 0x47, 0x2f, 0x00, 0x36, 0x1f, 0x50, 0x3c, 0x6b, 0xca, 0x67, 0xda, 0xa6,
	0xe2, 0xcc, 0xc5, 0x1b, 0xee, 0x97, 0x34, 0xe1, 0x66, 0x4b, 0xfa, 0x52, 0x24, 0xe8, 0x12, 0xba,
	0x25, 0x48, 0xba, 0xd3, 0x96, 0xee, 0xfc, 0x53, 0xb9, 0xb3, 0x19, 0x50, 0x5a, 0xa2, 0xc7, 0x55,
	0xd2, 0xff, 0xdd, 0x02, 0x7d, 0xb4, 0xa0, 0xcf, 0x28, 0xf8, 0x1f, 0x34, 0x1e, 0xc5, 0x2c, 0xe7,
	0x34, 0xce, 0xa4, 0x8c, 0x26, 0xae, 0x0a, 0x5b, 0x7d, 0x8d, 0x9a, 0xbe, 0x33, 0xd0, 0x97, 0x2c,
	0xcf, 0xd2, 0x24, 0x67, 0x01, 0x4f, 0x37, 0x3a, 0xa0, 0x2c, 0x91, 0x14, 0xfd, 0x0b, 0x1d, 0x96,
	0xe4, 0x41, 0x42, 0x63, 0x26, 0xb5, 0x68, 0x58, 0x65, 0x49, 0xee, 0xd1, 0x98, 0xd5, 0xbd, 0x69,
	0x3f, 0xf1, 0x66, 0x57, 0xa6, 0xfa, 0xb7, 0x32, 0xd1, 0x15, 0x74, 0xc3, 0x34, 0xe1, 0x2c, 0xe1,
	0x05, 0xb2, 0x23, 0x91, 0xaf, 0x2a, 0x64, 0xcd, 0x83, 0xe1, 0xa8, 0xe8, 0x2c, 0x58, 0xc2, 0x2a,
	0x41, 0x17, 0xa0, 0xe6, 0xc5, 0xd1, 0x98, 0x5a, 0x4f, 0x19, 0xe8, 0xe7, 0x66, 0x45, 0xf0, 0xf4,
	0x9a, 0xae, 0xf7, 0x70, 0xd9, 0x8a, 0x86, 0xd0, 0x8a, 0xc4, 0xc2, 0x9b, 0x20, 0x31, 0x27, 0x3b,
	0x77, 0x50, 0x21, 0x8a, 0x36, 0xd1, 0x4f, 0xc5, 0x2e, 0x9a, 0xfa, 0x6e, 0x7f, 0x7d, 0xc7, 0x45,
	0xbf, 0x6c, 0x43, 0x2f, 0x41, 0x0b, 0xd3, 0x38, 0x5e, 0x25, 0x11, 0x5f, 0x9b, 0x5d, 0xb1, 0x16,
	0xd7, 0x7b, 0xb8, 0x2a, 0x55, 0x2b, 0x73, 0x50, 0x5b, 0x99, 0xfe, 0x2f, 0x05, 0xf4, 0x9a, 0x50,
	0x64, 0xc2, 0x71, 0xb9, 0xf4, 0xa3, 0x89, 0x47, 0x6c, 0x8f, 0x94, 0x6b, 0x7f, 0x08, 0x40, 0xec,
	0xaf, 0x24, 0x98, 0x7e, 0xb4, 0x1c, 0xcf, 0x50, 0x90, 0x0e, 0xaa, 0x4f, 0x9c, 0xd1, 0x8d, 0x8d,
	0x8d, 0x7d, 0x04, 0xd0, 0xf6, 0x89, 0x45, 0x66, 0xbe, 0xd1, 0x40, 0x1a, 0xb4, 0x6c, 0x77, 0xf2,
	0xc1, 0x31, 0x9a, 0xe8, 0x14, 0x8e, 0x08, 0xb6, 0x3c, 0xdf, 0x1a, 0x11, 0x67, 0x22, 0x18, 0x5d,
	0xd7, 0xf2, 0xae, 0x8c, 0x16, 0x1a, 0xc0, 0x6b, 0xff, 0xd6, 0x27, 0xb6, 0x1b, 0xb8, 0xb6, 0xef,
	0x5b, 0x63, 0x7b, 0xfb, 0xb6, 0x29, 0x76, 0x3e, 0x5b, 0xc4, 0x0e, 0xc6, 0x78, 0x32, 0x9b, 0x1a,
	0x6d, 0xc1, 0xe6, 0xb8, 0xd6, 0xd8, 0x36, 0x54, 0x11, 0xca, 0x43, 0x34, 0x3a, 0xe8, 0x00, 0x34,
	0x41, 0x36, 0xf3, 0x1c, 0x72, 0x6b, 0x68, 0xe2, 0x54, 0x77, 0xe8, 0xc6, 0xd6, 0xd4, 0x80, 0xf7,
	0xda, 0xf6, 0x0f, 0xe4, 0xae, 0x2d, 0xed, 0x7b, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x56,
	0xc0, 0x4f, 0x1d, 0x05, 0x00, 0x00,
}
