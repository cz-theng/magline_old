package main

import(
	"encoding/json"
	"github.com/cz-it/magline/magline"
	"os"
)

type ConfigJsonWrapper struct {
	OuterAddr string
}

func LoadConfig(filePath string ) (err error) {
	fp,err := os.Open(filePath)
	if err != nil {
		magline.Logger.Error("Open Config file %s error: %s",filePath,err.Error())
		return
	}
	defer fp.Close()
	

	var config ConfigJsonWrapper
	decoder := json.NewDecoder(fp)
	if err= decoder.Decode(&config); err != nil {
		magline.Logger.Error("Decode Config Error:%s",err.Error())
		return
	}
	magline.Config.OuterAddr = config.OuterAddr
	magline.Logger.Debug("Load Config file %s Success",filePath)
	err = nil
	return
}



















