package entry

import (
	"testing"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/ds/user"
)

func TestMod(t *testing.T) {
	es, err := getElastic()
	if err != nil {
		t.Error(err)
		return
	}

	entry := user.NewEntry()
	entry.UserID = 1
	entry.Nickname = "iegad"
	err = ModEntry(entry, es)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(entry))
}
