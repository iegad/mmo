package cgi

import "sync"

var poolAddUserInfoReq = sync.Pool{
	New: func() interface{} {
		return &AddUserInfoReq{}
	},
}

func NewAddUserInfoReq() *AddUserInfoReq {
	return poolAddUserInfoReq.Get().(*AddUserInfoReq)
}

func DeleteAddUserInfoReq(obj *AddUserInfoReq) {
	poolAddUserInfoReq.Put(obj)
}

var poolAddUserInfoRsp = sync.Pool{
	New: func() interface{} {
		return &AddUserInfoRsp{}
	},
}

func NewAddUserInfoRsp() *AddUserInfoRsp {
	return poolAddUserInfoRsp.Get().(*AddUserInfoRsp)
}

func DeleteAddUserInfoRsp(obj *AddUserInfoRsp) {
	poolAddUserInfoRsp.Put(obj)
}

var poolAddUserListReq = sync.Pool{
	New: func() interface{} {
		return &AddUserListReq{}
	},
}

func NewAddUserListReq() *AddUserListReq {
	return poolAddUserListReq.Get().(*AddUserListReq)
}

func DeleteAddUserListReq(obj *AddUserListReq) {
	poolAddUserListReq.Put(obj)
}

var poolAdduserListRsp = sync.Pool{
	New: func() interface{} {
		return &AdduserListRsp{}
	},
}

func NewAdduserListRsp() *AdduserListRsp {
	return poolAdduserListRsp.Get().(*AdduserListRsp)
}

func DeleteAdduserListRsp(obj *AdduserListRsp) {
	poolAdduserListRsp.Put(obj)
}

var poolModifyUserInfoReq = sync.Pool{
	New: func() interface{} {
		return &ModifyUserInfoReq{}
	},
}

func NewModifyUserInfoReq() *ModifyUserInfoReq {
	return poolModifyUserInfoReq.Get().(*ModifyUserInfoReq)
}

func DeleteModifyUserInfoReq(obj *ModifyUserInfoReq) {
	poolModifyUserInfoReq.Put(obj)
}

var poolModifyUserInfoRsp = sync.Pool{
	New: func() interface{} {
		return &ModifyUserInfoRsp{}
	},
}

func NewModifyUserInfoRsp() *ModifyUserInfoRsp {
	return poolModifyUserInfoRsp.Get().(*ModifyUserInfoRsp)
}

func DeleteModifyUserInfoRsp(obj *ModifyUserInfoRsp) {
	poolModifyUserInfoRsp.Put(obj)
}

var poolRemoveUserInfoReq = sync.Pool{
	New: func() interface{} {
		return &RemoveUserInfoReq{}
	},
}

func NewRemoveUserInfoReq() *RemoveUserInfoReq {
	return poolRemoveUserInfoReq.Get().(*RemoveUserInfoReq)
}

func DeleteRemoveUserInfoReq(obj *RemoveUserInfoReq) {
	poolRemoveUserInfoReq.Put(obj)
}

var poolRemoveUserInfoRsp = sync.Pool{
	New: func() interface{} {
		return &RemoveUserInfoRsp{}
	},
}

func NewRemoveUserInfoRsp() *RemoveUserInfoRsp {
	return poolRemoveUserInfoRsp.Get().(*RemoveUserInfoRsp)
}

func DeleteRemoveUserInfoRsp(obj *RemoveUserInfoRsp) {
	poolRemoveUserInfoRsp.Put(obj)
}

var poolQueryUserInfoReq = sync.Pool{
	New: func() interface{} {
		return &QueryUserInfoReq{}
	},
}

func NewQueryUserInfoReq() *QueryUserInfoReq {
	return poolQueryUserInfoReq.Get().(*QueryUserInfoReq)
}

func DeleteQueryUserInfoReq(obj *QueryUserInfoReq) {
	poolQueryUserInfoReq.Put(obj)
}

