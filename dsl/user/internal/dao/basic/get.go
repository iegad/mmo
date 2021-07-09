package basic

import (
	"database/sql"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	ds "github.com/iegad/mmo/ds/user"
)

func GetBasicByID(userID int64, db *sql.DB) (*ds.Basic, error) {
	utils.Assert(userID > 0 && db != nil, "GetBasicByID in params is invalid")

	basic := ds.NewBasic()
	basic.Entry = ds.NewEntry()
	row := db.QueryRow("SELECT F_EMAIL,F_PHONE_NUM,F_GENDER,F_NICKNAME,F_AVATOR,F_CREATE_TIME,F_LAST_UPDATE,F_VER_CODE FROM `DB_USER`.`T_BASIC` WHERE F_USER_ID=?", userID)

	var (
		email, phoneNum sql.NullString
		err             error
	)

	for dwf := true; dwf; dwf = false {
		err = row.Scan(&email, &phoneNum, &basic.Entry.Gender, &basic.Entry.Nickname, &basic.Entry.Avator, &basic.CreateTime, &basic.LastUpdate, &basic.VerCode)
		if err != nil {
			log.Error(err)
			err = cgi.ErrMySQLInner
			break
		}

		if email.Valid {
			basic.Entry.Email = email.String
		}

		if phoneNum.Valid {
			basic.Entry.PhoneNum = phoneNum.String
		}

		basic.Entry.UserID = userID
	}

	if err != nil {
		ds.DeleteEntry(basic.Entry)
		ds.DeleteBasic(basic)
		return nil, err
	}

	return basic, nil
}
