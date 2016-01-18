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
	//ACKHeadLen is length of SYN message
	ACKHeadLen = 4
)

// ACKHead is ACK message's head
type ACKHead struct {
	Channel uint16
	Crypto  uint16
}

// Pack is implement of MessageHeader
func (h *ACKHead) Pack(buf *bytes.Buffer) (err error) {
	if buf.Cap()-buf.Len() < ACKHeadLen {
		err = proto.ErrBufLen
		return
	}
	binary.Write(buf, binary.LittleEndian, h.Channel)
	binary.Write(buf, binary.LittleEndian, h.Crypto)
	return
}

// Unpack is implement of MessageHeader
func (h *ACKHead) Unpack(buf *bytes.Buffer) error {
	return nil
}

// ACKBody is ACK message's body
type ACKBody struct {
}

// Pack is implement of MessageHeader
func (h *ACKBody) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *ACKBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

//ACK is ACK Message
type ACK struct {
	message.Message
}

//NewACK create a ACK message
func NewACK(channel uint16, crypto uint16) (msg *ACK) {
	head := &ACKHead{
		Channel: channel,
		Crypto:  crypto,
	}
	body := &ACKBody{}
	msg = &ACK{
		message.Message{
			Head: head,
			Body: body,
		},
	}

	return
}
