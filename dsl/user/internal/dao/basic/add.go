package basic

import (
	"database/sql"
	"time"

	"github.com/iegad/kraken/utils"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
	"github.com/olivere/elastic/v7"
)

func AddBasic(obj *ds.Basic, db *sql.DB, es *elastic.Client) error {
	utils.Assert(obj != nil && &obj.Entry != nil && db != nil && es != nil, "AddBasic in params is invalid")

	obj.CreateTime = time.Now().Unix()

	// Step 1: 开启MYSQL事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Step 2: 写入Basic 到 MYSQL
	res, err := tx.Exec("INSERT INTO `DB_USER`.`T_BASIC`(F_EMAIL,F_PHONE_NUM,F_GENDER,F_NICKNAME,F_AVATOR,F_CREATE_TIME) VALUES(?,?,?,?,?,?)",
		utils.IF_STR_EMPTY(obj.Entry.Email), utils.IF_STR_EMPTY(obj.Entry.PhoneNum), obj.Entry.Gender, obj.Entry.Nickname, obj.Entry.Avator, obj.CreateTime)

	for dwf := true; dwf; dwf = false {
		if err != nil {
			break
		}

		obj.Entry.UserID, err = res.LastInsertId()
		if err != nil {
			break
		}

		// Step 3: 写入Entry到 ES
		err = entry.AddEntry(obj.Entry, es)
		if err != nil {
			break
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
