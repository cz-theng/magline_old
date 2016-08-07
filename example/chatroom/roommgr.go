/**
* Author: CZ cz.devnet@gmail.com
 */

package main

import (
	"github.com/cz-it/magline/magknot"
	"sync"
)

type roommgr struct {
	rooms  map[string]*room
	knot   *magknot.MagKnot
	mtx    sync.Mutex
	agents map[*magknot.Agent]*room
}

func newRoomMgr(knot *magknot.MagKnot) (rm *roommgr, err error) {
	rm = &roommgr{
		rooms:  make(map[string]*room),
		knot:   knot,
		agents: make(map[*magknot.Agent]*room),
	}
	return
}

func (rm *roommgr) DelAgent(agent *magknot.Agent) (err error) {
	var ok bool
	rm.mtx.Lock()
	defer rm.mtx.Unlock()
	if _, ok = rm.agents[agent]; !ok {
		return
	}
	delete(rm.agents, agent)
	return
}

func (rm *roommgr) GetRoom(name string) (rom *room, err error) {
	var ok bool
	rm.mtx.Lock()
	defer rm.mtx.Unlock()
	if rom, ok = rm.rooms[name]; ok {
		return
	}
	rom = &room{
		name:    name,
		knot:    rm.knot,
		members: make(map[*magknot.Agent]*member),
	}
	rm.rooms[name] = rom
	return
}

func (rm *roommgr) AddAgent(agent *magknot.Agent, rom *room) (err error) {
	var ok bool
	rm.mtx.Lock()
	defer rm.mtx.Unlock()
	if _, ok = rm.agents[agent]; ok {
		return
	}
	rm.agents[agent] = rom
	return
}
