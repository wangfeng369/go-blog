package main

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"

	"github.com/wonderivan/logger"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   int    `form:"passwd" json:"passwd" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}

func init() {
	setting.Setup()
	models.Setup()
	logger.SetLogger("./conf/log.json")
}
func main() {
	// r := gin.Default()
	// r.POST("/ping", func(ctx *gin.Context) {
	// 	var user User
	// 	err := ctx.Bind(&user)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(user.Username)
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"code": 200,
	// 		"msg":  "success",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	//log.Printf("[info] start http server listening %s", endPoint)
	logger.Info("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
