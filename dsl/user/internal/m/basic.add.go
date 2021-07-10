package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/basic"
)

func (this_ *UserService) AddBasic(ctx *piper.Context, req *user.AddBasicReq, rsp *user.AddBasicRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "AddBasic in params is invalid")

	// Step 1: 入参检查
	if req.Basic == nil {
		rsp.Code = -1
		return nil
	}

	// Step 2: DAO调用
	err := basic.AddBasic(req.Basic, com.MySql, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -100
		return nil
	}

	// Step 3: 应答
	rsp.Basic = req.Basic
	return nil
}
