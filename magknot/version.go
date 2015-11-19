//Package magknot is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magknot

import (
	"fmt"
)

const (
	major = 0
	minor = 0
	patch = 2
)

// Version return maglined's version
func Version() string {
	return fmt.Sprintf("magknot[%d.%d.%d]", major, minor, patch)
}
