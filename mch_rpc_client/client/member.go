/**
 * @author mch
 */

package client

import (

	"mch_rpc_client/proto"
)

type Member struct {

}
func(m *Member) Login(method string,req proto.MemberReq,resp *proto.MemberResp) error {
	return call("member",method,req,resp)

}

