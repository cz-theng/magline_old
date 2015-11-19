/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"container/list"
	"fmt"
	"github.com/cz-it/magline/maglined"
	"github.com/cz-it/magline/maglined/proto"
	"net"
	"sync"
	"time"
)

const (
	// MKReadBufLen is read buffer length
	MKReadBufLen = 100 * 1024
)

//Message is  message
type Message struct {
	Seq  uint32
	Data []byte
}

//MagKnot is knot object
type MagKnot struct {
	conn      *net.UnixConn
	readBuf   []byte
	agents    map[uint32]*Agent
	mtx       sync.Mutex
	newAgents *list.List
}

//Init is initialize
func (knot *MagKnot) Init() (err error) {
	knot.readBuf = make([]byte, MKReadBufLen)
	knot.agents = make(map[uint32]*Agent)
	knot.newAgents = list.New()
	return
}

//Deinit is deinitialize
func (knot *MagKnot) Deinit() (err error) {
	return
}

func (knot *MagKnot) recvMsg() (err error) {
	fmt.Println("recvmsg")
	kmsg := &proto.KnotMessage{ReadBuf: knot.readBuf}
	kmsg.Init(knot.readBuf[:])
	err = kmsg.RecvAndUnpack(knot.conn)
	if err != nil {
		return
	}
	//fmt.Println("kmsg:", kmsg)

	if kmsg.CMD == proto.MKCMDReqNewAgent {
		return knot.dealNewAgent(kmsg)
	}
	//fmt.Printf("Get Node Message with agentid %d  data length :%d data: %s\n", kmsg.AgentID, kmsg.Length, string(kmsg.Body()))
	agent, ok := knot.agents[kmsg.AgentID]
	if !ok {
		err = ErrNoAgent
		return
	}
	msg := &Message{
		Seq: kmsg.Seq,
	}
	msg.Data = make([]byte, len(kmsg.Body()))
	copy(msg.Data, kmsg.Body())
	agent.pushMessage(msg)
	return
}

func (knot *MagKnot) routine() {
	for {
		fmt.Println("Routine")
		select {
		case <-time.After(200 * time.Millisecond):
			err := knot.recvMsg()
			if err != nil {
				fmt.Println("recv Msg Error ", err.Error())
			}
		}
	}
}

//Connect connect to maglined
func (knot *MagKnot) Connect(address string, timeout time.Duration) (err error) {
	addr, err := maglined.ParseAddr(address)
	if err != nil {
		//fmt.Println(err.Error())
		return
	}
	conn, err := net.Dial("unix", addr.IPPort)
	if err != nil {
		//fmt.Println(err.Error())
		return
	}
	knot.conn = conn.(*net.UnixConn)
	if err != nil {
		//fmt.Println(err.Error())
		return
	}
	//fmt.Println("Create unix doamin connect")
	msg := proto.KnotMessage{
		Magic:   0x01,
		Version: 0x01,
		CMD:     proto.MKCMDReqConn,
		Seq:     0x01,
		AgentID: 0x00,
		Length:  0,
	}
	msg.PackAndSend(nil, knot.conn)
	//fmt.Println("Send connect request!")
	rsp := proto.KnotMessage{ReadBuf: knot.readBuf}
	rsp.RecvAndUnpack(knot.conn)
	fmt.Printf("Connect Success with rsp cmd %d \n", rsp.CMD)
	go knot.routine()
	return
}

//AcceptAgent accept a new agent
func (knot *MagKnot) AcceptAgent(accepter func(uint32) bool) (agent *Agent, err error) {
	if knot.newAgents == nil {
		agent = nil
		err = ErrNoAgent
		return
	}
	knot.mtx.Lock()
	defer knot.mtx.Unlock()
	elem := knot.newAgents.Front()
	if elem == nil {
		agent = nil
		err = ErrNoAgent
		return
	}
	agent = elem.Value.(*Agent)
	knot.newAgents.Remove(elem)
	return
}
func (knot *MagKnot) dealNewAgent(kmsg *proto.KnotMessage) (err error) {
	var id uint32
	id = kmsg.AgentID
	fmt.Printf("Get Agent :%d \n", id)

	msg := proto.KnotMessage{
		Magic:   0x01,
		Version: 0x01,
		CMD:     proto.MKCMDRspNewAgent,
		Seq:     0x01,
		AgentID: id,
		Length:  0,
	}
	agent := &Agent{
		conn:    knot.conn,
		ID:      id,
		readBuf: make([]byte, MKReadBufLen),
	}
	agent.Init()
	knot.mtx.Lock()
	knot.newAgents.PushBack(agent)
	knot.agents[id] = agent
	knot.mtx.Unlock()
	msg.PackAndSend(nil, knot.conn)
	//fmt.Printf("Create a New Agent with id :%d \n", id)
	return
}

//New create a magknot
func New() *MagKnot {
	knot := new(MagKnot)
	return knot
}
