/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import (
	"encoding/binary"
	"github.com/cz-it/magline/maglined/utils"
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
func (fh *FrameHead) Init(buf []byte) {
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
	fh.CMD = binary.BigEndian.Uint16(framehead[2:4])
	fh.Seq = binary.BigEndian.Uint32(framehead[4:8])
	fh.Length = binary.BigEndian.Uint32(framehead[8:12])

	return
}

//Pack  pack frame head
func (fh *FrameHead) Pack(framehead []byte) (err error) {
	if framehead == nil {
		err = ErrFrameHeadBufNil
		utils.Logger.Error("Pack error :%s", err.Error())
		return
	}
	if cap(framehead) != MLFrameHeadLen {
		err = ErrFameHeadBufLen
		utils.Logger.Error("Pack error :%s", err.Error())
		return
	}
	framehead[0] = fh.Magic
	framehead[1] = fh.Version
	binary.BigEndian.PutUint16(framehead[2:4], fh.CMD)
	binary.BigEndian.PutUint32(framehead[4:8], fh.Seq)
	binary.BigEndian.PutUint32(framehead[8:12], fh.Length)
	return
}

//Frame is a proto message
type Frame struct {
	Head *FrameHead
	Body Messager
}

//UnpackFrameHead upack frame head from bytes.buffer
func UnpackFrameHead(buf []byte) (head *FrameHead, err error) {
	return
}

//UnpackFrameBody recv a frame from reader r and unpack it
func UnpackFrameBody(buf []byte) (body Messager, err error) {
	return
}
