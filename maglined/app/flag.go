package main

import (
	"flag"
	"fmt"
)

//Flag Flag
var Flag FlagInfo

//FlagInfo FlagInfo
type FlagInfo struct {
	Version bool
	Daemon  bool
	Config  string
}

func init() {
	flag.Usage = func() {
		fmt.Printf("Magline server\n")
		fmt.Println("Usage: server -[vdc]")
		flag.PrintDefaults()
	}

	flag.BoolVar(&Flag.Version, "v", false, "Show Magline's Version")
	flag.BoolVar(&Flag.Daemon, "d", false, "Start Magline as A Daemon")
	flag.StringVar(&Flag.Config, "c", "", "Config File of Magline")
	flag.Parse()

}
