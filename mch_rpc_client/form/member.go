/**
 * @author mch
 */

package form

type Member struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password"`
}
