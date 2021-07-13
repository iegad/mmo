package m

import (
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/iegad/mmo/dsl/user/internal/dao/personal"
)

func (this_ *UserService) GetPersonal(ctx *piper.Context, req *user.GetPersonalReq, rsp *user.GetPersonalRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "GetPersonal in parmas is invalid")

	if req.BirthBeg < 0 {
		rsp.Code = -1
		return nil
	}

	if req.BirthEnd < 0 {
		rsp.Code = -1
		return nil
	}

	if req.Gender > 0 && req.Gender != 1 && req.Gender != 2 {
		rsp.Code = -2
		return nil
	}

	if len(req.ID) > ds.MAX_ID {
		rsp.Code = -3
		return nil
	}

	if utf8.RuneCountInString(req.Name) > ds.MAX_NAME {
		rsp.Code = -4
		return nil
	}

	if len(req.Nationality) > ds.MAX_NATIONALITY {
		rsp.Code = -5
		return nil
	}

	if req.UserID < 0 {
		rsp.Code = -6
		return nil
	}

	if len(req.OrderBy) > 0 {
		if _, ok := dao.TUserPersonalFieldMap[req.OrderBy]; !ok {
			rsp.Code = -7
			return nil
		}
	}

	if req.UserID > 0 {
		pers, err := personal.GetPersonalByID(req.UserID, com.MySql)
		if err != nil {
			log.Error(err)
			rsp.Code = -10000
			return nil
		}

		rsp.PersonalList = append([]*ds.Personal{}, pers)
		rsp.Total = int64(len(rsp.PersonalList))
		return nil
	}

	dataList, total, err := personal.GetPersonal(req, com.MySql)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
		return nil
	}

	rsp.PersonalList = dataList
	rsp.Total = total
	return nil
}
