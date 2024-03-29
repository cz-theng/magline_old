//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"math"
	"sync"
)

//AgentMgr is agent manager
type AgentMgr struct {
	mtx     sync.Mutex
	agents  map[uint32]*Agent
	idGuard uint32
}

//NewAgentMgr create a agent manager
func NewAgentMgr(maxAgents int) (agentMgr *AgentMgr, err error) {
	agentMgr = new(AgentMgr)
	err = agentMgr.Init(maxAgents)
	if err != nil {
		err = ErrNewAgent
		return
	}
	return
}

var agentMgr *AgentMgr

//FindAgent find agent from agent manager
func (am *AgentMgr) FindAgent(id uint32) (agent *Agent, err error) {
	if a, ok := am.agents[id]; ok {
		err = nil
		agent = a
		return
	}
	err = ErrNoAgent
	return
}

//Init agnet manager's initialize
func (am *AgentMgr) Init(maxAgents int) error {
	am.mtx.Lock()
	defer am.mtx.Unlock()
	am.idGuard = uint32(math.Exp2(20)) | 512<<21
	am.agents = make(map[uint32]*Agent)
	return nil
}

//Alloc alloc a new agent
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
