package basic

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iegad/kraken/conf"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/olivere/elastic/v7"
)

func initES() error {
	var err error
	com.Elastic, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
	return err
}

func getDb() (*sql.DB, error) {
	return sql.Open("mysql", (&conf.MySql{
		Port:     3306,
		Addr:     "127.0.0.1",
		Database: "DB_USER",
		User:     "iegad",
		Pass:     "1111",
	}).String())
}

func TestAdd(t *testing.T) {
	err := initES()
	if err != nil {
		t.Error(err)
		return
	}

	db, err := getDb()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	basic := user.NewBasic()
	basic.Entry = user.NewEntry()
	basic.Entry.Nickname = "iegad"
	basic.Entry.PhoneNum = "123456789"

	err = AddBasic(basic, db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(basic))
}
