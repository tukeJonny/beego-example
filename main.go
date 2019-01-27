package main

import (
	"time"

	_ "github.com/tukejonny/beego-example/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// for develop (sqlite3)
	if err := orm.RegisterDriver("sqlite3", orm.DRSqlite); err != nil {
		panic(err)
	}
	if err := orm.RegisterDataBase("dev", "sqlite3", "file:beego.db"); err != nil {
		panic(err)
	}

	// for production (mysql)
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		panic(err)
	}
	if err := orm.RegisterDataBase("prod", "mysql", "root:"); err != nil {
		panic(err)
	}

	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.Local
}

func main() {
	db := orm.NewOrm()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		orm.Debug = true
		db.Using("dev")
	} else {
		db.Using("prod")
	}

	beego.Run()
}
