package maglined

/**
* Magline
*/

import (
)

func Start() {
	InitAgentMgr(Config.MaxConns)

	svr := Server{Addr:Config.OuterAddr}
	svr.Init(Config.MaxConns)
	err :=svr.ListenAndServe()
	if err != nil {
		Logger.Error("Start Magline Server error with s",err.Error())
	}
}









