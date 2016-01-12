/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import (
	"bytes"
)

// MessageHeader is message's head
type MessageHeader interface {
	Pack(buf *bytes.Buffer) (int, error)
	Unpack(buf []byte) error
}

// MessageBodyer is message's body
type MessageBodyer interface {
	Pack(buf *bytes.Buffer) (int, error)
	Unpack(buf []byte) error
}

// Messager is a interface of proto message
type Messager interface {
	Head() MessageHeader
	Body() MessageBodyer
	Pack(buf *bytes.Buffer) (int, error)
	Unpack(buf []byte) error
}
