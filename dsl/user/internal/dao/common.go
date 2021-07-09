package dao

import (
	"fmt"
)

var (
	N_USER_ENTRY = "user_entry"
	T_USER_BASIC = "`DB_USER`.`T_BASIC`"
)

func IDToString(userID int64) string {
	return fmt.Sprintf("%d", userID)
}
