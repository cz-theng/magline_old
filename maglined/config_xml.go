package main

import (
	"encoding/json"
	"os"
)

//ConfigJSONWrapper ConfigJsonWrapper
type ConfigJSONWrapper struct {
	OuterAddr string
	InnerAddr string
	MaxConns  int
}

//LoadConfig load configure
func LoadConfig(filePath string) (err error) {
	fp, err := os.Open(filePath)
	if err != nil {
		Logger.Error("Open Config file %s error: %s", filePath, err.Error())
		return
	}
	defer fp.Close()

	var config ConfigJSONWrapper
	decoder := json.NewDecoder(fp)
	if err = decoder.Decode(&config); err != nil {
		Logger.Error("Decode Config Error:%s", err.Error())
		return
	}
	Config.OuterAddr = config.OuterAddr
	Config.InnerAddr = config.InnerAddr
	Config.MaxConns = config.MaxConns
	Logger.Debug("Load Config file %s Success", filePath)
	err = nil
	return
}
