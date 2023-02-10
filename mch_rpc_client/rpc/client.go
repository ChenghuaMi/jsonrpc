/**
 * @author mch
 */

package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ServerCnf struct {
	*Config
}

func NewServerCfg(cfg *Config) *ServerCnf {
	return &ServerCnf{
		cfg,
	}
}
func(s *ServerCnf) Connect(servername string,certFile,keyFile string) (client *Client,err error) {
	if len(s.Servers) > 0 {
		cfg,ok := s.Servers[servername]
		if !ok {
			return &Client{},errors.New("服务名称不存在")
		}
		conn,err := s.Dial(cfg,certFile,keyFile)
		//conn,err := jsonrpc.Dial(cfg.Network,cfg.Address)
		if err != nil {
			return &Client{},errors.New("链接失败")
		}
		return &Client{conn: conn},nil
	}else {
		return &Client{},errors.New("服务不存在")
	}
}
func(s *ServerCnf) Dial(cfg Conf,certFile,keyFile string) (*rpc.Client,error) {
	if certFile == "" && keyFile == "" {
		return jsonrpc.Dial(cfg.Network,cfg.Address)
	}
	certPool := x509.NewCertPool()
	byt,_ := ioutil.ReadFile("rpc/key/server.pem")
	certPool.AppendCertsFromPEM(byt)
	conn,err := tls.Dial(cfg.Network,cfg.Address,&tls.Config{RootCAs: certPool,ServerName: "my_server_name"})
	return jsonrpc.NewClient(conn),err

}
type Client struct {
	conn *rpc.Client
}

func(c *Client) Call(method string,args interface{},reply interface{}) error {
	return c.conn.Call(method,args,reply)
}
