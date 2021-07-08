package entry

import "testing"

func TestRmv(t *testing.T) {
	es, err := getElastic()
	if err != nil {
		t.Error(err)
		return
	}

	err = RmvEntry(3, es)
	if err != nil {
		t.Error(err)
		return
	}
}
