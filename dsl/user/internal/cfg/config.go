package cfg

import (
	"os"

	"github.com/iegad/kraken/conf"
	"github.com/iegad/kraken/utils"
	"gopkg.in/yaml.v3"
)

var Instance *config

type config struct {
	Server *conf.Server `json:"server" yaml:"server"`
	Redis  *conf.Redis  `json:"redis"  yaml:"redis"`
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

	err = conf.CheckRedis(c.Redis)
	if err != nil {
		return err
	}

	if len(c.Server.ID) == 0 {
		c.Server.ID = utils.UUID_String()
		out, err := yaml.Marshal(c)
		if err != nil {
			return err
		}

		err = os.WriteFile(fname, out, 0755)
		if err != nil {
			return err
		}
	}

	Instance = c
	return nil
}
