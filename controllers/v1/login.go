package v1

import (
	"BeegoNaiveAdmin/controllers"
	"BeegoNaiveAdmin/errno"
	"BeegoNaiveAdmin/models"
	"BeegoNaiveAdmin/models/request"
	"BeegoNaiveAdmin/models/response"
	"BeegoNaiveAdmin/utils"
	"encoding/json"
	"fmt"
	"time"
)

// LoginController 登录api
type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
}

// Login 登录
// @Title 登录
// @Description 登录
// @Success 200 {object} controllers.Response
// @router /login [post]
func (c *LoginController) Login() {
	var authUser *request.AuthUser

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &authUser)

	if err != nil {
		fmt.Println(err)
		c.FailWithCodeMessage(errno.ErrLoginParameter.Code, errno.ErrLoginParameter.Message)
	} else {
		currentUser, e := models.GetUserByUsername(authUser.Username)

		if e != nil {
			c.FailWithCodeMessage(errno.ErrLoginNotExist.Code, errno.ErrLoginNotExist.Message)
		}

		if !utils.ComparePwd(currentUser.Password, []byte(authUser.Password)) {
			c.FailWithCodeMessage(errno.ErrPasswordError.Code, errno.ErrPasswordError.Message)
		} else {
			token, _ := utils.GenerateToken(currentUser, time.Hour*24*100)
			var login = new(response.Login)
			login.Token = token
			login.User = currentUser
			c.OkWithData(login)
		}
	}
}
