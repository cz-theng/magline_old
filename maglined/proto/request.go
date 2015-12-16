//Package proto is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package proto

import ()

//Requester is Request's interface
type Requester interface {
	Init()
	AgentID() uint32
	CMD() uint16
	Body() []byte
}

//Request is request object
type Request struct {
	CMD     uint16
	AgentID uint32
	Body    []byte
}

//Init is initialize
func (req *Request) Init() {
	req.CMD = uint16(0)
	req.AgentID = uint32(0)
	req.Body = req.Body[:0]
}

//CMD return request's cmd
func (req *Request) CMD() uint16 {
	return req.CMD
}

//AgentID return request's agent id
func (req *Request) AgentID() uint32 {
	return req.AgentID
}

//Body return request's  body
func (req *Request) Body() []byte {
	return req.Body
}
