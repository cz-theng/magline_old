/**
* Author: CZ cz.theng@gmail.com
 */

package magline

import (
	"github.com/cz-it/magline/proto/message"
	"github.com/cz-it/magline/proto/message/knot"
	"github.com/cz-it/magline/utils"
	"sync"
	"time"
)

const (
	//RopeReadBufLen is read buffer length
	RopeReadBufLen = 1024 * 100
	//RopeWriteBufLen is write buffer length
	RopeWriteBufLen = 1024 * 100
)

//Rope is backend agent
type Rope struct {
	Connection
	agents map[uint32]*Agent
	mtx    sync.Mutex
}

//AddAgent add a new agent
func (r *Rope) AddAgent(agent *Agent) (err error) {
	r.agents[agent.ID()] = agent
	return
}

//Init is initialize
func (r *Rope) Init() (err error) {
	err = r.Connection.Init(RopeReadBufLen, RopeWriteBufLen)
	if err != nil {
		return
	}
	r.agents = make(map[uint32]*Agent)
	r.dealMessage = r.DealMessage
	return
}

// SendNodeMessage send a magnode's message to magknot
func (r *Rope) SendNodeMessage(agentID uint32, data []byte) (err error) {
	msg := knot.NewNodeMsg(agentID, data)
	err = r.SendMessage(msg, 5*time.Second)
	return
}

//DealMessage implementation of Connectioner
func (r *Rope) DealMessage(msg message.Messager) (err error) {
	if msg == nil {
		err = ErrArg
		return
	}
	switch m := msg.(type) {
	case *knot.ConnReq:
		utils.Logger.Info("Get A Conn Request Message")
		err = r.dealConnReq(m)
		if err != nil {
			utils.Logger.Error("deal ConnReq error %s", err.Error())
		}

	case *knot.AgentArriveRsp:
		pbm := m.Body.(*knot.AgentArriveRspBody)
		utils.Logger.Debug("Got a AgentArriveRsp ")
		err = r.dealAgentArriveRsp(pbm)
		if err != nil {
			utils.Logger.Error("deal AgentArriveRsp error %s", err.Error())
		}
	default:
		utils.Logger.Error("Unknown Message type")
	}
	return
}

// SendArrive got a new agent
func (r *Rope) SendArrive(agentID uint32) (err error) {
	msg := knot.NewAgentArriveReq(agentID)
	err = r.SendMessage(msg, 5*time.Second)
	return
}

func (r *Rope) dealConnReq(connreq *knot.ConnReq) (err error) {
	rsp := knot.NewConnRsp([]byte("Lane"))
	err = r.SendMessage(rsp, 5*time.Second)
	return
}

func (r *Rope) dealAgentArriveRsp(agentArriveRsp *knot.AgentArriveRspBody) (err error) {
	if agentArriveRsp == nil {
		utils.Logger.Error("agentArriveRsp is null")
		return ErrArg
	}
	if agent, ok := r.agents[*agentArriveRsp.AgentID]; ok {
		err = agent.Confirm(*agentArriveRsp.Errno)
	} else {
		err = ErrArg
	}
	return
}
