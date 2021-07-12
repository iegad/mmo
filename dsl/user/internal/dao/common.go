package dao

import (
	"fmt"
)

const (
	N_USER_ENTRY       = "user_entry"
	N_USER_ARCHIVE_LOG = "user_archive_log"
	T_USER_BASIC       = "`DB_USER`.`T_BASIC`"
	T_USER_PERSONAL    = "`DB_USER`.`T_PERSONAL`"
)

var (
	TUserBasicFieldMap = map[string]string{
		"UserID":     "F_USER_ID",
		"Email":      "F_EMAIL",
		"PhoneNum":   "F_PHONE_NUM",
		"Gender":     "F_GENDER",
		"CreateTime": "F_CREATE_TIME",
		"LastUpdate": "F_LAST_UPDATE",
		"Nickname":   "F_NICKNAME",
		"VerCode":    "F_VER_CODE",
	}

	TUserPersonalFieldMap = map[string]string{
		"UserID":      "F_USER_ID",
		"Name":        "F_NAME",
		"Nationality": "F_NATIONALITY",
		"Gender":      "F_GENDER",
		"Birth":       "F_BIRTH",
		"ID":          "F_ID",
		"VerCode":     "F_VER_CODE",
	}
)

func IDToString(userID int64) string {
	return fmt.Sprintf("%d", userID)
}
