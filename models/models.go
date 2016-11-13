package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// LocalUser is for locally stored user accounts
type LocalUser struct {
	ID       int
	First    string
	Last     string
	Email    string `orm:"unique"`
	Password string
	RegKey   string
	RegDate  time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(LocalUser))
}
