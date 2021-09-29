package errno

import "net/http"

var (
	OK                  = &Errno{Code: http.StatusOK, Message: "OK"}
	InternalServerError = &Errno{Code: http.StatusInternalServerError, Message: "InternalServerError"}

	ErrLoginParameter   = &Errno{Code: 10101, Message: "用户名或密码无法解析。"}
	ErrLoginNotExist    = &Errno{Code: 20101, Message: "用户不存在"}
	ErrPasswordError    = &Errno{Code: 20102, Message: "密码错误"}
	ErrPasswordGenerate = &Errno{Code: 10102, Message: "生成密码出错"}
)
