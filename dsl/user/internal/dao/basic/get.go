package basic

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

const (
	query      = "SELECT F_USER_ID,F_EMAIL,F_PHONE_NUM,F_GENDER,F_NICKNAME,F_AVATOR,F_CREATE_TIME,F_LAST_UPDATE,F_VER_CODE FROM `DB_USER`.`T_BASIC` WHERE 1=1"
	queryCount = "SELECT COUNT(1) FROM `DB_USER`.`T_BASIC` WHERE 1=1"
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

func GetBasic(cond *user.GetBasicReq, db *sql.DB) ([]*ds.Basic, int64, error) {
	utils.Assert(cond != nil && db != nil, "GetBasic in params is invalid")

	var (
		sb  = &strings.Builder{}
		sbc = &strings.Builder{}
	)

	sb.WriteString(query)
	sbc.WriteString(queryCount)

	if cond.UserID > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_USER_ID=%d", cond.UserID))
		sbc.WriteString(fmt.Sprintf(" AND F_USER_ID=%d", cond.UserID))
	}

	if len(cond.PhoneNum) > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_PHONE_NUM='%s'", cond.PhoneNum))
		sbc.WriteString(fmt.Sprintf(" AND F_PHONE_NUM='%s'", cond.PhoneNum))
	}

	if len(cond.Email) > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_EMAIL='%s'", cond.Email))
		sbc.WriteString(fmt.Sprintf(" AND F_EMAIL='%s'", cond.Email))
	}

	if cond.Gender > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_GENDER=%d", cond.Gender))
		sbc.WriteString(fmt.Sprintf(" AND F_GENDER=%d", cond.Gender))
	}

	if cond.CreateTimeBeg > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_CREATE_TIME >= %d", cond.CreateTimeBeg))
		sbc.WriteString(fmt.Sprintf(" AND F_CREATE_TIME >= %d", cond.CreateTimeBeg))
	}

	if cond.CreateTimeEnd > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_CREATE_TIME < %d", cond.CreateTimeEnd))
		sbc.WriteString(fmt.Sprintf(" AND F_CREATE_TIME < %d", cond.CreateTimeEnd))
	}

	if cond.LastUpdateBeg > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_LAST_UPDATE >= %d", cond.LastUpdateBeg))
		sbc.WriteString(fmt.Sprintf(" AND F_LAST_UPDATE >= %d", cond.LastUpdateBeg))
	}

	if cond.LastUpdateEnd > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_LAST_UPDATE < %d", cond.LastUpdateEnd))
		sbc.WriteString(fmt.Sprintf(" AND F_LAST_UPDATE < %d", cond.LastUpdateEnd))
	}

	if len(cond.OrderBy) > 0 {
		sb.WriteString(fmt.Sprintf(" ORDER BY %s", dao.TUserBasicFieldMap[cond.OrderBy]))
		if cond.Desc {
			sb.WriteString(" DESC")
		}
	}

	if cond.Limit > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d OFFSET %d", cond.Limit, cond.Offset))
	}

	var (
		rows, err = db.Query(sb.String())
		dataList  = []*ds.Basic{}
		total     int64
	)

	if err != nil {
		return nil, -1, err
	}

	for dwf := true; dwf; dwf = false {
		for rows.Next() {
			var (
				basic           = ds.NewBasic()
				email, phoneNum sql.NullString
			)

			basic.Entry = ds.NewEntry()
			err = rows.Scan(&basic.Entry.UserID, &email, &phoneNum, &basic.Entry.Gender, &basic.Entry.Nickname, &basic.Entry.Avator, &basic.CreateTime, &basic.LastUpdate, &basic.VerCode)
			if err != nil {
				log.Error(err)
				break
			}

			if email.Valid {
				basic.Entry.Email = email.String
			}

			if phoneNum.Valid {
				basic.Entry.PhoneNum = phoneNum.String
			}

			dataList = append(dataList, basic)
		}

		rows.Close()
		if err != nil {
			break
		}

		row := db.QueryRow(sbc.String())
		err = row.Scan(&total)
	}

	if err != nil {
		for _, basic := range dataList {
			ds.DeleteBasic(basic)
		}

		return nil, -1, err
	}

	return dataList, total, nil
}
