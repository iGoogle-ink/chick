package model

// db model
type (
	MxUser struct {
		Name string `json:"name"`
	}
)

func (m *MxUser) TableName() string {
	return "mx_user"
}

type LoginRsp struct {
	Token string `json:"token"`
}

func (d *LoginRsp) CopyFromUser(user *MxUser) {
	d.Token = user.Name
}
