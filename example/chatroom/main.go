/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"flag"
	"fmt"
	"github.com/cz-it/serverkit/daemon"
)

func main() {

	if Flag.Version {
		fmt.Println("Cur Version:%s", Version())
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
	Init()
	if Flag.Daemon {
		daemon.Boot("/tmp/chatroom.lock", "/tmp/chatroom.pid", func() {
			Start()
		})
	} else {
		Start()
	}
	fmt.Println("Echo Server !")
}
