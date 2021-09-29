package utils

import (
	"BeegoNaiveAdmin/errno"

	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt 加密
func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", errno.ErrPasswordGenerate
	}
	return string(hash), nil
}

// ComparePwd 密码验证
func ComparePwd(hashPwd string, plainPwd []byte) bool {
	logs.Info(hashPwd)
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		logs.Error(err.Error())
		return false
	}

	return true
}
