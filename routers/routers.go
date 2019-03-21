package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	r.POST("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.POST("/tags", v1.GetTags)
		apiv1.POST("/addTags",v1.AddTag)
		apiv1.POST("/editTags",v1.EditTag)
	}

	return r
}
