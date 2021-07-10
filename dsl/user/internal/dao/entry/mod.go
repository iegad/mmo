package entry

import (
	"context"
	"encoding/json"
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func ModEntry(obj *ds.Entry, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "ModEntry in params is invalid")

	// Step 1: 入参检查
	if obj.UserID <= 0 {
		return cgi.ErrUserID
	}

	// Step 2: 获取原始数据
	raw, err := GetEntryByID(obj.UserID, es)
	if err != nil {
		log.Error(err)
		return err
	}

	// Step 3: 验证更新数据
	updata := map[string]interface{}{}

	// 头像验证
	if len(obj.Avator) > 0 {
		if len(obj.Avator) > ds.MAX_AVATOR {
			return cgi.ErrAvator
		}

		if obj.Avator != raw.Avator {
			updata["Avator"] = obj.Avator
		}
	}

	// 邮箱验证
	if len(obj.Email) > 0 {
		if utf8.RuneCountInString(obj.Email) > ds.MAX_EMAIL {
			return cgi.ErrEmail
		}

		if obj.Email != raw.Email {
			res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("Email", obj.Email)).Do(context.TODO())
			if err != nil {
				log.Error(err)
				return cgi.ErrESInner
			}

			if res.TotalHits() != 0 {
				return cgi.ErrEmailExists
			}

			updata["Email"] = obj.Email
		}
	}

	// 手机号验证
	if len(obj.PhoneNum) > 0 {
		if len(obj.PhoneNum) > ds.MAX_PHONE_NUM {
			return cgi.ErrPhoneNum
		}

		if obj.PhoneNum != raw.PhoneNum {
			found, err := existsPhoneNum(obj.PhoneNum, es)
			if err != nil {
				log.Error(err)
				return cgi.ErrESInner
			}

			if found {
				return cgi.ErrPhoneNumExists
			}

			updata["PhoneNum"] = obj.PhoneNum
		}
	}

	// 性别验证
	if obj.Gender > 0 {
		if obj.Gender > ds.MAX_GENDER {
			return cgi.ErrGender
		}

		if obj.Gender != raw.Gender {
			updata["Gender"] = obj.Gender
		}
	}

	// 昵称验证
	if len(obj.Nickname) > 0 {
		if utf8.RuneCountInString(obj.Nickname) > ds.MAX_NICKNAME {
			return cgi.ErrNickname
		}

		if obj.Nickname != raw.Nickname {
			updata["Nickname"] = obj.Nickname
		}
	}

	// 头像验证
	if len(obj.Avator) > 0 {
		if len(obj.Avator) > ds.MAX_AVATOR {
			return cgi.ErrAvator
		}

		if obj.Avator != raw.Avator {
			updata["Avator"] = obj.Avator
		}
	}

	if len(updata) == 0 {
		return cgi.ErrNoUpData
	}

	up, err := es.Update().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).Doc(updata).DocAsUpsert(true).FetchSource(true).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return cgi.ErrESInner
	}

	err = json.Unmarshal(up.GetResult.Source, obj)
	if err != nil {
		log.Error(err)
		return cgi.ErrESDataType
	}

	return nil
}
