/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"testing"
	"time"
)

var (
	Addr = "unix:///tmp/maglined"
)

func TestConnect(t *testing.T) {
	t.Log("Test MagKnot")
	knot := New()
	knot.Init()
	err := knot.Connect(Addr, 5000*time.Millisecond)
	if err != nil {
		t.Error("Connect error")
	}
	println("connected success!")
	for {
		knot.AcceptAgent(func(id uint32) bool { return true })
	}
}
