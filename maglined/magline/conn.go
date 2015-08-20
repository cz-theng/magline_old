package magline
/**
* Connection for client
*/
import (
	"net"
	"github.com/cz-it/magline/maglined"
)

type Connection struct {
	maglined.Connection
	RWC *net.TCPConn
}

func (conn *Connection)Serve() {
	
}




















