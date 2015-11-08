package maglined

/**
* Agent manager
 */

import (
	"errors"
	"math"
	"sync"

	"github.com/cz-it/magline/maglined/proto"
)

var (
	ENewAgent    = errors.New("New a Agent Error!")
	EREMOVE_TYPE = errors.New("Remove from List Error!")
	EINDEX       = errors.New("A Invalied Agent Index!")
	EIDLE_AGENT  = errors.New("It is a Idle Agent!")
	ENOAGENT     = errors.New("Don't Have Such a Agent")
)

type AgentMgr struct {
	mtx     sync.Mutex
	agents  map[uint32]*Agent
	idGuard uint32
}

func NewAgentMgr(maxAgents int) (agentMgr *AgentMgr, err error) {
	agentMgr = new(AgentMgr)
	err = agentMgr.Init(maxAgents)
	if err != nil {
		err = ENewAgent
		return
	}
	return
}

var agentMgr *AgentMgr

func (am *AgentMgr) FindAgent(id uint32) (agent *Agent, err error) {
	if a, ok := am.agents[id]; ok {
		err = nil
		agent = a
	}
	err = ENOAGENT
	return
}

func DealNewAgent(conn *Connection, req *Request) (err error) {
	Logger.Info("Deal a New Agent")
	agt, err := agentMgr.Alloc()
	rsp := &Response{
		CMD:     proto.MN_CMD_RSP_CONN,
		AgentID: uint32(agt.ID()),
		Body:    nil,
	}
	conn.SendResponse(rsp)
	return nil
}

func (am *AgentMgr) Init(maxAgents int) error {
	am.mtx.Lock()
	defer am.mtx.Unlock()
	am.idGuard = uint32(math.Exp2(20)) | 512<<21
	am.agents = make(map[uint32]*Agent)
	return nil
}

func (am *AgentMgr) Alloc() (agent *Agent, err error) {
	am.mtx.Lock()
	defer am.mtx.Unlock()
	id := am.tickGuard()
	agent = &Agent{id: id}
	am.agents[id] = agent
	return
}

func (am *AgentMgr) tickGuard() uint32 {
	am.idGuard++
	return am.idGuard
}
