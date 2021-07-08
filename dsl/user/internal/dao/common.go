package dao

import "github.com/iegad/kraken/security"

var (
	ES_INDEX    = "user_entry"
	PhoneNumKey = []byte("user_entry_phone_num")
	EmailKey    = []byte("user_entry_email")
)

func init() {
	PhoneNumKey = security.MD5(PhoneNumKey)
	EmailKey = security.MD5(EmailKey)
}

// func EncodePhoneNum()
