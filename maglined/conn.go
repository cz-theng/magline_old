package maglined

/**
* Connection for client
 */
import (
	"container/list"
	"github.com/cz-it/magline/maglined/proto"
	"io"
	"net"
	"time"
)

const (
	READ_BUF_SIZE = 10 * 1024
)

type Connection struct {
	RWC     *net.TCPConn
	ReadBuf []byte
	ID      int
	Elem    *list.Element
	AgentID uint32
	Server  *Server
}

func (conn *Connection) Init() error {
	conn.ReadBuf = make([]byte, READ_BUF_SIZE)
	return nil
}

func (conn *Connection) RecvRequest() (req *Request, err error) {
	Logger.Debug("RecvRequest with request and readbuf cap is %d", cap(conn.ReadBuf))
	protoData := new(proto.NodeProto)
	protoData.Init(conn.ReadBuf)
	err = protoData.RecvAndUnpack(conn.RWC)
	if err != nil {
		return nil, err
	}
	req = &Request{
		CMD:     protoData.CMD,
		AgentID: protoData.AgentID,
		Body:    protoData.Body(),
	}
	return
}

func (conn *Connection) SendResponse(rsp *Response) (err error) {
	protoData := new(proto.NodeProto)
	protoData.Init(rsp.Body)
	protoData.CMD = rsp.CMD
	protoData.Length = uint32(len(rsp.Body))
	protoData.AgentID = rsp.AgentID
	err = protoData.PackAndSend(conn.RWC)
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

func (conn *Connection) DealSendReq(req *Request) {
	Logger.Debug("Deal Send Req with req:%d ", req.AgentID)
	ag, err := conn.Server.AgentMgr.FindAgent(req.AgentID)
	if err != nil {
		Logger.Error("Find Agent Error %s", err.Error())
		if ag == nil {
			// timeout or something
		}
		// otherwith close conn
		Logger.Error("Close Connection of Agent %d", req.AgentID)
		conn.Close()
	}
	Logger.Debug("message req data  len is %d and data is %s", len(req.Body), string(req.Body))
	ag.DealRequest(req)
}

func (conn *Connection) Serve() {
	for {
		// deal timeout
		req, err := conn.RecvRequest()
		if err != nil {
			if err != io.EOF {
				Logger.Error("Connection[%v] Read Request Error:%s", conn, err.Error())
				time.Sleep(200 * time.Millisecond)
				continue
			}
			break
		}
		cmd := req.CMD
		Logger.Debug("Cmd is ", cmd)
		if cmd == proto.MN_CMD_REQ_CONN {
			conn.DealNewAgent(req)
		} else if cmd == proto.MN_CMD_MSG_NODE {
			conn.DealSendReq(req)
		} else {
			Logger.Error("Unknow CMD %d", cmd)
		}
	}
}
