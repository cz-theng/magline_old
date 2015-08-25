package agent
/**
* Agent manager 
*/

import (
	"errors"
	"sync"
	"container/list"
)

var (
	ENewAgent = errors.New("New a Agent Error!")
	EREMOVE_TYPE = errors.New("Remove from List Error!")
)

type AgentMgr struct {
	mtx sync.Mutex
	idleConns list.List
	AgentArray []*Agent
}

var agentMgr *AgentMgr

func InitAgentMgr(size int) (err error){
	agentMgr = new(AgentMgr)
	defer func (){
		err = ENewAgent
	}()

	agents := make([]*Agent,size)
	for i:=0; i<size; i++ {
		agents[i]= &Agent{}
	}
	agentMgr.Init(agents[:])
	return 
}


func (am *AgentMgr) Init(dataSplit []*Agent) (error) {
	am.mtx.Lock()
	defer am.mtx.Unlock()
	am.idleConns.Init()
	for i:=0; i<len(dataSplit); i++{
		am.idleConns.PushBack(i)
	}
	am.AgentArray = dataSplit
	return nil
}

func (am *AgentMgr) Recycle(agent *Agent) (error) {
	am.mtx.Lock()
	defer am.mtx.Unlock()
	id := agent.Index()
	am.idleConns.PushBack(id)
	return nil
}

func (am *AgentMgr) Alloc() (agent *Agent, err error) {
	if am.idleConns.Len() < 1 {
		return agent, nil
	}
	am.mtx.Lock()
	defer am.mtx.Unlock()
	top := am.idleConns.Front()
	if index,ok := am.idleConns.Remove(top).(int); ok {
		agent = am.AgentArray[index]
	} else {
		err = EREMOVE_TYPE
	}
	return 
}












