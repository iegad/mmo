package user

import (
	"sync"
)

var poolGetEntryReq = sync.Pool{
	New: func() interface{} {
		return &GetEntryReq{}
	},
}

func NewGetEntryReq() *GetEntryReq {
	return poolGetEntryReq.Get().(*GetEntryReq)
}

func DeleteGetEntryReq(req *GetEntryReq) {
	poolGetEntryReq.Put(req)
}

var poolGetEntryRsp = sync.Pool{
	New: func() interface{} {
		return &GetEntryRsp{}
	},
}

func NewGetEntryRsp() *GetEntryRsp {
	return poolGetEntryRsp.Get().(*GetEntryRsp)
}

func DeleteGetEntryRsp(rsp *GetEntryRsp) {
	poolGetEntryRsp.Put(rsp)
}

var poolAddBasicReq = sync.Pool{
	New: func() interface{} {
		return &AddBasicReq{}
	},
}

func NewAddBasicReq() *AddBasicReq {
	return poolAddBasicReq.Get().(*AddBasicReq)
}

func DeleteAddBasicReq(req *AddBasicReq) {
	poolAddBasicReq.Put(req)
}

var poolAddBasicRsp = sync.Pool{
	New: func() interface{} {
		return &AddBasicRsp{}
	},
}

func NewAddBasicRsp() *AddBasicRsp {
	return poolAddBasicRsp.Get().(*AddBasicRsp)
}

func DeleteAddBasicRsp(rsp *AddBasicRsp) {
	poolAddBasicRsp.Put(rsp)
}

var poolModBasicReq = sync.Pool{
	New: func() interface{} {
		return &ModBasicReq{}
	},
}

func NewModBasicReq() *ModBasicReq {
	return poolModBasicReq.Get().(*ModBasicReq)
}

func DeleteModBasicReq(req *ModBasicReq) {
	poolModBasicReq.Put(req)
}

var poolModBasicRsp = sync.Pool{
	New: func() interface{} {
		return &ModBasicRsp{}
	},
}

func NewModBasicRsp() *ModBasicRsp {
	return poolModBasicRsp.Get().(*ModBasicRsp)
}

func DeleteModBasicRsp(rsp *ModBasicRsp) {
	poolModBasicRsp.Put(rsp)
}

var poolRmvBasicReq = sync.Pool{
	New: func() interface{} {
		return &RmvBasicReq{}
	},
}

func NewRmvBasicReq() *RmvBasicReq {
	return poolRmvBasicReq.Get().(*RmvBasicReq)
}

func DeleteRmvBasicReq(req *RmvBasicReq) {
	poolRmvBasicReq.Put(req)
}

var poolRmvBasicRsp = sync.Pool{
	New: func() interface{} {
		return &RmvBasicRsp{}
	},
}

func NewRmvBasicRsp() *RmvBasicRsp {
	return poolRmvBasicRsp.Get().(*RmvBasicRsp)
}

func DeleteRmvBasicRsp(rsp *RmvBasicRsp) {
	poolRmvBasicRsp.Put(rsp)
}

var poolGetBasicReq = sync.Pool{
	New: func() interface{} {
		return &GetBasicReq{}
	},
}

func NewGetBasicReq() *GetBasicReq {
	return poolGetBasicReq.Get().(*GetBasicReq)
}

func DeleteGetBasicReq(req *GetBasicReq) {
	poolGetBasicReq.Put(req)
}

var poolGetBasicRsp = sync.Pool{
	New: func() interface{} {
		return &GetBasicRsp{}
	},
}

func NewGetBasicRsp() *GetBasicRsp {
	return poolGetBasicRsp.Get().(*GetBasicRsp)
}

func DeleteGetBasicRsp(rsp *GetBasicRsp) {
	poolGetBasicRsp.Put(rsp)
}
