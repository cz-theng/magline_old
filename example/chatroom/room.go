/**
* Author: CZ cz.devnet@gmail.com
 */

package main

import (
	"bytes"
	"fmt"
	"github.com/cz-it/magline/example/chatroom/proto"
	"github.com/cz-it/magline/magknot"
	protobuf "github.com/golang/protobuf/proto"
	"sync"
	"time"
)

type room struct {
	members map[*magknot.Agent]*member
	name    string
	idGuard uint32
	knot    *magknot.MagKnot
	mtx     sync.Mutex
}

func (r *room) DelMember(agent *magknot.Agent) (err error) {
	var ok bool
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok = r.members[agent]; !ok {
		return
	}
	delete(r.members, agent)
	return
}

func (r *room) AddMember(name string, agent *magknot.Agent) (mb *member, err error) {
	var ok bool
	if mb, ok = r.members[agent]; ok {
		return
	}
	mb = &member{
		name:  name,
		id:    r.seq(),
		agent: agent,
	}
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.members[agent] = mb
	return
}

func (r *room) BroadcastMessage(agent *magknot.Agent, data []byte) (err error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	for a, m := range r.members {
		if a == agent {
			continue
		}
		length := int32(len(data))
		dmsg := &proto.DownMessage{
			Message:  data,
			Length:   &length,
			NickName: &m.name,
		}
		mtype := proto.Message_DOWN_MESSAGE
		msg := &proto.Message{
			Type:        &mtype,
			DownMessage: dmsg,
		}
		d, err := protobuf.Marshal(msg)
		if err != nil {
			fmt.Errorf("Marshal error :%s", err.Error())
			return err
		}
		err = r.knot.SendMessage(m.agent, bytes.NewBuffer(d), 5*time.Second)
		if err != nil {
			fmt.Errorf("Send Message error")
		}
	}
	return
}

func (r *room) seq() uint32 {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.idGuard++
	return r.idGuard
}
