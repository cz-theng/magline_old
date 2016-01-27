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

//DealMessage implementation of Connectioner
func (r *Rope) DealMessage(msg message.Messager) (err error) {
	utils.Logger.Info("Deal Message")
	if msg == nil {
		err = ErrArg
		return
	}
	rsp := knot.NewConnRsp([]byte("Lane"))
	r.SendMessage(rsp, 5*time.Second)
	return
}
