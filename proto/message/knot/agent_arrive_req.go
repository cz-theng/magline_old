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

//AgentArriveReqHead is head of message
type AgentArriveReqHead struct {
}

// Pack is implement of MessageHeader
func (h *AgentArriveReqHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *AgentArriveReqHead) Unpack(buf *bytes.Buffer) (err error) {
	return
}

//AgentArriveReqBody is body of message SYN
type AgentArriveReqBody struct {
	pb.AgentArriveReq
}

// Pack is implement of MessageBodyer
func (h *AgentArriveReqBody) Pack(buf *bytes.Buffer) (err error) {
	buffer, err := protobuf.Marshal(&h.AgentArriveReq)
	if err != nil {
		return
	}
	_, err = buf.Write(buffer)
	return
}

// Unpack is implement of MessageBodyer
func (h *AgentArriveReqBody) Unpack(buf *bytes.Buffer) (err error) {
	buffer := buf.Next(buf.Len())
	err = protobuf.Unmarshal(buffer, &h.AgentArriveReq)
	return
}

// AgentArriveReq is AgentArriveReq Message from knot
type AgentArriveReq struct {
	message.Message
}

//NewAgentArriveReq new and init a AgentArriveReq message
func NewAgentArriveReq(agentID uint32) (msg *AgentArriveReq) {
	head := &AgentArriveReqHead{}
	body := &AgentArriveReqBody{
		pb.AgentArriveReq{
			AgentID: &agentID,
		},
	}
	msg = &AgentArriveReq{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
