package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func (this_ *UserService) QueryUserInfo(ctx *piper.Context, req *cgi.QueryUserInfoReq, rsp *cgi.QueryUserInfoRsp) error {
	utils.Assert(ctx != nil && req != nil, "QueryUserInfo in params invalid")

	if rsp == nil {
		return ErrRsp
	}

	for dwf := true; dwf; dwf = false {
		dataList, err := dao.QueryUserInfo(req, com.MySql)
		if err != nil {
			log.Error(err)
			rsp.Code = -100
			break
		}

		rsp.Total, err = dao.CountUserInfo(req, com.MySql)
		if err != nil {
			log.Error(err)
			rsp.Code = -101
			break
		}

		rsp.UserList = dataList
	}

	return nil
}
