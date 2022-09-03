package main

import (
	"fmt"
	"go_web/routes"
)
import "go_web/mysql_db"

//type UserInfo struct {
//	// 表明该属性与传递的json的user属性进行绑定
//	Username     string `json:"username" binding:"required"`
//	Age int `json:"age" binding:"required"`
//}


func main() {
	err := mysql_db.InitDB()
	if err != nil {
		fmt.Println("mysql 配置失败")
		return
	}
	r := routes.SetupRoute()
	r.Run()
}
