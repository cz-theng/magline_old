/**
* Author: CZ cz.theng@gmail.com
 */

package node

import (
	"bytes"
	"encoding/binary"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/proto/message"
)

const (
	//SYNHeadLen is length of SYN message
	SYNHeadLen = 6
)

//SYNHead is head of message SYN
type SYNHead struct {
	Protobuf uint16
	Channel  uint16
	Crypto   uint16
}

// Pack is implement of MessageHeader
func (h *SYNHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *SYNHead) Unpack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	if buf.Len() < SYNHeadLen {
		err = proto.ErrBufLen
		return
	}
	binary.Read(buf, binary.LittleEndian, &h.Protobuf)
	binary.Read(buf, binary.LittleEndian, &h.Channel)
	binary.Read(buf, binary.LittleEndian, &h.Crypto)
	return nil
}

//SYNBody is body of message SYN
type SYNBody struct {
}

// Pack is implement of MessageBodyer
func (h *SYNBody) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageBodyer
func (h *SYNBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

// SYN is syn
type SYN struct {
	message.Message
}

//NewSYN new and init a SYN message
func NewSYN(protobuf, channel, crypto uint16) (msg *SYN) {
	head := &SYNHead{
		Protobuf: protobuf,
		Channel:  channel,
		Crypto:   crypto,
	}
	body := &SYNBody{}
	msg = &SYN{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return msg
}
