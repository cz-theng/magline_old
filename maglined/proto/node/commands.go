/**
* Author: CZ cz.theng@gmail.com
 */

package proto

//MNFrameHeadLen is frame head length : 16byte
const MNFrameHeadLen = 16

// MagNodeCMD is Node's cmd type
type MagNodeCMD uint16

const (
	// MNMagic magnode's magic
	MNMagic = 0x7f
	// MNVersion magnode's version
	MNVersion = 101
	//MNCMDUnknown is unknown commands
	MNCMDUnknown = uint16(0x0000)

	//MNCMDReqConn is connect reqeust
	MNCMDReqConn = uint16(0x0001)
	//MNCMDRspConn is connect response
	MNCMDRspConn = uint16(0x0002)

	//MNCMDReqClose is close request
	MNCMDReqClose = uint16(0x0003)
	//MNCMDRspClose is close response
	MNCMDRspClose = uint16(0x0004)

	//MNCMDReqReconn is reconnect request
	MNCMDReqReconn = uint16(0x0005)
	//MNCMDRspReconn is reconnect response
	MNCMDRspReconn = uint16(0x0006)

	// MNCMDMsgNode is message from magnode
	MNCMDMsgNode = uint16(0x0007)
	// MNCMDMsgKnot is message from magknot
	MNCMDMsgKnot = uint16(0x0008)
)
