package magline

/**
* Magline
*/

import (

	"github.com/cz-it/magline/maglined"
	"github.com/cz-it/magline/maglined/agent"
)

func Start() {
	agent.InitAgentMgr(maglined.Config.MaxConns)

	svr := Server{Addr:maglined.Config.OuterAddr}
	svr.Init(maglined.Config.MaxConns)
	err :=svr.ListenAndServe()
	if err != nil {
		maglined.Logger.Error("Start Magline Server error with s",err.Error())
	}
}









