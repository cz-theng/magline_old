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
	//DiscardHeadLen is length of Request message
	DiscardHeadLen = 0
)

//DiscardHead is head of message Discard
type DiscardHead struct {
}

// Pack is implement of MessageHeader
func (h *DiscardHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *DiscardHead) Unpack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	return nil
}

//DiscardBody is body of message Discard
type DiscardBody struct {
}

// Pack is implement of MessageBodyer
func (b *DiscardBody) Pack(buf *bytes.Buffer) (err error) {
	if buf == nil {
		err = proto.ErrBufNil
		return
	}
	return
}

// Unpack is implement of MessageBodyer
func (b *DiscardBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

//Discard is Discard message
type Discard struct {
	message.Message
}

//NewDiscard create a Discard
func NewDiscard() (msg *Discard) {
	head := &DiscardHead{}
	body := &DiscardBody{}
	msg = &Discard{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return msg
}
