/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"bytes"
	"fmt"
	"github.com/cz-it/magline/example/chatroom/proto"
	"github.com/cz-it/magline/magknot"
	maglineproto "github.com/cz-it/magline/proto"
	protobuf "github.com/golang/protobuf/proto"
	"time"
)

type server struct {
	knot    *magknot.MagKnot
	addr    string
	roommgr *roommgr
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
	s.roommgr, err = newRoomMgr(s.knot)
	if err != nil {
		print("Init RoomMgr error:", err.Error())
	}
	return
}

func (s *server) dealEnterRoom(agent *magknot.Agent, req *proto.EnterRoomReq) {
	fmt.Printf("Enter with room %s  and nick name %s\n", req.GetRoomName(), req.GetNickName())
	rom, err := s.roommgr.GetRoom(req.GetRoomName())
	if err != nil {
		fmt.Printf("Get Room error:%s", err.Error())
		return
	}
	fmt.Println("after get room")
	err = s.roommgr.AddAgent(agent, rom)
	if err != nil {
		fmt.Printf("Add Agent error %s", err.Error())
		return
	}
	fmt.Println("after add agetn")
	_, err = rom.AddMember(req.GetNickName(), agent)
	if err != nil {
		fmt.Printf("Add Member error %s", err.Error())
		return
	}
	fmt.Println("after add member")
	fmt.Println("tick")
	var errno int32
	rsp := &proto.EnterRoomRsp{
		Error: &errno,
	}
	mtype := proto.Message_ENTER_ROOM_RSP
	msg := &proto.Message{
		Type:         &mtype,
		EnterRoomRsp: rsp,
	}
	d, err := protobuf.Marshal(msg)
	if err != nil {
		fmt.Printf("Marshal error :%s", err.Error())
		return
	}
	err = s.knot.SendMessage(agent, bytes.NewBuffer(d), 5*time.Second)
	if err != nil {
		fmt.Errorf("Send Message error")
		return
	}
	fmt.Println("send back enter rsp")
}

func (s *server) dealUpMessage(agent *magknot.Agent, msg *proto.UpMessage) {
	var rom *room
	var ok bool
	if rom, ok = s.roommgr.agents[agent]; !ok {
		fmt.Errorf("No Room for the agent!")
		return
	}

	err := rom.BroadcastMessage(agent, msg.Message)
	if err != nil {
		fmt.Errorf("Broadcast Message error %s", err.Error())
		return
	}
}

func (s *server) dealExitRoom(agent *magknot.Agent, req *proto.ExitRoomReq) {
	var errno int32
	rsp := &proto.ExitRoomRsp{
		Error: &errno,
	}
	mtype := proto.Message_EXIT_ROOM_RSP
	msg := &proto.Message{
		Type:        &mtype,
		ExitRoomRsp: rsp,
	}
	d, err := protobuf.Marshal(msg)
	if err != nil {
		fmt.Errorf("Marshal error :%s", err.Error())
		return
	}
	err = s.knot.SendMessage(agent, bytes.NewBuffer(d), 5*time.Second)
	if err != nil {
		fmt.Errorf("Send Message error")
	}
}

func (s *server) dealMessage(message *magknot.Message) {
	var msg proto.Message
	err := protobuf.Unmarshal(message.Data.Bytes(), &msg)
	if err != nil {
		fmt.Errorf("PROTOBUF: unmarshal message error-%s", err.Error())
	}
	switch msg.GetType() {
	case proto.Message_ENTER_ROOM_REQ:
		s.dealEnterRoom(message.Agent, msg.GetEnterRoomReq())
	case proto.Message_UP_MESSAGE:
		s.dealUpMessage(message.Agent, msg.GetUpMessage())
	case proto.Message_EXIT_ROOM_REQ:
		s.dealExitRoom(message.Agent, msg.GetExitRoomReq())
	default:
		fmt.Errorf("Unknown Message Type:%v", msg.Type)
	}
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
