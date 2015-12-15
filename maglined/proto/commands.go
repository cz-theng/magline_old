/**
* Author: CZ cz.theng@gmail.com
 */

package proto

//MLFrameHeadLen is frame head length : 16byte
const MLFrameHeadLen = 16

const (
	// MLMagic magnode's magic
	MLMagic = 0x7f
	// MLVersion magnode's version
	MLVersion = 101
	//MLCMDUnknown is unknown commands
	MLCMDUnknown = uint16(0x0000)
)
