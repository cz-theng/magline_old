/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	//"github.com/cz-it/magline/maglined/proto"
	//"io"
	"net"
	"sync"
	//"time"
)

const (
	//LaneReadBufLen is read buffer length
	LaneReadBufLen = 1024 * 100
	//LaneWriteBufLen is write buffer length
	LaneWriteBufLen = 1024 * 100
)

//Lane is backend agent
type Lane struct {
	RWC     *net.UnixConn
	agents  map[uint32]*Agent
	ReadBuf []byte
	seq     uint32
	mtx     sync.Mutex
}

//AddAgent add a new agent
func (l *Lane) AddAgent(agent *Agent) (err error) {
	l.agents[agent.ID()] = agent
	return
}

//SendNewAgent send a new agent's request
func (l *Lane) SendNewAgent(id uint32) (err error) {
	/*
		Logger.Info("Get New Agent with ID:%d", id)
		msg := &proto.KnotMessage{}
		msg.CMD = proto.MKCMDReqNewAgent
		msg.Length = 0
		msg.AgentID = id
		msg.Seq = l.tickSeq()
		msg.PackAndSend(nil, l.RWC)
	*/
	return
}

//Init is initialize
func (l *Lane) Init() (err error) {
	l.ReadBuf = make([]byte, LaneReadBufLen)
	l.seq = 0
	l.agents = make(map[uint32]*Agent)
	return
}

//ReadMsg read a message
/*
func (l *Lane) ReadMsg() (msg *proto.KnotMessage, err error) {
	msg = &proto.KnotMessage{ReadBuf: l.ReadBuf}
	err = msg.RecvAndUnpack(l.RWC)
	if err != nil && err != io.EOF {
		msg = nil
		return
	}
	return
}
*/

func (l *Lane) tickSeq() uint32 {
	l.mtx.Lock()
	l.seq++
	l.mtx.Unlock()
	return l.seq
}

//DealConnReq deal a request
/*
func (l *Lane) DealConnReq(msg *proto.KnotMessage) (err error) {
	Logger.Info("Lane dipatch a new Connection Request from Knot[]")
	rsp := &proto.KnotMessage{}
	rsp.CMD = proto.MKCMDRspConn
	rsp.Length = 0
	rsp.Seq = l.tickSeq()
	rsp.PackAndSend(nil, l.RWC)
	return
}
*/

/*
//DealNewAgentRsp deal a new agent's response
func (l *Lane) DealNewAgentRsp(msg *proto.KnotMessage) (err error) {
	Logger.Info("New Agent success , Reponse with agent id:%d", msg.AgentID)
	agent, ok := l.agents[msg.AgentID]
	if !ok {
		Logger.Error("Cant find agent %d in lane", msg.AgentID)
		return
	}
	agent.DealNewAgentRsp()
	return
}
*/
/*
//DealMsgK2N deal message from knot to node
func (l *Lane) DealMsgK2N(msg *proto.KnotMessage) (err error) {
	Logger.Debug("Send Message to Node by agent : %d", msg.AgentID)
	agent, ok := l.agents[msg.AgentID]
	if !ok {
		Logger.Error("No such Agent %d", msg.AgentID)
		return
	}
	err = agent.Send2Node(msg.Body())
	return
}
*/

//SendNodeMsg send message to node
func (l *Lane) SendNodeMsg(id uint32, data []byte) (err error) {
	/*
		Logger.Debug("Get Message form Node by agent id:%d", id)
		msg := &proto.KnotMessage{}
		msg.CMD = proto.MKCMDMsgN2K
		msg.Length = uint32(len(data))
		msg.Seq = l.tickSeq()
		msg.AgentID = id
		msg.PackAndSend(data, l.RWC)
		Logger.Debug("send node msg with length %d,data:%s", msg.Length, string(data))
	*/
	return

}

//Serve server the server
func (l *Lane) Serve() {
	for {
		// deal timeout
		/*
			msg, err := l.ReadMsg()
			if err != nil {
				time.Sleep(200 * time.Millisecond)
				if err != io.EOF {
					Logger.Error("Connection Read Request Error:%s", err.Error())
				}
				continue
			}
			if msg.CMD == proto.MKCMDReqConn {
				l.DealConnReq(msg)
			} else if msg.CMD == proto.MKCMDRspNewAgent {
				l.DealNewAgentRsp(msg)
			} else if msg.CMD == proto.MKCMDMsgK2N {
				l.DealMsgK2N(msg)
			} else {

			}
			Logger.Debug("get message: %v", msg.CMD)
		*/
	}
}
