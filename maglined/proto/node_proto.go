package proto

/**
* Proto for node
 */

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var (
	ELENGTH_TO_LONG = errors.New("Request's Length is Bigger than Read Buffer!")
)

type NodeProto struct {
	Proto
	Magic   uint8
	Version uint8
	CMD     uint16
	Seq     uint32
	AgentID uint32
	Length  uint32

	headBuf [16]byte
	readBuf []byte
}

func (np *NodeProto) Body() []byte {
	return np.readBuf[:np.Length]
}

func (np *NodeProto) Init(buf []byte) {
	np.readBuf = buf
}

func (np *NodeProto) RecvAndUnpack(rw io.ReadWriter) (err error) {
	if rw == nil {
		// TODO : add log here
		fmt.Println("rw is null")
	}
	len, err := io.ReadFull(rw, np.headBuf[:])
	if err != nil && err != io.EOF {
		return
	}
	if len != cap(np.headBuf) {
		return
	}
	np.Magic = np.headBuf[0]
	np.Version = np.headBuf[1]
	np.CMD = binary.BigEndian.Uint16(np.headBuf[2:4])
	np.Seq = binary.BigEndian.Uint32(np.headBuf[4:8])
	np.AgentID = binary.BigEndian.Uint32(np.headBuf[8:12])
	np.Length = binary.BigEndian.Uint32(np.headBuf[12:16])

	if np.Length > uint32(cap(np.readBuf)) {
		err = ELENGTH_TO_LONG
		return
	}

	fmt.Println("request cmd is %d, and body length %d", np.CMD, np.Length)
	len, err = io.ReadFull(rw, np.readBuf[:np.Length])
	if err != nil && err != io.EOF {
		return
	}
	return
}

func (np *NodeProto) PackAndSend(rw io.ReadWriter) (err error) {
	np.headBuf[0] = np.Magic
	np.headBuf[1] = np.Version
	binary.BigEndian.PutUint16(np.headBuf[2:4], np.CMD)
	binary.BigEndian.PutUint32(np.headBuf[4:8], np.Seq)
	binary.BigEndian.PutUint32(np.headBuf[8:12], np.AgentID)
	binary.BigEndian.PutUint32(np.headBuf[12:16], np.Length)
	rw.Write(np.headBuf[:])
	rw.Write(np.Body())
	return nil
}
