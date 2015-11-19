//Package maglined is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package maglined

import (
	"fmt"
	"github.com/cz-it/golangutils/log"
)

//Logger is logcat used by all app
var Logger *log.Logger

func init() {
	var err error
	Logger, err = log.NewFileLogger("log", "magline")
	if err != nil {
		fmt.Errorf("Create Logger Error\n")
		return
	}
	Logger.SetMaxFileSize(1024 * 1024 * 100) //100MB
	Logger.SetLevel(log.LDEBUG)
}
