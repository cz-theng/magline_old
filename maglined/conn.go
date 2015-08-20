package maglined
/**
* Basic Connection
*/

import (
//	"container/list"
)

type Connectioner interface {
	ID() int
	SetID(id int)
}

type Connection struct {
	id int
}

func (conn *Connection) ID()int {
	return conn.id
}

func (conn *Connection) SetID(id int) {
	conn.id =  id
}
