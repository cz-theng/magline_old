package magline

import (
	"fmt"
	"github.com/cz-it/golangutils/log"
)

var Logger *log.Logger

func init() {
	var err error
	Logger,err = log.NewFileLogger("log","magline")
	if err != nil {
		fmt.Errorf("Create Logger Error\n")
		return 
	}
	fmt.Println("Create Logger Success")
	Logger.SetMaxFileSize(1024*1024*100) //100MB
	Logger.SetLevel(log.LDEBUG)
}



















