package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func (this_ *UserService) GetUserSession(ctx *piper.Context, req *cgi.GetUserSessionReq, rsp *cgi.GetUserSessionRsp) error {
	utils.Assert(ctx != nil && req != nil, "GetUserSession in params invalid")

	if rsp == nil {
		return ErrRsp
	}

	for dwf := true; dwf; dwf = false {
		sess, err := dao.GetUserSess(req.UserID, com.Redis)
		if err != nil {
			log.Error(err)
			rsp.Code = -100
			break
		}

		rsp.UserSession = sess
	}

	return nil
}
