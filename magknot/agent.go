/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"github.com/cz-it/magline/maglined/proto"
	"net"
)

type Agent struct {
	ID      uint32
	conn    *net.UnixConn
	readBuf []byte
}

func (ag *Agent) Send(buf []byte) (err error) {
	msg := proto.KnotMessage{ReadBuf: ag.readBuf}
	err = msg.RecvAndUnpack(ag.conn)
	if err != nil {
		println("Recv And Unpack Error")
		return
	}
	return
}

func (ag *Agent) Recv() (buf []byte, err error) {
	msg := proto.KnotMessage{
		Magic:   0x01,
		Version: 0x01,
		CMD:     proto.MK_CMD_MSG_K2N,
		Seq:     0x01,
		AgentID: ag.ID,
		Length:  uint32(len(buf)),
	}
	msg.PackAndSend(buf, ag.conn)
	return
}

func (ag *Agent) Close() (err error) {
	return
}
