//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"bytes"
	"container/list"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/utils"
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

	seq      uint32
	protobuf uint16
	channel  uint16
	crypto   uint16
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
	conn.seq = 0
	conn.protobuf = 0
	conn.channel = 0
	conn.crypto = 0
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
		utils.Logger.Error("UnpackFrameBody Error with %s", err.Error())
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
		err = conn.dealSYN(head)
	case *proto.SessionReqHead:
		utils.Logger.Info("Get SessionReq ")
		err = conn.dealSessionReq(head)
	default:
		utils.Logger.Error("Unknown Message type")
	}
	return
}

func (conn *Connection) dealSessionReq(sq *proto.SessionReqHead) (err error) {
	utils.Logger.Info("Deal a SessionReq Message")
	agent, err := conn.Server.AgentMgr.Alloc()
	if err != nil {
		utils.Logger.Error("AgentManager create agent error!")
		return
	}
	rsp := proto.NewSessionRsp(agent.ID())
	err = conn.SendMessage(rsp, 5*time.Second)
	return
}

func (conn *Connection) dealSYN(syn *proto.SYNHead) (err error) {
	utils.Logger.Info("Deal SYN with protobuf: %d, key: %d, crypto: %d", syn.Protobuf, syn.Channel, syn.Crypto)
	conn.protobuf = syn.Protobuf
	conn.channel = syn.Channel
	conn.crypto = syn.Crypto
	ack := proto.NewACK(conn.channel, conn.crypto)
	err = conn.SendMessage(ack, 5*time.Second)
	return
}

func (conn *Connection) tickSeq() uint32 {
	conn.seq++
	return conn.seq
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
	head.Seq = conn.tickSeq()
	switch msg.(type) {
	case *proto.ACK:
		head.CMD = proto.MNCMDACK
	default:
		head.CMD = proto.MNCMDUnknown

	}
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

//Close close connection
func (conn *Connection) Close() error {
	return nil
}

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
