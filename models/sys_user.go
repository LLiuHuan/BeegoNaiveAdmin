package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(SysUser))
}

type SysUser struct {
	Id       string `json:"id" orm:"pk"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Enabled  int8   `json:"enabled"`
	Password string `json:"password"`
	Username string `json:"username" valid:"Required;"`
	//DeptId int32
	Phone string `json:"phone"`
	//JobId int32
	NickName    string     `json:"nickName"`
	Sex         string     `json:"sex"`
	Roles       []*SysRole `json:"roles" orm:"rel(m2m);rel_through(BeegoNaiveAdmin/models.SysUsersRoles)"`
	Jobs        *SysJob    `json:"job" orm:"column(job_id);rel(one)"`
	Depts       *SysDept   `json:"dept" orm:"column(dept_id);rel(one)"`
	Permissions []string   `orm:"-"`
	RoleIds     []string   `json:"roleIds" orm:"-"`
	BaseModel
}

// GetUserByUsername 根据用户名返回
func GetUserByUsername(name string) (v *SysUser, err error) {
	o := orm.NewOrm()
	user := &SysUser{}
	err = o.QueryTable(new(SysUser)).Filter("username", name).RelatedSel().One(user)
	if _, err = o.LoadRelated(user, "Roles"); err != nil {
		return nil, err
	}
	if err == nil {
		permissions, _ := FindByUserId(user.Id)
		user.Permissions = permissions
		return user, nil
	}

	return nil, err
}

func FindByUserId(id string) ([]string, error) {
	o := orm.NewOrm()
	var roles []SysRole
	_, err := o.Raw("SELECT r.* FROM sys_role r, sys_users_roles u "+
		"WHERE r.id = u.role_id AND u.user_id = ?", id).QueryRows(&roles)
	for k, _ := range roles {
		_, err = o.LoadRelated(&roles[k], "Menus")
	}

	var permissions []string

	for _, v := range roles {
		menus := v.Menus
		for _, m := range menus {
			if m.Permission == "" {
				continue
			}
			permissions = append(permissions, m.Permission)
		}
	}

	return permissions, err
}
