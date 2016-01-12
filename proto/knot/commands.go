/**
* Author: CZ cz.theng@gmail.com
 */

package proto

import ()

// MagKnotCMD is Knot's cmd type
type MagKnotCMD uint16

const (
	// MKMagic is magknot's magic
	MKMagic = 0x7f
	// MKVersion is magknot's version
	MKVersion = 0x01

	//MKCMDUnknown is unknown commands
	MKCMDUnknown = uint16(0x0000)

	//MKCMDReqConn is magknot's connect request
	MKCMDReqConn = uint16(0x0001)
	// MKCMDRspConn is magknot's connection response
	MKCMDRspConn = uint16(0x0002)

	//MKCMDMsgN2K is message from node to knot
	MKCMDMsgN2K = uint16(0x0003)
	//MKCMDMsgK2N is message from knot to node
	MKCMDMsgK2N = uint16(0x0004)

	// MKCMDReqClose is knot's close request
	MKCMDReqClose = uint16(0x0005)
	//MKCMDRspClose is knot's close response
	MKCMDRspClose = uint16(0x0006)

	//MKCMDReqNewAgent is new agent's request
	MKCMDReqNewAgent = uint16(0x0007)
	//MKCMDRspNewAgent is new agent's response
	MKCMDRspNewAgent = uint16(0x0008)
)
