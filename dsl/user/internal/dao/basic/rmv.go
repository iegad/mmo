package basic

import (
	"database/sql"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
	"github.com/olivere/elastic/v7"
)

func RmvBasic(userID int64, db *sql.DB, es *elastic.Client) error {
	utils.Assert(userID > 0 && db != nil && es != nil, "RmvBasic in params is invalid")

	tx, err := db.Begin()
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = tx.Exec("DELETE FROM `DB_USER`.`T_BASIC` WHERE F_USER_ID=?", userID)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM `DB_USER`.`T_PERSONAL` WHERE F_USER_ID=?", userID)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	err = entry.RmvEntry(userID, es)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
