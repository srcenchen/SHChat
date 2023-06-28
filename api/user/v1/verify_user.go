package v1

import "github.com/gogf/gf/v2/frame/g"

type VerifyReq struct {
	g.Meta   `method:"post" summary:"验证用户密码" tags:"用户"`
	Password string `v:"required#密码不能为空"`
}

type VerifyRes struct {
	g.Meta    `method:"post" summary:"验证用户密码返回"`
	IsSuccess bool `json:"isSuccess"`
}
