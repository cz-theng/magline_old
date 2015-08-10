package magline

/**
* Magline
*/

import (

	"github.com/cz-it/magline/maglined/server"
	"github.com/cz-it/magline/maglined/config"
)

func Start() {
	svr := server.Server{Addr:config.OuterAddr}
	svr.ListenAndServe()
}









