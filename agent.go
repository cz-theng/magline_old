//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"fmt"
	"github.com/cz-it/magline/proto/message"
)

//Agent is a client object
type Agent struct {
	line        *Line
	id          uint32
	rope        *Rope
	nodeMsgChan chan message.Messager
	knotMsgChan chan message.Messager
}

func (ag *Agent) deinit() {

}

//Init init a agent
func (ag *Agent) Init(line *Line, rope *Rope) error {
	ag.line = line
	ag.rope = rope
	ag.nodeMsgChan = make(chan message.Messager)
	ag.knotMsgChan = make(chan message.Messager)
	return nil
}

// ID return's agent's id
func (ag *Agent) ID() uint32 {
	return ag.id
}

// Discard discard an agent
func (ag *Agent) Discard() (err error) {
	err = ag.line.SendDiscard()
	ag.deinit()
	return
}

// Disconnect disconnect an agent
func (ag *Agent) Disconnect() (err error) {
	err = ag.rope.SendQuit(ag.ID())
	ag.deinit()
	return
}

// Confirm send a confirm message to client
func (ag *Agent) Confirm(errno int32) (err error) {
	err = ag.line.SendConfirm(errno)
	return
}

// Arrive send a arrive message to rope
func (ag *Agent) Arrive() (err error) {
	err = ag.rope.SendArrive(ag.ID())
	return
}

// DealNodeMessage deal message from a magnode
func (ag *Agent) DealNodeMessage(data []byte) (err error) {
	err = ag.rope.SendNodeMessage(ag.ID(), data)
	return
}

// DealKnotMessage deal messgae from a magknot
func (ag *Agent) DealKnotMessage(data []byte) (err error) {
	err = ag.line.SendKnotMessage(data)
	return
}

// Serve run a agent context
func (ag *Agent) Serve() {
	for {
		var msg message.Messager
		select {
		case msg = <-ag.nodeMsgChan:
			fmt.Println("node Message")
		case msg = <-ag.knotMsgChan:
			fmt.Println("knot message")
		default:
			fmt.Println("Default", msg)
			break
		}
	}
}

// OutputNode get a node's message
func (ag *Agent) OutputNode() <-chan message.Messager {
	return ag.nodeMsgChan
}

//OutputKnot get a knot's message
func (ag *Agent) OutputKnot() <-chan message.Messager {
	return ag.knotMsgChan
}

//DealMessage send a message to agent
func (ag *Agent) DealMessage(msg message.Messager) {

}
