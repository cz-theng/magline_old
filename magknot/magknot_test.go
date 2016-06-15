/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"bytes"
	"fmt"
	"github.com/cz-it/magline/proto"
	"testing"
	"time"
)

var (
	Addr = "unix:///tmp/maglined"
)

type TestDelegate struct {
	magknot *MagKnot
}

func (td *TestDelegate) OnAgentConnect(agentID uint32, result chan<- proto.MagKnotAgentStatus) {
	fmt.Printf("Agent %d is connected", agentID)
	result <- proto.MKASAccepted
}
func (td *TestDelegate) OnMessageArrive(agentID uint32, data *bytes.Buffer) {
	fmt.Printf("Got Agent %d's Message", agentID)
	td.magknot.SendMessage(agentID, data, 5*time.Second)
}
func (td *TestDelegate) OnAgentDisconnect(agentID uint32) {
	fmt.Printf("Agent %d is disconnect", agentID)
}

func TestConnect(t *testing.T) {
	t.Log("Test MagKnot")
	dlgt := &TestDelegate{}
	knot := New()
	knot.Init(dlgt)
	err := knot.Connect(Addr, 5000*time.Millisecond)
	if err != nil {
		fmt.Println("Connect error:%s", err.Error())
		return
	}
	println("connected success!")
}
