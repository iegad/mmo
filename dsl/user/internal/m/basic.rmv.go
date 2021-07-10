package m

import (
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/basic"
)

func (this_ *UserService) RmvBasic(ctx *piper.Context, req *user.RmvBasicReq, rsp *user.RmvBasicRsp) error {
	utils.Assert(ctx != nil && req != nil, "AddBasic in params is invalid")

	var (
		err error
		raw *ds.Basic
	)

	if rsp == nil {
		return cgi.ErrRsp
	}

	for dwf := true; dwf; dwf = false {
		raw, err = basic.GetBasicByID(req.UserID, com.MySql)
		if err != nil {
			break
		}

		err = basic.RmvBasic(req.UserID, com.MySql, com.Elastic)
		if err != nil {
			break
		}
	}

	if err != nil {
		log.Error(err)
		rsp.Code = -100
	} else {
		rsp.Basic = raw
	}

	return nil
}
