package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func (this_ *UserService) ModifyUserInfo(ctx *piper.Context, req *cgi.ModifyUserInfoReq, rsp *cgi.ModifyUserInfoRsp) error {
	utils.Assert(ctx != nil && req != nil, "ModifyUserInfo in params invalid")

	if rsp == nil {
		return ErrRsp
	}

	if req.UserInfo == nil {
		return ErrReq
	}

	for dwf := true; dwf; dwf = false {
		err := dao.ModifyUserInfo(req.UserInfo, com.MySql)
		if err != nil {
			log.Error(err)
			rsp.Code = -100
			break
		}

		rsp.UserInfo = req.UserInfo
	}

	return nil
}
