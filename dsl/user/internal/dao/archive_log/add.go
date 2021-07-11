package archive_log

import (
	"context"

	"github.com/iegad/kraken/utils"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func AddArchiveLog(obj *ds.ArchiveLog, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "AddArchiveLog in params is invalid")

	_, err := es.Index().Index(dao.N_USER_ARCHIVE_LOG).BodyJson(obj).Do(context.TODO())
	return err
}
