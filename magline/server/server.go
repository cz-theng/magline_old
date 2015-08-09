package server
/**
* Server.
*/

type Server struct {
   /**
    * Address to listen such as 
    * "tcp://114.1.0.1?keep-alive=true"
    * "tcp://114.1.0.1:80?keep-alive=false"
    * "udp://114.1.0.1:8088"
    */
	Addr string 
}

func (s *Server) ListenAndServe() error {
	return nil
}
