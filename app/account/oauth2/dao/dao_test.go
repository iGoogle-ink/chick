package dao

import (
	"os"
	"testing"

	"chick/app/account/oauth2/conf"
	"chick/pkg/log"
)

var d *Dao

func TestMain(m *testing.M) {
	os.Setenv("ACCOUNT_OAUTH_ENV", "prod")
	os.Setenv("ACCOUNT_OAUTH_CONF", "../cmd/oauth2.json")

	if err := conf.Parse(); err != nil {
		log.Panic(err)
	}
	d = New(conf.Conf)
	os.Exit(m.Run())
}
