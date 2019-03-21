package api

import (
	"net/http"

	"github.com/wonderivan/logger"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"gin-blog/service/authService"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" form:"username" json:"username" binding:"required"`
	Password string `valid:"Required; MaxSize(50)"  form:"password" json:"password" binding:"required"`
}

/*
@GetAuth 获取token
@params username 用户名
@params password 密码
*/
func GetAuth(c *gin.Context) {
	var user auth
	appG := app.Gin{C: c}
	httpCode, errCode := app.BindAndValid(c, &user)
	username := user.Username
	password := user.Password
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	authService := authService.Auth{UserName: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		logger.Error(err)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

}
