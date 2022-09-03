package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
import "go_web/mysql_db"

type UserInfo struct {
	Username string `json: username`
	Password string `json: password`
}


func SignUp(c *gin.Context) {
	var Params UserInfo
	c.ShouldBindJSON(&Params)
	// 写入数据库
	result, err := mysql_db.Db.Exec("INSERT INTO user(username,password)VALUES (?,?)","hzf","123456")
	fmt.Println(err)
	c.JSON(200, gin.H{
		"Username": Params.Username,
		"Password": Params.Password,
		"result": result,
	})
}

func SetupRoute() *gin.Engine {
	r := gin.Default()
	v1 = r.Group("/api/v1")
	v1.POST("/signup", SignUp)
	return r
}

