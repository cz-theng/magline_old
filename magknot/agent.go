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

//Agent is a agent object
type Agent struct {
	ID       uint32
	conn     *net.UnixConn
	readBuf  []byte
	mtx      sync.Mutex
	messages *list.List
}

//Init is initialize
func (ag *Agent) Init() {
	ag.messages = list.New()
}

func (ag *Agent) pushMessage(msg *Message) (err error) {
	ag.mtx.Lock()
	defer ag.mtx.Unlock()
	ag.messages.PushBack(msg)
	return
}

func (ag *Agent) popMessage() (msg *Message, err error) {
	if ag.messages == nil {
		msg = nil
		err = ErrEmptyMessage
		return
	}
	msgElem := ag.messages.Front()
	if msgElem == nil {
		msg = nil
		err = ErrEmptyMessage
		return
	}
	ag.mtx.Lock()
	defer ag.mtx.Unlock()
	msg = msgElem.Value.(*Message)
	ag.messages.Remove(msgElem)
	return
}

//Send will send data
func (ag *Agent) Send(buf []byte) (err error) {
	msg := proto.KnotMessage{
		Magic:   0x01,
		Version: 0x01,
		CMD:     proto.MKCMDMsgK2N,
		Seq:     0x01,
		AgentID: ag.ID,
		Length:  uint32(len(buf)),
	}
	err = msg.PackAndSend(buf, ag.conn)
	if err != nil {
		println("Send And Pack Error")
		return
	}
	return
}

//Recv will recv data
func (ag *Agent) Recv() (msg *Message, err error) {
	msg, err = ag.popMessage()
	return
}

//Close close the connection
func (ag *Agent) Close() (err error) {
	return
}
