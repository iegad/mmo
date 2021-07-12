package entry

import (
	"context"
	"reflect"
	"unicode/utf8"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

var (
	refType = reflect.TypeOf(&ds.Entry{})
)

// GetEntry 在ES中查询用户条目
//  @PS: 查询结果集中的Entry项不需要释放
func GetEntry(cond *user.GetEntryReq, es *elastic.Client) ([]*ds.Entry, int64, error) {
	utils.Assert(cond != nil && es != nil, "GetEntry in params is invalid")

	// Step 1: 构建查询条件
	query := elastic.NewBoolQuery()

	if cond.Gender > 0 {
		query.Must(elastic.NewTermQuery("Gender", cond.Gender))
	}

	if cond.UserID > 0 {
		query.Must(elastic.NewTermQuery("UserID", cond.UserID))
	}

	if len(cond.Key) > 0 {
		or := elastic.NewBoolQuery()

		klen := utf8.RuneCountInString(cond.Key)

		if klen <= ds.MAX_NICKNAME {
			or.Should(elastic.NewMatchPhrasePrefixQuery("Nickname", cond.Key))
		}

		if klen <= ds.MAX_PHONE_NUM {
			or.Should(elastic.NewMatchPhrasePrefixQuery("PhoneNum", cond.Key))
		}

		if klen <= ds.MAX_EMAIL {
			or.Should(elastic.NewMatchPhrasePrefixQuery("Email", cond.Key))
		}

		query.Must(or)
	}

	// Step 2: 构建搜索对象
	search := es.Search().Index(dao.N_USER_ENTRY).Query(query).From(int(cond.Offset))
	if cond.Limit > 0 {
		search.Size(int(cond.Limit))
	}

	result, err := search.Do(context.TODO())
	if err != nil {
		return nil, -1, err
	}

	// Step 3: 读取数据
	list := result.Each(refType)
	dataList := []*ds.Entry{}
	for _, item := range list {
		if t, ok := item.(*ds.Entry); ok {
			dataList = append(dataList, t)
		}
	}

	return dataList, result.TotalHits(), nil
}

// GetEntryByID 通过UserID在ES中查询用户条目
func GetEntryByID(userID int64, es *elastic.Client) (*ds.Entry, error) {
	utils.Assert(userID > 0 && es != nil, "GetEntryByID in params is invalid")

	res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("UserID", dao.IDToString(userID))).Do(context.TODO())
	if err != nil {
		return nil, err
	}

	if res.TotalHits() != 1 {
		return nil, nil
	}

	list := res.Each(refType)
	entry := list[0].(*ds.Entry)
	return entry, nil
}
