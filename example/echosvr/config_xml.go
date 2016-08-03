package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//ConfigJSONWrapper ConfigJsonWrapper
type ConfigJSONWrapper struct {
	OuterAddr string
	Addr      string
	MaxConns  int
	Crypto    string
	Channel   string
}

var config ConfigJSONWrapper

//LoadConfig load configure
func LoadConfig(filePath string) (err error) {
	fp, err := os.Open(filePath)
	if err != nil {
		fmt.Errorf("Open Config file %s error: %s", filePath, err.Error())
		return
	}
	defer fp.Close()

	decoder := json.NewDecoder(fp)
	if err = decoder.Decode(&config); err != nil {
		fmt.Errorf("Decode Config Error:%s", err.Error())
		return
	}
	fmt.Printf("Load Config file %s Success", filePath)
	err = nil
	return
}
