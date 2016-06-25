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

//ConnReqHead is head of message
type ConnReqHead struct {
}

// Pack is implement of MessageHeader
func (h *ConnReqHead) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *ConnReqHead) Unpack(buf *bytes.Buffer) (err error) {
	return
}

//ConnReqBody is body of message SYN
type ConnReqBody struct {
	pb.ConnReq
}

// Pack is implement of MessageBodyer
func (h *ConnReqBody) Pack(buf *bytes.Buffer) (err error) {
	buffer, err := protobuf.Marshal(&h.ConnReq)
	if err != nil {
		return
	}
	_, err = buf.Write(buffer)
	return
}

// Unpack is implement of MessageBodyer
func (h *ConnReqBody) Unpack(buf *bytes.Buffer) (err error) {
	buffer := buf.Next(buf.Len())
	err = protobuf.Unmarshal(buffer, &h.ConnReq)
	return
}

// ConnReq is ConnReq Message from knot
type ConnReq struct {
	message.Message
}

//NewConnReq new and init a ConnReq message
func NewConnReq(accessKey []byte) (msg *ConnReq) {
	head := &ConnReqHead{}
	body := &ConnReqBody{
		pb.ConnReq{
			AccessKey: accessKey,
		},
	}
	msg = &ConnReq{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
