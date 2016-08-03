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
	//SessionRspHeadLen is length of Response message
	SessionRspHeadLen = 4
)

// SessionRspHead is SessionRsp message's head
type SessionRspHead struct {
	AgentID uint32
}

// Pack is implement of MessageHeader
func (h *SessionRspHead) Pack(buf *bytes.Buffer) (err error) {
	if buf.Cap()-buf.Len() < SessionRspHeadLen {
		err = proto.ErrBufLen
		return
	}
	binary.Write(buf, binary.LittleEndian, h.AgentID)
	return
}

// Unpack is implement of MessageHeader
func (h *SessionRspHead) Unpack(buf *bytes.Buffer) error {
	return nil
}

// SessionRspBody is SessionRsp message's body
type SessionRspBody struct {
}

// Pack is implement of MessageHeader
func (h *SessionRspBody) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *SessionRspBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

// SessionRsp is SessionRsp Message
type SessionRsp struct {
	message.Message
}

//NewSessionRsp create a  SessionRsp message
func NewSessionRsp(id uint32) (msg *SessionRsp) {
	head := &SessionRspHead{
		AgentID: id,
	}
	body := &SessionRspBody{}
	msg = &SessionRsp{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
