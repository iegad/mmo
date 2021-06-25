package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/data"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func (this_ *UserService) AddUserList(ctx *piper.Context, req *cgi.AddUserListReq, rsp *cgi.AdduserListRsp) error {
	utils.Assert(ctx != nil && req != nil, "AddUserList in params invalid")

	if rsp == nil {
		return ErrRsp
	}

	if len(req.UserList) == 0 {
		return ErrReq
	}

	dataList := []*data.UserInfo{}
	for _, userInfo := range req.UserList {
		err := dao.AddUserInfo(userInfo, com.MySql)
		if err != nil {
			log.Error(err)
			continue
		}

		dataList = append(dataList, userInfo)
	}

	rsp.UserList = dataList
	if len(dataList) != len(req.UserList) {
		rsp.Code = -100
	}

	return nil
}
