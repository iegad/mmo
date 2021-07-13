package personal

import (
	"database/sql"
	"errors"
	"time"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/utils"
	ds "github.com/iegad/mmo/ds/user"
	"github.com/iegad/mmo/dsl/user/internal/dao/archive_log"
	"github.com/olivere/elastic/v7"
)

func ModPersonal(obj *ds.Personal, db *sql.DB, es *elastic.Client) error {
	utils.Assert(obj != nil && db != nil, "ModPersonal in parmas is invalid")

	var err error
	for i := 0; i < 3; i++ {
		err = updatePersonal(obj, db, es)
		if err == nil {
			break
		}
	}

	return err
}

func updatePersonal(obj *ds.Personal, db *sql.DB, es *elastic.Client) error {
	raw, err := GetPersonalByID(obj.UserID, db)
	if err != nil {
		return err
	}

	n := 0
	if obj.Birth == 0 {
		obj.Birth = raw.Birth
		n++
	}

	if obj.Gender == 0 {
		obj.Gender = raw.Gender
		n++
	}

	if len(obj.ID) == 0 {
		obj.ID = raw.ID
		n++
	}

	if len(obj.Name) == 0 {
		obj.Name = raw.Name
		n++
	}

	if len(obj.Nationality) == 0 {
		obj.Nationality = raw.Nationality
		n++
	}

	if n == 5 {
		return nil
	}

	obj.VerCode = raw.VerCode + 1

	res, err := db.Exec("UPDATE `DB_USER`.`T_PERSONAL` SET F_BIRTH=?,F_GENDER=?,F_ID=?,F_NAME=?,F_NATIONALITY=?,F_VER_CODE=? WHERE F_USER_ID=? AND F_VER_CODE=?",
		obj.Birth, obj.Gender, obj.ID, obj.Name, obj.Nationality, obj.VerCode, obj.UserID, raw.VerCode)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return errors.New("no row affected")
	}

	err = archive_log.AddArchiveLog(&ds.ArchiveLog{
		Type:        ds.ArchiveLogType_ArchiveLogType_Personal,
		VerCode:     obj.VerCode,
		UserID:      obj.UserID,
		ArchiveTime: time.Now().Unix(),
		Raw:         utils.PbToBytes(raw),
		Changed:     utils.PbToBytes(obj),
	}, es)
	if err != nil {
		log.Error(err)
	}

	return nil
}
