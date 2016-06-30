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
	//ErrorMsgHeadLen is length of SYN message
	ErrorMsgHeadLen = 4
)

// ErrorMsgHead is ErrorMsg message's head
type ErrorMsgHead struct {
	Status proto.Status
	ErrNO  proto.ErrNO
}

// Pack is implement of MessageHeader
func (h *ErrorMsgHead) Pack(buf *bytes.Buffer) (err error) {
	if buf.Cap()-buf.Len() < ErrorMsgHeadLen {
		err = proto.ErrBufLen
		return
	}
	binary.Write(buf, binary.LittleEndian, h.Status)
	binary.Write(buf, binary.LittleEndian, h.ErrNO)
	return
}

// Unpack is implement of MessageHeader
func (h *ErrorMsgHead) Unpack(buf *bytes.Buffer) error {
	return nil
}

// ErrorMsgBody is ErrorMsg message's body
type ErrorMsgBody struct {
}

// Pack is implement of MessageHeader
func (h *ErrorMsgBody) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *ErrorMsgBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

//ErrorMsg is ErrorMsg Message
type ErrorMsg struct {
	message.Message
}

//NewErrorMsg create a ErrorMsg message
func NewErrorMsg(status proto.Status, errno proto.ErrNO) (msg *ErrorMsg) {
	head := &ErrorMsgHead{
		Status: status,
		ErrNO:  errno,
	}
	body := &ErrorMsgBody{}
	msg = &ErrorMsg{
		message.Message{
			Head: head,
			Body: body,
		},
	}

	return
}
