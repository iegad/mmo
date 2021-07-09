package basic

import (
	"database/sql"
	"time"
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
)

func AddBasic(obj *user.Basic, db *sql.DB) error {
	utils.Assert(obj != nil && &obj.Entry != nil && db != nil, "AddBasic in params is invalid")

	if obj.Entry.UserID > 0 {
		return cgi.ErrUserID
	}

	if len(obj.Entry.Email) == 0 && len(obj.Entry.PhoneNum) == 0 {
		return cgi.ErrAccount
	}

	if len(obj.Entry.Nickname) > 0 && utf8.RuneCountInString(obj.Entry.Nickname) > 8 {
		return cgi.ErrNickname
	}

	if len(obj.Entry.Email) > 0 && utf8.RuneCountInString(obj.Entry.Email) > 50 {
		return cgi.ErrEmail
	}

	if len(obj.Entry.PhoneNum) > 0 && utf8.RuneCountInString(obj.Entry.PhoneNum) > 15 {
		return cgi.ErrPhoneNum
	}

	if obj.Entry.Gender == 0 {
		obj.Entry.Gender = 3
	}

	obj.CreateTime = time.Now().Unix()

	tx, err := db.Begin()
	if err != nil {
		log.Error(err)
		return cgi.ErrMySQLInner
	}

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

	err = entry.AddEntry(obj.Entry, com.Elastic)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return cgi.ErrESInner
	}

	return tx.Commit()
}
