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
