package m

import (
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/basic"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
)

// AddBasic 添加用户基础信息
//  @Response.Code:
//    0:         成功
//    -1 ~ -102: 参数错误
//    -10001:    邮箱已存在
//    -10002:    手机号已存在
//    -10000:    内部DB错误
func (this_ *UserService) AddBasic(ctx *piper.Context, req *user.AddBasicReq, rsp *user.AddBasicRsp) error {
	utils.Assert(ctx != nil && req != nil && rsp != nil, "AddBasic in params is invalid")

	for dwf := true; dwf; dwf = false {
		// Step 1: 入参检查
		if req.Basic == nil {
			rsp.Code = -100
			break
		}

		if req.Basic.Entry == nil {
			rsp.Code = -101
			break
		}

		if req.Basic.Entry.UserID != 0 {
			rsp.Code = -2
			break
		}

		if req.Basic.Entry.Gender == 0 {
			req.Basic.Entry.Gender = 3
		}

		if req.Basic.Entry.Gender < ds.MIN_GENDER || req.Basic.Entry.Gender > ds.MAX_GENDER {
			rsp.Code = -1
			break
		}

		if len(req.Basic.Entry.Email) == 0 && len(req.Basic.Entry.PhoneNum) == 0 {
			rsp.Code = -102
			break
		}

		if len(req.Basic.Entry.Email) > 0 && len(req.Basic.Entry.Email) > ds.MAX_EMAIL {
			rsp.Code = -3
			break
		}

		if utf8.RuneCountInString(req.Basic.Entry.Nickname) > ds.MAX_NICKNAME {
			rsp.Code = -4
			break
		}

		if len(req.Basic.Entry.PhoneNum) > 0 && len(req.Basic.Entry.PhoneNum) > ds.MAX_PHONE_NUM {
			rsp.Code = -5
			break
		}

		if len(req.Basic.Entry.Avator) > ds.MAX_AVATOR {
			rsp.Code = -6
			break
		}

		// Step 2: 检查邮箱是否被使用
		if len(req.Basic.Entry.Email) > 0 {
			found, err := entry.ExistsEmail(req.Basic.Entry.Email, com.Elastic)
			if err != nil {
				log.Error(err)
				rsp.Code = -10000
				break
			}

			if found {
				rsp.Code = -10001
				break
			}
		}

		// Step 3: 检查手机号是否被使用
		if len(req.Basic.Entry.PhoneNum) > 0 {
			found, err := entry.ExistsPhoneNum(req.Basic.Entry.PhoneNum, com.Elastic)
			if err != nil {
				log.Error(err)
				rsp.Code = -10000
			}

			if found {
				rsp.Code = -10002
				break
			}
		}

		// Step 4: 写入数据
		err := basic.AddBasic(req.Basic, com.MySql, com.Elastic)
		if err != nil {
			log.Error(err)
			rsp.Code = -10000
			break
		}

		// Step 5: 应答
		rsp.Basic = req.Basic
	}

	return nil
}
