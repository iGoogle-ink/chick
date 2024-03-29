package orm

import (
	"time"

	"chick/pkg/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MySQL mysql config.
type MySQL struct {
	DSN         string // data source name.
	Active      int    // pool
	Idle        int    // pool
	IdleTimeout int    // connect max life time. second
	ShowSQL     bool
}

func InitMySQL(c *MySQL) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		log.Panicf("failed to connect database error:%+v", err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.LogMode(c.ShowSQL)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	return db
}
