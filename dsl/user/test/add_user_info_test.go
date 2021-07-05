package test

import (
	"testing"

	"github.com/iegad/kraken/conf"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/cgi"
	"github.com/iegad/mmo/data"
)

func TestAddUserInfoTest(t *testing.T) {
	cli, err := piper.NewClient(&piper.ClientOption{
		Protocol: conf.PROTOCOL_TCP,
		Host:     "127.0.0.1:10000",
	})

	if err != nil {
		t.Error(err)
		return
	}

	defer cli.Close()

	req := cgi.NewAddUserInfoReq()
	req.UserInfo = data.NewUserInfo()
	req.UserInfo.Email = "3333@111.11221"
	req.UserInfo.Gender = 1

	rsp := cgi.NewAddUserInfoRsp()

	err = cli.Call("AddUserInfo", req, rsp)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(utils.PbToJson(rsp))

	cgi.DeleteAddUserInfoReq(req)
	cgi.DeleteAddUserInfoRsp(rsp)
}
