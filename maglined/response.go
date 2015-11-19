//Package maglined is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package maglined

//Response is response to knot
type Response struct {
	CMD     uint16
	AgentID uint32
	Body    []byte
}

//Init is initialize
func (rsp *Response) Init() {
	rsp.CMD = uint16(0)
	rsp.AgentID = uint32(0)
	rsp.Body = rsp.Body[:0]
}
