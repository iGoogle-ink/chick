package dao

import (
	"chick/app/account/oauth2/conf"
	"chick/pkg/orm"

	"github.com/jinzhu/gorm"
)

type Dao struct {
	DB *gorm.DB
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		DB: orm.InitMySQL(c.MySQL),
	}
	return
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}
