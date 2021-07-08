package cgi

import "errors"

var (
	ErrObj      = errors.New("obj is invalid")
	ErrDB       = errors.New("db is invalid")
	ErrESInner  = errors.New("elastic error")
	ErrNoUpData = errors.New("no data will be update")

	ErrEmail    = errors.New("email is invalid")
	ErrPhoneNum = errors.New("phoneNum is invalid")
	ErrAccount  = errors.New("account is invalid")
	ErrUserID   = errors.New("userID is invalid")
	ErrTimeout  = errors.New("更新数据超时")

	ErrCharID        = errors.New("charID is invalid")
	ErrInfo          = errors.New("info is invalid")
	ErrNickname      = errors.New("nickname is invalid")
	ErrRole          = errors.New("role is invalid")
	ErrGender        = errors.New("gender is invalid")
	ErrHairStyle     = errors.New("hairStyle is invalid")
	ErrHairColor     = errors.New("hairColor is invalid")
	ErrHeight        = errors.New("height is invalid")
	ErrBodyType      = errors.New("bodyType is invalid")
	ErrSkinColor     = errors.New("skinColor is invalid")
	ErrBasic         = errors.New("basic is invalid")
	ErrEXP           = errors.New("EXP is invalid")
	ErrLEV           = errors.New("LEV is invalid")
	ErrSTA           = errors.New("STA is invalid")
	ErrSTR           = errors.New("STR is invalid")
	ErrAGI           = errors.New("AGI is invalid")
	ErrCON           = errors.New("CON is invalid")
	ErrINT           = errors.New("INT is invalid")
	ErrLCK           = errors.New("LCK is invalid")
	ErrUserNotExists = errors.New("user not exists")
)
