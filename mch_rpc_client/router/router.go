/**
 * @author mch
 */

package router

import "github.com/gin-gonic/gin"
func InitRouter() (*gin.Engine) {
	g := gin.Default()
	for _,rt := range routes {
		rt(g)
	}

	return g
}

type route func(g *gin.Engine)

var routes []route

func InsertRouters(rts ...route) {
	routes = append(routes,rts...)
}
