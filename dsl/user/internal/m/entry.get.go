package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
)

// GetEntry 搜索用户条目
func (this_ *UserService) GetEntry(ctx *piper.Context, req *user.GetEntryReq, rsp *user.GetEntryRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "GetEntry in params is invalid")

	if req.UserID > 0 {
		entry, err := entry.GetEntryByID(req.UserID, com.Elastic)
		if err != nil {
			log.Error(err)
			rsp.Code = -10000
		} else {
			rsp.EntryList = append([]*ds.Entry{}, entry)
			rsp.Total = int64(len(rsp.EntryList))
		}

		return nil
	}

	if req.UserID <= 0 && req.Gender == 0 && len(req.Key) == 0 && req.Limit == 0 {
		rsp.Code = -100
		return nil
	}

	if req.Gender > 0 && (req.Gender > ds.MAX_GENDER || req.Gender < ds.MIN_GENDER) {
		rsp.Code = -1
		return nil
	}

	if req.UserID < 0 {
		rsp.Code = -2
		return nil
	}

	if req.Offset < 0 {
		rsp.Code = -3
	}

	if req.Limit < 0 {
		rsp.Code = -4
		return nil
	}

	dataList, total, err := entry.GetEntry(req, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -10000
		return nil
	}

	rsp.EntryList = dataList
	rsp.Total = total

	return nil
}
