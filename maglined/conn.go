package maglined
/**
* Connection for client
*/
import (
	"net"

	"github.com/cz-it/magline/maglined/proto"
)

type Connection struct {
	RWC *net.TCPConn
	ReadBuf []byte
	ID int
	protoData *proto.NodeProto
}

func (conn *Connection) RecvRequest() (*Request, error) {
	conn.protoData.Init()
	err := conn.protoData.Unpack(conn.RWC)
	if err != nil {
		return nil,err
	}
	req := &Request{
		CMD:conn.protoData.CMD,
		AgentID:conn.protoData.AgentID,
		Body:conn.protoData.Body(),
	}
	return req, nil
}

func (conn *Connection) SendResponse(rsp *Response) (err error) {
	conn.protoData.Init()
	conn.protoData.CMD = CMD_MN_CONN_RSP
	conn.protoData.AgentID = rsp.AgentID
	err = conn.protoData.Pack(conn.RWC)
	return 
}

func (conn *Connection) Close() (error ) {
	return nil
}

func (conn *Connection)Serve() {
	
	for {
		// deal timeout
		req, err := conn.RecvRequest()
		if err != nil {
			Logger.Error("Connection Read Request Error !")
			break
		}
		cmd := req.CMD
		if cmd == CMD_MN_CONN_REQ {
			DealNewAgent(conn, req)
			continue
		} 
		ag,err := Find(int(req.AgentID))
		if err != nil  {
			if ag == nil {
				// timeout or something
			}
			// otherwith close conn
			conn.Close()
		}
		ag.DealRequest(req)
	}
}






















