package conf

import (
	"os"

	"chick/pkg/orm"
	"chick/pkg/store"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	Conf = &Config{}
)

type Config struct {
	Addrs   []string
	MySQL   *orm.MySQL
	Redis   *orm.Redis
	Clients []*store.ClientInfo
	Name    string
}

func init() {

}

// 解析配置文件
func Parse() error {
	env := os.Getenv("MICRO_OAUTH_ENV")
	filePath := os.Getenv("MICRO_OAUTH_CONF")
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
