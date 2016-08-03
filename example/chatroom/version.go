/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"fmt"
)

const (
	major = 0
	minor = 1
	patch = 0
)

// Version return server's version
func Version() string {
	return fmt.Sprintf("ChatroomSvr[%d.%d.%d]", major, minor, patch)
}
