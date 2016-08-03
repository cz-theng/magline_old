/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"fmt"
	"github.com/cz-it/magline/example/chatroom/proto"
	"github.com/cz-it/magline/magknot"
	maglineproto "github.com/cz-it/magline/proto"
	protobuf "github.com/golang/protobuf/proto"
	"time"
)

type server struct {
	knot *magknot.MagKnot
	addr string
}

var chatroomServer server

// Start a server
func Start() {
	chatroomServer.start()
}

// Init init a server
func Init() {
	chatroomServer.init()
}

func (s *server) init() (err error) {
	s.addr = config.Addr
	s.knot = magknot.New()
	err = s.knot.Init()
	if err != nil {
		print("Init knot error :", err.Error())
	}
	return
}

func (s *server) dealEnterRoom(req *proto.EnterRoomReq) {

}

func (s *server) dealMessage(message *magknot.Message) {
	var msg proto.Message
	err := protobuf.Unmarshal(message.Data.Bytes(), &msg)
	if err != nil {
		fmt.Errorf("PROTOBUF: unmarshal message error-%s", err.Error())
	}
	switch msg.GetType() {
	case proto.Message_ENTER_ROOM_REQ:
		s.dealEnterRoom(msg.GetEnterRoomReq())
	default:
		fmt.Errorf("Unknown Message Type:%v", msg.Type)
	}
	//err := s.knot.SendMessage(msg.Agent, msg.Data, 5*time.Second)
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
			err := s.knot.Accept(agent, maglineproto.NewAgentSucc)
			if err != nil {
				fmt.Errorf("Accept Agent[%d] error %s \n", agent.ID, err.Error())
			}
		case msg := <-s.knot.MessageArriveChan:
			fmt.Printf("Agent %d send message[%s] with length %d \n", msg.Agent.ID, string(msg.Data.Bytes()), msg.Data.Len())
			s.dealMessage(msg)
		case agent := <-s.knot.AgentQuitChan:
			fmt.Printf("Agent %d is disconnect\n", agent.ID)

		}
	}
}
