/**
 * @author mch
 */

package main

import  (
	_ "mch_rpc_client/init"
	"mch_rpc_client/router"
)
func main() {
	router.InitRouter().Run(":3001")
}
