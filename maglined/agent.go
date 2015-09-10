package maglined
/**
* Agent.
*/

type Agent struct {
	index int
	id    int
}

func (ag *Agent) ID() (int) {
	return ag.id
}

func (ag *Agent) Index() (int) { 
	return ag.index
}


func (ag *Agent) DealRequest(req *Request)(err error) {
	Logger.Info("Deal a Client Request!")
	return nil
}









