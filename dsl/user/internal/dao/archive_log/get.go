package archive_log

import (
	"context"
	"reflect"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

var (
	refType = reflect.TypeOf(&ds.ArchiveLog{})
)

func GetArchiveLog(cond *user.GetArchiveLogReq, es *elastic.Client) ([]*ds.ArchiveLog, int64, error) {
	utils.Assert(cond != nil && es != nil, "GetArchiveLog in params is invalid")

	var (
		query   = elastic.NewBoolQuery()
		hasCond = false
	)

	if cond.UserID > 0 {
		query.Must(elastic.NewTermQuery("UserID", cond.UserID))
		hasCond = true
	}

	if cond.VerCode > 0 {
		query.Must(elastic.NewTermQuery("VerCode", cond.VerCode))
		hasCond = true
	}

	if cond.ArchiveTimeBeg > 0 {
		query.Must(elastic.NewRangeQuery("ArchiveTime").Gte(cond.ArchiveTimeBeg))
		hasCond = true
	}

	if cond.ArchiveTimeEnd > 0 {
		query.Must(elastic.NewRangeQuery("ArchiveTime").Lt(cond.ArchiveTimeEnd))
		hasCond = true
	}

	if !hasCond {
		return nil, -1, cgi.ErrNoCond
	}

	search := es.Search().Index(dao.N_USER_ARCHIVE_LOG).Query(query).From(int(cond.Offset))
	if cond.Limit > 0 {
		search.Size(int(cond.Limit))
	}

	res, err := search.Do(context.TODO())
	if err != nil {
		return nil, -1, err
	}

	list := res.Each(refType)
	dataList := []*ds.ArchiveLog{}
	for _, item := range list {
		if t, ok := item.(*ds.ArchiveLog); ok {
			dataList = append(dataList, t)
		}
	}

	return dataList, res.TotalHits(), nil
}
