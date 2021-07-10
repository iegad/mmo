package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/basic"
)

func (this_ *UserService) ModBasic(ctx *piper.Context, req *user.ModBasicReq, rsp *user.ModBasicRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "ModBasic in params is invalid")

	if req.Basic == nil {
		rsp.Code = -1
		return nil
	}

	err := basic.ModBasic(req.Basic, com.MySql, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -100
		return nil
	}

	rsp.Basic = req.Basic
	return nil
}
