package entry

import (
	"context"

	"github.com/iegad/kraken/utils"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func ModEntry(obj *ds.Entry, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "ModEntry in params is invalid")

	updata := map[string]interface{}{}

	// 头像验证
	if len(obj.Avator) > 0 {
		updata["Avator"] = obj.Avator
	}

	// 邮箱验证
	if len(obj.Email) > 0 {
		updata["Email"] = obj.Email
	}

	// 手机号验证
	if len(obj.PhoneNum) > 0 {
		updata["PhoneNum"] = obj.PhoneNum
	}

	// 性别验证
	if obj.Gender > 0 {
		updata["Gender"] = obj.Gender
	}

	// 昵称验证
	if len(obj.Nickname) > 0 {
		updata["Nickname"] = obj.Nickname
	}

	if len(updata) == 0 {
		return nil
	}

	_, err := es.Update().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).Doc(updata).DocAsUpsert(true).FetchSource(true).Do(context.TODO())
	return err
}
