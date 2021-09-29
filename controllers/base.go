package controllers

import (
	"BeegoNaiveAdmin/errno"
	"BeegoNaiveAdmin/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var uuid, _ = utils.CreateWorker(0, 0)

// utils.ConvertToString(uuid.NextId())

func Result(code int, data interface{}, msg string) Response {
	var r Response
	r.Code = code
	r.Data = data
	r.Msg = msg
	return r
}

func (c *BaseController) Ok() {
	c.Data["json"] = Result(errno.OK.Code, map[string]interface{}{}, errno.OK.Message)
	c.ServeJSON()
}

func (c *BaseController) OkWithData(data interface{}) {
	c.Data["json"] = Result(errno.OK.Code, data, errno.OK.Message)
	c.ServeJSON()
}

func (c *BaseController) OkWithMessage(msg string) {
	c.Data["json"] = Result(errno.OK.Code, map[string]interface{}{}, msg)
	c.ServeJSON()
}

func (c *BaseController) OkWithCodeMessage(code int, msg string) {
	c.Data["json"] = Result(code, map[string]interface{}{}, msg)
	c.ServeJSON()
}

func (c *BaseController) OkWithDetailed(data interface{}, msg string) {
	c.Data["json"] = Result(errno.OK.Code, data, msg)
	c.ServeJSON()
}

func (c *BaseController) Fail() {
	c.Data["json"] = Result(errno.InternalServerError.Code, map[string]interface{}{}, errno.InternalServerError.Message)
	c.ServeJSON()
}

func (c *BaseController) FailWithData(data interface{}) {
	c.Data["json"] = Result(errno.InternalServerError.Code, data, errno.InternalServerError.Message)
	c.ServeJSON()
}

func (c *BaseController) FailWithMessage(msg string) {
	c.Data["json"] = Result(errno.InternalServerError.Code, map[string]interface{}{}, msg)
	c.ServeJSON()
}

func (c *BaseController) FailWithCodeMessage(code int, msg string) {
	c.Data["json"] = Result(code, map[string]interface{}{}, msg)
	c.ServeJSON()
}

func (c *BaseController) FailWithDetailed(data interface{}, msg string) {
	c.Data["json"] = Result(errno.InternalServerError.Code, data, msg)
	c.ServeJSON()
}
