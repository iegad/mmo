package entry

import (
	"context"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func RmvEntry(userID int64, es *elastic.Client) error {
	utils.Assert(userID > 0 && es != nil, "RmvEntry in params is invalid")

	_, err := es.Delete().Index(dao.N_USER_ENTRY).Id(dao.IDToString(userID)).Do(context.TODO())
	return err
}
