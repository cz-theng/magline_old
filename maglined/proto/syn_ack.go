/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import (
	"bytes"
	"encoding/binary"
)

const (
	//SYNLen is length of SYN message
	SYNLen = 6
	//ACKLen is length of ACK message
	ACKLen = 4
)

//NewSYN create a syn
func NewSYN() *SYN {
	return nil
}

//SYNHead is head of message SYN
type SYNHead struct {
	Protobuf uint16
	Channel  uint16
	Crypto   uint16
}

// Pack is implement of MessageHeader
func (h *SYNHead) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *SYNHead) Unpack(buf []byte) error {
	return nil
}

//SYNBody is body of message SYN
type SYNBody struct {
}

// Pack is implement of MessageBodyer
func (h *SYNBody) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageBodyer
func (h *SYNBody) Unpack(buf []byte) error {
	return nil
}

// SYN is syn
type SYN struct {
	head *SYNHead
	body *SYNBody
}

// Head is implement of Messager
func (syn *SYN) Head() MessageHeader {
	return syn.head
}

// Body is implement of Messager
func (syn *SYN) Body() MessageBodyer {
	return syn.body
}

// Pack is implement of Messager
func (syn *SYN) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of Messager
func (syn *SYN) Unpack(buf []byte) error {
	return nil
}

// UnpackSYN unpack a syn message from buf
func UnpackSYN(buf []byte) (syn *SYN, err error) {
	if buf == nil {
		err = ErrBufNil
		return
	}
	if len(buf) != SYNLen {
		err = ErrBufLen
		return
	}
	head := new(SYNHead)
	head.Protobuf = binary.LittleEndian.Uint16(buf[:2])
	head.Channel = binary.LittleEndian.Uint16(buf[2:4])
	head.Crypto = binary.LittleEndian.Uint16(buf[4:6])
	syn = &SYN{
		head: head,
		body: nil}
	return
}

//ACK is ACK Message
type ACK struct {
	head *ACKHead
	body *ACKBody
}

// Head is implement of Messager
func (ack *ACK) Head() MessageHeader {
	return ack.head
}

// Body is implement of Messager
func (ack *ACK) Body() MessageBodyer {
	return ack.body
}

// Unpack is implement of Messager
func (ack *ACK) Unpack(buf []byte) error {
	return nil
}

// ACKBody is ACK message's body
type ACKBody struct {
}

// Pack is implement of MessageHeader
func (h *ACKBody) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *ACKBody) Unpack(buf []byte) error {
	return nil
}

// ACKHead is ACK message's head
type ACKHead struct {
	Channel uint16
	Crypto  uint16
}

// Pack is implement of MessageHeader
func (h *ACKHead) Pack(buf *bytes.Buffer) (length int, err error) {
	return
}

// Unpack is implement of MessageHeader
func (h *ACKHead) Unpack(buf []byte) error {
	return nil
}

//NewACK create a ACK message
func NewACK(channel uint16, crypto uint16) (msg Messager) {
	head := &ACKHead{
		Channel: channel,
		Crypto:  crypto,
	}
	body := &ACKBody{}
	msg = &ACK{
		head: head,
		body: body,
	}
	return
}

// Pack pack ack to buffer
func (ack *ACK) Pack(buf *bytes.Buffer) (length int, err error) {
	// should check buf's len
	if buf.Cap()-buf.Len() < ACKLen {
		err = ErrBufLen
		return
	}
	binary.Write(buf, binary.LittleEndian, ack.head.Channel)
	binary.Write(buf, binary.LittleEndian, ack.head.Crypto)
	length = ACKLen
	return
}
