package server
/**
* Server.
*/

import (
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
}

func (s *Server) ListenerAndServe() error {
	s.ListenAndServeTCP(nil)
}

func (s *Server) ListenAndServeTCP(l net.Listener) error {
	defer l.Close()
	var tempDelay time.Duration // how long to sleep on accept failure

	for {
		rw, e := l.Accept()
		if e != nil {
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
		c, err := srv.newConn(rw)
		if err != nil {
			continue
		}
		go c.serve()
	}
	return nil
}
