package dao

import "errors"

var (
	ErrObj     = errors.New("obj is invalid")
	ErrDB      = errors.New("db is invalid")
	ErrAccount = errors.New("account is invalid")
	ErrUserID  = errors.New("userID is invalid")
)
