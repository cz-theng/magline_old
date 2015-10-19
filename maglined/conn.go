package maglined

/**
* Connection for client
 */
import (
	"container/list"
	"net"

	"github.com/cz-it/magline/maglined/proto"
)

type Connection struct {
	RWC       *net.TCPConn
	ReadBuf   []byte
	ID        int
	protoData *proto.NodeProto
	Elem      *list.Element
	AgentID   uint32
	Server    *Server
}

func (conn *Connection) RecvRequest() (*Request, error) {
	conn.protoData.Init()
	err := conn.protoData.Unpack(conn.RWC)
	if err != nil {
		return nil, err
	}
	req := &Request{
		CMD:     conn.protoData.CMD,
		AgentID: conn.protoData.AgentID,
		Body:    conn.protoData.Body(),
	}
	return req, nil
}

func (conn *Connection) SendResponse(rsp *Response) (err error) {
	conn.protoData.Init()
	conn.protoData.CMD = proto.MN_CMD_RSP_CONN
	conn.protoData.AgentID = rsp.AgentID
	err = conn.protoData.Pack(conn.RWC)
	return
}

func (conn *Connection) Close() error {
	return nil
}

func (conn *Connection) Serve() {
	for {
		// deal timeout
		req, err := conn.RecvRequest()
		if err != nil {
			Logger.Error("Connection Read Request Error !")
			break
		}
		cmd := req.CMD
		if cmd == proto.MN_CMD_REQ_CONN {
			DealNewAgent(conn, req)
			continue
		}
		ag, err := conn.Server.AgentMgr.FindAgent(req.AgentID)
		if err != nil {
			if ag == nil {
				// timeout or something
			}
			// otherwith close conn
			conn.Close()
		}
		ag.DealRequest(req)
	}
}
