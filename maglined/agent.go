//Package maglined is a daemon process for connection layer

/**
* Author: CZ cz.theng@gmail.com
 */
package main

import (
	"github.com/cz-it/magline/maglined/proto"
	"github.com/cz-it/magline/maglined/utils"
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
func (ag *Agent) DealConnReq(req proto.Requester) (err error) {
	utils.Logger.Info("Deal New Agent[%d]'s Connection ", ag.id)
	if ag.lane == nil {
		utils.Logger.Info("There is no magknot")
		return
	}
	ag.lane.AddAgent(ag)
	ag.lane.SendNewAgent(ag.id)
	return
}

//DealNewAgentRsp deal a new agent's response for knot
func (ag *Agent) DealNewAgentRsp() (err error) {
	utils.Logger.Info("Agent Confirm New Agent ID: %d", ag.id)
	rsp := &Response{}
	rsp.Init()
	rsp.CMD = 1 //proto.MNCMDRspConn
	rsp.AgentID = ag.id
	ag.conn.SendResponse(rsp)
	return
}

//DealNodeMsg deal a message form node
func (ag *Agent) DealNodeMsg(data []byte) (err error) {
	if ag.lane == nil {
		utils.Logger.Error("Agent %d 's lane is nil ", ag.ID())
	}
	err = ag.lane.SendNodeMsg(ag.ID(), data)
	if err != nil {
		utils.Logger.Error("Send to Node %d error %s", ag.ID(), err.Error())
	}
	return
}

//Send2Node send message to node
func (ag *Agent) Send2Node(data []byte) (err error) {
	utils.Logger.Debug("Send data %s to node %d", string(data), ag.id)
	rsp := &Response{}
	rsp.Init()
	rsp.CMD = 1 //proto.MNCMDMsgKnot
	rsp.AgentID = ag.id
	rsp.Body = data
	ag.conn.SendResponse(rsp)
	return
}

// DealRequest deal a client's request
func (ag *Agent) DealRequest(req proto.Requester) (err error) {
	utils.Logger.Info("Deal a Client Request! with cmd %d", req.CMD)
	/*
		if req.CMD == node.MNCMDReqConn {
			err = ag.DealConnReq(req)
		} else if req.CMD == proto.MNCMDMsgNode {
			err = ag.DealNodeMsg(req.Body)
		}
		if err != nil {
			utils.Logger.Error("Deal Request Error %v", req)
		}
	*/
	return
}
