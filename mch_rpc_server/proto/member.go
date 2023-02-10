/**
 * @author mch
 */

package proto

type MemberReq struct {
	Username string
	Password string
}

type MemberResp struct {
	Id string
	Username string
	Age int
}
