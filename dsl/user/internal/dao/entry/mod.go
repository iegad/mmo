package entry

import (
	"context"
	"encoding/json"

	"github.com/iegad/kraken/utils"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func ModEntry(obj *ds.Entry, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "ModEntry in params is invalid")

	// Step 2: 获取原始数据
	raw, err := GetEntryByID(obj.UserID, es)
	if err != nil {
		return err
	}

	// Step 3: 验证更新数据
	updata := map[string]interface{}{}

	// 头像验证
	if len(obj.Avator) > 0 && obj.Avator != raw.Avator {
		updata["Avator"] = obj.Avator
	}

	// 邮箱验证
	if len(obj.Email) > 0 && obj.Email != raw.Email {
		updata["Email"] = obj.Email
	}

	// 手机号验证
	if len(obj.PhoneNum) > 0 && obj.PhoneNum != raw.PhoneNum {
		updata["PhoneNum"] = obj.PhoneNum
	}

	// 性别验证
	if obj.Gender > 0 && obj.Gender != raw.Gender {
		updata["Gender"] = obj.Gender
	}

	// 昵称验证
	if len(obj.Nickname) > 0 && obj.Nickname != raw.Nickname {
		updata["Nickname"] = obj.Nickname
	}

	// 头像验证
	if len(obj.Avator) > 0 && obj.Avator != raw.Avator {
		updata["Avator"] = obj.Avator
	}

	if len(updata) == 0 {
		return nil
	}

	up, err := es.Update().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).Doc(updata).DocAsUpsert(true).FetchSource(true).Do(context.TODO())
	if err != nil {
		return err
	}

	err = json.Unmarshal(up.GetResult.Source, obj)
	if err != nil {
		return err
	}

	return nil
}
