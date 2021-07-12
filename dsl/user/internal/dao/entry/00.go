package entry

import (
	"context"

	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func ExistsPhoneNum(phoneNum string, es *elastic.Client) (bool, error) {
	res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("PhoneNum", phoneNum)).Do(context.TODO())
	if err != nil {
		return false, err
	}

	return res.TotalHits() > 0, nil
}

func ExistsEmail(email string, es *elastic.Client) (bool, error) {
	res, err := es.Search().Index(dao.N_USER_ENTRY).Query(elastic.NewTermQuery("Email", email)).Do(context.TODO())
	if err != nil {
		return false, err
	}

	return res.TotalHits() > 0, nil
}

func ExistsUserID(userID int64, es *elastic.Client) (bool, error) {
	res, err := es.Exists().Index(dao.N_USER_ENTRY).Id(dao.IDToString(userID)).Do(context.TODO())
	if err != nil {
		return false, err
	}

	return res, nil
}
