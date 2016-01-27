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

//ConnRspHead is head of message SYN
type ConnRspHead struct {
	pb.ConnRsp
}

// Pack is implement of MessageHeader
func (h *ConnRspHead) Pack(buf *bytes.Buffer) (err error) {
	buffer, err := protobuf.Marshal(&h.ConnRsp)
	if err != nil {
		return
	}
	_, err = buf.Write(buffer)
	return
}

// Unpack is implement of MessageHeader
func (h *ConnRspHead) Unpack(buf *bytes.Buffer) (err error) {
	buffer := buf.Next(buf.Len())
	err = protobuf.Unmarshal(buffer, &h.ConnRsp)
	return
}

//ConnRspBody is body of message SYN
type ConnRspBody struct {
}

// Pack is implement of MessageBodyer
func (h *ConnRspBody) Pack(buf *bytes.Buffer) (err error) {
	return
}

// Unpack is implement of MessageBodyer
func (h *ConnRspBody) Unpack(buf *bytes.Buffer) error {
	return nil
}

// ConnRsp is ConnRsp Message from knot
type ConnRsp struct {
	message.Message
}

//NewConnRsp new and init a ConnRsp message
func NewConnRsp(accessKey []byte) (msg *ConnRsp) {
	head := &ConnRspHead{
		pb.ConnRsp{},
	}
	body := &ConnRspBody{}
	msg = &ConnRsp{
		message.Message{
			Head: head,
			Body: body,
		},
	}
	return
}
