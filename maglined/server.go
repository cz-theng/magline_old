//Package maglined is a daemon process for connection layer
/**
* Author: CZ cz.theng@gmail.com
 */
package main

import (
	"net"
	"os"
	"time"
)

//BackendServer is backend server
type BackendServer struct {
	Addr   string
	Bridge *Bridge
}

//Dispatch Dispatch a lane
func (bs *BackendServer) Dispatch() (lane *Lane, err error) {
	lane, err = bs.Bridge.Dispatch()
	return
}

//Init is initionlize
func (bs *BackendServer) Init() (err error) {
	bs.Bridge = &Bridge{}
	return
}

//AcceptAndServe accept and serve
func (bs *BackendServer) AcceptAndServe(ln *net.UnixListener) {
	for {
		rw, err := ln.AcceptUnix()
		if err != nil {
			return
		}
		lane, err := bs.Bridge.Alloc(rw)
		lane.Init()
		if err != nil {
			return
		}
		go lane.Serve()
	}
}

// Listen listen the port
func (bs *BackendServer) Listen() (err error) {
	//var tempDelay time.Duration // how long to sleep on accept failure

	addr, err := ParseAddr(bs.Addr)
	if err != nil {
		return err
	}
	Logger.Debug("net is %s and ipport %s", addr.Network, addr.IPPort)

	if addr.Network != "unix" {
		Logger.Error("Error Inner Address, Should be unix://")
		err = ErrAddr
		return
	}
	os.Remove(addr.IPPort)
	l, err := net.Listen("unix", addr.IPPort)
	if err != nil {
		Logger.Error(err.Error())
		return
	}
	ln := l.(*net.UnixListener)
	go bs.AcceptAndServe(ln)
	return
}

//Server is frontend server
type Server struct {
	Addr     string
	ConnPool *ConnPool
	AgentMgr *AgentMgr
	Backend  *BackendServer
}

//Init is server's init
func (svr *Server) Init(maxConns int) (err error) {
	svr.ConnPool, err = NewMLConnPool(maxConns)
	if err != nil {
		Logger.Error("New Magline Connection Pool Error!")
		return
	}
	return
}

//ListenAndServe is server's listen and serve
func (svr *Server) ListenAndServe() error {
	Logger.Debug("ListenAndServe with addr %s", svr.Addr)
	addr, err := ParseAddr(svr.Addr)
	if err != nil {
		return err
	}
	Logger.Debug("net is %s and ipport %s", addr.Network, addr.IPPort)
	if addr.Network == "tcp" {
		ln, err := net.Listen("tcp", addr.IPPort)
		if err != nil {
			return err
		}

		svr.listenAndServeTCP(ln.(*net.TCPListener), addr.Kpal)
	}
	return nil
}

func (svr *Server) newConn(rwc *net.TCPConn) (conn *Connection, err error) {
	conn, err = svr.ConnPool.Alloc()
	if err != nil {
		return
	}
	conn.RWC = rwc
	conn.Server = svr

	return
}

func (svr *Server) listenAndServeTCP(l *net.TCPListener, kpal bool) error {
	defer l.Close()
	var tempDelay time.Duration // how long to sleep on accept failure

	for {
		rw, e := l.AcceptTCP()
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
				Logger.Error("[TCP]: Accept error: %v; retrying in %v", e, tempDelay)
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
