package maglined

/**
* Net Utils
 */
import (
	"errors"
	"strings"
)

var (
	EURL     = errors.New("Invaliad URL!")
	ENETWORK = errors.New("Unknown Network Type!")
)

type Addr struct {
	Network string
	IPPort  string
	Kpal    bool
}

func ParseAddr(url string) (addr Addr, err error) {
	if url[:6] == "udp://" {
		addr.Network = "udp"
		addr.IPPort = url[6:]
	} else if url[:6] == "tcp://" {
		addr.Network = "tcp"
		if strings.Contains(url, "keep-alive=false") {
			addr.Kpal = false
		} else {
			addr.Kpal = true
		}
		addr.IPPort = url[6:]
	} else if url[:7] == "unix://" {
		addr.Network = "unix"
		addr.IPPort = url[7:]
	} else {
		addr.Network = "unknown"
		addr.IPPort = "unknown:unknown"
		err = ENETWORK
		return
	}
	return
}
