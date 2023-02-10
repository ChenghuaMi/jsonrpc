/**
 * @author mch
 */

package api

import (
	"github.com/gin-gonic/gin"
	"mch_rpc_client/client"
	"mch_rpc_client/form"
	"mch_rpc_client/proto"
)

type Member struct {

}
var MemberService = new(Member)
func(m *Member) Login(ctx *gin.Context) {
	req := new(form.Member)
	if err := ctx.ShouldBind(req);err != nil {
		ctx.JSON(500,gin.H{
			"msg": err.Error(),
		})
		return
	}
	var resp proto.MemberResp
	err := client.RpcClient.Member.Login("Member.Login",proto.MemberReq{
		Username: req.Username,
		Password: req.Password,
	},&resp)
	if err != nil {
		ctx.JSON(500,gin.H{
			"rpc_error": err.Error(),
			"code": 10000,
		})
		return
	}
	ctx.JSON(200,gin.H{
		"data": resp,
		"code": 0,
	})
	return

}
