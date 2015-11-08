/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"errors"
)

var (
	// ErrTimeout timeout error
	ErrTimeout = errors.New("Do things timeout such as send/recv/connect")
)
