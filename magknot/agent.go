/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"net"
	"time"
)

type Agent struct {
	ID   uint32
	conn *net.Conn
}

func (ag *Agent) Send(buf []byte, timeout time.Duration) (err error) {

	return
}

func (ag *Agent) Recv(timeout time.Duration) (buf []byte, err error) {
	return
}

func (ag *Agent) Close() (err error) {
	return
}
