package personal

import (
	"database/sql"

	"github.com/iegad/kraken/utils"
	ds "github.com/iegad/mmo/ds/user"
)

func AddPersonal(obj *ds.Personal, db *sql.DB) error {
	utils.Assert(obj != nil && db != nil, "AddPersonal in params is invalid")

	_, err := db.Exec("INSERT INTO `DB_USER`.`T_PERSONAL`(F_USER_ID,F_NAME,F_NATIONALITY,F_GENDER,F_BIRTH,F_ID) VALUES(?,?,?,?,?,?)",
		obj.UserID, obj.Name, obj.Nationality, obj.Gender, obj.Birth, obj.ID)

	return err
}
