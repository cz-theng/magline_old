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
	mtx sync.Mutex
	idleConns list.List
	ConnArray []*Connection
}

func (cp *ConnPool) Init(dataSplit []*Connection) (error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	cp.idleConns.Init()
	for i:=0; i<len(dataSplit); i++{
		cp.idleConns.PushBack(i)
	}
	cp.ConnArray = dataSplit
	return nil
}

func (cp *ConnPool) Recycle(conn *Connection) (error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	id := conn.ID
	cp.idleConns.PushBack(id)
	//conn.Uninit()
	return nil
}

func (cp *ConnPool) Alloc() (conn *Connection, err error) {
	if cp.idleConns.Len() < 1 {
		//TODO: Log here
		return conn, nil
	}
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	top := cp.idleConns.Front()
	if index,ok := cp.idleConns.Remove(top).(int); ok {
		conn = cp.ConnArray[index]
	} else {
		err = EREMOVE_TYPE
	}
	return 
}

func NewMLConnPool(size int)(mlConnPool *ConnPool, err error) {
	defer func (){
		err = ENewConn
	}()

	conns := make([]*Connection,size)
	mlConnPool = new(ConnPool)
	for i:=0; i<size; i++ {
		conn := &Connection{}
		conn.ReadBuf = make([]byte, 1024)
		conns[i] = conn
	}
	mlConnPool.Init(conns[:])
	return 
}


















