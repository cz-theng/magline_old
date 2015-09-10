package proto

/**
* Proto for node
*/

import (
	"io"
	"encoding/binary"
	"errors"
)

var (
	ELENGTH_TO_LONG = errors.New("Request's Length is Bigger than Read Buffer!")
)

type NodeProto struct {
	Proto
	Magic uint8
	Version uint8
	CMD uint16
	Seq uint32
	AgentID uint32
	Length uint32

	headBuf [16]byte
	readBuf []byte
}

func (np *NodeProto) Body() ([]byte) {
	return np.readBuf[:np.Length]
}

func (np *NodeProto) Init() {
	
}

func (np *NodeProto) Unpack(rw io.ReadWriter) (err error) {
	len, err := io.ReadFull(rw, np.headBuf[:])
	if err != nil {
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

	len, err = io.ReadFull(rw, np.readBuf[:np.Length])
	if err != nil {
		return 
	}
	return
}


func (np *NodeProto) Pack(rw io.ReadWriter) (err error) {
	return nil
}




















