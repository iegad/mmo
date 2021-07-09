package basic

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iegad/kraken/conf"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/ds/user"
	"github.com/olivere/elastic/v7"
)

func getES() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
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
	es, err := getES()
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

	err = AddBasic(basic, db, es)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(basic))
}
