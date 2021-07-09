package basic

import "testing"

func TestRmv(t *testing.T) {
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

	err = RmvBasic(1, db, es)
	if err != nil {
		t.Error(err)
		return
	}
}
