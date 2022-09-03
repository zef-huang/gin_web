package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "go_web/mysql_db"

type ResponseData struct {
	Code int64
	Message string
	Data interface{}
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    1000,
		Message: "success",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}


type UserInfo struct {
	Username string `json: username`
	Password string `json: password`
}


func SignUp(c *gin.Context) {
	var Params UserInfo
	c.ShouldBindJSON(&Params)
	// 写入数据库
	_, err := mysql_db.Db.Exec("INSERT INTO user(username,password)VALUES (?,?)","hzf","123456")
	fmt.Println(err)
	ResponseSuccess(c, nil)
}

func Login(c *gin.Context) {
	var Params UserInfo
	c.ShouldBindJSON(&Params)
	// 写入数据库
	UserRecord, err := mysql_db.Db.Exec("SELECT password FROM user WHERE username=? LIMIT 1", Params.Username)
	fmt.Println(UserRecord, err)

	ResponseSuccess(c,gin.H{
		"user_id": 1, //js识别的最大值：id值大于1<<53-1  int64: i<<63-1
		"user_name": "hzf",
		"access_token": "user.AccessToke",
		"refresh_token": "user.RefreshToken",
	})
}

func SetupRoute() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.POST("/signup", SignUp)
	v1.POST("/login", Login)

	return r
}

