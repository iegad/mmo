package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	common "github.com/iegad/mmo/ds"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/login_log"
)

func (this_ *UserService) GetLoginLog(ctx *piper.Context, req *user.GetLoginLogReq, rsp *user.GetLoginLogRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "AddLoginLog in params is invalid")

	if req.OSType <= common.OSType_OSType_Invalid || req.OSType > common.OSType_OSType_Linux {
		rsp.Code = -1
		return nil
	}

	if req.DeviceType < 0 {
		rsp.Code = -2
		return nil
	}

	if req.UserID < 0 {
		rsp.Code = -3
		return nil
	}

	if req.LoginTimeBeg < 0 {
		rsp.Code = -4
		return nil
	}

	if req.LoginTimeEnd < 0 {
		rsp.Code = -5
		return nil
	}

	if req.Offset < 0 {
		rsp.Code = -6
		return nil
	}

	if req.Limit < 0 {
		rsp.Code = -7
		return nil
	}

	dataList, total, err := login_log.GetLoginLog(req, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
		return nil
	}

	rsp.LogList = dataList
	rsp.Total = total
	return nil
}
