/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"time"
)

//Agenter is agenter
type Agenter interface {
	KickOut() (err error)
	SendMsg(data []byte, timeout time.Duration) (err error)
}

//Agent is implement of Agenter
type Agent struct {
}