var poolQueryUserInfoRsp = sync.Pool{
	New: func() interface{} {
		return &QueryUserInfoRsp{}
	},
}

func NewQueryUserInfoRsp() *QueryUserInfoRsp {
	return poolQueryUserInfoRsp.Get().(*QueryUserInfoRsp)
}

func DeleteQueryUserInfoRsp(obj *QueryUserInfoRsp) {
	poolQueryUserInfoRsp.Put(obj)
}

var poolExistsUserInfoReq = sync.Pool{
	New: func() interface{} {
		return &ExistsUserInfoReq{}
	},
}

func NewExistsUserInfoReq() *ExistsUserInfoReq {
	return poolExistsUserInfoReq.Get().(*ExistsUserInfoReq)
}

func DeleteExistsUserInfoReq(obj *ExistsUserInfoReq) {
	poolExistsUserInfoReq.Put(obj)
}

var poolExistsUserInfoRsp = sync.Pool{
	New: func() interface{} {
		return &ExistsUserInfoRsp{}
	},
}

func NewExistsUserInfoRsp() *ExistsUserInfoRsp {
	return poolExistsUserInfoRsp.Get().(*ExistsUserInfoRsp)
}

func DeleteExistsUserInfoRsp(obj *ExistsUserInfoRsp) {
	poolExistsUserInfoRsp.Put(obj)
}

var poolSetUserSessionReq = sync.Pool{
	New: func() interface{} {
		return &SetUserSessionReq{}
	},
}

func NewSetUserSessionReq() *SetUserSessionReq {
	return poolSetUserSessionReq.Get().(*SetUserSessionReq)
}

func DeleteSetUserSessionReq(obj *SetUserSessionReq) {
	poolSetUserSessionReq.Put(obj)
}

var poolSetUserSessionRsp = sync.Pool{
	New: func() interface{} {
		return &SetUserSessionRsp{}
	},
}

func NewSetUserSessionRsp() *SetUserSessionRsp {
	return poolSetUserSessionRsp.Get().(*SetUserSessionRsp)
}

func DeleteSetUserSessionRsp(obj *SetUserSessionRsp) {
	poolSetUserSessionRsp.Put(obj)
}

var poolGetUserSessionReq = sync.Pool{
	New: func() interface{} {
		return &GetUserSessionReq{}
	},
}

func NewGetUserSessionReq() *GetUserSessionReq {
	return poolGetUserSessionReq.Get().(*GetUserSessionReq)
}

func DeleteGetUserSessionReq(obj *GetUserSessionReq) {
	poolGetUserSessionReq.Put(obj)
}

var poolGetUserSessionRsp = sync.Pool{
	New: func() interface{} {
		return &GetUserSessionRsp{}
	},
}

func NewGetUserSessionRsp() *GetUserSessionRsp {
	return poolGetUserSessionRsp.Get().(*GetUserSessionRsp)
}

func DeleteGetUserSessionRsp(obj *GetUserSessionRsp) {
	poolGetUserSessionRsp.Put(obj)
}

var poolRemoveUserSessionReq = sync.Pool{
	New: func() interface{} {
		return &RemoveUserSessionReq{}
	},
}

func NewRemoveUserSessionReq() *RemoveUserSessionReq {
	return poolRemoveUserSessionReq.Get().(*RemoveUserSessionReq)
}

func DeleteRemoveUserSessionReq(obj *RemoveUserSessionReq) {
	poolRemoveUserSessionReq.Put(obj)
}

var poolRemoveUserSessionRsp = sync.Pool{
	New: func() interface{} {
		return &RemoveUserSessionRsp{}
	},
}

func NewRemoveUserSessionRsp() *RemoveUserSessionRsp {
	return poolRemoveUserSessionRsp.Get().(*RemoveUserSessionRsp)
}

func DeleteRemoveUserSessionRsp(obj *RemoveUserSessionRsp) {
	poolRemoveUserSessionRsp.Put(obj)
}
