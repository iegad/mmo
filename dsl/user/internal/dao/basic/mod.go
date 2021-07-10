package basic

import (
	"database/sql"
	"time"
	"unicode/utf8"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/ds/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao/entry"
	"github.com/olivere/elastic/v7"
)

func ModBasic(obj *ds.Basic, db *sql.DB, es *elastic.Client) error {
	utils.Assert(obj != nil && &obj.Entry != nil && db != nil && es != nil, "ModBasic in params is invalid")

	if obj.Entry.UserID <= 0 {
		return cgi.ErrUserID
	}

	var (
		err error
		raw *ds.Basic
	)

	for i := 0; i < 3; i++ {
		raw, err = GetBasicByID(obj.Entry.UserID, db)
		if err != nil {
			return err
		}

		err = updateBasic(obj, raw, db, es)
		if err == nil || err != cgi.ErrNoAffected {
			break
		}
	}

	if err != nil {
		return err
	}

	// TODO: 记录存档日志
	return err
}

func updateBasic(obj, raw *user.Basic, db *sql.DB, es *elastic.Client) error {
	var err error

	for dwf := true; dwf; dwf = false {
		changed := false

		if len(obj.Entry.Email) > 0 {
			if len(obj.Entry.Email) > ds.MAX_EMAIL {
				err = cgi.ErrEmail
				break
			}

			if obj.Entry.Email != raw.Entry.Email {
				found, err := existsEmail(obj.Entry.Email, db)
				if err != nil {
					log.Error(err)
					err = cgi.ErrMySQLInner
					break
				}

				if found {
					err = cgi.ErrEmailExists
					break
				}

				changed = true
			}
		} else {
			obj.Entry.Email = raw.Entry.Email
		}

		if len(obj.Entry.PhoneNum) > 0 {
			if len(obj.Entry.PhoneNum) > ds.MAX_PHONE_NUM {
				err = cgi.ErrPhoneNum
				break
			}

			if obj.Entry.PhoneNum != raw.Entry.PhoneNum {
				found, err := existsPhoneNum(obj.Entry.PhoneNum, db)
				if err != nil {
					log.Error(err)
					err = cgi.ErrMySQLInner
					break
				}

				if found {
					err = cgi.ErrPhoneNumExists
					break
				}

				changed = true
			}
		} else {
			obj.Entry.PhoneNum = raw.Entry.PhoneNum
		}

		if obj.Entry.Gender > 0 {
			if obj.Entry.Gender < ds.MIN_GENDER || obj.Entry.Gender > ds.MAX_GENDER {
				err = cgi.ErrGender
				break
			}

			if obj.Entry.Gender != raw.Entry.Gender {
				changed = true
			}
		} else {
			obj.Entry.Gender = raw.Entry.Gender
		}

		if len(obj.Entry.Nickname) > 0 {
			if utf8.RuneCountInString(obj.Entry.Nickname) > ds.MAX_NICKNAME {
				err = cgi.ErrNickname
				break
			}

			if obj.Entry.Nickname != raw.Entry.Nickname {
				changed = true
			}
		} else {
			obj.Entry.Nickname = raw.Entry.Nickname
		}

		if len(obj.Entry.Avator) > 0 {
			if len(obj.Entry.Avator) > ds.MAX_AVATOR {
				err = cgi.ErrAvator
				break
			}

			if obj.Entry.Avator != raw.Entry.Avator {
				changed = true
			}
		} else {
			obj.Entry.Avator = raw.Entry.Avator
		}

		if !changed {
			err = cgi.ErrNoUpData
			break
		}

		obj.LastUpdate = time.Now().Unix()
		obj.CreateTime = raw.CreateTime
		obj.VerCode = raw.VerCode + 1

		tx, err := db.Begin()
		if err != nil {
			log.Error(err)
			err = cgi.ErrMySQLInner
			break
		}

		res, err := tx.Exec("UPDATE `DB_USER`.`T_BASIC` SET F_EMAIL=?,F_PHONE_NUM=?,F_GENDER=?,F_NICKNAME=?,F_AVATOR=?,F_VER_CODE=? WHERE F_USER_ID=? AND F_VER_CODE=?",
			utils.IF_STR_EMPTY(obj.Entry.Email), utils.IF_STR_EMPTY(obj.Entry.PhoneNum), obj.Entry.Gender, obj.Entry.Nickname, obj.Entry.Avator, obj.VerCode, obj.Entry.UserID, raw.VerCode)
		if err != nil {
			log.Error(err)
			err = cgi.ErrMySQLInner
			tx.Rollback()
			break
		}

		affected, err := res.RowsAffected()
		if err != nil {
			log.Error(err)
			err = cgi.ErrMySQLInner
			tx.Rollback()
			break
		}

		if affected != 1 {
			err = cgi.ErrNoAffected
			tx.Rollback()
			break
		}

		err = entry.ModEntry(obj.Entry, es)
		if err != nil {
			tx.Rollback()
			break
		}

		tx.Commit()
	}

	return err
}
