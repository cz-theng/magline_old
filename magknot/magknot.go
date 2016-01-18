/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"github.com/cz-it/magline"
	"net"
	"time"
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
	hdl  Handler
	conn *net.UnixConn
}

//Init is init
func (knot *MagKnot) Init() (err error) {
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
	//fmt.Println("Create unix doamin connect")
	/*
		msg := proto.KnotMessage{
			Magic:   0x01,
			Version: 0x01,
			CMD:     proto.MKCMDReqConn,
			Seq:     0x01,
			AgentID: 0x00,
			Length:  0,
		}
		msg.PackAndSend(nil, knot.conn)
		//fmt.Println("Send connect request!")
		rsp := proto.KnotMessage{ReadBuf: knot.readBuf}
		rsp.RecvAndUnpack(knot.conn)
		fmt.Printf("Connect Success with rsp cmd %d \n", rsp.CMD)
		go knot.routine()
	*/
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
