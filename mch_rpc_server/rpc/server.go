/**
 * @author mch
 */

package rpc

import (
	"crypto/tls"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Server struct {
	tlsoption tlsoption
}
func NewServer(fns ...ChangeTlsOptionImpl) *Server {
	defaultOption := tlsoption{}
	for _,fn := range fns {
		fn.apply(&defaultOption)
	}
	return &Server{tlsoption: defaultOption}
}
func(s *Server) Register(sev interface{}) {
	rpc.Register(sev)
}
func(s *Server) Listen(addr string) (net.Listener,error) {
	if !s.tlsoption.openssl {
		return net.Listen("tcp",addr)
	}
	cert,err := tls.LoadX509KeyPair(s.tlsoption.certFile,s.tlsoption.keyFile)
	if err != nil {
		return nil,err
	}
	cnf := tls.Config{Certificates: []tls.Certificate{cert}}
	return tls.Listen("tcp",addr,&cnf)

}
func(s *Server) Run(addr string) error {
	l,err := s.Listen(addr)
	if err != nil {
		return err
	}
	for  {
		conn,err := l.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
