package maglined

/**
* Request for client
 */

import ()

type Request struct {
	CMD     uint16
	AgentID uint32
	Body    []byte
}
