package main

import (
	"fmt"
	"flag"
	
	"github.com/cz-it/magline/maglined/server"
	"github.com/cz-it/magline/maglined"
	"github.com/cz-it/golangutils/daemon"
)

func main() {
	if Flag.Version {
		fmt.Println("Cur Version:%s",maglined.Version())
		return
	}

	println("[Testing]:Main....")
	
	if Flag.Config == "" {
		flag.Usage()
		return
	}

	if err := LoadConfig(Flag.Config); err != nil{
		println("Loading Config Error")
		return
	}
	
	if Flag.Daemon {
		daemon.Boot("/tmp/magline.lock","/tmp/magline.pid")
	}
	
	svr := &server.Server{Addr:maglined.Config.OuterAddr}
	svr.ListenAndServe()
	
	println("[Testing]:End")
}

















