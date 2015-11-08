package maglined

/**
* Agent.
 */
import (
	"github.com/cz-it/magline/maglined/proto"
)

//Agent is a client object
type Agent struct {
	conn *Connection
	id   uint32
}

// ID return's agent's id
func (ag *Agent) ID() uint32 {
	return ag.id
}

//DealConnReq deal connnection reqeuest
func (ag *Agent) DealConnReq(req *Request) (err error) {
	rsp := new(Response)
	rsp.Init()
	rsp.AgentID = ag.ID()
	rsp.CMD = proto.MN_CMD_RSP_CONN
	ag.conn.SendResponse(rsp)
	return
}

// DealRequest deal a client's request
func (ag *Agent) DealRequest(req *Request) (err error) {
	Logger.Info("Deal a Client Request! with cmd %d", req.CMD)
	if req.CMD == proto.MN_CMD_REQ_CONN {
		ag.DealConnReq(req)
	}
	return nil
}
