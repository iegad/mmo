package basic

import (
	"database/sql"
	"time"
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
	"github.com/olivere/elastic/v7"
)

func AddBasic(obj *ds.Basic, db *sql.DB, es *elastic.Client) error {
	utils.Assert(obj != nil && &obj.Entry != nil && db != nil && es != nil, "AddBasic in params is invalid")

	// Step 1: 入参检查
	if obj.Entry.UserID > 0 {
		return cgi.ErrUserID
	}

	if len(obj.Entry.Email) == 0 && len(obj.Entry.PhoneNum) == 0 {
		return cgi.ErrAccount
	}

	if len(obj.Entry.Nickname) > 0 && utf8.RuneCountInString(obj.Entry.Nickname) > ds.MAX_NICKNAME {
		return cgi.ErrNickname
	}

	if len(obj.Entry.Avator) > ds.MAX_AVATOR {
		return cgi.ErrAvator
	}

	if len(obj.Entry.Email) > 0 {
		if len(obj.Entry.Email) > ds.MAX_EMAIL {
			return cgi.ErrEmail
		}

		found, err := existsEmail(obj.Entry.Email, db)
		if err != nil {
			log.Error(err)
			return cgi.ErrMySQLInner
		}

		if found {
			return cgi.ErrEmailExists
		}
	}

	if len(obj.Entry.PhoneNum) > 0 {
		if len(obj.Entry.PhoneNum) > ds.MAX_PHONE_NUM {
			return cgi.ErrPhoneNum
		}

		exists, err := existsPhoneNum(obj.Entry.PhoneNum, db)
		if err != nil {
			log.Error(err)
			return cgi.ErrMySQLInner
		}

		if exists {
			return cgi.ErrPhoneNum
		}
	}

	if obj.Entry.Gender > ds.MAX_GENDER || obj.Entry.Gender < ds.MIN_GENDER {
		return cgi.ErrGender
	}

	if obj.Entry.Gender == 0 {
		obj.Entry.Gender = 3
	}

	obj.CreateTime = time.Now().Unix()

	// Step 2: 开启MYSQL事务
	tx, err := db.Begin()
	if err != nil {
		log.Error(err)
		return cgi.ErrMySQLInner
	}

	// Step 3: 写入Basic 到 MYSQL
	res, err := tx.Exec("INSERT INTO `DB_USER`.`T_BASIC`(F_EMAIL,F_PHONE_NUM,F_GENDER,F_NICKNAME,F_AVATOR,F_CREATE_TIME) VALUES(?,?,?,?,?,?)",
		utils.IF_STR_EMPTY(obj.Entry.Email), utils.IF_STR_EMPTY(obj.Entry.PhoneNum), obj.Entry.Gender, obj.Entry.Nickname, obj.Entry.Avator, obj.CreateTime)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return cgi.ErrMySQLInner
	}

	obj.Entry.UserID, err = res.LastInsertId()
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return cgi.ErrMySQLInner
	}

	// Step 4: 写入Entry到 ES
	err = entry.AddEntry(obj.Entry, es)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return cgi.ErrESInner
	}

	// Step 5: 提交事务
	return tx.Commit()
}

func existsPhoneNum(phoneNum string, db *sql.DB) (bool, error) {
	count := 0
	row := db.QueryRow("SELECT COUNT(1) FROM `DB_USER`.`T_BASIC` WHERE F_PHONE_NUM=?", phoneNum)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func existsEmail(email string, db *sql.DB) (bool, error) {
	count := 0
	row := db.QueryRow("SELECT COUNT(1) FROM `DB_USER`.`T_BASIC` WHERE F_EMAIL=?", email)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
