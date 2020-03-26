package conf

import (
	"flag"

	"chick/pkg/orm"
	"chick/pkg/store"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	env      string
	filePath string
	Conf     = &Config{}
)

type Config struct {
	Addrs   []string
	MySQL   *orm.MySQL
	Redis   *orm.Redis
	Clients []*store.ClientInfo
	Name    string
}

func init() {
	flag.StringVar(&env, "env", "", "env or prod")
	flag.StringVar(&filePath, "conf", "", "conf file path")
}

// 解析配置文件
func Parse() error {
	flag.Parse()
	if filePath == "" {
		return errors.New("load conf path fail")
	}
	viper.SetConfigFile(filePath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.UnmarshalKey(env, Conf)
	if err != nil {
		return err
	}
	return nil
}
