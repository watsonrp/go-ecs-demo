package main

import (
	"fmt"
	_ "techlabs/models"
	_ "techlabs/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", "tlUser:1qazxcde32ws@/tldb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	name := "default"
	force := false
	verbose := false
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
