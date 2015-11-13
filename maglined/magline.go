package maglined

/**
* Magline
 */

import (
	"fmt"
	"os"
)

func Exit() {
	os.Exit(-1)
}

func Start() {
	var err error

	agentMgr, err := NewAgentMgr(Config.MaxConns)
	if err != nil {
		fmt.Println("Create Agent Manager Error")
		Logger.Error("Create Agent Manager Error")
		Exit()
	}
	backend := &BackendServer{
		Addr: Config.InnerAddr,
	}
	backend.Init()
	err = backend.Listen()
	if err != nil {
		fmt.Println("Start Magline Backend Server error with s", err.Error())
		Logger.Error("Start Magline Backend Server error with s", err.Error())
		Exit()
	}

	svr := Server{
		Addr:     Config.OuterAddr,
		AgentMgr: agentMgr,
		Backend:  backend,
	}
	svr.Init(Config.MaxConns)
	err = svr.ListenAndServe()
	if err != nil {
		fmt.Println("Start Magline Server error with s", err.Error())
		Logger.Error("Start Magline Server error with s", err.Error())
		Exit()
	}
}
