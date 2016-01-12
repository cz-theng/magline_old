/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import (
	"bytes"
	"encoding/binary"
	"github.com/cz-it/magline/utils"
)

// FrameHead is frame head info
type FrameHead struct {
	Magic   uint8
	Version uint8
	CMD     uint16
	Seq     uint32
	Length  uint32
}

//Init is initionlize
func (fh *FrameHead) Init() {
	fh.Magic = MLMagic
	fh.Version = MLVersion
	fh.CMD = MLCMDUnknown
	fh.Seq = 0
	fh.Length = 0
}

//Unpack unpack framehead
func (fh *FrameHead) Unpack(framehead []byte) (err error) {
	if framehead == nil {
		err = ErrFrameHeadBufNil
		utils.Logger.Error("Unpack error :%s", err.Error())
		return
	}
	if len(framehead) != MLFrameHeadLen {
		err = ErrFameHeadBufLen
		utils.Logger.Error("Unpack error :%s", err.Error())
		return
	}
	fh.Magic = framehead[0]
	fh.Version = framehead[1]
	fh.CMD = binary.LittleEndian.Uint16(framehead[2:4])
	fh.Seq = binary.LittleEndian.Uint32(framehead[4:8])
	fh.Length = binary.LittleEndian.Uint32(framehead[8:12])
	return
}

//Pack  pack frame head
func (fh *FrameHead) Pack(buf *bytes.Buffer) (length int, err error) {
	if buf == nil {
		err = ErrFrameHeadBufNil
		utils.Logger.Error("Pack error :%s", err.Error())
		return
	}
	if buf.Cap()-buf.Len() < MLFrameHeadLen {
		err = ErrFameHeadBufLen
		utils.Logger.Error("Pack error :%s", err.Error())
		return
	}
	buf.WriteByte(fh.Magic)
	buf.WriteByte(fh.Version)
	binary.Write(buf, binary.LittleEndian, fh.CMD)
	binary.Write(buf, binary.LittleEndian, fh.Seq)
	binary.Write(buf, binary.LittleEndian, fh.Length)
	length = ACKLen
	return
}

//Frame is a proto message
type Frame struct {
	Head *FrameHead
	Body Messager
}

// Pack pack a frame message
func (frame *Frame) Pack(buf *bytes.Buffer) (length int, err error) {
	offset, err := frame.Head.Pack(buf)
	length += offset
	if err != nil {
		return
	}
	offset, err = frame.Body.Pack(buf)
	length += offset
	binary.LittleEndian.PutUint16(buf.Bytes()[buf.Len()-offset-4:buf.Len()-offset], uint16(offset))
	if err != nil {
		return
	}
	return
}

//UnpackFrameHead upack frame head from bytes.buffer
func UnpackFrameHead(buf []byte) (head *FrameHead, err error) {
	head = new(FrameHead)
	head.Init()
	err = head.Unpack(buf)
	return
}

//UnpackFrameBody recv a frame from reader r and unpack it
func UnpackFrameBody(cmd uint16, buf []byte) (body Messager, err error) {
	switch cmd {
	case MNCMDSYN:
		body, err = UnpackSYN(buf)
	case MNCMDSeesionReq:
		body, err = UnpackSessionReq(buf)
	default:
		err = ErrUnknownCMD
	}
	return
}
