package entry

import (
	"testing"

	"github.com/iegad/kraken/utils"
)

func TestGet(t *testing.T) {
	es, err := getElastic()
	if err != nil {
		t.Error(err)
		return
	}

	dataList, total, err := GetEntry(&Condition{
		Key: "111222",
	}, es)

	if err != nil {
		t.Error(err)
	}

	t.Logf("total: %d\n", total)
	for _, entry := range dataList {
		t.Log(utils.PbToJson(entry))
	}
}
