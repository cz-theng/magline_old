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
	conn        *Connection
	id          uint32
	rope        *Rope
	nodeMsgChan chan message.Messager
	knotMsgChan chan message.Messager
}

//Init init a agent
func (ag *Agent) Init(conn *Connection, rope *Rope, id uint32) {
	ag.conn = conn
	ag.id = id
	ag.rope = rope
	ag.nodeMsgChan = make(chan message.Messager)
	ag.knotMsgChan = make(chan message.Messager)
}

// ID return's agent's id
func (ag *Agent) ID() uint32 {
	return ag.id
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
