package dao

import (
	"fmt"
	"strings"
)

var (
	N_USER_ENTRY = "user_entry"
)

func IDToString(userID int64) string {
	return fmt.Sprintf("%d", userID)
}

func EncodePhoneNum(phoneNum string) string {
	sb := strings.Builder{}
	for _, char := range phoneNum {
		sb.WriteRune(char ^ 0xFF)
	}

	return sb.String()
}

func DecodePhoneNum(phoneNum string) string {
	return EncodePhoneNum(phoneNum)
}

func EncodeEmail(email string) string {
	sb := strings.Builder{}
	for _, char := range email {
		sb.WriteRune(char ^ 0xCC)
	}

	return sb.String()
}

func DecodeEmail(email string) string {
	return EncodeEmail(email)
}
