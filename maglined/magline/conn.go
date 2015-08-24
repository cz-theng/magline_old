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

func (conn *Connection) readRequest() (*Request, error) {
	return nil,nil
}

func (conn *Connection)Serve() {
	
	for {
		req, err := conn.readRequest()
		if err != nil {
			maglined.Logger.Error("Connection Read Request Error !")
			break
		}
		
	}
}




















