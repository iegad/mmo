package entry

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/ds/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func ModEntry(obj *ds.Entry, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "ModEntry in params is invalid")

	if obj.UserID <= 0 {
		return cgi.ErrUserID
	}

	res, err := es.Search().Index(dao.ES_INDEX).Query(elastic.NewTermQuery("UserID", obj.UserID)).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return cgi.ErrESInner
	}

	if res.TotalHits() != 1 {
		return cgi.ErrUserID
	}

	var (
		list    = res.Each(refType)
		raw, ok = list[0].(*user.Entry)
		updata  = map[string]interface{}{}
		changed = false
	)

	if !ok {
		log.Error("%v type is not user.Entry", list)
		return cgi.ErrESInner
	}

	if len(obj.Avator) > 0 && obj.Avator != raw.Avator {
		updata["Avator"] = obj.Avator
		changed = true
	}

	if len(obj.Email) > 0 && obj.Email != raw.Email {
		updata["Email"] = obj.Email
		changed = true
	}

	if len(obj.PhoneNum) > 0 && obj.PhoneNum != raw.PhoneNum {
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

	up, err := es.Update().Index(dao.ES_INDEX).Id(fmt.Sprintf("%d", obj.UserID)).Doc(updata).DocAsUpsert(true).FetchSource(true).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return cgi.ErrESInner
	}

	json.Unmarshal(up.GetResult.Source, obj)
	return nil
}
