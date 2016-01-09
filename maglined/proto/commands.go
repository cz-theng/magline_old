/**
* Author: CZ cz.theng@gmail.com
 */

package proto

//MLFrameHeadLen is frame head length : 12byte
const MLFrameHeadLen = 12

const (
	// MLMagic magnode's magic
	MLMagic = 0x7f
	// MLVersion magnode's version
	MLVersion = 101
	//MLCMDUnknown is unknown commands
	MLCMDUnknown = uint16(0x0000)
)
