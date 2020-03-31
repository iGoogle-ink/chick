package conf

import (
	"chick/pkg/orm"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	Conf = &Config{}
)

type Config struct {
	Name           string
	Port           string
	Addrs          []string
	TokenExpiresIn int //单位:秒
	MySQL          *orm.MySQL
	Redis          *orm.Redis
}

// 解析配置文件
func Parse() error {

	env := os.Getenv("OAUTH_ENV")
	filePath := os.Getenv("OAUTH_CONF")
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
