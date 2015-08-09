package main

import (
	"github.com/cz-it/magline/magline"
)

func main() {
	println("[Testing]:Main....")
	if err := LoadConfig("config/config.json"); err != nil{
		println("Loading Config Error")
		return
	}
	
	magline.Start()

	println("[Testing]:End")
}
