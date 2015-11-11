/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"testing"
)

var (
	Addr = "unix:///tmp/maglined"
)

func TestConnect(t *testing.T) {
	t.Log("Test MagKnot")
	knot := New()
	knot.Init()
	knot.Connect(Addr, timeout)
}
