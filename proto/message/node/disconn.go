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
	//DisconnReqHeadLen is length of Request message
	DisconnReqHeadLen = 0
)

//DisconnReqHead is head of message DisconnReq
type DisconnReqHead struct {
}

// Pack is implement of MessageHeader
func (h *DisconnReqHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *DisconnReqHead) Unpack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	return nil
}

//DisconnReqBody is body of message DisconnReq
type DisconnReqBody struct {
}

// Pack is implement of MessageBodyer
func (b *DisconnReqBody) Pack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	return
}

// Unpack is implement of MessageBodyer
func (b *DisconnReqBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

//DisconnReq is DisconnReq message
type DisconnReq struct {
	message.Message
}

//NewDisconnReq create a DisconnReq
func NewDisconnReq() (msg *DisconnReq) {
	head := &DisconnReqHead{}
	body := &DisconnReqBody{}
	msg = &DisconnReq{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return msg
}
