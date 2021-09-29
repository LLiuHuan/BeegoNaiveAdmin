package utils

import (
	"BeegoNaiveAdmin/models"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/core/logs"
	"github.com/dgrijalva/jwt-go"
)

type userClaims struct {
	models.JwtUser
	jwt.StandardClaims
}

var (
	verifyKey string
)

func init() {
	verifyKey, _ = beego.AppConfig.String("jwt_token")
}

const bearerLength = len("Bearer ")

func GenerateToken(m *models.SysUser, d time.Duration) (string, error) {
	m.Password = ""
	//m.Permissions = []string{}
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Id:        m.Id,
		Issuer:    "YshopGo",
	}

	var jwtUser = models.JwtUser{
		Id:       m.Id,
		Avatar:   m.Avatar,
		Email:    m.Email,
		Username: m.Username,
		Phone:    m.Phone,
		NickName: m.NickName,
		Sex:      m.Sex,
		Dept:     m.Depts.Name,
		Job:      m.Jobs.Name,
	}

	uClaims := userClaims{
		StandardClaims: stdClaims,
		JwtUser:        jwtUser,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	tokenString, err := token.SignedString([]byte(verifyKey))
	if err != nil {
		logs.Info(err)
	}

	return tokenString, err
}
