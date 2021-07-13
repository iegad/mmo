package m

import (
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/personal"
)

func (this_ *UserService) ModPersonal(ctx *piper.Context, req *user.ModPersonalReq, rsp *user.ModPersonalRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "ModPersonal in parmas is invalid")

	if req.Personal == nil {
		rsp.Code = -100
		return nil
	}

	if req.Personal.Gender > 0 && req.Personal.Gender > 2 {
		rsp.Code = -1
		return nil
	}

	if req.Personal.UserID <= 0 {
		rsp.Code = -3
		return nil
	}

	if req.Personal.Birth < 0 {
		rsp.Code = -4
		return nil
	}

	if utf8.RuneCountInString(req.Personal.Name) > ds.MAX_NAME {
		rsp.Code = -5
		return nil
	}

	if len(req.Personal.ID) > ds.MAX_ID {
		rsp.Code = -6
		return nil
	}

	if len(req.Personal.Nationality) > ds.MAX_NATIONALITY {
		rsp.Code = -7
		return nil
	}

	err := personal.ModPersonal(req.Personal, com.MySql, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
	}

	return nil
}
