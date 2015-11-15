package proto

type MagNodeCMD uint16

const (
	MN_MAGIC   = 0x7f
	MN_VERSION = 0x01

	MN_CMD_UNKNOWN = uint16(0x0000)

	MN_CMD_REQ_CONN = uint16(0x0001)
	MN_CMD_RSP_CONN = uint16(0x0002)

	MN_CMD_REQ_SEND = uint16(0x0003)
	MN_CMD_RSP_SEND = uint16(0x0004)

	MN_CMD_REQ_RECV = uint16(0x0005)
	MN_CMD_RSP_RECV = uint16(0x0006)

	MN_CMD_REQ_CLOSE = uint16(0x0007)
	MN_CMD_RSP_CLOSE = uint16(0x0008)

	MN_CMD_REQ_RECONN = uint16(0x0009)
	MN_CMD_RSPREQCONN = uint16(0x000a)

	MN_CMD_MSG_NODE = uint16(0x000b)
	MN_CMD_MSG_KNOT = uint16(0x000c)
)

type MagKnotCMD uint16

const (
	MK_MAGIC   = 0x7f
	MK_VERSION = 0x01

	MK_CMD_UNKNOWN = uint16(0x0000)

	MK_CMD_REQ_CONN = uint16(0x0001)
	MK_CMD_RSP_CONN = uint16(0x0002)

	MK_CMD_MSG_N2K = uint16(0x0003)
	MK_CMD_MSG_K2N = uint16(0x0004)

	MK_CMD_REQ_CLOSE = uint16(0x0005)
	MK_CMD_RSP_CLOSE = uint16(0x0006)

	MK_CMD_REQ_NEWAGENT = uint16(0x0007)
	MK_CMD_RSP_NEWAGENT = uint16(0x0008)
)
