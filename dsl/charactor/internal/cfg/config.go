package cfg

import (
	"os"

	"github.com/iegad/kraken/conf"
	"gopkg.in/yaml.v3"
)

var Instance *config

type config struct {
	Server *conf.Server `json:"server" yaml:"server"`
	Etcd   *conf.Etcd   `json:"etcd"   yaml:"etcd"`
	MySql  *conf.MySql  `json:"mysql"  yaml:"mysql"`
}

func Init(fname string) error {
	data, err := os.ReadFile(fname)
	if err != nil {
		return err
	}

	c := &config{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	err = conf.CheckServer(c.Server)
	if err != nil {
		return err
	}

	err = conf.CheckEtcd(c.Etcd)
	if err != nil {
		return err
	}

	err = conf.CheckMySql(c.MySql)
	if err != nil {
		return err
	}

	Instance = c
	return nil
}
