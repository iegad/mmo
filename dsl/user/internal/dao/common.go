package dao

import (
	"fmt"
)

const (
	N_USER_ENTRY = "user_entry"
	T_USER_BASIC = "`DB_USER`.`T_BASIC`"
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
)

func IDToString(userID int64) string {
	return fmt.Sprintf("%d", userID)
}
