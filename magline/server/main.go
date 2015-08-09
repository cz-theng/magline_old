package main

import (
	"fmt"
	"flag"
	
	"github.com/cz-it/magline/magline"
	"github.com/cz-it/golangutils/daemon"
)

func main() {
	if Flag.Version {
		fmt.Println("Cur Version:%s",magline.Version())
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
	
	magline.Start()

	println("[Testing]:End")
}










