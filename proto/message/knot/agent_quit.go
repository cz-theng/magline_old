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

//AgentQuitHead is head of message
type AgentQuitHead struct {
}

// Pack is implement of MessageHeader
func (h *AgentQuitHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *AgentQuitHead) Unpack(buf *bytes.Buffer) (err error) {
	return
}

//AgentQuitBody is body of message SYN
type AgentQuitBody struct {
	pb.AgentQuit
}

// Pack is implement of MessageBodyer
func (b *AgentQuitBody) Pack(buf *bytes.Buffer) (err error) {
	buffer, err := protobuf.Marshal(&b.AgentQuit)
	if err != nil {
		return
	}
	_, err = buf.Write(buffer)
	return
}

// Unpack is implement of MessageBodyer
func (b *AgentQuitBody) Unpack(buf *bytes.Buffer) (err error) {
	buffer := buf.Next(buf.Len())
	err = protobuf.Unmarshal(buffer, &b.AgentQuit)
	return
}

// AgentQuit is AgentQuit Message from knot
type AgentQuit struct {
	message.Message
}

//NewAgentQuit new and init a AgentQuit message
func NewAgentQuit(agentID uint32) (msg *AgentQuit) {
	head := &AgentQuitHead{}
	body := &AgentQuitBody{
		pb.AgentQuit{
			AgentID: &agentID,
		},
	}
	msg = &AgentQuit{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
