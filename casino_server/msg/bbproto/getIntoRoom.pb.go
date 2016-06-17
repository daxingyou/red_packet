// Code generated by protoc-gen-go.
// source: getIntoRoom.proto
// DO NOT EDIT!

package bbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetIntoRoom struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	RoomId           *int32  `protobuf:"varint,2,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,3,opt,name=roomType" json:"roomType,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetIntoRoom) Reset()                    { *m = GetIntoRoom{} }
func (m *GetIntoRoom) String() string            { return proto.CompactTextString(m) }
func (*GetIntoRoom) ProtoMessage()               {}
func (*GetIntoRoom) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *GetIntoRoom) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *GetIntoRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *GetIntoRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func init() {
	proto.RegisterType((*GetIntoRoom)(nil), "bbproto.GetIntoRoom")
}

var fileDescriptor2 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4f, 0x2d, 0xf1,
	0xcc, 0x2b, 0xc9, 0x0f, 0xca, 0xcf, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f,
	0x4a, 0x02, 0x33, 0x94, 0xec, 0xb9, 0xb8, 0xdd, 0x11, 0xb2, 0x42, 0x7c, 0x5c, 0x6c, 0xa5, 0xc5,
	0xa9, 0x45, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x20, 0x7e, 0x11, 0x50, 0x1c, 0xc8,
	0x67, 0x02, 0xf2, 0x59, 0x85, 0x04, 0xb8, 0x38, 0x40, 0xfc, 0x90, 0xca, 0x82, 0x54, 0x09, 0x66,
	0x90, 0x08, 0x20, 0x00, 0x00, 0xff, 0xff, 0x39, 0x4f, 0xb8, 0xb1, 0x5d, 0x00, 0x00, 0x00,
}
