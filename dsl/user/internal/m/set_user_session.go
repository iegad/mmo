package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func (this_ *UserService) SetUserSession(ctx *piper.Context, req *cgi.SetUserSessionReq, rsp *cgi.SetUserSessionRsp) error {
	utils.Assert(ctx != nil && req != nil, "RemoveUserSession in params invalid")

	if rsp == nil {
		return ErrRsp
	}

	if req.UserSession == nil {
		return ErrReq
	}

	for dwf := true; dwf; dwf = false {
		err := dao.SetUserSess(req.UserSession, com.Redis)
		if err != nil {
			log.Error(err)
			rsp.Code = -100
			break
		}

		rsp.UserSession = req.UserSession
	}

	return nil
}
