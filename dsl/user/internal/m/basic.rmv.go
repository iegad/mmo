package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/basic"
)

// RmvBasic 移除用户基础信息
func (this_ *UserService) RmvBasic(ctx *piper.Context, req *user.RmvBasicReq, rsp *user.RmvBasicRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "RmvBasic in params is invalid")

	raw, err := basic.GetBasicByID(req.UserID, com.MySql)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
		return nil
	}

	err = basic.RmvBasic(req.UserID, com.MySql, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -10001
		return nil
	}

	rsp.Basic = raw
	return nil
}
