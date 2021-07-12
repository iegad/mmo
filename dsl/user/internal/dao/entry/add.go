package entry

import (
	"context"

	"github.com/iegad/kraken/utils"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

// AddEntry 新增用户条目
func AddEntry(obj *ds.Entry, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "AddEntry in params invalid")

	_, err := es.Index().Index(dao.N_USER_ENTRY).Id(dao.IDToString(obj.UserID)).BodyJson(obj).Do(context.TODO())
	return err
}
