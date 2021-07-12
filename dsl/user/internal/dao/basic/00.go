package basic

import (
	"database/sql"
)

func ExistsPhoneNum(phoneNum string, db *sql.DB) (bool, error) {
	count := 0
	row := db.QueryRow("SELECT COUNT(1) FROM `DB_USER`.`T_BASIC` WHERE F_PHONE_NUM=?", phoneNum)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func ExistsEmail(email string, db *sql.DB) (bool, error) {
	count := 0
	row := db.QueryRow("SELECT COUNT(1) FROM `DB_USER`.`T_BASIC` WHERE F_EMAIL=?", email)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
