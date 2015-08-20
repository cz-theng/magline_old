package maglined
/**
* Conn Pool
*/
import (
	"container/list"
	"errors"
	"sync"
)

var (
	EREMOVE_TYPE = errors.New("list.Remov with Error Type!")
)

type ConnPooler interface {
	Recycle (conn Connectioner) (bool)
	Alloc() (conn Connectioner )
	Init(dataSplit []Connectioner)(bool)
}

type ConnPool struct {
	mtx sync.Mutex
	idleConns list.List
	ConnArray []Connectioner
}

func (cp *ConnPool) Init(dataSplit []Connectioner) (error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	cp.idleConns.Init()
	for i:=0; i<len(dataSplit); i++{
		cp.idleConns.PushBack(i)
	}
	cp.ConnArray = dataSplit
	return nil
}

func (cp *ConnPool) Recycle(conn Connectioner) (error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	id := conn.ID()
	cp.idleConns.PushBack(id)
	//conn.Uninit()
	return nil
}

func (cp *ConnPool) Alloc() (conn Connectioner, err error) {
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











