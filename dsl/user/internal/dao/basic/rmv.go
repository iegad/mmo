package basic

import (
	"database/sql"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
	"github.com/olivere/elastic/v7"
)

func RmvBasic(userID int64, db *sql.DB, es *elastic.Client) error {
	utils.Assert(userID > 0 && db != nil && es != nil, "RmvBasic in params is invalid")

	raw, err := GetBasicByID(userID, db)
	if err != nil {
		log.Error(err)
		return cgi.ErrMySQLInner
	}

	if raw == nil {
		return cgi.ErrUserNotExists
	}

	tx, err := db.Begin()
	if err != nil {
		log.Error(err)
		return cgi.ErrMySQLInner
	}

	for dwf := true; dwf; dwf = false {
		res, err := tx.Exec("DELETE FROM `DB_USER`.`T_BASIC` WHERE F_USER_ID=?", userID)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			break
		}

		affected, err := res.RowsAffected()
		if err != nil {
			log.Error(err)
			tx.Rollback()
			break
		}

		if affected != 1 {
			err = cgi.ErrNoAffected
			tx.Rollback()
			break
		}

		err = entry.RmvEntry(userID, es)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			break
		}

		tx.Commit()
	}

	if err != nil {
		return err
	}

	return err
}
