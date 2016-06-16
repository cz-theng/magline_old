package main

import (
	"encoding/json"
	"github.com/cz-it/magline"
	"github.com/cz-it/magline/proto"
	"github.com/cz-it/magline/utils"
	"os"
)

//ConfigJSONWrapper ConfigJsonWrapper
type ConfigJSONWrapper struct {
	OuterAddr string
	InnerAddr string
	MaxConns  int
	Crypto    string
	Channel   string
}

//LoadConfig load configure
func LoadConfig(filePath string) (err error) {
	fp, err := os.Open(filePath)
	if err != nil {
		utils.Logger.Error("Open Config file %s error: %s", filePath, err.Error())
		return
	}
	defer fp.Close()

	var config ConfigJSONWrapper
	decoder := json.NewDecoder(fp)
	if err = decoder.Decode(&config); err != nil {
		utils.Logger.Error("Decode Config Error:%s", err.Error())
		return
	}
	magline.Config.OuterAddr = config.OuterAddr
	magline.Config.InnerAddr = config.InnerAddr
	magline.Config.MaxConns = config.MaxConns

	switch config.Crypto {
	case "none":
		magline.Config.Crypto = proto.CryptoNone
	case "aes128":
		magline.Config.Crypto = proto.CryptoAES128
	}

	switch config.Channel {
	case "none":
		magline.Config.Channel = proto.ChanNone
	case "salt":
		magline.Config.Channel = proto.ChanSalt
	case "dh":
		magline.Config.Channel = proto.ChanDH
	}

	utils.Logger.Debug("Load Config file %s Success", filePath)
	err = nil
	return
}
