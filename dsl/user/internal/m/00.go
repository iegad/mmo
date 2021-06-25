package m

import (
	"errors"

	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/nw/server"
	"github.com/iegad/kraken/piper"
)

var (
	ErrCtx = errors.New("piper.context is invalid")
	ErrReq = errors.New("req is invalid")
	ErrRsp = errors.New("rsp is invalid")
)

type UserService struct {
}

func (this_ *UserService) OnConnected(conn server.IConn) {
	log.Info("%s has connected", conn.RemoteAddr().String())
}

func (this_ *UserService) OnDisconnect(conn server.IConn) {
	log.Info("%s has disconnected", conn.RemoteAddr().String())
}

func (this_ *UserService) OnStop(svr *piper.Server) {
	log.Info("server is stopped")
}

func (this_ *UserService) OnRun(svr *piper.Server) {
	log.Info("server is running")
}

func (this_ *UserService) Decode(c server.IConn, data []byte) ([]byte, error) {
	return data, nil
}

func (this_ *UserService) Encode(c server.IConn, data []byte) ([]byte, error) {
	return data, nil
}
