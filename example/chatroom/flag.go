package main

import (
	"flag"
	"fmt"
)

//Flag Flag
var Flag FlagInfo

//FlagInfo FlagInfo
type FlagInfo struct {
	Version    bool
	Daemon     bool
	Config     string
	CPUProfile string
}

func init() {
	flag.Usage = func() {
		fmt.Printf("Chatroom server\n")
		fmt.Println("Usage: server -[vdcp]")
		flag.PrintDefaults()
	}

	flag.BoolVar(&Flag.Version, "v", false, "Show Chatroom's Version")
	flag.BoolVar(&Flag.Daemon, "d", false, "Start Chatroom as A Daemon")
	flag.StringVar(&Flag.CPUProfile, "cpuprof", "", "Start CPU  Profile for Chatroom")
	flag.StringVar(&Flag.Config, "c", "", "Config File of Chatroom")
	flag.Parse()

}
