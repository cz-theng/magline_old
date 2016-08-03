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
	//SessionSeqHeadLen is length of Request message
	SessionSeqHeadLen = 0
)

//SessionReqHead is head of message SessionReq
type SessionReqHead struct {
}

// Pack is implement of MessageHeader
func (h *SessionReqHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *SessionReqHead) Unpack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	return nil
}

//SessionReqBody is body of message SessionReq
type SessionReqBody struct {
}

// Pack is implement of MessageBodyer
func (h *SessionReqBody) Pack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	return
}

// Unpack is implement of MessageBodyer
func (h *SessionReqBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

//SessionReq is sessionreq message
type SessionReq struct {
	message.Message
}

//NewSessionReq create a SessionReq
func NewSessionReq() (msg *SessionReq) {
	head := &SessionReqHead{}
	body := &SessionReqBody{}
	msg = &SessionReq{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return msg
}
