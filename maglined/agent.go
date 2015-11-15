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
	lane *Lane
}

// ID return's agent's id
func (ag *Agent) ID() uint32 {
	return ag.id
}

//DealConnReq deal connnection reqeuest
func (ag *Agent) DealConnReq(req *Request) (err error) {
	Logger.Info("Deal New Agent[%d]'s Connection ", ag.id)
	ag.lane.AddAgent(ag)
	ag.lane.SendNewAgent(ag.id)
	return
}

func (ag *Agent) DealNewAgentRsp() (err error) {
	Logger.Info("Agent Confirm New Agent ID: %d", ag.id)
	rsp := &Response{}
	rsp.Init()
	rsp.CMD = proto.MN_CMD_RSP_CONN
	rsp.AgentID = ag.id
	ag.conn.SendResponse(rsp)
	return
}
func (ag *Agent) DealNodeMsg(data []byte) (err error) {
	ag.lane.SendNodeMsg(ag.ID(), data)
	return
}

func (ag *Agent) Send2Node(data []byte) (err error) {
	rsp := &Response{}
	rsp.Init()
	rsp.CMD = proto.MN_CMD_MSG_KNOT
	rsp.AgentID = ag.id
	rsp.Body = data
	ag.conn.SendResponse(rsp)
	return
}

// DealRequest deal a client's request
func (ag *Agent) DealRequest(req *Request) (err error) {
	Logger.Info("Deal a Client Request! with cmd %d", req.CMD)
	if req.CMD == proto.MN_CMD_REQ_CONN {
		err = ag.DealConnReq(req)
	} else if req.CMD == proto.MN_CMD_MSG_NODE {
		err = ag.DealNodeMsg(req.Body)
	}
	if err != nil {
		Logger.Error("Deal Request Error %v", req)
	}
	return
}
