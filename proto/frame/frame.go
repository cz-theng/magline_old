/**
* Author: CZ cz.theng@gmail.com
 */

package frame

import (
	"bytes"
	"encoding/binary"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/proto/message"
	"github.com/cz-it/magline/proto/message/knot"
	"github.com/cz-it/magline/proto/message/node"
)

// Head is frame head info
type Head struct {
	Magic   uint8
	Version uint8
	CMD     uint16
	Seq     uint32
	Length  uint32
}

//Init is initionlize
func (fh *Head) Init() {
	fh.Magic = proto.MLMagic
	fh.Version = proto.MLVersion
	fh.CMD = proto.MLCMDUnknown
	fh.Seq = 0
	fh.Length = 0
}

//Unpack unpack framehead
func (fh *Head) Unpack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrFrameHeadBufNil
		return
	}
	if buf.Len() < proto.MLFrameHeadLen {
		err = proto.ErrFameHeadBufLen
		return
	}
	fh.Magic, err = buf.ReadByte()
	fh.Version, err = buf.ReadByte()
	binary.Read(buf, binary.LittleEndian, &fh.CMD)
	binary.Read(buf, binary.LittleEndian, &fh.Seq)
	binary.Read(buf, binary.LittleEndian, &fh.Length)
	return
}

//Pack  pack frame head
func (fh *Head) Pack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrFrameHeadBufNil
		return
	}
	if buf.Cap()-buf.Len() < proto.MLFrameHeadLen {
		err = proto.ErrFameHeadBufLen
		return
	}
	buf.WriteByte(fh.Magic)
	buf.WriteByte(fh.Version)
	binary.Write(buf, binary.LittleEndian, fh.CMD)
	binary.Write(buf, binary.LittleEndian, fh.Seq)
	binary.Write(buf, binary.LittleEndian, fh.Length)
	return
}

//Frame is a proto message
type Frame struct {
	Head *Head
	Body message.Messager
}

// Pack pack a frame message
func (frame *Frame) Pack(buf *bytes.Buffer) (err error) {
	idx := buf.Len()
	err = frame.Head.Pack(buf)
	if err != nil {
		return
	}
	err = frame.Body.Pack(buf)
	binary.LittleEndian.PutUint16(buf.Bytes()[idx+8:idx+proto.MLFrameHeadLen], uint16(buf.Len()-idx-proto.MLFrameHeadLen))
	if err != nil {
		return
	}
	return
}

//UnpackHead upack frame head from bytes.buffer
func UnpackHead(buf *bytes.Buffer) (head *Head, err error) {
	head = new(Head)
	head.Init()
	err = head.Unpack(buf)
	return
}

//UnpackBody unpack a specific
func UnpackBody(cmd uint16, buf *bytes.Buffer) (body message.Messager, err error) {
	switch cmd {
	case proto.MNCMDSYN:
		body = node.NewSYN(proto.BufProtoBin, uint16(proto.ChanNone), uint16(proto.CryptoNone))
	case proto.MNCMDSeesionReq:
		body = node.NewSessionReq()
	case proto.MKCMDConnReq:
		body = knot.NewConnReq(nil)
	case proto.MKCMDConnRsp:
		body = knot.NewConnRsp(nil)
	case proto.MKCMDAgentArriveReq:
		body = knot.NewAgentArriveReq(0)
	case proto.MKCMDAgentArriveRsp:
		body = knot.NewAgentArriveRsp(0, 0)
	case proto.MNCMDNodeMsg:
		body = node.NewNodeMsg(nil)
	case proto.MKCMDNodeMsg:
		body = knot.NewNodeMsg(0, nil)
	case proto.MKCMDKnotMsg:
		body = knot.NewKnotMsg(0, nil)
	case proto.MNCMDDisconnReq:
		body = node.NewDisconnReq()
	case proto.MKCMDAgentQuit:
		body = knot.NewAgentQuit(0)
	case proto.MKCMDDiscardAgent:
		body = knot.NewDiscardAgent(0)
	default:
		err = proto.ErrUnknownCMD
	}
	if err == nil {
		err = body.Unpack(buf)
	}
	return
}
