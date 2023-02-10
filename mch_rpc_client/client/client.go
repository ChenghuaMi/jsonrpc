/**
 * @author mch
 */

package client

import (
	"fmt"
	"mch_rpc_client/global"
	"mch_rpc_client/proto"
)

type rpcClient struct {
	Member Member
}
var RpcClient = new(rpcClient)

func call(serverName string,method string,req proto.MemberReq,resp *proto.MemberResp) error {
	cli,err := global.ServerCnf.Connect(serverName,"rpc/key/server.pem", "rpc/key/server.key")
	if err != nil {
		fmt.Println("error:",err)
		return err
	}
	return cli.Call(method,req,resp)
}

