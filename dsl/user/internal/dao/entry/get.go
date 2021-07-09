package entry

import (
	"context"
	"reflect"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

type Condition user.GetEntryReq

var (
	refType = reflect.TypeOf(&ds.Entry{})
)

func GetEntry(cond *Condition, es *elastic.Client) ([]*ds.Entry, int64, error) {
	utils.Assert(cond != nil && es != nil, "GetEntry in params is invalid")

	query := elastic.NewBoolQuery()
	if cond.Gender > 0 {
		query = query.Must(elastic.NewTermQuery("Gender", cond.Gender))
	}

	if cond.UserID > 0 {
		query = query.Must(elastic.NewTermQuery("UserID", cond.UserID))
	}

	if len(cond.Key) > 0 {
		or := elastic.NewBoolQuery()
		or.Should(elastic.NewMatchPhrasePrefixQuery("Nickname", cond.Key))
		or.Should(elastic.NewMatchPhrasePrefixQuery("PhoneNum", cond.Key))
		or.Should(elastic.NewMatchPhrasePrefixQuery("Email", cond.Key))

		query = query.Must(or)
	}

	search := es.Search().Index(dao.N_USER_ENTRY).Query(query)
	if cond.Limit > 0 {
		search = search.From(int(cond.Offset)).Size(int(cond.Limit))
	}

	result, err := search.Do(context.TODO())
	if err != nil {
		return nil, -1, err
	}

	list := result.Each(refType)
	dataList := []*ds.Entry{}
	for _, item := range list {
		if t, ok := item.(*ds.Entry); ok {
			dataList = append(dataList, t)
		}
	}

	return dataList, result.TotalHits(), nil
}

func GetEntryByID(userID int64, es *elastic.Client) (*ds.Entry, error) {
	utils.Assert(userID > 0 && es != nil, "GetEntryByID in params is invalid")

	res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("UserID", dao.IDToString(userID))).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return nil, cgi.ErrESInner
	}

	if res.TotalHits() != 1 {
		return nil, cgi.ErrUserID
	}

	var (
		list      = res.Each(refType)
		entry, ok = list[0].(*ds.Entry)
	)

	if !ok {
		log.Error("user_entry type is invalid")
		return nil, cgi.ErrESInner
	}

	return entry, nil
}
