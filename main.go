package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "quickstart/routers"
)


func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql",
	//	"debian-sys-maint:tFKrxhcOVXiDRPiN@tcp(18.224.3.47:3306)/default?charset=utf8")
}

func main() {
	//err := orm.RunSyncdb("default", true, true)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//o := orm.NewOrm()
	//o.Using("orm_test")

	//user := new(models.User)
	//user.Name = "larryxing"
	//fmt.Println(o.Insert(user))

	beego.Run()
}

