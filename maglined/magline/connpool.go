package magline
/**
* Connector Magnager for client
*/

import (
	"errors"

	"github.com/cz-it/magline/maglined"
)

var (
	ENewConn = errors.New("New Connection Error!")
)

func NewMLConnPool(size int)(mlConnPool *maglined.ConnPool, err error) {
	defer func (){
		err = ENewConn
	}()

	conns := make([]maglined.Connectioner,size)
	mlConnPool = new(maglined.ConnPool)
	for i:=0; i<size; i++ {
		conns[i]= &Connection{}
	}
	mlConnPool.Init(conns[:])
	return 
}


















