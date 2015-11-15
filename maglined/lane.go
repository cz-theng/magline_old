/**
* Author: CZ cz.theng@gmail.com
 */

package maglined

import (
	"github.com/cz-it/magline/maglined/proto"
	"net"
	"sync"
)

const (
	LANE_READBUF_LEN  = 1024 * 100
	LANE_WRITEBUF_LEN = 1024 * 100
)

type Lane struct {
	RWC     *net.UnixConn
	agents  map[uint32]*Agent
	ReadBuf []byte
	seq     uint32
	mtx     sync.Mutex
}

func (l *Lane) AddAgent(agent *Agent) (err error) {
	l.agents[agent.ID()] = agent
	return
}
func (l *Lane) SendNewAgent(id uint32) (err error) {
	Logger.Info("Get New Agent with ID:%d", id)
	msg := &proto.KnotMessage{}
	msg.CMD = proto.MK_CMD_REQ_NEWAGENT
	msg.Length = 0
	msg.AgentID = id
	msg.Seq = l.TickSeq()
	msg.PackAndSend(nil, l.RWC)
	return
}

func (l *Lane) Init() (err error) {
	l.ReadBuf = make([]byte, LANE_READBUF_LEN)
	l.seq = 0
	l.agents = make(map[uint32]*Agent)
	return
}

func (l *Lane) ReadMsg() (msg *proto.KnotMessage, err error) {
	msg = &proto.KnotMessage{ReadBuf: l.ReadBuf}
	err = msg.RecvAndUnpack(l.RWC)
	if err != nil {
		msg = nil
		return
	}
	return
}

func (l *Lane) TickSeq() uint32 {
	l.mtx.Lock()
	l.seq += 1
	l.mtx.Unlock()
	return l.seq
}

func (l *Lane) DealConnReq(msg *proto.KnotMessage) (err error) {
	Logger.Info("Lane dipatch a new Connection Request from Knot[]")
	rsp := &proto.KnotMessage{}
	rsp.CMD = proto.MK_CMD_RSP_CONN
	rsp.Length = 0
	rsp.Seq = l.TickSeq()
	rsp.PackAndSend(nil, l.RWC)
	return
}

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

func (l *Lane) DealMsgK2N(msg *proto.KnotMessage) (err error) {
	Logger.Debug("Send Message to Node by agent : %d", msg.AgentID)
	agent, ok := l.agents[msg.AgentID]
	if !ok {
		return
	}
	err = agent.Send2Node(msg.Body())
	return
}

func (l *Lane) SendNodeMsg(id uint32, data []byte) (err error) {
	Logger.Debug("Get Message form Node by agent id:", 0)
	rsp := &proto.KnotMessage{}
	rsp.CMD = proto.MK_CMD_MSG_N2K
	rsp.Length = uint32(len(data))
	rsp.Seq = l.TickSeq()
	rsp.AgentID = id
	rsp.PackAndSend(data, l.RWC)
	return

}
func (l *Lane) Serve() {
	for {
		// deal timeout
		msg, err := l.ReadMsg()
		if err != nil {
			Logger.Error("Connection Read Request Error:%s", err.Error())
			continue
		}
		if msg.CMD == proto.MK_CMD_REQ_CONN {
			l.DealConnReq(msg)
		} else if msg.CMD == proto.MK_CMD_RSP_NEWAGENT {
			l.DealNewAgentRsp(msg)
		} else if msg.CMD == proto.MK_CMD_MSG_K2N {
			l.DealMsgK2N(msg)
		} else {

		}
		Logger.Debug("get message: %v", msg.CMD)
	}
}
