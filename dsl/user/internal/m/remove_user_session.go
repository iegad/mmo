package m

import (
	"encoding/hex"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func (this_ *UserService) RemoveUserSession(ctx *piper.Context, req *cgi.RemoveUserSessionReq, rsp *cgi.RemoveUserSessionRsp) error {
	utils.Assert(ctx != nil && req != nil, "RemoveUserSession in params invalid")

	if rsp == nil {
		return ErrRsp
	}

	if req.UserID <= 0 ||
		len(req.Token) != 16 {
		return ErrReq
	}

	for dwf := true; dwf; dwf = false {
		sess, err := dao.GetUserSess(req.UserID, com.Redis)
		if err != nil {
			log.Error(err)
			rsp.Code = -100
			break
		}

		if hex.EncodeToString(sess.Token) != hex.EncodeToString(req.Token) {
			rsp.Code = -101
			break
		}

		err = dao.RemoveUserSess(req.UserID, com.Redis)
		if err != nil {
			log.Error(err)
			rsp.Code = -102
			break
		}

		rsp.UserSession = sess
	}

	return nil
}
