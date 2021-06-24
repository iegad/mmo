package cfg

import (
	"os"

	"github.com/iegad/kraken/conf"
	"gopkg.in/yaml.v3"
)

var Instance *config

type config struct {
	Server *conf.Server `json:"server" yaml:"server"`
	Redis  *conf.Redis  `json:"redis"  yaml:"redis"`
	Etcd   *conf.Etcd   `json:"etcd"   yaml:"etcd"`
	MySql  *conf.MySql  `json:"mysql"  yaml:"mysql"`
}

func Init(fname ...string) error {
	file := "config.yaml"
	if len(fname) > 0 && len(fname[0]) > 0 {
		file = fname[0]
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	c := &config{}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	Instance = c

	return nil
}
