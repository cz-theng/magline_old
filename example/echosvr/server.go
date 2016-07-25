/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"fmt"
	"github.com/cz-it/magline/magknot"
	"github.com/cz-it/magline/proto"
	"time"
)

type server struct {
	knot *magknot.MagKnot
	addr string
}

var echoServer server

// Start a server
func Start() {
	echoServer.addr = config.Addr
	echoServer.start()
}

func (s *server) Init() (err error) {
	s.knot = magknot.New()
	err = s.knot.Init()
	if err != nil {
		print("Init knot error :", err.Error())
	}
	return
}

func (s *server) start() {
	err := s.knot.Connect(s.addr, 5000*time.Millisecond)
	if err != nil {
		fmt.Println("Connect error:%s", err.Error())
		return
	}
	println("connected success!")
	s.knot.Go()
	for {
		select {
		case agent := <-s.knot.AgentArriveChan:
			fmt.Printf("Agent %d is connected \n", agent.ID)
			err := s.knot.Accept(agent, proto.NewAgentSucc)
			if err != nil {
				fmt.Errorf("Accept Agent[%d] error %s \n", agent.ID, err.Error())
			}
		case msg := <-s.knot.MessageArriveChan:
			fmt.Printf("Agent %d send message[%s] with length %d \n", msg.Agent.ID, string(msg.Data.Bytes()), msg.Data.Len())
			err := s.knot.SendMessage(msg.Agent, msg.Data, 5*time.Second)
			if err != nil {
				fmt.Errorf("Send Message with error %s \n", err.Error())
			}
			fmt.Println("Send Back Echo Message Success ")
			//knot.DiscardAgent(msg.Agent)
		case agent := <-s.knot.AgentQuitChan:
			fmt.Printf("Agent %d is disconnect\n", agent.ID)

		}
	}
}
