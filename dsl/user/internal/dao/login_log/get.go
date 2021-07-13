package login_log

import (
	"context"
	"reflect"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

var (
	refType = reflect.TypeOf(&ds.LoginLog{})
)

func GetLoginLog(cond *user.GetLoginLogReq, es *elastic.Client) ([]*ds.LoginLog, int64, error) {
	utils.Assert(cond != nil && es != nil, "GetLoginLog in params is invalid")

	query := elastic.NewBoolQuery()
	if cond.UserID > 0 {
		query.Must(elastic.NewTermQuery("UserID", cond.UserID))
	}

	if cond.LoginTimeBeg > 0 {
		query.Must(elastic.NewRangeQuery("LoginTime").Gte(cond.LoginTimeBeg))
	}

	if cond.LoginTimeEnd > 0 {
		query.Must(elastic.NewRangeQuery("LoginTime").Lt(cond.LoginTimeEnd))
	}

	if cond.OSType > 0 {
		query.Must(elastic.NewTermQuery("OSType", cond.OSType))
	}

	if cond.DeviceType > 0 {
		query.Must(elastic.NewTermQuery("DeviceType", cond.DeviceType))
	}

	if len(cond.MountAddr) > 0 {
		query.Must(elastic.NewTermQuery("MountAddr", cond.MountAddr))
	}

	search := es.Search().Index(dao.N_USER_LOGIN_LOG).Query(query).From(int(cond.Offset))
	if cond.Limit > 0 {
		search.Size(int(cond.Limit))
	}

	res, err := search.Do(context.TODO())
	if err != nil {
		return nil, -1, err
	}

	list := res.Each(refType)
	dataList := []*ds.LoginLog{}
	for _, item := range list {
		if t, ok := item.(*ds.LoginLog); ok {
			dataList = append(dataList, t)
		}
	}

	return dataList, res.TotalHits(), nil
}
