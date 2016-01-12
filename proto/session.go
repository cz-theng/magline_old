/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import (
	"bytes"
	"encoding/binary"
)

const (
	//SessionSeqLen is length of Request message
	SessionSeqLen = 0
	//SessionRspLen is length of Response message
	SessionRspLen = 4
)

//SessionReqHead is head of message SessionReq
type SessionReqHead struct {
}

// Pack is implement of MessageHeader
func (h *SessionReqHead) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *SessionReqHead) Unpack(buf []byte) error {
	return nil
}

//SessionReqBody is body of message SessionReq
type SessionReqBody struct {
}

// Pack is implement of MessageBodyer
func (h *SessionReqBody) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageBodyer
func (h *SessionReqBody) Unpack(buf []byte) error {
	return nil
}

//SessionReq is sessionreq message
type SessionReq struct {
	head *SessionReqHead
	body *SessionReqBody
}

// Head is implement of Messager
func (sq *SessionReq) Head() MessageHeader {
	return sq.head
}

// Body is implement of Messager
func (sq *SessionReq) Body() MessageBodyer {
	return sq.body
}

// Pack is implement of Messager
func (sq *SessionReq) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of Messager
func (sq *SessionReq) Unpack(buf []byte) error {
	return nil
}

// UnpackSessionReq unpack a session req message from buf
func UnpackSessionReq(buf []byte) (sq *SessionReq, err error) {
	if buf == nil {
		err = ErrBufNil
		return
	}
	if len(buf) == 0 {
		sq = &SessionReq{
			head: nil,
			body: nil,
		}
	}
	return
}

// SessionRspHead is SessionRsp message's head
type SessionRspHead struct {
	AgentID uint32
}

// Pack is implement of MessageHeader
func (h *SessionRspHead) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *SessionRspHead) Unpack(buf []byte) error {
	return nil
}

// SessionRspBody is SessionRsp message's body
type SessionRspBody struct {
}

// Pack is implement of MessageHeader
func (h *SessionRspBody) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *SessionRspBody) Unpack(buf []byte) error {
	return nil
}

//NewSessionRsp create a  SessionRsp message
func NewSessionRsp(id uint32) (msg Messager) {
	head := &SessionRspHead{
		AgentID: id,
	}
	body := &SessionRspBody{}
	msg = &SessionRsp{
		head: head,
		body: body,
	}
	return
}

// SessionRsp is SessionRsp Message
type SessionRsp struct {
	head *SessionRspHead
	body *SessionRspBody
}

// Head is implement of Messager
func (sp *SessionRsp) Head() MessageHeader {
	return sp.head
}

// Body is implement of Messager
func (sp *SessionRsp) Body() MessageBodyer {
	return sp.body
}

// Unpack is implement of Messager
func (sp *SessionRsp) Unpack(buf []byte) error {
	return nil
}

// Pack pack ack to buffer
func (sp *SessionRsp) Pack(buf *bytes.Buffer) (length int, err error) {
	// should check buf's len
	if buf.Cap()-buf.Len() < SessionRspLen {
		err = ErrBufLen
		return
	}
	binary.Write(buf, binary.LittleEndian, sp.head.AgentID)
	length = SessionRspLen
	return
}
