package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iegad/kraken/conf"
	"github.com/iegad/kraken/log"
	"github.com/iegad/kraken/piper"
	"github.com/iegad/mmo/dsl/user/internal/cfg"
	"github.com/iegad/mmo/dsl/user/internal/com"
	"github.com/iegad/mmo/dsl/user/internal/m"
)

func main() {
	fname, err := conf.DefaultFile()
	if err != nil {
		log.Fatal("获取配置文件失败: %v", err)
	}

	err = cfg.Init(fname)
	if err != nil {
		log.Fatal("加载配置失败: %v", err)
	}

	err = com.Init()
	if err != nil {
		log.Fatal("组件初始化失败: %v", err)
	}

	svc, err := piper.NewServer(&m.UserService{}, &piper.ServerOption{
		ID:        cfg.Instance.Server.ID,
		Protocol:  cfg.Instance.Server.Protocol,
		Service:   cfg.Instance.Server.Service,
		EtcdHosts: cfg.Instance.Etcd.Hosts,
		Host:      cfg.Instance.Server.Host,
		MaxConn:   cfg.Instance.Server.MaxConn,
		Timeout:   cfg.Instance.Server.Timeout,
	})
	if err != nil {
		log.Fatal("初始化服务失败: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-done
		svc.Stop()
	}()

	err = svc.Run()
	if err != nil {
		log.Fatal(err)
	}
}
