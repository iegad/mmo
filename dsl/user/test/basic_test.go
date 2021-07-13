package test

import (
	"testing"
	"time"

	"github.com/iegad/kraken/conf"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi/user"
	ds "github.com/iegad/mmo/ds/user"
)

func getClient() (*piper.Client, error) {
	return piper.NewClient(&piper.ClientOption{
		Protocol:  conf.PROTOCOL_TCP,
		Service:   "userService",
		EtcdHosts: []string{"127.0.0.1:2379"},
	})
}

func TestAddBasic(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	req := &user.AddBasicReq{
		Basic: &ds.Basic{
			Entry: &ds.Entry{
				Email:    "iegad2021@qq.com",
				Nickname: "iegad",
			},
		},
	}

	rsp := &user.AddBasicRsp{}

	err = pc.Call("AddBasic", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}

func TestModBasic(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	req := &user.ModBasicReq{
		Basic: &ds.Basic{
			Entry: &ds.Entry{
				UserID:   1,
				Gender:   1,
				PhoneNum: "15827323665",
			},
		},
	}
	rsp := &user.ModBasicRsp{}

	err = pc.Call("ModBasic", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}

func TestRmvBasic(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	req := &user.RmvBasicReq{
		UserID: 1,
	}

	rsp := &user.RmvBasicRsp{}

	err = pc.Call("RmvBasic", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}

func TestGetBasic(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	req := &user.GetBasicReq{}
	rsp := &user.GetBasicRsp{}

	err = pc.Call("GetBasic", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}

func TestGetEntry(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	req := &user.GetEntryReq{Key: "iegad"}
	rsp := &user.GetEntryRsp{}

	err = pc.Call("GetEntry", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}

func TestGetArchiveLog(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	req := &user.GetArchiveLogReq{UserID: 1}
	rsp := &user.GetArchiveLogRsp{}

	err = pc.Call("GetArchiveLog", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}

func TestAddPersonal(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	birth, _ := time.Parse("2006-01-02", "1987-07-22")

	req := &user.AddPersonalReq{
		Personal: &ds.Personal{
			UserID:      1,
			Name:        "肖琪",
			Nationality: "+86",
			ID:          "420107198707221515",
			Gender:      1,
			Birth:       birth.Unix(),
		},
	}
	rsp := &user.AddPersonalRsp{}

	err = pc.Call("AddPersonal", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}

func TestModPersonal(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	req := &user.ModPersonalReq{
		Personal: &ds.Personal{
			UserID:      1,
			Nationality: "+852",
		},
	}
	rsp := &user.ModPersonalRsp{}

	err = pc.Call("ModPersonal", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}

func TestGetPersonal(t *testing.T) {
	pc, err := getClient()
	if err != nil {
		t.Error(err)
		return
	}

	defer pc.Close()

	req := &user.GetPersonalReq{}
	rsp := &user.GetPersonalRsp{}

	err = pc.Call("GetPersonal", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))
}
