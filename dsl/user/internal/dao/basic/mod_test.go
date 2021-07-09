package basic

import (
	"testing"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/ds/user"
)

func TestMod(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	es, err := getES()
	if err != nil {
		t.Error(err)
		return
	}

	basic := user.NewBasic()
	basic.Entry = user.NewEntry()
	basic.Entry.UserID = 1
	basic.Entry.Email = "iegad@vip.qq.com"
	basic.Entry.PhoneNum = "15827323665"

	err = ModBasic(basic, db, es)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(basic))
}
