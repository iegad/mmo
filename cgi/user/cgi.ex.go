package user

import (
	"sync"
)

var poolAddEntryReq = sync.Pool{
	New: func() interface{} {
		return &AddEntryReq{}
	},
}

func NewAddEntryReq() *AddEntryReq {
	return poolAddEntryReq.Get().(*AddEntryReq)
}

func DeleteAddEntryReq(req *AddEntryReq) {
	poolAddEntryReq.Put(req)
}

var poolAddEntryRsp = sync.Pool{
	New: func() interface{} {
		return &AddEntryRsp{}
	},
}

func NewAddEntryRsp() *AddEntryRsp {
	return poolAddEntryRsp.Get().(*AddEntryRsp)
}

func DeleteAddEntryRsp(rsp *AddEntryRsp) {
	poolAddEntryRsp.Put(rsp)
}

var poolModEntryReq = sync.Pool{
	New: func() interface{} {
		return &ModEntryReq{}
	},
}

func NewModEntryReq() *ModEntryReq {
	return poolModEntryReq.Get().(*ModEntryReq)
}

func DeleteModEntryReq(req *ModEntryReq) {
	poolModEntryReq.Put(req)
}

var poolModEntryRsp = sync.Pool{
	New: func() interface{} {
		return &ModEntryRsp{}
	},
}

func NewModEntryRsp() *ModEntryRsp {
	return poolModEntryRsp.Get().(*ModEntryRsp)
}

func DeleteModEntryRsp(rsp *ModEntryRsp) {
	poolModEntryRsp.Put(rsp)
}

var poolRmvEntryReq = sync.Pool{
	New: func() interface{} {
		return &RmvEntryReq{}
	},
}

func NewRmvEntryReq() *RmvEntryReq {
	return poolRmvEntryReq.Get().(*RmvEntryReq)
}

func DeleteRmvEntryReq(req *RmvEntryReq) {
	poolRmvEntryReq.Put(req)
}

var poolRmvEntryRsp = sync.Pool{
	New: func() interface{} {
		return &RmvEntryRsp{}
	},
}

func NewRmvEntryRsp() *RmvEntryRsp {
	return poolRmvEntryRsp.Get().(*RmvEntryRsp)
}

func DeleteRmvEntryRsp(rsp *RmvEntryRsp) {
	poolRmvEntryRsp.Put(rsp)
}

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
