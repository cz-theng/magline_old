//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"container/list"
	"sync"
)

//ConnPool is connection poll
type ConnPool struct {
	mtx   sync.Mutex
	conns list.List
}

//Init is initialize
func (cp *ConnPool) Init() error {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return nil
}

//Alloc is a allocater
func (cp *ConnPool) Alloc() (conn *Connection, err error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	conn = new(Connection)
	err = conn.Init()
	if err != nil {
		return
	}
	cp.conns.PushBack(conn)
	return
}

//Release will reuse a connection
func (cp *ConnPool) Release(conn *Connection) (err error) {
	cp.conns.Remove(conn.Elem)
	err = nil
	return
}

//NewMLConnPool is ConnPoll creater
func NewMLConnPool(size int) (conn *ConnPool, err error) {
	conn = new(ConnPool)
	err = conn.Init()
	return
}
