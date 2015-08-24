package maglined
/**
* Base Request 
*/

type Requester interface {
	CMD (uint8)
	Data([]byte)
}

type Request struct {

}
