package main

import (
	"flag"
	"fmt"
	"github.com/cz-it/golangutils/daemon"
	"github.com/cz-it/magline"
	"os"
	"runtime/pprof"
	"time"
)

func fab(f int32) int32 {
	time.Sleep(20 * time.Millisecond)
	if f == 0 {
		return 0
	} else if f == 1 {
		return 1
	} else {
		return fab(f-1) + fab(f-2)
	}
}

func main() {

	if Flag.Version {
		fmt.Println("Cur Version:%s", magline.Version())
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

	if Flag.CPUProfile != "" {
		for _, p := range pprof.Profiles() {
			fmt.Println("Profile:", p.Name())
		}
		w, err := os.Create(Flag.CPUProfile)
		if err != nil {
			fmt.Printf("Create Profile Error %s \n", err.Error())
		}
		pprof.StartCPUProfile(w)
		defer w.Close()
		stoper := time.NewTimer(60 * time.Second)
		go func() {
			<-stoper.C
			println("profile done")
			pprof.StopCPUProfile()
			w.Close()
			os.Exit(0)
		}()
		//fab(100)
		magline.Start()
	} else if Flag.Daemon {
		daemon.Boot("/tmp/magline.lock", "/tmp/magline.pid", func() {
			magline.Start()
		})
	} else {
		magline.Start()
	}

	println("[Testing]:End")
}
