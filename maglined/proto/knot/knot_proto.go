/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import (
	"encoding/binary"
	"io"
)

//KnotMessage is knot's proto
type KnotMessage struct {
	proto
	Magic   uint8
	Version uint8
	CMD     uint16
	Seq     uint32
	AgentID uint32
	Length  uint32

	headBuf [16]byte
	ReadBuf []byte
}

//Body get message's body
func (km *KnotMessage) Body() []byte {
	return km.ReadBuf[:km.Length]
}

//Init is initionlize
func (km *KnotMessage) Init(buf []byte) {
	km.Magic = MKMagic
	km.Version = MKVersion
	km.CMD = MKCMDUnknown
	km.Seq = 0
	km.AgentID = 0
	km.Length = 0
	km.ReadBuf = buf
}

//RecvAndUnpack recv and unpack message
func (km *KnotMessage) RecvAndUnpack(rw io.ReadWriter) (err error) {
	if rw == nil {
		// TODO : add log here
		//fmt.Println("rw is null")
	}
	len, err := io.ReadFull(rw, km.headBuf[:])
	if err == io.EOF && len == 16 {
		err = nil
	}
	if err != nil {
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
		err = ErrRequestTooLong
		return
	}

	//fmt.Println("Knot:request cmd is %d, and body length %d", km.CMD, km.Length)
	len, err = io.ReadFull(rw, km.ReadBuf[:km.Length])
	if err == io.EOF && len == int(km.Length) {
		err = nil
	}
	if err != nil {
		return
	}
	return
}

//PackAndSend pack and send message
func (km *KnotMessage) PackAndSend(data []byte, rw io.ReadWriter) (err error) {
	km.headBuf[0] = km.Magic
	km.headBuf[1] = km.Version
	binary.BigEndian.PutUint16(km.headBuf[2:4], km.CMD)
	binary.BigEndian.PutUint32(km.headBuf[4:8], km.Seq)
	binary.BigEndian.PutUint32(km.headBuf[8:12], km.AgentID)
	binary.BigEndian.PutUint32(km.headBuf[12:16], km.Length)
	rw.Write(km.headBuf[:])
	if data != nil {
		rw.Write(data)
	}
	return nil
}
