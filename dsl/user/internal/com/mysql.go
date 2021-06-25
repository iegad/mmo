package com

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iegad/kraken/utils"
	"github.com/iegad/mmo/dsl/user/internal/cfg"
)

var MySql *sql.DB

func initMySql() error {
	utils.Assert(cfg.Instance != nil, "配置文件未加载")

	var err error
	MySql, err = sql.Open("mysql", cfg.Instance.MySql.String())
	return err
}
