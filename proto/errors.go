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

	// ErrFrameHeadBufNil is FrameHead Buffer is nil
	ErrFrameHeadBufNil = errors.New("FrameHead Buffer is nil")

	//ErrFameHeadBufLen is FrameHead Buffer's Length is Invalid,Should be 16 bytes
	ErrFameHeadBufLen = errors.New("FrameHead Buffer's Length is Invalid,Should be 16 bytes")

	//ErrUnknownCMD is a unknown command
	ErrUnknownCMD = errors.New("A Unknown Command")

	// ErrBufNil is FrameHead Buffer is nil
	ErrBufNil = errors.New("Buffer is nil")

	// ErrBufLen is buffer's length error
	ErrBufLen = errors.New("Buffer's length is not right")
)
