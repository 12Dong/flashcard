package main

import (
	_ "flashcard/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	initEnv()
	beego.Run()
}

func initEnv() {
	var dataBaseUrl string = os.Getenv("databaseurl")
	var username string = os.Getenv("username")
	var password string = os.Getenv("password")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+dataBaseUrl+")/flashcard?charset=utf8")

}
