/**
* Author: CZ cz.theng@gmail.com
 */

package node

import (
	"bytes"
	//	"encoding/binary"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/proto/message"
)

const (
	//KnotMsgHeadLen is length of Request message
	KnotMsgHeadLen = 0
)

//KnotMsgHead is head of message KnotMsg
type KnotMsgHead struct {
}

// Pack is implement of MessageHeader
func (h *KnotMsgHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *KnotMsgHead) Unpack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	return nil
}

//KnotMsgBody is body of message KnotMsg
type KnotMsgBody struct {
	Payload []byte
}

// Pack is implement of MessageBodyer
func (b *KnotMsgBody) Pack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	_, err = buf.Write(b.Payload)
	return
}

// Unpack is implement of MessageBodyer
func (b *KnotMsgBody) Unpack(buf *bytes.Buffer) error {
	b.Payload = buf.Bytes()
	return nil
}

//KnotMsg is KnotMsg message
type KnotMsg struct {
	message.Message
}

//NewKnotMsg create a KnotMsg
func NewKnotMsg(payload []byte) (msg *KnotMsg) {
	head := &KnotMsgHead{}
	body := &KnotMsgBody{
		Payload: payload,
	}
	msg = &KnotMsg{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return msg
}
