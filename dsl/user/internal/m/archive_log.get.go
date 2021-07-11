package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/archive_log"
)

func (this_ *UserService) GetArchiveLog(ctx *piper.Context, req *user.GetArchiveLogReq, rsp *user.GetArchiveLogRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "GetArchiveLog in params is invalid")

	dataList, total, err := archive_log.GetArchiveLog(req, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -100
		return nil
	}

	rsp.LogList = dataList
	rsp.Total = total
	return nil
}
