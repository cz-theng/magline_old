//Package maglined is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package main

import (
	"math"
	"sync"

	"github.com/cz-it/magline/maglined/proto"
	"github.com/cz-it/magline/maglined/utils"
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

//DealNewAgent deal a new agent
func DealNewAgent(conn *Connection, req proto.Requester) (err error) {
	utils.Logger.Info("Deal a New Agent")
	agt, err := agentMgr.Alloc()
	rsp := &Response{
		CMD:     1, //proto.MNCMDRspConn,
		AgentID: uint32(agt.ID()),
		Body:    nil,
	}
	conn.SendResponse(rsp)
	return nil
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
