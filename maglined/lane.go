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
	//WriteBuf []byte
}

func (l *Lane) AddAgent(agent *Agent) (err error) {
	l.agents[agent.ID()] = agent
	return
}
func (l *Lane) SendNewAgent(id uint32) (err error) {
	msg := &proto.KnotMessage{}
	msg.CMD = proto.MK_CMD_REQ_NEWAGENT
	msg.Length = 0
	msg.Seq = l.TickSeq()
	msg.PackAndSend(l.RWC)
	return
}

func (l *Lane) Init() (err error) {
	l.ReadBuf = make([]byte, LANE_READBUF_LEN)
	//l.WriteBuf = make([]byte, LANE_WRITEBUF_LEN)
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
	rsp := &proto.KnotMessage{}
	rsp.CMD = proto.MK_CMD_RSP_CONN
	rsp.Length = 0
	rsp.Seq = l.TickSeq()
	rsp.PackAndSend(l.RWC)
	return
}

func (l *Lane) DealNewAgentRsp(msg *proto.KnotMessage) (err error) {
	agent, ok := l.agents[msg.AgentID]
	if !ok {
		return
	}
	agent.DealNewAgentRsp()
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
		} else {

		}
		Logger.Debug("get message: %v", msg.CMD)
	}
}
