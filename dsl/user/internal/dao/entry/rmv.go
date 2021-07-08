package entry

import (
	"context"

	"github.com/iegad/kraken/log"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func RmvEntry(userID int64, es *elastic.Client) error {
	_, err := es.Delete().Index(dao.N_USER_ENTRY).Id(dao.IDToString(userID)).Do(context.TODO())
	if err != nil {
		log.Error(err)
		return cgi.ErrESInner
	}

	return nil
}
