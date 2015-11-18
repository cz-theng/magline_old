package main

import (
	"flag"
	"fmt"

	"github.com/cz-it/golangutils/daemon"
	"github.com/cz-it/magline/maglined"
)

func main() {
	if Flag.Version {
		fmt.Println("Cur Version:%s", maglined.Version())
		return
	}

	if Flag.Config == "" {
		flag.Usage()
		return
	}

	if err := LoadConfig(Flag.Config); err != nil {
		println("Loading Config Error")
		return
	}

	if Flag.Daemon {
		daemon.Boot("/tmp/magline.lock", "/tmp/magline.pid", func() {
			maglined.Start()
		})
	} else {
		maglined.Start()
	}

	println("[Testing]:End")
}
