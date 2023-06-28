// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (
	"context"
	
	"SHChat/api/user/v1"
)

type IUserV1 interface {
	Verify(ctx context.Context, req *v1.VerifyReq) (res *v1.VerifyRes, err error)
}


