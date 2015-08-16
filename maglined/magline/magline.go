package magline

/**
* Magline
*/

import (

	"github.com/cz-it/magline/maglined"
	"github.com/cz-it/magline/maglined/server"
)

func Start() {
	svr := server.Server{Addr:maglined.config.OuterAddr}
	svr.ListenAndServe()
}









