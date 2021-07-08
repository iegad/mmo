package entry

import (
	"context"
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func AddEntry(obj *user.Entry, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "AddEntry in params invalid")

	if obj.UserID <= 0 {
		return cgi.ErrUserID
	}

	if obj.Gender < 1 || obj.Gender > 3 {
		return cgi.ErrGender
	}

	if len(obj.Email) == 0 && len(obj.PhoneNum) == 0 {
		return cgi.ErrAccount
	}

	if len(obj.Email) > 0 && utf8.RuneCountInString(obj.Email) > 50 {
		return cgi.ErrEmail
	}

	if len(obj.PhoneNum) > 0 && utf8.RuneCountInString(obj.PhoneNum) > 15 {
		return cgi.ErrPhoneNum
	}

	hasIndex, err := es.IndexExists(dao.N_USER_ENTRY).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return cgi.ErrESInner
	}

	if hasIndex {
		exists, err := es.Exists().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).Do(context.TODO())
		if err != nil {
			log.Error(err)
			return cgi.ErrESInner
		}

		if exists {
			return cgi.ErrUserID
		}
	}

	if len(obj.Email) > 0 {
		obj.Email = dao.EncodeEmail(obj.Email)

		if hasIndex {
			res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("Email", obj.Email)).Do(context.TODO())
			if err != nil {
				log.Error(err)
				return cgi.ErrESInner
			}

			if res.TotalHits() != 0 {
				return cgi.ErrEmail
			}
		}
	}

	if len(obj.PhoneNum) > 0 {
		obj.PhoneNum = dao.EncodePhoneNum(obj.PhoneNum)

		if hasIndex {
			res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("PhoneNum", obj.PhoneNum)).Do(context.TODO())
			if err != nil {
				log.Error(err)
				return cgi.ErrESInner
			}

			if res.TotalHits() != 0 {
				return cgi.ErrPhoneNum
			}
		}
	}

	_, err = es.Index().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).BodyJson(obj).Do(context.TODO())
	return err
}
