//Package proto is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package proto

import ()

//Requester is Request's interface
type Requester interface {
	Unpack([]byte) error
	AgentID() uint32
	CMD() uint16
	Body() []byte
}

//Request is request object
type Request struct {
	cmd     uint16
	agentID uint32
	body    []byte
}

//Init is initialize
func (req *Request) Init() {
	req.cmd = uint16(0)
	req.agentID = uint32(0)
	req.body = req.body[:0]
}

//CMD return request's cmd
func (req *Request) CMD() uint16 {
	return req.cmd
}

//AgentID return request's agent id
func (req *Request) AgentID() uint32 {
	return req.agentID
}

//Body return request's  body
func (req *Request) Body() []byte {
	return req.body
}
