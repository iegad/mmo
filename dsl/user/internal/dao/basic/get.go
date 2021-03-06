package basic

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

const (
	query      = "SELECT F_USER_ID,F_EMAIL,F_PHONE_NUM,F_GENDER,F_NICKNAME,F_AVATOR,F_CREATE_TIME,F_LAST_UPDATE,F_VER_CODE FROM `DB_USER`.`T_BASIC` WHERE 1=1"
	queryCount = "SELECT COUNT(1) FROM `DB_USER`.`T_BASIC` WHERE 1=1"
)

// GetBasicByID 通过UserID获取用户基础信息
func GetBasicByID(userID int64, db *sql.DB) (*ds.Basic, error) {
	utils.Assert(userID > 0 && db != nil, "GetBasicByID in params is invalid")

	var (
		basic = &ds.Basic{
			Entry: &ds.Entry{},
		}
		email, phoneNum sql.NullString
	)

	row := db.QueryRow("SELECT F_EMAIL,F_PHONE_NUM,F_GENDER,F_NICKNAME,F_AVATOR,F_CREATE_TIME,F_LAST_UPDATE,F_VER_CODE FROM `DB_USER`.`T_BASIC` WHERE F_USER_ID=?", userID)
	err := row.Scan(&email, &phoneNum, &basic.Entry.Gender, &basic.Entry.Nickname, &basic.Entry.Avator, &basic.CreateTime, &basic.LastUpdate, &basic.VerCode)
	if err != nil {
		return nil, err
	}

	if email.Valid {
		basic.Entry.Email = email.String
	}

	if phoneNum.Valid {
		basic.Entry.PhoneNum = phoneNum.String
	}

	basic.Entry.UserID = userID
	return basic, nil
}

// GetBasic 获取基础信息
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

	var total int64
	row := db.QueryRow(sbc.String())
	err := row.Scan(&total)
	if err != nil {
		return nil, -1, err
	}

	if total == 0 {
		return nil, 0, nil
	}

	dataList := []*ds.Basic{}
	rows, err := db.Query(sb.String())
	if err != nil {
		return nil, -1, err
	}

	for rows.Next() {
		var (
			item = &ds.Basic{
				Entry: &ds.Entry{},
			}
			email, phoneNum sql.NullString
		)

		err = rows.Scan(&item.Entry.UserID, &email, &phoneNum, &item.Entry.Gender, &item.Entry.Nickname, &item.Entry.Avator, &item.CreateTime, &item.LastUpdate, &item.VerCode)
		if err != nil {
			break
		}

		if email.Valid {
			item.Entry.Email = email.String
		}

		if phoneNum.Valid {
			item.Entry.PhoneNum = phoneNum.String
		}

		dataList = append(dataList, item)
	}

	rows.Close()
	if err != nil {
		return nil, -1, err
	}

	return dataList, total, nil
}
