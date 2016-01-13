/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"time"
)

// Handler is handler
type Handler interface {
	NewAgent(agent *Agent)
	RecvMsg(agent *Agent, data []byte)
	Quit(agent *Agent)
	Timeout()
	Close()
}

//MagKnot is magknot
type MagKnot struct {
	hdl Handler
}

//Init is init
func (knot *MagKnot) Init() (err error) {
	return
}

//Connect connect to maglined
func (knot *MagKnot) Connect(address string, timeout time.Duration) (err error) {
	return
}

// Serve is main routine
func (knot *MagKnot) Serve() {

}

//New create a magknot
func New(hdl Handler) (knot *MagKnot) {
	if hdl == nil {
		return nil
	}
	knot = new(MagKnot)
	knot.hdl = hdl
	return
}
