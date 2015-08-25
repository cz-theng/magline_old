package agent
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










