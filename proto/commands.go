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

//ChannelType is channel type
type ChannelType uint16

const (
	// ChanNone is channel none
	ChanNone ChannelType = ChannelType(0x01)

	// ChanSalt is channel salt
	ChanSalt ChannelType = ChannelType(0x01 << 1)

	// ChanDH is channel dh
	ChanDH ChannelType = ChannelType(0x01 << 2)
)

//CryptoType is crypto's type
type CryptoType uint16

const (
	//CryptoNone is none crypto
	CryptoNone CryptoType = CryptoType(0x01)

	//CryptoAES128 is aes128 crypto
	CryptoAES128 CryptoType = CryptoType(0x01 << 2)
)

//BufSeqType is buffer sequence type
type BufSeqType uint16

const (
	//BufProtoBin is custom proto
	BufProtoBin = uint16(0x01)
	//BufProtoBuffer is protobuf
	BufProtoBuffer = uint16(0x01 << 1)
)

// MagNodeCMD is Node's cmd type
type MagNodeCMD uint16

const (
	//MNCMDSYN is connect reqeust
	MNCMDSYN = uint16(0x0001)
	//MNCMDACK is connect response
	MNCMDACK = uint16(0x0002)

	//MNCMDSeesionReq is close request
	MNCMDSeesionReq = uint16(0x0003)
	//MNCMDSessionRsp is close response
	MNCMDSessionRsp = uint16(0x0004)

	//MNCMDAuthReq is reconnect request
	MNCMDAuthReq = uint16(0x0005)
	//MNCMDAuthRsp is reconnect response
	MNCMDAuthRsp = uint16(0x0006)

	// MNCMDConfirm is message from magknot
	MNCMDConfirm = uint16(0x0008)

	//MNCMDNodeMsg is message from magnode
	MNCMDNodeMsg = uint16(0x0009)

	//MNCMDKnotMsg is message from magknot
	MNCMDKnotMsg = uint16(0x000A)
)

// MagKnotCMD is Knot's cmd type
type MagKnotCMD uint16

const (
	//MKCMDConnReq is ConnReq cmd
	MKCMDConnReq = uint16(0x1001)

	//MKCMDConnRsp is ConnRsp cmd
	MKCMDConnRsp = uint16(0x1002)

	//MKCMDAgentArriveReq is new agent request
	MKCMDAgentArriveReq = uint16(0x1003)

	// MKCMDAgentArriveRsp is new agent response
	MKCMDAgentArriveRsp = uint16(0x1004)

	//MKCMDNodeMsg is message from magnode
	MKCMDNodeMsg = uint16(0x1009)

	//MKCMDKnotMsg is message from magknot
	MKCMDKnotMsg = uint16(0x100A)
)

// MagKnotAgentStatus is Agent's status
type MagKnotAgentStatus uint16

const (
	//MKASAccepted accept the new agent
	MKASAccepted MagKnotAgentStatus = MagKnotAgentStatus(0x2001)

	//MKASRefused refuse the new agent
	MKASRefused MagKnotAgentStatus = MagKnotAgentStatus(0x2002)
)

// ErrNO is proto's error number
type ErrNO int32

const (
	//NewAgentSucc is a new agent
	NewAgentSucc = ErrNO(0)
)
