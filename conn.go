//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"bytes"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/proto/frame"
	"github.com/cz-it/magline/proto/message"
	"github.com/cz-it/magline/proto/message/knot"
	"github.com/cz-it/magline/proto/message/node"
	"github.com/cz-it/magline/utils"
	"io"
	"sync"
	"time"
)

//Connectioner is connection API
type Connectioner interface {
	// DealMessage deal a message from connection
	DealMessage(msg message.Messager) (err error)
}

//Connection is connection object
type Connection struct {
	RWC         io.ReadWriter
	ReadBuf     *bytes.Buffer
	WriteBuf    *bytes.Buffer
	Server      *Server
	seq         uint32
	recvChan    chan message.Messager
	Agent       *Agent
	wg          sync.WaitGroup
	recvable    chan bool
	dealMessage func(message.Messager) error
}

//Init is initialize
func (conn *Connection) Init(rbs, wbs int) error {
	rbuf := make([]byte, rbs)
	if rbuf == nil {
		utils.Logger.Error("make bytes error")
		return ErrNewBuffer
	}
	wbuf := make([]byte, wbs)
	if wbuf == nil {
		utils.Logger.Error("make bytes error")
		return ErrNewBuffer
	}
	conn.ReadBuf = bytes.NewBuffer(rbuf)
	conn.ReadBuf.Reset()
	conn.WriteBuf = bytes.NewBuffer(wbuf)
	conn.WriteBuf.Reset()
	conn.seq = 0
	conn.recvable = make(chan bool)
	conn.recvChan = make(chan message.Messager)
	return nil
}

//RecvMessage Recv a request message
func (conn *Connection) RecvMessage(timeout time.Duration) (msg message.Messager, err error) {
	var frameHead *frame.Head
	priBufLen := conn.ReadBuf.Len()
	if priBufLen <= proto.MLFrameHeadLen { // havn't got a complete framehead
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(proto.MLFrameHeadLen-priBufLen))
		if err != nil {
			if err == io.EOF {
				//err = ErrClose
			}
			utils.Logger.Error("CopyN error in head with %s", err.Error())
			return
		}
	}
	frameHead, err = frame.UnpackHead(conn.ReadBuf)
	if err != nil {
		utils.Logger.Error("Unpack Header with %s", err.Error())
		return
	}
	utils.Logger.Debug("Get FrameHead %v", frameHead)
	if priBufLen > proto.MLFrameHeadLen { // have some body but not complete
		utils.Logger.Debug("Have some body but not complete %d", frameHead.Length-uint32(priBufLen-proto.MLFrameHeadLen))
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(frameHead.Length-uint32(priBufLen-proto.MLFrameHeadLen)))
	} else {
		utils.Logger.Debug("Get Data %d", frameHead.Length)
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(frameHead.Length))
	}
	if err != nil {
		if err == io.EOF {
			err = ErrClose
		}
		utils.Logger.Error("CopyN with body error with %s", err.Error())
		return
	}
	msg, err = frame.UnpackBody(frameHead.CMD, conn.ReadBuf)
	if err != nil {
		utils.Logger.Error("UnpackFrameBody Error with %s", err.Error())
		return
	}
	conn.ReadBuf.Reset()
	return
}

// SendMessage send a message frame
func (conn *Connection) SendMessage(msg message.Messager, timeout time.Duration) (err error) {
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
	head := new(frame.Head)
	head.Init()
	head.Seq = conn.tickSeq()
	switch msg.(type) {
	case *node.ACK:
		head.CMD = proto.MNCMDACK
	case *node.SessionRsp:
		head.CMD = proto.MNCMDSessionRsp
	case *knot.ConnRsp:
		head.CMD = proto.MKCMDConnRsp
	case *knot.AgentArriveReq:
		head.CMD = proto.MKCMDAgentArriveReq
	case *node.Confirm:
		head.CMD = proto.MNCMDConfirm
	case *knot.NodeMsg:
		head.CMD = proto.MKCMDNodeMsg
	case *node.KnotMsg:
		head.CMD = proto.MNCMDKnotMsg
	case *node.ErrorMsg:
		head.CMD = proto.MNCMDErrorMsg
	case *knot.AgentQuit:
		head.CMD = proto.MKCMDAgentQuit
	case *node.Discard:
		head.CMD = proto.MNCMDDiscard
	default:
		head.CMD = proto.MLCMDUnknown

	}
	frame := frame.Frame{
		Head: head,
		Body: msg,
	}
	utils.Logger.Debug("before pack buf is %d", conn.WriteBuf.Len())
	err = frame.Pack(conn.WriteBuf)
	utils.Logger.Debug("after pack buf is %d", conn.WriteBuf.Len())
	if err != nil {
		utils.Logger.Error("Pack frame with error %s", err.Error())
		return
	}

	// Send current package
	s, err := io.CopyN(conn.RWC, conn.WriteBuf, int64(conn.WriteBuf.Len()))
	if err != nil {
		utils.Logger.Error("Send  current packge data error with %s", err.Error())
		return
	}
	utils.Logger.Debug("CopyN with %d", s)
	return
}

//Close close connection
func (conn *Connection) Close() error {
	return nil
}

func (conn *Connection) recvRoutine() {
	for {
		select {
		case r := <-conn.recvable:
			if r {
				conn.wg.Done()
				return
			}
		default:
		}
		msg, err := conn.RecvMessage(5 * time.Second)
		if err != nil {
			if err == io.EOF {
				utils.Logger.Info("Client Close Connection")
				close(conn.recvChan)
				break
			}
		} else {
			utils.Logger.Debug("send msg to recvChan channel")
			conn.recvChan <- msg
		}
	}
}

//Serve serve a server
func (conn *Connection) Serve() {
	conn.wg.Add(1)
	go conn.recvRoutine()
	for {
		select {
		case msg, ok := <-conn.recvChan:
			if !ok {
				utils.Logger.Info("Conn is close")
				conn.recvable <- false
				break
			}
			utils.Logger.Debug("Recv Message %v", msg)
			if msg == nil {
				utils.Logger.Info("Client Close Connection")
				conn.recvable <- false
				break
			}
			utils.Logger.Debug("Recv Message %v", msg)
			err := conn.dealMessage(msg)
			if err != nil {
				utils.Logger.Error("DealMessage error with %s", err.Error())
			}
			utils.Logger.Info("after deal message")
		}
	}
	conn.wg.Wait()
}

func (conn *Connection) tickSeq() uint32 {
	conn.seq++
	return conn.seq
}
