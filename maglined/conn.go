package maglined

/**
* Connection for client
 */
import (
	"container/list"
	"net"

	"github.com/cz-it/magline/maglined/proto"
)

const (
	READ_BUF_SIZE = 10 * 1024
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

func (conn *Connection) Init() error {
	conn.ReadBuf = make([]byte, READ_BUF_SIZE)
	conn.protoData = new(proto.NodeProto)
	return nil
}

func (conn *Connection) RecvRequest() (*Request, error) {
	conn.protoData.Init(conn.ReadBuf)
	err := conn.protoData.RecvAndUnpack(conn.RWC)
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
	conn.protoData.Init(rsp.Body)
	conn.protoData.CMD = rsp.CMD
	conn.protoData.AgentID = rsp.AgentID
	err = conn.protoData.PackAndSend(conn.RWC)
	return
}

func (conn *Connection) Close() error {
	return nil
}

func (conn *Connection) DealNewAgent(req *Request) {
	agent, err := conn.Server.AgentMgr.Alloc()
	if err != nil {
		Logger.Error("Alloc Agent Error")
		return
	}
	agent.DealRequest(req)
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
			print("MN_CMD_REQ_CONN")
			conn.DealNewAgent(req)
			continue
		} else {
			Logger.Error("Unknow CMD %d", cmd)
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
