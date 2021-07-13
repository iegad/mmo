package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	"github.com/iegad/mmo/ds"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/login_log"
)

func (this_ *UserService) AddLoginLog(ctx *piper.Context, req *user.AddLoginLogReq, rsp *user.AddLoginLogRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "AddLoginLog in params is invalid")

	if req.LoginLog == nil {
		rsp.Code = -100
		return nil
	}

	if req.LoginLog.OSType <= 0 || req.LoginLog.OSType > ds.OSType_OSType_Linux {
		rsp.Code = -1
		return nil
	}

	if req.LoginLog.DeviceType < 0 {
		rsp.Code = -2
		return nil
	}

	if req.LoginLog.UserID <= 0 {
		rsp.Code = 3
		return nil
	}

	if req.LoginLog.LoginTime <= 0 {
		rsp.Code = -4
		return nil
	}

	if len(req.LoginLog.Token) != 16 {
		rsp.Code = -5
		return nil
	}

	if len(req.LoginLog.MacAddr) != 6 {
		rsp.Code = -6
		return nil
	}

	if len(req.LoginLog.OperatingSystem) == 0 {
		rsp.Code = -7
		return nil
	}

	if len(req.LoginLog.DeviceUniqueIdentifier) == 0 {
		rsp.Code = -8
		return nil
	}

	if len(req.LoginLog.MountAddr) == 0 {
		rsp.Code = -9
		return nil
	}

	if req.LoginLog.Entry == nil {
		rsp.Code = -10
		return nil
	}

	err := login_log.AddLoginLog(req.LoginLog, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
	}

	return nil
}
