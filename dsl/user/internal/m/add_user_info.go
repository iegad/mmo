package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func (this_ *UserService) AddUserInfo(ctx *piper.Context, req *cgi.AddUserInfoReq, rsp *cgi.AddUserInfoRsp) error {
	utils.Assert(ctx != nil && req != nil, "AddUserInfo in params invalid")

	if rsp == nil {
		return ErrRsp
	}

	if req.UserInfo == nil {
		return ErrReq
	}

	for dwf := true; dwf; dwf = false {
		err := dao.AddUserInfo(req.UserInfo, com.MySql)
		if err != nil {
			log.Error(err)
			rsp.Code = -100
			break
		}

		rsp.UserInfo = req.UserInfo
	}

	return nil
}
