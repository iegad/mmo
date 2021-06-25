package m

import (
	"github.com/iegad/kraken/piper"
	"github.com/iegad/mmo/cgi"
)

func checkAddUserInfoReq(req *cgi.AddUserInfoReq) bool {
	if req == nil {
		return false
	}

	if req.UserInfo == nil {
		return false
	}

	if len(req.UserInfo.Email) == 0 && len(req.UserInfo.PhoneNum) == 0 {
		return false
	}

	return true
}

func (this_ *UserService) AddUserInfo(ctx *piper.Context, req *cgi.AddUserInfoReq, rsp *cgi.AddUserInfoRsp) error {
	if ctx == nil {
		return ErrCtx
	}

	if req == nil {
		return ErrReq
	}

	if rsp == nil {
		return ErrRsp
	}

	if !checkAddUserInfoReq(req) {
		return ErrReq
	}

	return nil
}
