package cgi

import "errors"

var (
	ErrRsp        = errors.New("none response")
	ErrESInner    = errors.New("elastic error")          // ES返回错误
	ErrMySQLInner = errors.New("mysql error")            // MYSQL 返回错误
	ErrNoUpData   = errors.New("no data will be update") // 没有需要更新的数据
	ErrNoAffected = errors.New("no row affected")        // DML操作受影响的行数为0

	ErrUserID        = errors.New("userID is invalid")       // UserID 无效
	ErrUserNotExists = errors.New("user not exists")         // 用户不存在
	ErrUserExists    = errors.New("user has already exists") // 用户已存在

	ErrEmail          = errors.New("email is invalid")               // EMAIL 无效
	ErrEmailExists    = errors.New("email has already exists")       // EMAIL已存在
	ErrPhoneNum       = errors.New("phoneNum is invalid")            // 手机号无效
	ErrPhoneNumExists = errors.New("phoneNum has alread exists")     // 手机号已存在
	ErrAccount        = errors.New("account is invalid")             // 账号(邮箱或手机号)无效
	ErrKey            = errors.New("key is invalid")                 // 搜索用户条目时的KEY
	ErrNoCond         = errors.New("query data with none condition") // 没有搜索条件
	ErrESDataType     = errors.New("elastic data type is invalid")   // ES存储的数据类型不匹配

	ErrAvator  = errors.New("avator is invalid")
	ErrTimeout = errors.New("更新数据超时")

	ErrCharID    = errors.New("charID is invalid")
	ErrInfo      = errors.New("info is invalid")
	ErrNickname  = errors.New("nickname is invalid")
	ErrRole      = errors.New("role is invalid")
	ErrGender    = errors.New("gender is invalid")
	ErrHairStyle = errors.New("hairStyle is invalid")
	ErrHairColor = errors.New("hairColor is invalid")
	ErrHeight    = errors.New("height is invalid")
	ErrBodyType  = errors.New("bodyType is invalid")
	ErrSkinColor = errors.New("skinColor is invalid")
	ErrBasic     = errors.New("basic is invalid")
	ErrEXP       = errors.New("EXP is invalid")
	ErrLEV       = errors.New("LEV is invalid")
	ErrSTA       = errors.New("STA is invalid")
	ErrSTR       = errors.New("STR is invalid")
	ErrAGI       = errors.New("AGI is invalid")
	ErrCON       = errors.New("CON is invalid")
	ErrINT       = errors.New("INT is invalid")
	ErrLCK       = errors.New("LCK is invalid")
)
