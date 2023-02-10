/**
 * @author mch
 */

package init

import "mch_rpc_client/core"

func init() {
	core.RegisterRoute()
	core.InitRpc()
}
