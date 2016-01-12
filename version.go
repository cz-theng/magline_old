//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"fmt"
)

const (
	major = 0
	minor = 1
	patch = 0
)

// Version return maglined's version
func Version() string {
	return fmt.Sprintf("maglined[%d.%d.%d]", major, minor, patch)
}
