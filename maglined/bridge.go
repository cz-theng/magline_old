/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"container/list"
	"net"
	"sync"
)

//Bridge is lane manager
type Bridge struct {
	mtx   sync.Mutex
	lanes list.List
}

//Alloc alloc a lane
func (b *Bridge) Alloc(rwc *net.UnixConn) (lane *Lane, err error) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	lane = &Lane{
		RWC: rwc,
	}
	b.lanes.PushBack(lane)
	return
}

//Dispatch dispatch a lane
func (b *Bridge) Dispatch() (lane *Lane, err error) {
	elem := b.lanes.Front()
	if elem == nil {
		return
	}
	lane, ok := elem.Value.(*Lane)
	if !ok {
		lane = nil
		return
	}
	return
}
