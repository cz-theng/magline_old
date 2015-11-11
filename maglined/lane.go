/**
* Author: CZ cz.theng@gmail.com
 */

package maglined

import (
	"net"
)

type Lane struct {
	RWC    *net.UnixConn
	agents map[uint32]*Agent
}

func (l *Lane) ReadMsg() (msg string, err error) {
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
	}
}
