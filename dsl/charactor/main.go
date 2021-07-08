package main

import (
	"github.com/iegad/kraken/conf"
	"github.com/iegad/kraken/log"
	"github.com/iegad/mmo/dsl/charactor/internal/cfg"
	"github.com/iegad/mmo/dsl/charactor/internal/com"
)

func main() {
	fname, err := conf.DefaultFile()
	if err != nil {
		log.Fatal("获取配置文件失败: %v", err)
	}

	err = cfg.Init(fname)
	if err != nil {
		log.Fatal("加载配置文件失败: %v", err)
	}

	if len(cfg.Instance.Server.LogPath) > 0 {
		log.SetPath(cfg.Instance.Server.LogPath)
	}

	err = com.Init()
	if err != nil {
		log.Fatal("组件加载失败: %v", err)
	}

}
