package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/archive_log"
)

// GetArchiveLog 获取存档日志
func (this_ *UserService) GetArchiveLog(ctx *piper.Context, req *user.GetArchiveLogReq, rsp *user.GetArchiveLogRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "GetArchiveLog in params is invalid")

	if req.UserID <= 0 {
		rsp.Code = -3
		return nil
	}

	if req.ArchiveTimeBeg < 0 {
		rsp.Code = -4
		return nil
	}

	if req.ArchiveTimeEnd < 0 {
		rsp.Code = -5
		return nil
	}

	if req.Limit < 0 {
		rsp.Code = -6
		return nil
	}

	if req.Offset < 0 {
		rsp.Code = -7
		return nil
	}

	dataList, total, err := archive_log.GetArchiveLog(req, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
		return nil
	}

	rsp.LogList = dataList
	rsp.Total = total
	return nil
}
