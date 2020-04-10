package conf

import (
	"os"

	"chick/pkg/config"
	"chick/pkg/orm"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	Conf = &Config{}
)

type Config struct {
	Name           string
	HttpPort       string
	TokenExpiresIn int // 单位:秒
	Registry       *config.Registry
	Broker         *config.Broker
	MySQL          *orm.MySQL
	Redis          *orm.Redis
}

// 解析配置文件
func Parse() error {
	env := os.Getenv("ACCOUNT_OAUTH_ENV")
	filePath := os.Getenv("ACCOUNT_OAUTH_CONF")
	if filePath == "" {
		return errors.New("load config path fail")
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
