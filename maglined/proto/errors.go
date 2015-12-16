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
)
