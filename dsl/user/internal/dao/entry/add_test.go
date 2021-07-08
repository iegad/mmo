package entry

import (
	"testing"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
	"github.com/olivere/elastic/v7"
)

func getElastic() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false))
}

func TestAdd(t *testing.T) {
	entry := user.NewEntry()
	entry.Nickname = "123456789"
	entry.PhoneNum = "111222333"
	entry.UserID = 4
	entry.Gender = 1

	es, err := getElastic()
	if err != nil {
		t.Error(err)
		return
	}

	err = AddEntry(entry, es)
	if err != nil {
		t.Error(err)
		return
	}

	entry.PhoneNum = dao.DecodePhoneNum(entry.PhoneNum)
	t.Log(utils.PbToJson(entry))
}
