package user

import (
	"SHChat/internal/dao"
)

func VerifyUser(password string) bool {
	return dao.GetUserInfoByPasswd(password) != 0
}
