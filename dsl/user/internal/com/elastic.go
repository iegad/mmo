package com

import (
	"fmt"
	"strings"

	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/dsl/user/internal/cfg"
	"github.com/olivere/elastic/v7"
)

var Elastic *elastic.Client

func initElastic() error {
	utils.Assert(cfg.Instance != nil, "配置文件未加载")

	var (
		hosts = []string{}
		err   error
	)

	for _, host := range cfg.Instance.Elastic.Hosts {
		if strings.Index(host, "http") != 0 {
			hosts = append(hosts, fmt.Sprintf("http://%s/", host))
		}
	}

	Elastic, err = elastic.NewClient(elastic.SetURL(hosts...), elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return nil

}
