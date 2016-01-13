/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"fmt"
	"testing"
	"time"
)

var (
	Addr = "unix:///tmp/maglined"
)

type ServerHandler struct {
}

func (svr *ServerHandler) NewAgent(agent *Agent) {

}
func (svr *ServerHandler) RecvMsg(agent *Agent, data []byte) {

}
func (svr *ServerHandler) Quit(agent *Agent) {

}
func (svr *ServerHandler) Timeout() {

}
func (svr *ServerHandler) Close() {

}

func TestConnect(t *testing.T) {
	t.Log("Test MagKnot")
	knot := New(&ServerHandler{})
	knot.Init()
	err := knot.Connect(Addr, 5000*time.Millisecond)
	if err != nil {
		fmt.Println("Connect error")
		return
	}
	println("connected success!")
	knot.Serve()
}
