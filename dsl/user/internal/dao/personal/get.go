package personal

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao"
)

const (
	query      = "SELECT F_USER_ID,F_NAME,F_NATIONALITY,F_GENDER,F_BIRTH,F_ID,F_VER_CODE FROM `DB_USER`.`T_PERSONAL` WHERE 1=1"
	queryCount = "SELECT COUNT(1) FROM `DB_USER`.`T_PERSONAL` WHERE 1=1"
)

func GetPersonal(cond *user.GetPersonalReq, db *sql.DB) ([]*ds.Personal, int64, error) {
	utils.Assert(cond != nil && db != nil, "GetPersonal in parmas is invalid")

	sb := strings.Builder{}
	sbc := strings.Builder{}

	sb.WriteString(query)
	sbc.WriteString(queryCount)

	if cond.BirthBeg > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_BIRTH >= %d", cond.BirthBeg))
		sbc.WriteString(fmt.Sprintf(" AND F_BIRTH >= %d", cond.BirthBeg))
	}

	if cond.BirthEnd > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_BIRTH < %d", cond.BirthEnd))
		sbc.WriteString(fmt.Sprintf(" AND F_BIRTH < %d", cond.BirthEnd))
	}

	if cond.Gender > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_GENDER=%d", cond.Gender))
		sbc.WriteString(fmt.Sprintf(" AND F_GENDER=%d", cond.Gender))
	}

	if len(cond.ID) > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_ID='%s'", cond.ID))
		sbc.WriteString(fmt.Sprintf(" AND F_ID='%s'", cond.ID))
	}

	if len(cond.Name) > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_NAME='%s'", cond.Name))
		sbc.WriteString(fmt.Sprintf(" AND F_NAME='%s'", cond.Name))
	}

	if len(cond.Nationality) > ds.MAX_NATIONALITY {
		sb.WriteString(fmt.Sprintf(" AND F_NATIONALITY='%s'", cond.Nationality))
		sbc.WriteString(fmt.Sprintf(" AND F_NATIONALITY='%s'", cond.Nationality))
	}

	if cond.UserID > 0 {
		sb.WriteString(fmt.Sprintf(" AND F_USER_ID=%d", cond.UserID))
		sbc.WriteString(fmt.Sprintf(" AND F_USER_ID=%d", cond.UserID))
	}

	if len(cond.OrderBy) > 0 {
		sb.WriteString(fmt.Sprintf(" ORDER BY %s", dao.TUserPersonalFieldMap[cond.OrderBy]))
		if cond.Desc {
			sb.WriteString(" DESC")
		}
	}

	if cond.Limit > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d", cond.Limit))

		if cond.Offset > 0 {
			sb.WriteString(fmt.Sprintf(" OFFSET %d", cond.Offset))
		}
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

	rows, err := db.Query(sb.String())
	if err != nil {
		log.Error(err)
		return nil, -1, err
	}

	dataList := []*ds.Personal{}

	for rows.Next() {
		item := &ds.Personal{}

		err = rows.Scan(&item.UserID, &item.Name, &item.Nationality, &item.Gender, &item.Birth, &item.ID, &item.VerCode)
		if err != nil {
			break
		}

		dataList = append(dataList, item)
	}

	rows.Close()
	if err != nil {
		return nil, -1, err
	}

	return dataList, total, nil
}

func GetPersonalByID(userID int64, db *sql.DB) (*ds.Personal, error) {
	utils.Assert(userID > 0 && db != nil, "GetPersonalByID in parmas is invalid")

	personal := &ds.Personal{}
	row := db.QueryRow("SELECT F_NAME,F_NATIONALITY,F_GENDER,F_BIRTH,F_ID,F_VER_CODE FROM `DB_USER`.`T_PERSONAL` WHERE F_USER_ID=?", userID)
	err := row.Scan(&personal.Name, &personal.Nationality, &personal.Gender, &personal.Birth, &personal.ID, &personal.VerCode)
	if err != nil {
		return nil, err
	}

	personal.UserID = userID
	return personal, nil
}
