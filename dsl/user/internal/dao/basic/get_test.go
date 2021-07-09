package basic

import (
	"testing"

	"github.com/iegad/mmo/cgi/user"
)

func TestGet(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	dataList, total, err := GetBasic(user.NewGetBasicReq(), db)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(total)
	t.Log(dataList)
}
