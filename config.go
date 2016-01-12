//Package magline is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package magline

import ()

//Config is config info
var Config ConfigInfo

//ConfigInfo is config info's data
type ConfigInfo struct {
	OuterAddr string
	InnerAddr string
	MaxConns  int
}
