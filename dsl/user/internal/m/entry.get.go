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

func (this_ *UserService) GetEntry(ctx *piper.Context, req *user.GetEntryReq, rsp *user.GetEntryRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "GetEntry in params is invalid")

	if req.UserID > 0 {
		entry, err := entry.GetEntryByID(req.UserID, com.Elastic)
		if err != nil {
			log.Error(err)
			rsp.Code = -100
		} else {
			rsp.EntryList = append([]*ds.Entry{}, entry)
			rsp.Total = int64(len(rsp.EntryList))
		}

		return nil
	}

	dataList, total, err := entry.GetEntry(req, com.Elastic)
	if err != nil {
		log.Error(err)
		rsp.Code = -100
	} else {
		rsp.EntryList = dataList
		rsp.Total = total
	}

	return nil
}
