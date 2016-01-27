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
		return ErrNewBuffer
	}
	wbuf := make([]byte, wbs)
	if wbuf == nil {
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
	if priBufLen <= proto.MLFrameHeadLen {
		_, err = io.CopyN(conn.ReadBuf, conn.RWC, int64(proto.MLFrameHeadLen-priBufLen))
		if err != nil {
			if err == io.EOF {
				//err = ErrClose
			}
			//utils.Logger.Error("CopyN error in head with %s", err.Error())
			return
		}
	}
	frameHead, err = frame.UnpackHead(conn.ReadBuf)
	if err != nil {
		utils.Logger.Error("Unpack Header with %s", err.Error())
		return
	}
	utils.Logger.Debug("Get FrameHead %v", frameHead)
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
	msg, err = frame.UnpackBody(frameHead.CMD, conn.ReadBuf)
	if err != nil {
		utils.Logger.Error("UnpackFrameBody Error with %s", err.Error())
		return
	}
	conn.ReadBuf.Reset()
	return
}

func (conn *Connection) tickSeq() uint32 {
	conn.seq++
	return conn.seq
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
	case *knot.ConnRsp:
		head.CMD = proto.MKCMDConnRsp
	default:
		head.CMD = proto.MLCMDUnknown

	}
	frame := frame.Frame{
		Head: head,
		Body: msg,
	}
	err = frame.Pack(conn.WriteBuf)
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
		//utils.Logger.Debug("Recv Message %v", msg)
		if err != nil {
			/*
				if err == ErrClose {
					utils.Logger.Info("Client Close Connection")
					conn.recvChan <- nil
				} else {
					utils.Logger.Error("RecvMessage with error %s", err.Error())
				}
			*/
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
		var msg message.Messager
		select {
		case msg = <-conn.recvChan:
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
			/*
				case msg = <-conn.Agent.nodeMsgChan:
					conn.SendMessage(msg, 5*time.Second)
			*/
		}
	}
	conn.wg.Wait()
}
