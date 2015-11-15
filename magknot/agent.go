/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"container/list"
	"github.com/cz-it/magline/maglined/proto"
	"net"
	"sync"
)

type Agent struct {
	ID       uint32
	conn     *net.UnixConn
	readBuf  []byte
	mtx      sync.Mutex
	messages *list.List
}

func (ag *Agent) Init() {
	ag.messages = list.New()
}

func (ag *Agent) PushMessage(msg *Message) (err error) {
	ag.mtx.Lock()
	ag.messages.PushBack(msg)
	ag.mtx.Unlock()
	return
}

func (ag *Agent) PopMessage() (msg *Message, err error) {
	msgElem := ag.messages.Front()
	if msgElem == nil {
		msg = nil
		err = ErrEmptyMessage
	}
	ag.mtx.Lock()
	msg = msgElem.Value.(*Message)
	ag.messages.Remove(msgElem)
	ag.mtx.Unlock()
	return
}

func (ag *Agent) Send(buf []byte) (err error) {
	msg := proto.KnotMessage{ReadBuf: ag.readBuf}
	err = msg.RecvAndUnpack(ag.conn)
	if err != nil {
		println("Recv And Unpack Error")
		return
	}
	return
}

func (ag *Agent) Recv() (msg *Message, err error) {
	msg, err = ag.PopMessage()
	return
}

func (ag *Agent) Close() (err error) {
	return
}
