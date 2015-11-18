/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import (
	"errors"
)

var (
	// ErrRequestTooLong is Request's Length is Bigger than Read Buffer!
	ErrRequestTooLong = errors.New("Request's Length is Bigger than Read Buffer!")
)
