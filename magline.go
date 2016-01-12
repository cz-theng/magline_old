//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import (
	"fmt"
	"github.com/cz-it/magline/utils"
	"os"
)

func exit() {
	os.Exit(-1)
}

//Start the two servers
func Start() {
	var err error

	agentMgr, err := NewAgentMgr(Config.MaxConns)
	if err != nil {
		fmt.Println("Create Agent Manager Error")
		utils.Logger.Error("Create Agent Manager Error")
		exit()
	}
	backend := &BackendServer{
		Addr: Config.InnerAddr,
	}
	backend.Init()
	err = backend.Listen()
	if err != nil {
		fmt.Println("Start Magline Backend Server error with s", err.Error())
		utils.Logger.Error("Start Magline Backend Server error with s", err.Error())
		exit()
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
		utils.Logger.Error("Start Magline Server error with s", err.Error())
		exit()
	}
}
