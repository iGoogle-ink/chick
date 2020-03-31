package model

import "time"

type (
	CloudUser struct {
		Id        int `gorm:"primary_key"`
		Uname     string
		Passwd    string
		Phone     string
		IsDeleted int
		Mtime     time.Time
	}
)

func (m *CloudUser) TableName() string {
	return "cloud_user"
}

func (m *CloudUser) CopyFrom(req *RegisterReq) {
	m.Uname = req.Uname
	m.Passwd = req.Passwd
	m.Phone = req.Phone
}

type RegisterReq struct {
	Uname  string `form:"uname" json:"uname"`
	Passwd string `form:"passwd" json:"passwd"`
	Phone  string `form:"phone" json:"phone"`
}

type LoginReq struct {
	Uname  string `form:"uname" json:"uname"`
	Passwd string `form:"passwd" json:"passwd"`
}
