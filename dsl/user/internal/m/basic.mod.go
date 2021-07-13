package m

import (
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/basic"
)

// ModBasic 修改用户基础信息
func (this_ *UserService) ModBasic(ctx *piper.Context, req *user.ModBasicReq, rsp *user.ModBasicRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "ModBasic in params is invalid")

	if req.Basic == nil {
		rsp.Code = -100
		return nil
	}

	if req.Basic.Entry == nil {
		rsp.Code = -101
		return nil
	}

	if req.Basic.Entry.Gender < ds.MIN_GENDER || req.Basic.Entry.Gender > ds.MAX_GENDER {
		rsp.Code = -1
		return nil
	}

	if req.Basic.Entry.UserID <= 0 {
		rsp.Code = -2
		return nil
	}

	if len(req.Basic.Entry.Email) > ds.MAX_EMAIL {
		rsp.Code = -3
		return nil
	}

	if utf8.RuneCountInString(req.Basic.Entry.Nickname) > ds.MAX_NICKNAME {
		rsp.Code = -4
		return nil
	}

	if len(req.Basic.Entry.PhoneNum) > ds.MAX_PHONE_NUM {
		rsp.Code = -5
		return nil
	}

	if len(req.Basic.Entry.Avator) > ds.MAX_AVATOR {
		rsp.Code = -6
		return nil
	}

	err := basic.ModBasic(req.Basic, com.MySql, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
		return nil
	}

	rsp.Basic = req.Basic
	return nil
}
