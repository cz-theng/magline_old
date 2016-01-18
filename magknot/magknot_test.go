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

func (svr *ServerHandler) OnNewAgent(agent *Agent) {

}
func (svr *ServerHandler) OnRecvMsg(agent *Agent, data []byte) {

}
func (svr *ServerHandler) OnAgentQuit(agent *Agent) {

}
func (svr *ServerHandler) OnTimeout() {

}
func (svr *ServerHandler) OnClose() {

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
