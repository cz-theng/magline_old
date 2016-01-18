/**
* Author: CZ cz.theng@gmail.com
 */

package message

import (
	"bytes"
)

// Header is message's head
type Header interface {
	Pack(buf *bytes.Buffer) error
	Unpack(buf *bytes.Buffer) error
}

// Bodyer is message's body
type Bodyer interface {
	Pack(buf *bytes.Buffer) error
	Unpack(buf *bytes.Buffer) error
}

// Messager is a interface of proto message
type Messager interface {
	Pack(buf *bytes.Buffer) error
	Unpack(buf *bytes.Buffer) error
}

// Message is base class of Messager
type Message struct {
	Head Header
	Body Bodyer
}

// Pack is implement of Messager
func (m *Message) Pack(buf *bytes.Buffer) (err error) {
	err = m.Head.Pack(buf)
	if err != nil {
		return
	}
	err = m.Body.Pack(buf)
	if err != nil {
		return
	}
	return nil
}

// Unpack is implement of Messager
func (m *Message) Unpack(buf *bytes.Buffer) (err error) {
	err = m.Head.Unpack(buf)
	if err != nil {
		return
	}
	err = m.Body.Unpack(buf)
	if err != nil {
		return
	}
	return nil
}
