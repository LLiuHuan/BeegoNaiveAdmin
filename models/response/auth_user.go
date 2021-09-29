package response

import "BeegoNaiveAdmin/models"

type Login struct {
	Token string          `json:"token"`
	User  *models.SysUser `json:"user"`
}
