/**
 * @author mch
 */

package rpc

type Config struct {
	Servers map[string]Conf
}

type Conf struct {
	Network string
	Address string
}
