/**
 * @author mch
 */

package core

import (
	"mch_rpc_client/global"
	"mch_rpc_client/rpc"
)

func InitRpc() {
	cfg := rpc.Config{
		Servers: map[string]rpc.Conf{
			"member": {
				Network: "tcp",
				Address: ":3000",
						},
		},
	}
	global.ServerCnf = rpc.NewServerCfg(&cfg)
}
