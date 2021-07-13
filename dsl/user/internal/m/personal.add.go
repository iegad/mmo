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

func (this_ *UserService) AddPersonal(ctx *piper.Context, req *user.AddPersonalReq, rsp *user.AddPersonalRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "AddPersonal in params is invalid")

	if req.Personal == nil {
		rsp.Code = -1
		return nil
	}

	if req.Personal.UserID <= 0 {
		rsp.Code = -2
		return nil
	}

	if req.Personal.Birth <= 0 {
		rsp.Code = -3
		return nil
	}

	if req.Personal.Gender < ds.MIN_GENDER || req.Personal.Gender > ds.MAX_GENDER {
		rsp.Code = -4
		return nil
	}

	if utf8.RuneCountInString(req.Personal.Name) > ds.MAX_NAME || len(req.Personal.Name) == 0 {
		rsp.Code = -5
		return nil
	}

	if len(req.Personal.Nationality) == 0 || len(req.Personal.Nationality) > ds.MAX_NATIONALITY {
		rsp.Code = -6
		return nil
	}

	if len(req.Personal.ID) == 0 {
		rsp.Code = -7
		return nil
	}

	err := personal.AddPersonal(req.Personal, com.MySql)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
	}

	return nil
}
