package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/basic"
)

func (this_ *UserService) GetBasic(ctx *piper.Context, req *user.GetBasicReq, rsp *user.GetBasicRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "GetBasic in params is invalid")

	if req.UserID > 0 {
		basic, err := basic.GetBasicByID(req.UserID, com.MySql)
		if err != nil {
			log.Error(err)
			rsp.Code = -10000
		} else {
			rsp.BasicList = append([]*ds.Basic{}, basic)
			rsp.Total = int64(len(rsp.BasicList))
		}

		return nil
	}

	if req.UserID < 0 {
		rsp.Code = -3
		return nil
	}

	if req.CreateTimeBeg < 0 {
		rsp.Code = -4
		return nil
	}

	if req.CreateTimeEnd < 0 {
		rsp.Code = -5
		return nil
	}

	if req.LastUpdateBeg < 0 {
		rsp.Code = -6
		return nil
	}

	if req.LastUpdateEnd < 0 {
		rsp.Code = -7
		return nil
	}

	if req.Offset < 0 {
		rsp.Code = -8
		return nil
	}

	if req.Limit < 0 {
		rsp.Code = -9
		return nil
	}

	if len(req.PhoneNum) > ds.MAX_PHONE_NUM {
		rsp.Code = -11
		return nil
	}

	if len(req.Email) > ds.MAX_EMAIL {
		rsp.Code = -12
		return nil
	}

	dataList, total, err := basic.GetBasic(req, com.MySql)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
	} else {
		rsp.BasicList = dataList
		rsp.Total = total
	}

	return nil
}
