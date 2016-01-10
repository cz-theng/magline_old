//Package maglined is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package main

import (
	"bytes"
	"container/list"
	"github.com/cz-it/magline/maglined/proto"
	"github.com/cz-it/magline/maglined/utils"
	"io"
	"net"
	"time"
)

const (
	//ReadBufSize is read buffer size
	ReadBufSize = 10 * 1024

	//WriteBufSize is write buffer size
	WriteBufSize = 10 * 1024
)

//Connection is connection object
type Connection struct {
	RWC      *net.TCPConn
	ReadBuf  *bytes.Buffer
	WriteBuf *bytes.Buffer
	ID       int
	Elem     *list.Element
	AgentID  uint32
	Server   *Server
}

//Init is initialize
func (conn *Connection) Init() error {
	rbuf := make([]byte, ReadBufSize)
	if rbuf == nil {
		return ErrNewBuffer
	}
	wbuf := make([]byte, WriteBufSize)
	if wbuf == nil {
		return ErrNewBuffer
	}
	conn.ReadBuf = bytes.NewBuffer(rbuf)
	conn.ReadBuf.Reset()
	conn.WriteBuf = bytes.NewBuffer(wbuf)
	conn.WriteBuf.Reset()
	return nil
}

//RecvMessage Recv a request message
func (conn *Connection) RecvMessage(timeout time.Duration) (msg proto.Messager, err error) {
	var frameHead *proto.FrameHead
	priBufLen := conn.ReadBuf.Len()
	utils.Logger.Debug("priBufLen is %d", priBufLen)
	if priBufLen <= proto.MLFrameHeadLen {
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(proto.MLFrameHeadLen-priBufLen))
		if err != nil {
			if err == io.EOF {
				err = ErrClose
			}
			utils.Logger.Error("CopyN error in head with %s", err.Error())
			return
		}
	}
	frameHead, err = proto.UnpackFrameHead(conn.ReadBuf.Bytes()[:proto.MLFrameHeadLen])
	if err != nil {
		// unpack errro
		utils.Logger.Error("Unpack Header with %s", err.Error())
	}
	utils.Logger.Info("Get FrameHead %v", frameHead)
	if priBufLen > proto.MLFrameHeadLen {
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(frameHead.Length-uint32(priBufLen-proto.MLFrameHeadLen)))
	} else {
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(frameHead.Length))
	}
	if err != nil {
		if err == io.EOF {
			err = ErrClose
		}
		utils.Logger.Error("CopyN with body error with %s", err.Error())
		return
	}
	msg, err = proto.UnpackFrameBody(frameHead.CMD, conn.ReadBuf.Bytes()[proto.MLFrameHeadLen:proto.MLFrameHeadLen+frameHead.Length])
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
	case *proto.SYNHead:
		utils.Logger.Info("Get A SYN Message")
		err = conn.DealSYN(head)
	default:
		utils.Logger.Error("Unknown Message type")
	}
	return
}

//DealSYN deal SYN Message
func (conn *Connection) DealSYN(syn *proto.SYNHead) (err error) {
	utils.Logger.Info("Deal SYN with protobuf: %d, key: %d, crypto: %d", syn.Protobuf, syn.Channel, syn.Crypto)
	ack := proto.NewACK(1, 1)
	err = conn.SendMessage(ack, 5*time.Second)
	return
}

// SendMessage send a message frame
func (conn *Connection) SendMessage(msg proto.Messager, timeout time.Duration) (err error) {
	// Send residual data
	priBufLen := conn.WriteBuf.Len()
	if priBufLen > 0 {
		_, err = io.CopyN(conn.RWC, conn.WriteBuf, int64(conn.WriteBuf.Len()))
	}
	if err != nil {
		utils.Logger.Error("Send residual data error with %s", err.Error())
		return
	}

	// Pack data
	head := new(proto.FrameHead)
	head.Init()
	head.Seq = 2
	head.CMD = proto.MNCMDACK

	frame := proto.Frame{
		Head: head,
		Body: msg,
	}
	_, err = frame.Pack(conn.WriteBuf)
	if err != nil {
		utils.Logger.Error("Pack frame with error %s", err.Error())
		return
	}

	// Send current package
	_, err = io.CopyN(conn.RWC, conn.WriteBuf, int64(conn.WriteBuf.Len()))
	if err != nil {
		utils.Logger.Error("Send  current packge data error with %s", err.Error())
		return
	}
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
			if err == ErrClose {
				utils.Logger.Info("Client Close Connection")
				break
			} else {
				utils.Logger.Error("Connection[%v] Read Request Error:%s", conn, err.Error())
				time.Sleep(200 * time.Millisecond)
				continue
			}
		}
		err = conn.DealMessage(msg)
		if err != nil {
			utils.Logger.Error("DealMessage error with %s", err.Error())
		}
	}
}
