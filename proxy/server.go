package proxy

import (
	"io"
	"net"
	"xd/lib/log"
)

type Server struct {
	remoteName  string
	DCCListener net.Listener
	Dial        func(string, string) (net.Conn, error)
}

func (s *Server) Serve(l net.Listener) (err error) {
	for err == nil {
		var c net.Conn
		c, err = l.Accept()
		if err == nil {
			go s.handleAccept(c)
		}
	}
	return
}

func (s *Server) handleAccept(c net.Conn) {
	log.Debugf("new connection from %s", c.RemoteAddr())
	conn, err := s.Dial("tcp", s.remoteName)
	if err != nil {
		log.Warnf("failed to connect: %s", err.Error())
		c.Close()
		return
	}

	readDone := make(chan error)
	writeDone := make(chan error)

	copyConn := func(dst io.Writer, src io.Reader, chnl chan error) {
		var buff [1024]byte
		_, e := io.CopyBuffer(dst, src, buff[:])
		chnl <- e
	}

	go copyConn(c, conn, readDone)
	go copyConn(conn, c, writeDone)
	select {
	case err = <-readDone:
		c.Close()
		conn.Close()
	case err = <-writeDone:
		c.Close()
		conn.Close()
	}
}

func NewServer(name string) *Server {
	return &Server{
		remoteName: name,
	}
}
