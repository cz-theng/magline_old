package maglined

/**
* Connector Magnager for client
 */

import (
	"container/list"
	"errors"
	"sync"
)

var (
	ENewConn = errors.New("New Connection Error!")
)

type ConnPool struct {
	mtx   sync.Mutex
	conns list.List
}

func (cp *ConnPool) Init() error {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return nil
}

func (cp *ConnPool) Alloc() (conn *Connection, err error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	conn = new(Connection)
	cp.conns.PushBack(conn)
	err = nil
	return
}

func (cp *ConnPool) Release(conn *Connection) (err error) {
	cp.conns.Remove(conn.Elem)
	err = nil
	return
}

func NewMLConnPool(size int) (conn *ConnPool, err error) {
	conn = new(ConnPool)
	err = conn.Init()
	return
}
