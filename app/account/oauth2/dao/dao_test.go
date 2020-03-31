package dao

import (
	"os"
	"testing"

	"chick/app/account/oauth2/conf"
)

var d *Dao

func TestMain(m *testing.M) {
	os.Setenv("OAUTH_ENV", "prod")
	os.Setenv("OAUTH_CONF", "../cmd/oauth2.json")
	if err := conf.Parse(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	os.Exit(m.Run())
}
