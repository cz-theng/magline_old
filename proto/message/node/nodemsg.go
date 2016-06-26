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
	//NodeMsgHeadLen is length of Request message
	NodeMsgHeadLen = 0
)

//NodeMsgHead is head of message NodeMsg
type NodeMsgHead struct {
}

// Pack is implement of MessageHeader
func (h *NodeMsgHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *NodeMsgHead) Unpack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	return nil
}

//NodeMsgBody is body of message NodeMsg
type NodeMsgBody struct {
	Payload []byte
}

// Pack is implement of MessageBodyer
func (b *NodeMsgBody) Pack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	_, err = buf.Read(b.Payload)
	return
}

// Unpack is implement of MessageBodyer
func (b *NodeMsgBody) Unpack(buf *bytes.Buffer) error {
	b.Payload = buf.Bytes()
	return nil
}

//NodeMsg is NodeMsg message
type NodeMsg struct {
	message.Message
}

//NewNodeMsg create a NodeMsg
func NewNodeMsg(payload []byte) (msg *NodeMsg) {
	head := &NodeMsgHead{}
	body := &NodeMsgBody{
		Payload: payload,
	}
	msg = &NodeMsg{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return msg
}
