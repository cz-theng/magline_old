/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"github.com/cz-it/magline/maglined"
	"github.com/cz-it/magline/maglined/proto"
	"net"
	"time"
)

const (
	MK_READBUF_LEN = 100 * 1024
)

type MagKnot struct {
	conn    *net.UnixConn
	readBuf []byte
}

func (knot *MagKnot) Init() (err error) {
	knot.readBuf = make([]byte, MK_READBUF_LEN)
	return
}

func (knot *MagKnot) Deinit() (err error) {
	return
}

func (knot *MagKnot) Connect(address string, timeout time.Duration) (err error) {
	addr, err := maglined.ParseAddr(address)
	if err != nil {
		println(err.Error())
		return
	}
	conn, err := net.Dial("unix", addr.IPPort)
	knot.conn = conn.(*net.UnixConn)
	if err != nil {
		println(err.Error())
		return
	}
	msg := proto.KnotMessage{
		Magic:   0x01,
		Version: 0x01,
		CMD:     proto.MK_CMD_REQ_CONN,
		Seq:     0x01,
		AgentID: 0x00,
		Length:  0,
	}
	msg.PackAndSend(knot.conn)
	rsp := proto.KnotMessage{ReadBuf: knot.readBuf}
	rsp.RecvAndUnpack(knot.conn)
	println("rsp cmd", rsp.CMD)
	return
}

func (knot *MagKnot) Close() (err error) {
	return
}

func (knot *MagKnot) Send(buf []byte, timeout uint32) (err error) {
	return
}
func (knot *MagKnot) AcceptAgent(accepter func(uint32) bool) {
	var id uint32
	rsp := proto.KnotMessage{ReadBuf: knot.readBuf}
	err := rsp.RecvAndUnpack(knot.conn)
	if err != nil {
		println("Recv And Unpack Error")
		return
	}
	id = rsp.AgentID
	if !accepter(id) {
		return
	}

	msg := proto.KnotMessage{
		Magic:   0x01,
		Version: 0x01,
		CMD:     proto.MK_CMD_RSP_NEWAGENT,
		Seq:     0x01,
		AgentID: id,
		Length:  0,
	}
	msg.PackAndSend(knot.conn)
	return
}

func (knot *MagKnot) Recv(timeout uint32) (data []byte, err error) {
	return
}

func New() *MagKnot {
	knot := new(MagKnot)
	return knot
}
