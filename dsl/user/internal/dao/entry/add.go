package entry

import (
	"context"
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

// AddEntry 新增用户条目
func AddEntry(obj *ds.Entry, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "AddEntry in params invalid")

	// Step 1: 入参检查
	if obj.UserID < ds.MIN_USER_ID {
		return cgi.ErrUserID
	}

	if obj.Gender < ds.MIN_GENDER || obj.Gender > ds.MAX_GENDER {
		return cgi.ErrGender
	}

	if len(obj.Email) == 0 && len(obj.PhoneNum) == 0 {
		return cgi.ErrAccount
	}

	if len(obj.Email) > 0 && len(obj.Email) > ds.MAX_EMAIL {
		return cgi.ErrEmail
	}

	if len(obj.PhoneNum) > 0 && len(obj.PhoneNum) > ds.MAX_PHONE_NUM {
		return cgi.ErrPhoneNum
	}

	if utf8.RuneCountInString(obj.Nickname) > ds.MAX_NICKNAME {
		return cgi.ErrNickname
	}

	if len(obj.Avator) > 500 {
		return cgi.ErrAvator
	}

	// Step 2: 判断UserID是否存在
	exists, err := es.Exists().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return cgi.ErrESInner
	}

	if exists {
		return cgi.ErrUserExists
	}

	// Step 3: 检查邮箱是否被使用
	if len(obj.Email) > 0 {
		found, err := ExistsEmail(obj.Email, es)
		if err != nil {
			log.Error(err)
			return cgi.ErrESInner
		}

		if found {
			return cgi.ErrEmailExists
		}
	}

	// Step 4: 检查手机号是否被使用
	if len(obj.PhoneNum) > 0 {
		found, err := ExistsPhoneNum(obj.PhoneNum, es)
		if err != nil {
			log.Error(err)
			return cgi.ErrESInner
		}

		if found {
			return cgi.ErrPhoneNumExists
		}
	}

	// Step 5: 写入ES
	_, err = es.Index().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).BodyJson(obj).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return cgi.ErrESInner
	}

	return err
}
