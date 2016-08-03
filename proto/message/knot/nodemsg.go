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

//NodeMsgHead is head of message
type NodeMsgHead struct {
}

// Pack is implement of MessageHeader
func (h *NodeMsgHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *NodeMsgHead) Unpack(buf *bytes.Buffer) (err error) {
	return
}

//NodeMsgBody is body of message SYN
type NodeMsgBody struct {
	pb.NodeMsg
}

// Pack is implement of MessageBodyer
func (b *NodeMsgBody) Pack(buf *bytes.Buffer) (err error) {
	buffer, err := protobuf.Marshal(&b.NodeMsg)
	if err != nil {
		return
	}
	_, err = buf.Write(buffer)
	return
}

// Unpack is implement of MessageBodyer
func (b *NodeMsgBody) Unpack(buf *bytes.Buffer) (err error) {
	buffer := buf.Next(buf.Len())
	err = protobuf.Unmarshal(buffer, &b.NodeMsg)
	return
}

// NodeMsg is NodeMsg Message from knot
type NodeMsg struct {
	message.Message
}

//NewNodeMsg new and init a NodeMsg message
func NewNodeMsg(agentID uint32, payload []byte) (msg *NodeMsg) {
	head := &NodeMsgHead{}
	body := &NodeMsgBody{
		pb.NodeMsg{
			AgentID: &agentID,
			Payload: payload,
		},
	}
	msg = &NodeMsg{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
