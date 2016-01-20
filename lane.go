/**
* Author: CZ cz.theng@gmail.com
 */

package magline

import (
	"bytes"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/proto/frame"
	"github.com/cz-it/magline/proto/message"
	"github.com/cz-it/magline/proto/message/node"
	"github.com/cz-it/magline/utils"
	"io"
	"net"
	"sync"
	"time"
)

const (
	//LaneReadBufLen is read buffer length
	LaneReadBufLen = 1024 * 100
	//LaneWriteBufLen is write buffer length
	LaneWriteBufLen = 1024 * 100
)

//Lane is backend agent
type Lane struct {
	RWC      *net.UnixConn
	ReadBuf  *bytes.Buffer
	WriteBuf *bytes.Buffer
	agents   map[uint32]*Agent
	seq      uint32
	mtx      sync.Mutex
}

//AddAgent add a new agent
func (l *Lane) AddAgent(agent *Agent) (err error) {
	l.agents[agent.ID()] = agent
	return
}

//Init is initialize
func (l *Lane) Init() (err error) {
	rbuf := make([]byte, ReadBufSize)
	if rbuf == nil {
		return ErrNewBuffer
	}
	wbuf := make([]byte, WriteBufSize)
	if wbuf == nil {
		return ErrNewBuffer
	}
	l.ReadBuf = bytes.NewBuffer(rbuf)
	l.ReadBuf.Reset()
	l.WriteBuf = bytes.NewBuffer(wbuf)
	l.WriteBuf.Reset()
	l.seq = 0
	l.agents = make(map[uint32]*Agent)
	return
}

// SendMessage send a message frame
func (l *Lane) SendMessage(msg message.Messager, timeout time.Duration) (err error) {
	// Send residual data
	priBufLen := l.WriteBuf.Len()
	if priBufLen > 0 {
		_, err = io.CopyN(l.RWC, l.WriteBuf, int64(l.WriteBuf.Len()))
	}
	if err != nil {
		utils.Logger.Error("Send residual data error with %s", err.Error())
		return
	}

	// Pack data
	head := new(frame.Head)
	head.Init()
	head.Seq = l.tickSeq()
	switch msg.(type) {
	case *node.ACK:
		head.CMD = proto.MNCMDACK
	default:
		head.CMD = proto.MNCMDUnknown

	}
	frame := frame.Frame{
		Head: head,
		Body: msg,
	}
	err = frame.Pack(l.WriteBuf)
	if err != nil {
		utils.Logger.Error("Pack frame with error %s", err.Error())
		return
	}

	// Send current package
	_, err = io.CopyN(l.RWC, l.WriteBuf, int64(l.WriteBuf.Len()))
	if err != nil {
		utils.Logger.Error("Send  current packge data error with %s", err.Error())
		return
	}
	return
}

//RecvMessage Recv a request message
func (l *Lane) RecvMessage(timeout time.Duration) (msg message.Messager, err error) {
	var frameHead *frame.Head
	priBufLen := l.ReadBuf.Len()
	utils.Logger.Debug("priBufLen is %d", priBufLen)
	if priBufLen <= proto.MLFrameHeadLen {
		_, err = io.CopyN(l.ReadBuf, l.RWC, int64(proto.MLFrameHeadLen-priBufLen))
		if err != nil {
			if err == io.EOF {
				err = ErrClose
			}
			utils.Logger.Error("CopyN error in head with %s", err.Error())
			return
		}
	}
	frameHead, err = frame.UnpackHead(l.ReadBuf)
	if err != nil {
		// unpack errro
		utils.Logger.Error("Unpack Header with %s", err.Error())
	}
	utils.Logger.Info("Get FrameHead %v", frameHead)
	if priBufLen > proto.MLFrameHeadLen {
		_, err = io.CopyN(l.ReadBuf, l.RWC, int64(frameHead.Length-uint32(priBufLen-proto.MLFrameHeadLen)))
	} else {
		_, err = io.CopyN(l.ReadBuf, l.RWC, int64(frameHead.Length))
	}
	if err != nil {
		if err == io.EOF {
			err = ErrClose
		}
		utils.Logger.Error("CopyN with body error with %s", err.Error())
		return
	}
	msg, err = frame.UnpackBody(frameHead.CMD, l.ReadBuf)
	if err != nil {
		utils.Logger.Error("UnpackFrameBody Error with %s", err.Error())
		return
	}
	l.ReadBuf.Reset()
	utils.Logger.Debug("read one message success!")
	return
}

// DealMessage deal a message from connection
func (l *Lane) DealMessage(msg message.Messager) (err error) {
	if msg == nil {
		err = ErrArg
		return
	}
	utils.Logger.Info("Deal Message")
	return
}

func (l *Lane) tickSeq() uint32 {
	l.mtx.Lock()
	l.seq++
	l.mtx.Unlock()
	return l.seq
}

/*
//DealMsgK2N deal message from knot to node
func (l *Lane) DealMsgK2N(msg *proto.KnotMessage) (err error) {
	Logger.Debug("Send Message to Node by agent : %d", msg.AgentID)
	agent, ok := l.agents[msg.AgentID]
	if !ok {
		Logger.Error("No such Agent %d", msg.AgentID)
		return
	}
	err = agent.Send2Node(msg.Body())
	return
}
*/

//SendNodeMsg send message to node
func (l *Lane) SendNodeMsg(id uint32, data []byte) (err error) {
	/*
		Logger.Debug("Get Message form Node by agent id:%d", id)
		msg := &proto.KnotMessage{}
		msg.CMD = proto.MKCMDMsgN2K
		msg.Length = uint32(len(data))
		msg.Seq = l.tickSeq()
		msg.AgentID = id
		msg.PackAndSend(data, l.RWC)
		Logger.Debug("send node msg with length %d,data:%s", msg.Length, string(data))
	*/
	return

}

//Serve server the server
func (l *Lane) Serve() {
	for {
		msg, err := l.RecvMessage(5 * time.Second)
		if err != nil {
			if err == ErrClose {
				utils.Logger.Info("Knot Close Connection")
				break
			} else {
				utils.Logger.Error("Connection[%v] Read Request Error:%s", l, err.Error())
				time.Sleep(200 * time.Millisecond)
				continue
			}
		}
		err = l.DealMessage(msg)
		if err != nil {
			utils.Logger.Error("DealMessage error with %s", err.Error())
		}
	}
}
