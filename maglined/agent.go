package maglined

/**
* Agent.
 */
import ()

type Agent struct {
	id uint32
}

func (ag *Agent) ID() uint32 {
	return ag.id
}

func (ag *Agent) DealRequest(req *Request) (err error) {
	Logger.Info("Deal a Client Request!")
	return nil
}
