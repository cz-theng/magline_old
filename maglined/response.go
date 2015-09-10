package maglined
/**
* response for client
*/

import (
)


type Response  struct {
	CMD uint16
	AgentID uint32
	Body []byte
}

func (rsp *Response) Init() {
	rsp.CMD = uint16(0)
	rsp.AgentID = uint32(0)
	rsp.Body=rsp.Body[:0]
}

