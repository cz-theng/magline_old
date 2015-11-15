/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"fmt"
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
	agents  map[uint32]*Agent
}

func (knot *MagKnot) Init() (err error) {
	knot.readBuf = make([]byte, MK_READBUF_LEN)
	knot.agents = make(map[uint32]*Agent)
	return
}

func (knot *MagKnot) Deinit() (err error) {
	return
}

func (knot *MagKnot) Connect(address string, timeout time.Duration) (err error) {
	addr, err := maglined.ParseAddr(address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	conn, err := net.Dial("unix", addr.IPPort)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	knot.conn = conn.(*net.UnixConn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Create unix doamin connect")
	msg := proto.KnotMessage{
		Magic:   0x01,
		Version: 0x01,
		CMD:     proto.MK_CMD_REQ_CONN,
		Seq:     0x01,
		AgentID: 0x00,
		Length:  0,
	}
	msg.PackAndSend(nil, knot.conn)
	fmt.Println("Send connect request!")
	rsp := proto.KnotMessage{ReadBuf: knot.readBuf}
	rsp.RecvAndUnpack(knot.conn)
	fmt.Printf("Connect Success with rsp cmd", rsp.CMD)
	return
}

func (knot *MagKnot) AcceptAgent(accepter func(uint32) bool) (agent *Agent, err error) {
	var id uint32
	rsp := proto.KnotMessage{ReadBuf: knot.readBuf}
	err = rsp.RecvAndUnpack(knot.conn)
	if err != nil {
		fmt.Printf("Recv And Unpack Error :%s", err.Error())
		return
	}
	id = rsp.AgentID
	fmt.Printf("Get Agent :%d", id)
	if !accepter(id) {
		fmt.Printf("id[%d] is not acceptable", id)
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
	agent = &Agent{
		conn:    knot.conn,
		ID:      id,
		readBuf: make([]byte, MK_READBUF_LEN),
	}
	msg.PackAndSend(nil, knot.conn)
	fmt.Printf("Create a New Agent with id :%d", id)
	return
}

func New() *MagKnot {
	knot := new(MagKnot)
	return knot
}
