/**
* Author: CZ cz.theng@gmail.com
 */

package magknot

import (
	"errors"
)

var (
	//ErrClose is Close Connection
	ErrClose = errors.New("Close Connection")

	//ErrUnknownCMD is  a Unknown CMD message
	ErrUnknownCMD = errors.New("Unknown CMD message")

	// ErrTimeout timeout error
	ErrTimeout = errors.New("Do things timeout such as send/recv/connect")

	// ErrEmpty read buffer is empty
	ErrEmpty = errors.New("Buffer is emtpy")

	//ErrNoAgent is No such a Agent
	ErrNoAgent = errors.New("No such a Agent")

	//ErrEmptyMessage is No Messages
	ErrEmptyMessage = errors.New("No Messages")

	//ErrNewBuffer is create buffer error
	ErrNewBuffer = errors.New("Create Buffer Error")
)
