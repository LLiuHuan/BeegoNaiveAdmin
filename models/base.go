package models

import "time"

type BaseModel struct {
	UpdateTime time.Time `orm:"auto_now_add;type(datetime)" json:"updateTime"`
	CreateTime time.Time `orm:"auto_now;type(datetime)" json:"createTime"`
	IsDel      int8      `json:"isDel"`
}
