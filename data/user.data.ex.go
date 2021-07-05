package data

import "sync"

const TUserInfo = "T_USER_INFO"

var (
	TUserInfoFieldMap = map[string]string{}
)

func init() {
	TUserInfoFieldMap["UserID"] = "F_AID"
	TUserInfoFieldMap["Email"] = "F_EMAIL"
	TUserInfoFieldMap["PhoneNum"] = "F_PHONE_NUM"
	TUserInfoFieldMap["Gender"] = "F_GENDER"
	TUserInfoFieldMap["Nickname"] = "F_NICKNAME"
	TUserInfoFieldMap["Avator"] = "F_AVATOR"
	TUserInfoFieldMap["CreateTime"] = "F_CREATE_TIME"
	TUserInfoFieldMap["LastUpdate"] = "F_LAST_UPDATE"
	TUserInfoFieldMap["Ver"] = "F_VER"
	TUserInfoFieldMap["VerCode"] = "F_VER_CODE"
}

var poolUserInfo = sync.Pool{
	New: func() interface{} {
		return &UserInfo{}
	},
}

func NewUserInfo() *UserInfo {
	return poolUserInfo.Get().(*UserInfo)
}

func DeleteUserInfo(obj *UserInfo) {
	poolUserInfo.Put(obj)
}

var poolUserSession = sync.Pool{
	New: func() interface{} {
		return &UserSession{}
	},
}

func NewUserSession() *UserSession {
	return poolUserSession.Get().(*UserSession)
}

func DeleteUserSession(obj *UserSession) {
	poolUserSession.Put(obj)
}
