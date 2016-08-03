// Code generated by protoc-gen-go.
// source: knot.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	knot.proto

It has these top-level messages:
	ConnReq
	ConnRsp
	AgentArriveReq
	AgentArriveRsp
	DisconnReq
	DisconnRsp
	NodeMsg
	KnotMsg
	AgentQuit
	DiscardAgent
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// Connect to Maglined
type ConnReq struct {
	AccessKey        []byte `protobuf:"bytes,1,opt,name=AccessKey" json:"AccessKey,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ConnReq) Reset()                    { *m = ConnReq{} }
func (m *ConnReq) String() string            { return proto.CompactTextString(m) }
func (*ConnReq) ProtoMessage()               {}
func (*ConnReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConnReq) GetAccessKey() []byte {
	if m != nil {
		return m.AccessKey
	}
	return nil
}

type ConnRsp struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ConnRsp) Reset()                    { *m = ConnRsp{} }
func (m *ConnRsp) String() string            { return proto.CompactTextString(m) }
func (*ConnRsp) ProtoMessage()               {}
func (*ConnRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// *
// New Agent from Maglined
type AgentArriveReq struct {
	AgentID          *uint32 `protobuf:"varint,1,opt,name=AgentID" json:"AgentID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AgentArriveReq) Reset()                    { *m = AgentArriveReq{} }
func (m *AgentArriveReq) String() string            { return proto.CompactTextString(m) }
func (*AgentArriveReq) ProtoMessage()               {}
func (*AgentArriveReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AgentArriveReq) GetAgentID() uint32 {
	if m != nil && m.AgentID != nil {
		return *m.AgentID
	}
	return 0
}

type AgentArriveRsp struct {
	AgentID          *uint32 `protobuf:"varint,1,opt,name=AgentID" json:"AgentID,omitempty"`
	Errno            *int32  `protobuf:"varint,2,opt,name=errno" json:"errno,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AgentArriveRsp) Reset()                    { *m = AgentArriveRsp{} }
func (m *AgentArriveRsp) String() string            { return proto.CompactTextString(m) }
func (*AgentArriveRsp) ProtoMessage()               {}
func (*AgentArriveRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AgentArriveRsp) GetAgentID() uint32 {
	if m != nil && m.AgentID != nil {
		return *m.AgentID
	}
	return 0
}

func (m *AgentArriveRsp) GetErrno() int32 {
	if m != nil && m.Errno != nil {
		return *m.Errno
	}
	return 0
}

// *
// Disconnect from Maglined
type DisconnReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DisconnReq) Reset()                    { *m = DisconnReq{} }
func (m *DisconnReq) String() string            { return proto.CompactTextString(m) }
func (*DisconnReq) ProtoMessage()               {}
func (*DisconnReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DisconnRsp struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DisconnRsp) Reset()                    { *m = DisconnRsp{} }
func (m *DisconnRsp) String() string            { return proto.CompactTextString(m) }
func (*DisconnRsp) ProtoMessage()               {}
func (*DisconnRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// *
// Message from Agent on Maglined
type NodeMsg struct {
	AgentID          *uint32 `protobuf:"varint,1,opt,name=AgentID" json:"AgentID,omitempty"`
	Payload          []byte  `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NodeMsg) Reset()                    { *m = NodeMsg{} }
func (m *NodeMsg) String() string            { return proto.CompactTextString(m) }
func (*NodeMsg) ProtoMessage()               {}
func (*NodeMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NodeMsg) GetAgentID() uint32 {
	if m != nil && m.AgentID != nil {
		return *m.AgentID
	}
	return 0
}

func (m *NodeMsg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// *
// Send Message to Agent on Maglined
type KnotMsg struct {
	AgentID          *uint32 `protobuf:"varint,1,opt,name=AgentID" json:"AgentID,omitempty"`
	Payload          []byte  `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KnotMsg) Reset()                    { *m = KnotMsg{} }
func (m *KnotMsg) String() string            { return proto.CompactTextString(m) }
func (*KnotMsg) ProtoMessage()               {}
func (*KnotMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *KnotMsg) GetAgentID() uint32 {
	if m != nil && m.AgentID != nil {
		return *m.AgentID
	}
	return 0
}

func (m *KnotMsg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// *
// Agent Quit from Maglined
type AgentQuit struct {
	AgentID          *uint32 `protobuf:"varint,1,opt,name=AgentID" json:"AgentID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AgentQuit) Reset()                    { *m = AgentQuit{} }
func (m *AgentQuit) String() string            { return proto.CompactTextString(m) }
func (*AgentQuit) ProtoMessage()               {}
func (*AgentQuit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AgentQuit) GetAgentID() uint32 {
	if m != nil && m.AgentID != nil {
		return *m.AgentID
	}
	return 0
}

// *
// Discard a agent
type DiscardAgent struct {
	AgentID          *uint32 `protobuf:"varint,1,opt,name=AgentID" json:"AgentID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DiscardAgent) Reset()                    { *m = DiscardAgent{} }
func (m *DiscardAgent) String() string            { return proto.CompactTextString(m) }
func (*DiscardAgent) ProtoMessage()               {}
func (*DiscardAgent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DiscardAgent) GetAgentID() uint32 {
	if m != nil && m.AgentID != nil {
		return *m.AgentID
	}
	return 0
}

func init() {
	proto.RegisterType((*ConnReq)(nil), "pb.ConnReq")
	proto.RegisterType((*ConnRsp)(nil), "pb.ConnRsp")
	proto.RegisterType((*AgentArriveReq)(nil), "pb.AgentArriveReq")
	proto.RegisterType((*AgentArriveRsp)(nil), "pb.AgentArriveRsp")
	proto.RegisterType((*DisconnReq)(nil), "pb.DisconnReq")
	proto.RegisterType((*DisconnRsp)(nil), "pb.DisconnRsp")
	proto.RegisterType((*NodeMsg)(nil), "pb.NodeMsg")
	proto.RegisterType((*KnotMsg)(nil), "pb.KnotMsg")
	proto.RegisterType((*AgentQuit)(nil), "pb.AgentQuit")
	proto.RegisterType((*DiscardAgent)(nil), "pb.DiscardAgent")
}

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xce, 0xcb, 0x2f,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe1, 0x62, 0x77, 0xce,
	0xcf, 0xcb, 0x0b, 0x4a, 0x2d, 0x14, 0x12, 0xe4, 0xe2, 0x74, 0x4c, 0x4e, 0x4e, 0x2d, 0x2e, 0xf6,
	0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x51, 0xe2, 0x84, 0xca, 0x16, 0x17, 0x28, 0x29,
	0x72, 0xf1, 0x39, 0xa6, 0xa7, 0xe6, 0x95, 0x38, 0x16, 0x15, 0x65, 0x96, 0xa5, 0x82, 0xd4, 0xf3,
	0x73, 0xb1, 0x83, 0x45, 0x3c, 0x5d, 0xc0, 0xaa, 0x79, 0x95, 0x0c, 0x50, 0x95, 0x14, 0x17, 0x60,
	0x28, 0x11, 0xe2, 0xe5, 0x62, 0x4d, 0x2d, 0x2a, 0xca, 0xcb, 0x97, 0x60, 0x02, 0x72, 0x59, 0x95,
	0x78, 0xb8, 0xb8, 0x5c, 0x32, 0x8b, 0x93, 0x21, 0x0e, 0x40, 0xe6, 0x01, 0x2d, 0xd4, 0xe6, 0x62,
	0xf7, 0xcb, 0x4f, 0x49, 0xf5, 0x2d, 0x4e, 0xc7, 0x34, 0x06, 0x28, 0x50, 0x90, 0x58, 0x99, 0x93,
	0x9f, 0x98, 0x02, 0x36, 0x88, 0x07, 0xa4, 0xd8, 0x1b, 0xe8, 0x31, 0xe2, 0x14, 0xcb, 0x00, 0x3d,
	0x0a, 0x52, 0x11, 0x58, 0x9a, 0x59, 0x82, 0xe9, 0x0b, 0x79, 0x2e, 0x1e, 0x90, 0x2b, 0x12, 0x8b,
	0x52, 0xc0, 0xe2, 0x18, 0x0a, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96, 0x6c, 0x9c, 0x76, 0x43,
	0x01, 0x00, 0x00,
}
