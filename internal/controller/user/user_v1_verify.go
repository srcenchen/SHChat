package user

import (
	"SHChat/internal/logic/user"
	"context"

	"SHChat/api/user/v1"
)

func (c *ControllerV1) Verify(ctx context.Context, req *v1.VerifyReq) (res *v1.VerifyRes, err error) {
	res = &v1.VerifyRes{
		IsSuccess: user.VerifyUser(req.Password),
	}
	return
}
