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
	for {
		buf, err := agent.Recv(time.Second * 5)
		if err != nil {
			continue
		}
		fmt.Println("Recv Message:", string(buf))
		err = agent.Send(buf, time.Second*5)
		if err != nil {
			continue
		}
		fmt.Println("Send Message:", string(buf))
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
			fmt.Println(err.Error())
			continue
		}
		go DealAgent(agent)
	}
}
