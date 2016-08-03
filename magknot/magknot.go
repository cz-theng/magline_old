/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"bytes"
	"github.com/cz-it/magline"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/proto/frame"
	"github.com/cz-it/magline/proto/message"
	knotproto "github.com/cz-it/magline/proto/message/knot"
	"io"
	"net"
	"time"
)

const (
	//ReadBufSize is read buffer size
	ReadBufSize = 10240

	// WriteBufSize is write buffer size
	WriteBufSize = 10240
)

// Message is buffer with agent's ID
type Message struct {
	Agent *Agent
	Data  *bytes.Buffer
}

// Agent is a client agent
type Agent struct {
	ID uint32
}

// MagKnot is magknot
type MagKnot struct {
	seq               uint32
	conn              *net.UnixConn
	ReadBuf           *bytes.Buffer
	WriteBuf          *bytes.Buffer
	AgentArriveChan   chan *Agent
	AgentQuitChan     chan *Agent
	MessageArriveChan chan *Message
	agents            map[uint32]*Agent
}

//New create a magknot
func New() (knot *MagKnot) {
	knot = new(MagKnot)
	return
}

//Init is init
func (knot *MagKnot) Init() (err error) {
	rbuf := make([]byte, ReadBufSize)
	if rbuf == nil {
		return ErrNewBuffer
	}
	wbuf := make([]byte, WriteBufSize)
	if wbuf == nil {
		return ErrNewBuffer
	}
	knot.ReadBuf = bytes.NewBuffer(rbuf)
	knot.ReadBuf.Reset()
	knot.WriteBuf = bytes.NewBuffer(wbuf)
	knot.WriteBuf.Reset()
	knot.seq = 0
	knot.AgentArriveChan = make(chan *Agent)
	knot.AgentQuitChan = make(chan *Agent)
	knot.MessageArriveChan = make(chan *Message)
	knot.agents = make(map[uint32]*Agent)
	return
}

//Connect connect to maglined
func (knot *MagKnot) Connect(address string, timeout time.Duration) (err error) {
	err = knot.connect(address, timeout)
	return
}

// Accept accept a new arriving agent
func (knot *MagKnot) Accept(agent *Agent, errno proto.ErrNO) (err error) {
	msg := knotproto.NewAgentArriveRsp(agent.ID, int32(errno))
	err = knot.sendMessage(msg, 5*time.Second)
	return
}

// Go serve asynchronoursly
func (knot *MagKnot) Go() {
	go knot.reciver()
}

//SendMessage send a message to agent with agentID
func (knot *MagKnot) SendMessage(agent *Agent, data *bytes.Buffer, timeout time.Duration) (err error) {
	msg := knotproto.NewKnotMsg(agent.ID, data.Bytes())
	err = knot.sendMessage(msg, timeout)
	return
}

// DiscardAgent kick an agent with agnetID off
func (knot *MagKnot) DiscardAgent(agent *Agent) (err error) {
	msg := knotproto.NewDiscardAgent(agent.ID)
	err = knot.sendMessage(msg, 5*time.Second)
	return
}

func (knot *MagKnot) recvMessage(timeout time.Duration) (msg message.Messager, err error) {
	var frameHead *frame.Head
	priBufLen := knot.ReadBuf.Len()
	if priBufLen <= proto.MLFrameHeadLen {
		_, err = io.CopyN(knot.ReadBuf, knot.conn, int64(proto.MLFrameHeadLen-priBufLen))
		if err != nil {
			if err == io.EOF {
				err = ErrClose
			}
			return
		}
	}
	frameHead, err = frame.UnpackHead(knot.ReadBuf)
	logcat.Debug("framehead is %v", frameHead)
	if err != nil {
		logcat.Debug("framehead is %v", frameHead)
		logcat.Error("unpack head err is %s", err.Error())
		// unpack errro
	}
	if priBufLen > proto.MLFrameHeadLen {
		_, err = io.CopyN(knot.ReadBuf, knot.conn, int64(frameHead.Length-uint32(priBufLen-proto.MLFrameHeadLen)))
	} else {
		_, err = io.CopyN(knot.ReadBuf, knot.conn, int64(frameHead.Length))
	}
	if err != nil {
		if err == io.EOF {
			err = ErrClose
		}
		return
	}
	msg, err = frame.UnpackBody(frameHead.CMD, knot.ReadBuf)
	if err != nil {
		logcat.Error("unpackbody error %s", err.Error())
		return
	}
	knot.ReadBuf.Reset()
	return
}

