/**
* Author: CZ cz.theng@gmail.com
 */

package knot

import (
	"bytes"
	"github.com/cz-it/magline/proto/message"
	"github.com/cz-it/magline/proto/message/knot/pb"
	protobuf "github.com/golang/protobuf/proto"
)

//KnotMsgHead is head of message
type KnotMsgHead struct {
}

// Pack is implement of MessageHeader
func (h *KnotMsgHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *KnotMsgHead) Unpack(buf *bytes.Buffer) (err error) {
	return
}

//KnotMsgBody is body of message SYN
type KnotMsgBody struct {
	pb.KnotMsg
}

// Pack is implement of MessageBodyer
func (b *KnotMsgBody) Pack(buf *bytes.Buffer) (err error) {
	buffer, err := protobuf.Marshal(&b.KnotMsg)
	if err != nil {
		return
	}
	_, err = buf.Write(buffer)
	return
}

// Unpack is implement of MessageBodyer
func (b *KnotMsgBody) Unpack(buf *bytes.Buffer) (err error) {
	buffer := buf.Next(buf.Len())
	err = protobuf.Unmarshal(buffer, &b.KnotMsg)
	return
}

// KnotMsg is KnotMsg Message from knot
type KnotMsg struct {
	message.Message
}

//NewKnotMsg new and init a KnotMsg message
func NewKnotMsg(agentID uint32, payload []byte) (msg *KnotMsg) {
	head := &KnotMsgHead{}
	body := &KnotMsgBody{
		pb.KnotMsg{
			AgentID: &agentID,
			Payload: payload,
		},
	}
	msg = &KnotMsg{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
