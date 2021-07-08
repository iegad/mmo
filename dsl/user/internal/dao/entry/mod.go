package entry

import (
	"context"
	"encoding/json"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func ModEntry(obj *ds.Entry, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "ModEntry in params is invalid")

	if obj.UserID <= 0 {
		return cgi.ErrUserID
	}

	raw, err := GetEntryByID(obj.UserID, es)
	if err != nil {
		log.Error(err)
		return err
	}

	var (
		updata  = map[string]interface{}{}
		changed = false
	)

	if len(obj.Avator) > 0 && obj.Avator != raw.Avator {
		updata["Avator"] = obj.Avator
		changed = true
	}

	if len(obj.Email) > 0 && obj.Email != raw.Email {
		obj.Email = dao.EncodeEmail(obj.Email)
		res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("Email", obj.Email)).Do(context.TODO())
		if err != nil {
			log.Error(err)
			return cgi.ErrESInner
		}

		if res.TotalHits() != 0 {
			return cgi.ErrEmail
		}

		updata["Email"] = obj.Email
		changed = true
	}

	if len(obj.PhoneNum) > 0 && obj.PhoneNum != raw.PhoneNum {
		obj.PhoneNum = dao.EncodePhoneNum(obj.PhoneNum)
		res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("PhoneNum", obj.PhoneNum)).Do(context.TODO())
		if err != nil {
			log.Error(err)
			return cgi.ErrESInner
		}

		if res.TotalHits() != 0 {
			return cgi.ErrPhoneNum
		}

		updata["PhoneNum"] = obj.PhoneNum
		changed = true
	}

	if obj.Gender > 0 && obj.Gender != raw.Gender {
		updata["Gender"] = obj.Gender
		changed = true
	}

	if len(obj.Nickname) > 0 && obj.Nickname != raw.Nickname {
		updata["Nickname"] = obj.Nickname
		changed = true
	}

	if !changed {
		return cgi.ErrNoUpData
	}

	up, err := es.Update().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).Doc(updata).DocAsUpsert(true).FetchSource(true).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return cgi.ErrESInner
	}

	json.Unmarshal(up.GetResult.Source, obj)
	return nil
}
