package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"gin-blog/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.PostForm("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}