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

//DiscardAgentHead is head of message
type DiscardAgentHead struct {
}

// Pack is implement of MessageHeader
func (h *DiscardAgentHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *DiscardAgentHead) Unpack(buf *bytes.Buffer) (err error) {
	return
}

//DiscardAgentBody is body of message SYN
type DiscardAgentBody struct {
	pb.DiscardAgent
}

// Pack is implement of MessageBodyer
func (b *DiscardAgentBody) Pack(buf *bytes.Buffer) (err error) {
	buffer, err := protobuf.Marshal(&b.DiscardAgent)
	if err != nil {
		return
	}
	_, err = buf.Write(buffer)
	return
}

// Unpack is implement of MessageBodyer
func (b *DiscardAgentBody) Unpack(buf *bytes.Buffer) (err error) {
	buffer := buf.Next(buf.Len())
	err = protobuf.Unmarshal(buffer, &b.DiscardAgent)
	return
}

// DiscardAgent is DiscardAgent Message from knot
type DiscardAgent struct {
	message.Message
}

//NewDiscardAgent new and init a DiscardAgent message
func NewDiscardAgent(agentID uint32) (msg *DiscardAgent) {
	head := &DiscardAgentHead{}
	body := &DiscardAgentBody{
		pb.DiscardAgent{
			AgentID: &agentID,
		},
	}
	msg = &DiscardAgent{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
