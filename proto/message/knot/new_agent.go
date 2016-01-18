/**
* Author: CZ cz.theng@gmail.com
 */

package knot

import (
	"bytes"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/proto/knot/pb"
	protobuf "github.com/golang/protobuf/proto"
)

const ()

//NewAgentReqHead is head of message NewAgentReq
type NewAgentReqHead struct {
	head pb.NewAgentReq
}

// Pack is implement of MessageHeader
func (h *NewAgentReqHead) Pack(buf *bytes.Buffer) (length int, err error) {
	data, err := protobuf.Marshal(&h.head)
	if err != nil {
		return
	}
	buf.Write(data)
	length = len(data)
	return
}

// Unpack is implement of MessageHeader
func (h *NewAgentReqHead) Unpack(buf []byte) error {
	return nil
}

//NewAgentReqBody is body of message NewAgentReq
type NewAgentReqBody struct {
}

// Pack is implement of MessageBodyer
func (h *NewAgentReqBody) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageBodyer
func (h *NewAgentReqBody) Unpack(buf []byte) error {
	return nil
}

// NewAgentReq is NewAgentReq
type NewAgentReq struct {
	head *NewAgentReqHead
	body *NewAgentReqBody
}

// Head is implement of Messager
func (nar *NewAgentReq) Head() proto.MessageHeader {
	return nar.head
}

// Body is implement of Messager
func (nar *NewAgentReq) Body() proto.MessageBodyer {
	return nar.body
}

// Pack is implement of Messager
func (nar *NewAgentReq) Pack(buf *bytes.Buffer) (length int, err error) {
	l, err := nar.head.Pack(buf)
	if err != nil {
		length += l
		return
	}
	length += l
	l, err = nar.body.Pack(buf)
	if err != nil {
		length += l
		return
	}
	length += l
	return
}

// Unpack is implement of Messager
func (nar *NewAgentReq) Unpack(buf []byte) error {
	return nil
}

//NewNewAgentReq create a NewAgentReq Message
func NewNewAgentReq(uuid []byte) (msg proto.Messager) {
	head := &NewAgentReqHead{
		head: pb.NewAgentReq{
			Uuid: uuid,
		},
	}
	body := &NewAgentReqBody{}
	msg = &NewAgentReq{
		head: head,
		body: body,
	}
	return
}

//NewAgentRsqHead is head of message NewAgentRsq
type NewAgentRsqHead struct {
}

// Pack is implement of MessageHeader
func (h *NewAgentRsqHead) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *NewAgentRsqHead) Unpack(buf []byte) error {
	return nil
}

//NewAgentRsqBody is body of message NewAgentReq
type NewAgentRsqBody struct {
}

// Pack is implement of MessageBodyer
func (h *NewAgentRsqBody) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageBodyer
func (h *NewAgentRsqBody) Unpack(buf []byte) error {
	return nil
}

// NewAgentRsq is NewAgentRsq
type NewAgentRsq struct {
	head *NewAgentRsqHead
	body *NewAgentRsqBody
}

// Head is implement of Messager
func (nar *NewAgentRsq) Head() proto.MessageHeader {
	return nar.head
}

// Body is implement of Messager
func (nar *NewAgentRsq) Body() proto.MessageBodyer {
	return nar.body
}

// Pack is implement of Messager
func (nar *NewAgentRsq) Pack(buf *bytes.Buffer) (length int, err error) {
	l, err := nar.head.Pack(buf)
	if err != nil {
		length += l
		return
	}
	length += l
	l, err = nar.body.Pack(buf)
	if err != nil {
		length += l
		return
	}
	length += l
	return
}

// Unpack is implement of Messager
func (nar *NewAgentRsq) Unpack(buf []byte) error {
	return nil
}

// UnpackNewAgentRsq unpack a NewAgentRsq message from buf
func UnpackNewAgentRsq(buf []byte) (nar *NewAgentRsq, err error) {
	return
}
