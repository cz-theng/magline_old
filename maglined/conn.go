package maglined

/**
* Connection for client
 */
import (
	"container/list"
	"github.com/cz-it/magline/maglined/proto"
	"net"
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
	Logger.Debug("RecvRequest with request ")
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
	Logger.Debug("DealNewAgent with req %v", req)
	agent, err := conn.Server.AgentMgr.Alloc()
	if err != nil {
		Logger.Error("Alloc Agent Error")
		return
	}
	agent.conn = conn
	agent.lane, err = conn.Server.Backend.Dispatch()
	if err != nil {
		return
	}
	agent.DealRequest(req)
}

func (conn *Connection) Serve() {
	for {
		// deal timeout
		req, err := conn.RecvRequest()
		if err != nil {
			Logger.Error("Connection[%v] Read Request Error:%s", conn, err.Error())
			break
		}
		cmd := req.CMD
		Logger.Debug("Cmd is ", cmd)
		if cmd == proto.MN_CMD_REQ_CONN {
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
