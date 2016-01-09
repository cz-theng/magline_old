//Package maglined is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package main

import (
	"bytes"
	"container/list"
	"github.com/cz-it/magline/maglined/proto"
	"github.com/cz-it/magline/maglined/proto/node"
	"github.com/cz-it/magline/maglined/utils"
	"io"
	"net"
	"time"
)

const (
	//ReadBufSize is read buffer size
	ReadBufSize = 10 * 1024
)

//Connection is connection object
type Connection struct {
	RWC     *net.TCPConn
	ReadBuf *bytes.Buffer
	ID      int
	Elem    *list.Element
	AgentID uint32
	Server  *Server
}

//Init is initialize
func (conn *Connection) Init() error {
	buf := make([]byte, ReadBufSize)
	if buf == nil {
		return ErrNewBuffer
	}
	conn.ReadBuf = bytes.NewBuffer(buf)
	return nil
}

//RecvMessage Recv a request message
func (conn *Connection) RecvMessage(timeout time.Duration) (msg proto.Messager, err error) {
	var frameHead *proto.FrameHead
	priBufLen := conn.ReadBuf.Len()
	if priBufLen <= proto.MLFrameHeadLen {
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(proto.MLFrameHeadLen-priBufLen))
		if err != nil {
			utils.Logger.Error("CopyN Error with %s", err.Error())
		}
		if err != nil {
			utils.Logger.Error("Head not complete")
			return
		}
	}
	frameHead, err = proto.UnpackFrameHead(conn.ReadBuf.Bytes()[:proto.MLFrameHeadLen])
	if err != nil {
		print(frameHead)
	}
	if priBufLen > proto.MLFrameHeadLen {
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(frameHead.Length-uint32(priBufLen-proto.MLFrameHeadLen)))
	} else {
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(frameHead.Length))
	}
	if err != nil {
		utils.Logger.Error("CopyN with body error")
		return
	}
	msg, err = proto.UnpackFrameBody(conn.ReadBuf.Bytes()[proto.MLFrameHeadLen:frameHead.Length])
	if err != nil {
		utils.Logger.Error("UnpackFrameBody Error")
		return
	}
	conn.ReadBuf.Reset()
	utils.Logger.Debug("read one message success!")
	return
}

// DealMessage deal a message from connection
func (conn *Connection) DealMessage(msg proto.Messager) (err error) {
	if msg == nil {
		err = ErrArg
		return
	}
	switch head := msg.Head().(type) {
	case *node.SYNHead:
		utils.Logger.Info("Get A SYN Message")
		err = conn.DealSYN(head)
	default:
		utils.Logger.Error("Unknown Message type")
	}
	return
}

//DealSYN deal SYN Message
func (conn *Connection) DealSYN(syn *node.SYNHead) (err error) {
	utils.Logger.Info("Deal SYN with protobuf: %d, key: %d, crypto: %d", syn.Protobuf, syn.Key, syn.Crypto)
	return
}

/*
//SendResponse Send a response
func (conn *Connection) SendResponse(rsp *Response) (err error) {
	protoData := new(proto.NodeProto)
	protoData.Init(rsp.Body)
	protoData.CMD = rsp.CMD
	protoData.Length = uint32(len(rsp.Body))
	protoData.AgentID = rsp.AgentID
	err = protoData.PackAndSend(conn.RWC)
	return
}
*/

//Close close connection
func (conn *Connection) Close() error {
	return nil
}

/*
//DealNewAgent deal a new agent
func (conn *Connection) DealNewAgent(req proto.Requester) {
	Logger.Debug("DealNewAgent with req %v", req)
	agent, err := conn.Server.AgentMgr.Alloc()
	if err != nil {
		Logger.Error("Alloc Agent Error")
		return
	}
	agent.conn = conn
	agent.lane, err = conn.Server.Backend.Dispatch()
	if err != nil {
		return
	}
	agent.DealRequest(req)
}
*/

/*
//DealSendReq deal send request
func (conn *Connection) DealSendReq(req proto.Requester) {
	Logger.Debug("Deal Send Req with req:%d ", req.AgentID)
	ag, err := conn.Server.AgentMgr.FindAgent(req.AgentID)
	if err != nil {
		Logger.Error("Find Agent Error %s", err.Error())
		if ag == nil {
			// timeout or something
		}
		// otherwith close conn
		Logger.Error("Close Connection of Agent %d", req.AgentID)
		conn.Close()
	}
	Logger.Debug("message req data  len is %d and data is %s", len(req.Body), string(req.Body))
	ag.DealRequest(req)
}
*/

//Serve serve a server
func (conn *Connection) Serve() {
	for {
		// deal timeout
		msg, err := conn.RecvMessage(5 * time.Second)
		if err != nil {
			if err != io.EOF {
				utils.Logger.Error("Connection[%v] Read Request Error:%s", conn, err.Error())
				time.Sleep(200 * time.Millisecond)
				continue
			}
			break
		}
		err = conn.DealMessage(msg)
		if err != nil {
			utils.Logger.Error("DealMessage error with %s", err.Error())
		}
	}
}
