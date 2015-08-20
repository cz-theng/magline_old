package magline
/**
* Server.
*/

import (
	"net"
	"time"

	"github.com/cz-it/magline/maglined"
)

type Server struct {
   /**
    * Address to listen such as 
    * "tcp://114.1.0.1?keep-alive=true"
    * "tcp://114.1.0.1:80?keep-alive=false"
    * "udp://114.1.0.1:8088"
    */
	Addr string 

	/**
     * Coonection pool for client
    */
	ConnPool *maglined.ConnPool
}

func (svr *Server) Init (maxConns int) (err error) {
	svr.ConnPool, err = NewMLConnPool(maxConns)
	if err != nil {
		maglined.Logger.Error("New Magline Connection Pool Error!")
		return 
	}
	return 
}

func (svr *Server) ListenAndServe() error {
	maglined.Logger.Debug("ListenAndServe with addr %s",svr.Addr)
	addr, err := maglined.ParseAddr(svr.Addr)
	if err != nil {
		return err
	}
	maglined.Logger.Debug("net is %s and ipport %s",addr.Network,addr.IPPort)
	if addr.Network == "tcp" {
		ln, err := net.Listen("tcp", addr.IPPort)
		if err != nil {
			return err
		}
		
		svr.ListenAndServeTCP(ln.(*net.TCPListener),addr.Kpal)
	}
	return nil
}

func (svr *Server) newConn(rwc *net.TCPConn) (conn *Connection, err error) {
	c, err := svr.ConnPool.Alloc()
	if err != nil {
		return
	}
	conn = c.(*Connection)
	conn.RWC = rwc

	return 
}

func (svr *Server) ListenAndServeTCP(l *net.TCPListener, kpal bool) error {
	defer l.Close()
	var tempDelay time.Duration // how long to sleep on accept failure
	
	for {
		rw, e := l.AcceptTCP()
		println(rw)
		if e != nil {
			if kpal {
				rw.SetKeepAlive(true)
				//rw.SetKeepAlivePeriod()
			}

			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				maglined.Logger.Error("[TCP]: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0

		c, err := svr.newConn(rw)
		if err != nil {
			continue
		}
		go c.Serve()
	}
	return nil
}
