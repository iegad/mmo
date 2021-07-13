package login_log

import (
	"context"

	"github.com/iegad/kraken/utils"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func AddLoginLog(obj *ds.LoginLog, es *elastic.Client) error {
	utils.Assert(obj != nil && es != nil, "AddLoginLog in parmas is invalid")

	_, err := es.Index().Index(dao.N_USER_LOGIN_LOG).BodyJson(obj).Do(context.TODO())
	return err
}
