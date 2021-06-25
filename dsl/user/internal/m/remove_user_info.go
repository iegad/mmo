package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func (this_ *UserService) RemoveUserInfo(ctx *piper.Context, req *cgi.RemoveUserInfoReq, rsp *cgi.RemoveUserInfoRsp) error {
	utils.Assert(ctx != nil && req != nil, "RemoveUserInfo in params invalid")

	if rsp == nil {
		return ErrRsp
	}

	if len(req.UserIDs) == 0 {
		return ErrReq
	}

	affected := int32(0)

	for dwf := true; dwf; dwf = false {
		for _, userID := range req.UserIDs {
			err := dao.RemoveUserInfo(userID, com.MySql)
			if err != nil {
				log.Error(err)
				continue
			}

			affected++
		}
	}

	if affected != int32(len(req.UserIDs)) {
		rsp.Code = -100
	}

	rsp.AffectedRows = int32(affected)
	return nil
}
