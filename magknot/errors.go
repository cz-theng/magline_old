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

	// ErrEmpty read buffer is empty
	ErrEmpty = errors.New("Buffer is emtpy")

	//ErrNoAgent is No such a Agent
	ErrNoAgent = errors.New("No such a Agent")

	//ErrEmptyMessage is No Messages
	ErrEmptyMessage = errors.New("No Messages")
)
