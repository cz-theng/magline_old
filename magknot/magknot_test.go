/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
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

func logHooker(log string) {

}

func TestConnect(t *testing.T) {
	t.Log("Test MagKnot")
	knot := New()
	knot.Init()
	err := knot.Connect(Addr, 5000*time.Millisecond)
	if err != nil {
		fmt.Println("Connect error:%s", err.Error())
		return
	}
	println("connected success!")
	knot.Go()
	for {
		select {
		case agent := <-knot.AgentArriveChan:
			fmt.Printf("Agent %d is connected \n", agent.ID)
			err := knot.Accept(agent, proto.NewAgentSucc)
			if err != nil {
				fmt.Errorf("Accept Agent[%d] error %s \n", agent.ID, err.Error())
			}
		case msg := <-knot.MessageArriveChan:
			fmt.Printf("Agent %d send message[%s] with length %d \n", msg.Agent.ID, string(msg.Data.Bytes()), msg.Data.Len())
			err := knot.SendMessage(msg.Agent, msg.Data, 5*time.Second)
			if err != nil {
				fmt.Errorf("Send Message with error %s \n", err.Error())
			}
			fmt.Println("Send Back Echo Message Success ")
		case agent := <-knot.AgentDisconnectChan:
			fmt.Printf("Agent %d is disconnect\n", agent.ID)

		}
	}
}
