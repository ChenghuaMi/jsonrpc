/**
 * @author mch
 */

package main

import (
	"mch_rpc_server/rpc"
	"mch_rpc_server/server"
)
func main() {
	var optionFuncs []rpc.ChangeTlsOptionImpl
	optionFuncs = append(optionFuncs,rpc.AddChangeTlsOptionFunc("rpc/key/server.pem","rpc/key/server.key"))
	serv := rpc.NewServer(optionFuncs...)
	serv.Register(server.MemberServer)
	serv.Run(":3000")
}
