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

func DealAgent(agent *Agent) {
	fmt.Printf("Test a new Agent %d \n", agent.ID)
	for {
		msg, err := agent.Recv()
		if err != nil {
			if err == ErrEmptyMessage {
				time.Sleep(time.Millisecond * 200)
			}
			continue
		}
		fmt.Println("Recv Message:", string(msg.Data))
		err = agent.Send(msg.Data)
		if err != nil {
			continue
		}
		fmt.Println("Send Message:", string(msg.Data))
	}
	agent.Close()
}

func TestConnect(t *testing.T) {
	t.Log("Test MagKnot")
	knot := New()
	knot.Init()
	err := knot.Connect(Addr, 5000*time.Millisecond)
	if err != nil {
		fmt.Println("Connect error")
		return
	}
	println("connected success!")
	for {
		agent, err := knot.AcceptAgent(func(id uint32) bool { return true })
		if err != nil {
			if err == ErrNoAgent {
				time.Sleep(time.Millisecond * 200)
			} else {
				fmt.Println(err.Error())
			}
			continue
		}
		go DealAgent(agent)
	}
}
