package magline
/**
* Request for client
*/

import (
	"github.com/cz-it/magline/maglined"
)

type Request struct {
	maglined.Request
}


func (req *Request)CMD() (uint8) {
	return 0
}

func (req *Request) Data() ([]byte) {
	return []byte("")
}










