package main

import(
	"encoding/json"
	"github.com/cz-it/magline/magline"
	"os"
)

func LoadConfig(filePath string ) (err error) {
	fp,err := os.Open(filePath)
	if err != nil {
		magline.Logger.Error("Open Config file %s error: %s",filePath,err.Error())
		return
	}
	defer fp.Close()
	
	decoder := json.NewDecoder(fp)
	if err= decoder.Decode(&magline.Config); err != nil {
		magline.Logger.Error("Decode Config Error:%s",err.Error())
		return
	}

	magline.Logger.Debug("Load Config file %s Success",filePath)
	err = nil
	return
}



















