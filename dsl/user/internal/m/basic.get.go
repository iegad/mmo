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
			rsp.Code = -100
		} else {
			rsp.BasicList = append([]*ds.Basic{}, basic)
			rsp.Total = int64(len(rsp.BasicList))
		}

		return nil
	}

	dataList, total, err := basic.GetBasic(req, com.MySql)
	if err != nil {
		log.Error(err)
		rsp.Code = -100
	} else {
		rsp.BasicList = dataList
		rsp.Total = total
	}

	return nil
}
