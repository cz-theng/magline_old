/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import ()

// MessageHeader is message's head
type MessageHeader interface {
}

// MessageBodyer is message's body
type MessageBodyer interface {
}

// Messager is a interface of proto message
type Messager interface {
	Head() MessageHeader
	Body() MessageBodyer
}
