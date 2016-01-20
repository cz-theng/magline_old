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

//
const (
	//ReadBufSize is read buffer size
	ReadBufSize = 10240
	// WriteBufSize is write buffer size
	WriteBufSize = 10240
)

// Handler is handler
type Handler interface {
	OnNewAgent(agent *Agent)
	OnRecvMsg(agent *Agent, data []byte)
	OnAgentQuit(agent *Agent)
	OnTimeout()
	OnClose()
}

//MagKnot is magknot
type MagKnot struct {
	hdl      Handler
	seq      uint32
	conn     *net.UnixConn
	ReadBuf  *bytes.Buffer
	WriteBuf *bytes.Buffer
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
	return
}

//RecvMessage Recv a request message
func (knot *MagKnot) RecvMessage(timeout time.Duration) (msg message.Messager, err error) {
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
	if err != nil {
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
		return
	}
	knot.ReadBuf.Reset()
	return
}

//SendMessage send a message
func (knot *MagKnot) SendMessage(msg message.Messager, timeout time.Duration) (err error) {
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
		head.CMD = proto.MNCMDUnknown

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

//Connect connect to maglined
func (knot *MagKnot) Connect(address string, timeout time.Duration) (err error) {
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
	err = knot.SendMessage(connreq, 5*time.Second)
	if err != nil {
		return
	}
	msg, err := knot.RecvMessage(5 * time.Second)
	if err != nil {
		return
	}
	switch m := msg.(type) {
	case *knotproto.ConnRsp:
		print(m)
	default:
		err = ErrUnknownCMD
	}

	return
}

// Serve is main routine
func (knot *MagKnot) Serve() {

}

// ServeAsync is main routine
func (knot *MagKnot) ServeAsync() {
	go knot.Serve()
}

//New create a magknot
func New(hdl Handler) (knot *MagKnot) {
	if hdl == nil {
		return nil
	}
	knot = new(MagKnot)
	knot.hdl = hdl
	return
}

func (knot *MagKnot) tickSeq() uint32 {
	knot.seq++
	return knot.seq
}
