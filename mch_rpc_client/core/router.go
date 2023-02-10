/**
 * @author mch
 */
package core

import (
	"github.com/gin-gonic/gin"
	"mch_rpc_client/api"
	"mch_rpc_client/router"
)

func RegisterRoute() {
	router.InsertRouters(Login)
}

func Login(g *gin.Engine) {
	group := g.Group("login")
	{
		group.GET("/login",api.MemberService.Login)
	}

}

