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
	//ConfirmHeadLen is length of SYN message
	ConfirmHeadLen = 4
)

// ConfirmHead is Confirm message's head
type ConfirmHead struct {
	ErrNO proto.ErrNO
}

// Pack is implement of MessageHeader
func (h *ConfirmHead) Pack(buf *bytes.Buffer) (err error) {
	if buf.Cap()-buf.Len() < ConfirmHeadLen {
		err = proto.ErrBufLen
		return
	}
	binary.Write(buf, binary.LittleEndian, h.ErrNO)
	return
}

// Unpack is implement of MessageHeader
func (h *ConfirmHead) Unpack(buf *bytes.Buffer) error {
	return nil
}

// ConfirmBody is Confirm message's body
type ConfirmBody struct {
}

// Pack is implement of MessageHeader
func (h *ConfirmBody) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *ConfirmBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

//Confirm is Confirm Message
type Confirm struct {
	message.Message
}

//NewConfirm create a Confirm message
func NewConfirm(errno proto.ErrNO) (msg *Confirm) {
	head := &ConfirmHead{
		ErrNO: errno,
	}
	body := &ConfirmBody{}
	msg = &Confirm{
		message.Message{
			Head: head,
			Body: body,
		},
	}

	return
}
