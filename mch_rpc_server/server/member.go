/**
 * @author mch
 */

package server

import (
	"fmt"
	"mch_rpc_server/proto"
)

type Member struct {

}
var MemberServer = new(Member)
func(m *Member) Login(req proto.MemberReq,resp *proto.MemberResp) error {
	fmt.Println("login .......")
	fmt.Printf("req %#v\n",req)
	*resp = proto.MemberResp{
		Id:       "10000000000",
		Username: "michenghua",
		Age:      0,
	}
	return nil
}