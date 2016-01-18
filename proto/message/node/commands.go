/**
* Author: CZ cz.theng@gmail.com
 */

package node

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

//ChannelType is channel type
type ChannelType uint16

const (
	// ChanNone is channel none
	ChanNone = uint16(0x01)
	// ChanSalt is channel salt
	ChanSalt = uint16(0x01 << 1)
	// ChanDH is channel dh
	ChanDH = uint16(0x01 << 2)
)

//CryptoType is crypto's type
type CryptoType uint16

const (
	//CryptoNone is none crypto
	CryptoNone = uint16(0x01)
	//CryptoAES128 is aes128 crypto
	CryptoAES128 = uint16(0x01 << 2)
)

//BufSeqType is buffer sequence type
type BufSeqType uint16

const (
	//BufProtoBin is custom proto
	BufProtoBin = uint16(0x01)
	//BufProtoBuffer is protobuf
	BufProtoBuffer = uint16(0x01 << 1)
)
