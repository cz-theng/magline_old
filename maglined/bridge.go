/**
* Author: CZ cz.theng@gmail.com
 */

package maglined

import (
	"container/list"
	"net"
	"sync"
)

type Bridge struct {
	mtx   sync.Mutex
	lanes list.List
}

func (b *Bridge) Alloc(rwc *net.UnixConn) (lane *Lane, err error) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	lane = &Lane{
		RWC: rwc,
	}
	b.lanes.PushBack(lane)
	return
}

func (b *Bridge) Shunt(agent *Agent) (err error) {
	return
}
