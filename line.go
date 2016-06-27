//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"container/list"
	"github.com/cz-it/magline/proto"
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
	svr *Server
	Connection
	ID      int
	Elem    *list.Element
	AgentID uint32
	agent   *Agent

	protobuf uint16
	channel  uint16
	crypto   uint16
}

//Init is initialize
func (l *Line) Init(svr *Server) error {
	if svr == nil {
		return ErrArg
	}
	err := l.Connection.Init(ReadBufSize, WriteBufSize)
	if err != nil {
		utils.Logger.Error("Init Connection Error with %s", err.Error())
		return err
	}
	l.svr = svr
	l.protobuf = 0
	l.channel = 0
	l.crypto = 0
	l.dealMessage = l.DealMessage
	return nil
}

//NewLine create and init a line
func NewLine() (l *Line, err error) {
	l = new(Line)
	return
}

// SendConfirm send a confirm message to client
func (l *Line) SendConfirm(errno int32) (err error) {
	msg := node.NewConfirm(proto.ErrNO(errno))
	err = l.SendMessage(msg, 5*time.Second)
	if err != nil {
		utils.Logger.Error("Send confirm error %s", err.Error())
	}
	return
}

// SendKnotMessage send a knot message to magnode
func (l *Line) SendKnotMessage(data []byte) (err error) {
	utils.Logger.Debug("send a knot message %s", string(data))
	msg := node.NewKnotMsg(data)
	err = l.SendMessage(msg, 5*time.Second)
	if err != nil {
		utils.Logger.Error("send KnotMessage error %s", err.Error())
	}
	utils.Logger.Debug("send a knot message %s", string(data))
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
	case *node.NodeMsg:
		utils.Logger.Info("Get a Message from Agent")
		err = l.dealNodeMessage(m)
	default:
		utils.Logger.Error("Unknown Message type")
	}
	return
}

func (l *Line) dealNodeMessage(msg *node.NodeMsg) (err error) {
	body := msg.Body.(*node.NodeMsgBody)
	err = l.agent.DealNodeMessage(body.Payload)
	return
}

func (l *Line) dealSYN(syn *node.SYN) (err error) {
	head := syn.Head.(*node.SYNHead)
	utils.Logger.Debug("Deal SYN with protobuf: %d, key: %d, crypto: %d", head.Protobuf, head.Channel, head.Crypto)
	l.protobuf = head.Protobuf
	l.channel = head.Channel
	l.crypto = head.Crypto
	ack := node.NewACK(l.channel, l.crypto)
	err = l.SendMessage(ack, 5*time.Second)
	return
}

func (l *Line) dealSessionReq(sq *node.SessionReq) (err error) {
	utils.Logger.Info("Deal a SessionReq Message")
	agent, err := l.Server.AgentMgr.Alloc()
	if err != nil {
		utils.Logger.Error("AgentManager create agent error!")
		return
	}
	rope, err := l.svr.Backend.Dispatch()
	if err != nil {
		utils.Logger.Error("Dispatch rope error with %s", err.Error())
		// should return send an error message
		return
	}
	err = rope.AddAgent(agent)
	if err != nil {
		utils.Logger.Error("Add Agent Error %s", err.Error())
		return
	}
	err = agent.Init(l, rope)
	if err != nil {
		utils.Logger.Error("Agent init error with %s", err.Error())
		// should send an error message
		return
	}
	l.agent = agent
	err = agent.Arrive()
	if err != nil {
		utils.Logger.Error("rope[%v] send agent arrive error ", rope, err.Error())
	}

	rsp := node.NewSessionRsp(agent.ID())
	err = l.SendMessage(rsp, 5*time.Second)
	if err != nil {
		utils.Logger.Error("l.SendMessage error with %s", err.Error())
		return
	}
	return
}
