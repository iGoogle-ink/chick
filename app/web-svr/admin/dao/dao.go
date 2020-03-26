package dao

import (
	"chick/api/oauth2"
	"chick/app/web-svr/admin/conf"
	"chick/pkg/orm"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
)

type Dao struct {
	DB        *gorm.DB
	oauthGrpc oauth2.Oauth2Service
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		DB:        orm.InitMySQL(c.MySQL),
		oauthGrpc: newOauth2Client("micro.service.oauth2.client", "micro.service.oauth2"),
	}
	return
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}

func newOauth2Client(cliName, srvName string) oauth2.Oauth2Service {
	client := micro.NewService(micro.Name(cliName))
	client.Init()

	oauthCli := oauth2.NewOauth2Service(srvName, client.Client())
	return oauthCli
}