func (knot *MagKnot) sendMessage(msg message.Messager, timeout time.Duration) (err error) {
	// Send residual data
	priBufLen := knot.WriteBuf.Len()
	if priBufLen > 0 {
		_, err = io.CopyN(knot.conn, knot.WriteBuf, int64(knot.WriteBuf.Len()))
	}
	if err != nil {
		return
	}

	// Pack data
	head := new(frame.Head)
	head.Init()
	head.Seq = knot.tickSeq()
	switch msg.(type) {
	case *knotproto.ConnReq:
		head.CMD = proto.MKCMDConnReq
	case *knotproto.AgentArriveRsp:
		head.CMD = proto.MKCMDAgentArriveRsp
	case *knotproto.KnotMsg:
		head.CMD = proto.MKCMDKnotMsg
	case *knotproto.DiscardAgent:
		head.CMD = proto.MKCMDDiscardAgent
	default:
		head.CMD = proto.MLCMDUnknown
	}
	frame := frame.Frame{
		Head: head,
		Body: msg,
	}
	err = frame.Pack(knot.WriteBuf)
	if err != nil {
		return
	}

	// Send current package
	_, err = io.CopyN(knot.conn, knot.WriteBuf, int64(knot.WriteBuf.Len()))
	if err != nil {
		return
	}
	return
}

func (knot *MagKnot) connect(address string, timeout time.Duration) (err error) {
	addr, err := magline.ParseAddr(address)
	if err != nil {
		return
	}
	conn, err := net.Dial("unix", addr.IPPort)
	if err != nil {
		return
	}
	knot.conn = conn.(*net.UnixConn)
	if err != nil {
		return
	}

	connreq := knotproto.NewConnReq([]byte("abcdefghijklmn"))
	err = knot.sendMessage(connreq, 5*time.Second)
	if err != nil {
		return
	}
	msg, err := knot.recvMessage(5 * time.Second)
	if err != nil {
		return
	}
	switch m := msg.(type) {
	case *knotproto.ConnRsp:
		logcat.Debug("Get Conn Response \n")
		knot.dealConnRsp(m)
	default:
		err = ErrUnknownCMD
	}

	return
}

func (knot *MagKnot) dealConnRsp(connRsp *knotproto.ConnRsp) error {
	return nil
}

func (knot *MagKnot) reciver() {
	for {
		msg, err := knot.recvMessage(5 * time.Second)
		if err != nil {
			logcat.Error("recv message with error %s", err.Error())
			return
		}
		switch m := msg.(type) {
		case *knotproto.AgentArriveReq:
			pbm := m.Body.(*knotproto.AgentArriveReqBody)
			logcat.Debug("Get New Agent with ID %d\n", *pbm.AgentID)
			knot.dealNewAgent(pbm)
		case *knotproto.NodeMsg:
			pbm := m.Body.(*knotproto.NodeMsgBody)
			logcat.Debug("Get a message from Agent %d ", pbm.AgentID)
			knot.dealMessage(pbm)
		case *knotproto.AgentQuit:
			pbm := m.Body.(*knotproto.AgentQuitBody)
			logcat.Debug("Get a Quit Message")
			knot.dealAgentQuit(pbm)
		default:
			err = ErrUnknownCMD
		}

	}
}

func (knot *MagKnot) dealAgentQuit(quitReq *knotproto.AgentQuitBody) (err error) {
	agent := knot.agents[*quitReq.AgentID]
	knot.AgentQuitChan <- agent
	return
}
func (knot *MagKnot) dealMessage(nodeMsg *knotproto.NodeMsgBody) (err error) {
	logcat.Debug("deal message from node %v", nodeMsg)
	msg := Message{
		Agent: knot.agents[*nodeMsg.AgentID],
		Data:  bytes.NewBuffer(nodeMsg.Payload),
	}
	knot.MessageArriveChan <- &msg
	return
}

func (knot *MagKnot) dealNewAgent(agentArriveReq *knotproto.AgentArriveReqBody) error {
	agent := &Agent{
		ID: *agentArriveReq.AgentID,
	}
	knot.agents[*agentArriveReq.AgentID] = agent
	knot.AgentArriveChan <- agent
	return nil
}

func (knot *MagKnot) tickSeq() uint32 {
	knot.seq++
	return knot.seq
}
