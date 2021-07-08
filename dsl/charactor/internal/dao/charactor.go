package dao

import (
	"database/sql"

	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/data"
)

func AddCharactor(obj *data.Charactor, db *sql.DB) error {
	return nil
}

func ModifyCharactor(obj *data.Charactor, db *sql.DB) error {
	return nil
}

func RemoveCharactor(userID, charID int64, db *sql.DB) error {
	return nil
}

func QueryCharactor(cond *cgi.QueryCharactorReq, db *sql.DB) ([]*data.Charactor, error) {
	return nil, nil
}

func QueryOneCharactor(cond *cgi.QueryCharactorReq, db *sql.DB) (*data.Charactor, error) {
	return nil, nil
}

func CountCharactor(cond *cgi.QueryCharactorReq, db *sql.DB) (int64, error) {
	return 0, nil
}
