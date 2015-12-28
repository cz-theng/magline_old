/**
* Author: CZ cz.theng@gmail.com
 */

package node

import (
	"github.com/cz-it/magline/maglined/proto"
)

//NewSYN create a syn
func NewSYN() *SYN {
	syn = new(SYN)
	return syn
}

// SYN is syn
type SYN struct {
	proto.Request
	Seq     uint32
	Buffers uint16
	Channel uint16
	Crypto  uint16
}

// Unpack unpack data to syn
func (syn *SYN) Unpack(data []byte) error {

}
