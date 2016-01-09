/**
* Author: CZ cz.theng@gmail.com
 */

package node

import (
	"github.com/cz-it/magline/maglined/proto"
)

//NewSYN create a syn
func NewSYN() *SYN {
	return nil
}

//SYNHead is head of message SYN
type SYNHead struct {
	Protobuf uint16
	Key      uint16
	Crypto   uint16
}

//SYNBody is body of message SYN
type SYNBody struct {
}

// SYN is syn
type SYN struct {
	head *SYNHead
	body *SYNBody
}

// Head is implement of proto.Messager
func (syn *SYN) Head() proto.MessageHeader {
	return syn.head
}

// Body is implement of proto.Messager
func (syn *SYN) Body() proto.MessageBodyer {
	return syn.body
}
