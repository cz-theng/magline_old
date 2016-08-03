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

//AgentArriveRspHead is head of message
type AgentArriveRspHead struct {
}

// Pack is implement of MessageHeader
func (h *AgentArriveRspHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *AgentArriveRspHead) Unpack(buf *bytes.Buffer) (err error) {
	return
}

//AgentArriveRspBody is body of message SYN
type AgentArriveRspBody struct {
	pb.AgentArriveRsp
}

// Pack is implement of MessageBodyer
func (h *AgentArriveRspBody) Pack(buf *bytes.Buffer) (err error) {
	buffer, err := protobuf.Marshal(&h.AgentArriveRsp)
	if err != nil {
		return
	}
	_, err = buf.Write(buffer)
	return
}

// Unpack is implement of MessageBodyer
func (h *AgentArriveRspBody) Unpack(buf *bytes.Buffer) (err error) {
	buffer := buf.Next(buf.Len())
	err = protobuf.Unmarshal(buffer, &h.AgentArriveRsp)
	return
}

// AgentArriveRsp is AgentArriveRsp Message from knot
type AgentArriveRsp struct {
	message.Message
}

//NewAgentArriveRsp new and init a AgentArriveRsp message
func NewAgentArriveRsp(agentID uint32, errno int32) (msg *AgentArriveRsp) {
	head := &AgentArriveRspHead{}
	body := &AgentArriveRspBody{
		pb.AgentArriveRsp{
			Errno:   &errno,
			AgentID: &agentID,
		},
	}
	msg = &AgentArriveRsp{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
