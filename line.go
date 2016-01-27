//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"container/list"
	"github.com/cz-it/magline/proto/message"
	"github.com/cz-it/magline/proto/message/node"
	"github.com/cz-it/magline/utils"
	"time"
)

const (
	//ReadBufSize is read buffer size
	ReadBufSize = 10 * 1024

	//WriteBufSize is write buffer size
	WriteBufSize = 10 * 1024
)

//Line is connection object
type Line struct {
	Connection
	ID      int
	Elem    *list.Element
	AgentID uint32

	protobuf uint16
	channel  uint16
	crypto   uint16
}

//Init is initialize
func (l *Line) Init() error {
	err := l.Connection.Init(ReadBufSize, WriteBufSize)
	if err != nil {
		return err
	}
	l.protobuf = 0
	l.channel = 0
	l.crypto = 0
	l.dealMessage = l.DealMessage
	return nil
}

//NewLine create and init a line
func NewLine() (l *Line, err error) {
	l = new(Line)
	err = l.Init()
	return
}

//DealMessage implementation of Connectioner
func (l *Line) DealMessage(msg message.Messager) (err error) {
	if msg == nil {
		err = ErrArg
		return
	}
	switch m := msg.(type) {
	case *node.SYN:
		utils.Logger.Info("Get A SYN Message")
		err = l.dealSYN(m)
	case *node.SessionReq:
		utils.Logger.Info("Get SessionReq ")
		err = l.dealSessionReq(m)
	default:
		utils.Logger.Error("Unknown Message type")
	}
	return
}

func (l *Line) dealSessionReq(sq *node.SessionReq) (err error) {
	utils.Logger.Info("Deal a SessionReq Message")
	agent, err := l.Server.AgentMgr.Alloc()
	if err != nil {
		utils.Logger.Error("AgentManager create agent error!")
		return
	}
	rsp := node.NewSessionRsp(agent.ID())
	err = l.SendMessage(rsp, 5*time.Second)
	return
}

func (l *Line) dealSYN(syn *node.SYN) (err error) {
	head := syn.Head.(*node.SYNHead)
	utils.Logger.Info("Deal SYN with protobuf: %d, key: %d, crypto: %d", head.Protobuf, head.Channel, head.Crypto)
	l.protobuf = head.Protobuf
	l.channel = head.Channel
	l.crypto = head.Crypto
	ack := node.NewACK(l.channel, l.crypto)
	err = l.SendMessage(ack, 5*time.Second)
	return
}
