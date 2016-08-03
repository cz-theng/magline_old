/**
* Author: CZ cz.theng@gmail.com
 */

package magline

import (
	"container/list"
	"net"
	"sync"
)

//RopeMgr is rope manager
type RopeMgr struct {
	mtx   sync.Mutex
	ropes list.List
}

//Init init RopeMgr
func (rm *RopeMgr) Init() (err error) {
	return nil
}

//Alloc alloc a lane
func (rm *RopeMgr) Alloc(rwc *net.UnixConn) (rope *Rope, err error) {
	rm.mtx.Lock()
	defer rm.mtx.Unlock()
	rope = new(Rope)
	rope.RWC = rwc
	rm.ropes.PushBack(rope)
	return
}

//Dispatch dispatch a lane
func (rm *RopeMgr) Dispatch() (rope *Rope, err error) {
	elem := rm.ropes.Front()
	if elem == nil {
		return
	}
	rope, ok := elem.Value.(*Rope)
	if !ok {
		rope = nil
		return
	}
	return
}

//NewMKRopeMgr create and init a RopeMgr
func NewMKRopeMgr(max int) (rm *RopeMgr, err error) {
	rm = new(RopeMgr)
	err = rm.Init()
	return
}
