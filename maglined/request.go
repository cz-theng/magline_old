//Package maglined is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package maglined

import ()

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
