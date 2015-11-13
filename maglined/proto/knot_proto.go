package proto

/**
* Proto for knot
 */

import (
	"encoding/binary"
	//"errors"
	"fmt"
	"io"
)

var (
//ELENGTH_TO_LONG = errors.New("Request's Length is Bigger than Read Buffer!")
)

type KnotMessage struct {
	Proto
	Magic   uint8
	Version uint8
	CMD     uint16
	Seq     uint32
	AgentID uint32
	Length  uint32

	headBuf [16]byte
	ReadBuf []byte
}

func (km *KnotMessage) Body() []byte {
	return km.ReadBuf[:km.Length]
}

func (km *KnotMessage) Init(buf []byte) {
	km.Magic = MK_MAGIC
	km.Version = MK_VERSION
	km.CMD = MK_CMD_UNKNOWN
	km.Seq = 0
	km.Length = 0
	km.ReadBuf = buf
}

func (km *KnotMessage) RecvAndUnpack(rw io.ReadWriter) (err error) {
	if rw == nil {
		// TODO : add log here
		fmt.Println("rw is null")
	}
	len, err := io.ReadFull(rw, km.headBuf[:])
	if err != nil && err != io.EOF {
		return
	}
	if len != cap(km.headBuf) {
		return
	}
	km.Magic = km.headBuf[0]
	km.Version = km.headBuf[1]
	km.CMD = binary.BigEndian.Uint16(km.headBuf[2:4])
	km.Seq = binary.BigEndian.Uint32(km.headBuf[4:8])
	km.AgentID = binary.BigEndian.Uint32(km.headBuf[8:12])
	km.Length = binary.BigEndian.Uint32(km.headBuf[12:16])

	if km.Length > uint32(cap(km.ReadBuf)) {
		err = ELENGTH_TO_LONG
		return
	}

	fmt.Println("Knot:request cmd is %d, and body length %d", km.CMD, km.Length)
	len, err = io.ReadFull(rw, km.ReadBuf[:km.Length])
	if err != nil && err != io.EOF {
		return
	}
	return
}

func (km *KnotMessage) PackAndSend(rw io.ReadWriter) (err error) {
	km.headBuf[0] = km.Magic
	km.headBuf[1] = km.Version
	binary.BigEndian.PutUint16(km.headBuf[2:4], km.CMD)
	binary.BigEndian.PutUint32(km.headBuf[4:8], km.Seq)
	binary.BigEndian.PutUint32(km.headBuf[8:12], km.AgentID)
	binary.BigEndian.PutUint32(km.headBuf[12:16], km.Length)
	rw.Write(km.headBuf[:])
	rw.Write(km.Body())
	return nil
}
