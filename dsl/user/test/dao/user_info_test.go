package dao

import (
	"database/sql"
	"testing"

	"github.com/iegad/kraken/store"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/data"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

func getDB() (*sql.DB, error) {
	return store.NewMySql(&store.MySqlOption{
		Host:     "127.0.0.1:3306",
		User:     "iegad",
		Pass:     "1111",
		Database: "DB_BASIC",
	})
}

func TestAddUserInfo(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	uinfo := data.NewUserInfo()
	uinfo.PhoneNum = "1234567890"
	uinfo.Nickname = "iegad"
	uinfo.Gender = 1

	err = dao.AddUserInfo(uinfo, db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(uinfo))
}

func TestModUserInfo(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	res, err := dao.QueryUserInfo(&cgi.QueryUserInfoReq{
		PhoneNum: "1234567890",
	}, db)

	if err != nil {
		t.Error(err)
		return
	}

	if len(res) != 1 {
		t.Error("查询失败")
		return
	}

	t.Log(utils.PbToJson(res[0]))

	res[0].Email = "111.11@222.222"
	err = dao.ModifyUserInfo(res[0], db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(res[0]))

	err = dao.RemoveUserInfo(res[0].UserID, db)
	if err != nil {
		t.Error(err)
		return
	}
}
