/**
* Author: CZ cz.theng@gmail.com
 */

package maglined

import (
	"github.com/cz-it/magline/maglined/proto"
	"net"
)

type Lane struct {
	RWC     *net.UnixConn
	agents  map[uint32]*Agent
	ReadBuf []byte
}

func (l *Lane) Init() (err error) {
	l.ReadBuf = make([]byte, READ_BUF_SIZE)
	return
}

func (l *Lane) ReadMsg() (msg *proto.KnotMessage, err error) {
	msg = &proto.KnotMessage{}
	err = msg.RecvAndUnpack(l.RWC)
	if err != nil {
		msg = nil
		return
	}
	return
}

func (l *Lane) Serve() {
	for {
		// deal timeout
		msg, err := l.ReadMsg()
		if err != nil {
			Logger.Error("Connection Read Request Error:%s", err.Error())
			continue
		}
		Logger.Debug("get message: %v", msg.CMD)
	}
}
