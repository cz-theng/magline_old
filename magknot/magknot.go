/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"bytes"
	"fmt"
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

//Delegate is delegate for a knot
type Delegate interface {
	OnAgentConnect(agentID uint32, result chan<- proto.MagKnotAgentStatus)
	OnMessageArrive(agentID uint32, data *bytes.Buffer)
	OnAgentDisconnect(agentID uint32)
}

//MagKnot is magknot
type MagKnot struct {
	seq      uint32
	conn     *net.UnixConn
	ReadBuf  *bytes.Buffer
	WriteBuf *bytes.Buffer
	delegate Delegate
}

//New create a magknot
func New() (knot *MagKnot) {
	knot = new(MagKnot)
	return
}

//Init is init
func (knot *MagKnot) Init(delegate Delegate) (err error) {
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
	knot.delegate = delegate
	return
}

//Connect connect to maglined
func (knot *MagKnot) Connect(address string, timeout time.Duration) (err error) {
	err = knot.connect(address, timeout)
	return
}

//SendMessage send a message to agent with agentID
func (knot *MagKnot) SendMessage(agentID uint32, data *bytes.Buffer, timeout time.Duration) (err error) {
	return
}

//Kickoff kick an agent with agnetID off
func (knot *MagKnot) Kickoff(agentID uint32) (err error) {
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
	fmt.Println("framehead is ", frameHead)
	if err != nil {
		fmt.Println("framehead is ", frameHead)
		fmt.Println("err is ", err)
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
		fmt.Println("unpackbody error ", err)
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
		//fmt.Println(err.Error())
		return
	}
	conn, err := net.Dial("unix", addr.IPPort)
	if err != nil {
		//fmt.Println(err.Error())
		return
	}
	knot.conn = conn.(*net.UnixConn)
	if err != nil {
		//fmt.Println(err.Error())
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
		fmt.Println(m)
	default:
		err = ErrUnknownCMD
	}

	return
}

func (knot *MagKnot) tickSeq() uint32 {
	knot.seq++
	return knot.seq
}
