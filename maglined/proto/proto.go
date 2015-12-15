package proto

// Protoer is protoer
type Protoer interface {
	ReadHead()
	ReadBody()
	WriteResponse([]byte) error
}

// Proto is proto
type Proto struct {
}
