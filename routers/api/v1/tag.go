package v1

import (
	"fmt"
	"gin-blog/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/Unknwon/com"

	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	tagservice "gin-blog/service/tag_service"
)

type GetTagForm struct {
	Name  string `form:"name" json:"name" binding:"required"`
	State int    `form:"state" json:"state" binding:"required"`
}

//获取多个文章标签
func GetTags(ctx *gin.Context) {
	var tag GetTagForm
	err := ctx.Bind(&tag)
	if err != nil {
		fmt.Println(err)
	}
	name := tag.Name
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	fmt.Println(name)
	if name != "" {
		maps["name"] = name
	}

	// var state int = -1
	arg := tag.State
	if arg != -1 {
		// 	state = com.StrTo(arg).MustInt()
		maps["state"] = arg
	}

	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(ctx), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type addTagForm struct {
	Name     string `form:"name" json:"name" valid:"Required;MaxSize(100)" binding:"required"`
	State    int    `form:"state" json:"state" valid:"Range(0,1)" binding:"required"`
	CreateBy string `form:"createBy" json:"createBy" valid:"Required;MaxSize(100)" binding:"required"`
}

//@AddTag 添加标签
func AddTag(ctx *gin.Context) {
	var (
		addTag addTagForm
		appG   = app.Gin{C: ctx}
	)
	httpCode, errCode := app.BindAndValid(ctx, &addTag)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	tagService := tagservice.Tag{
		Name:      addTag.Name,
		State:     addTag.State,
		CreatedBy: addTag.CreateBy,
	}
	isExist, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if isExist {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}
	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditTagForm struct {
	ID         int    `form:"id" json:"id" valid:"Required;Min(1)" binding:"required"`
	Name       string `form:"name" json:"name" valid:"MaxSize(100)"`
	ModifiedBy string `form:"modifiedBy" json:"modifiedBy" valid:"MaxSize(100)"`
	State      int    `form:"state" json:"state" valid:"Range(0,1)"`
}

func EditTag(ctx *gin.Context) {
	var (
		editTagFrom EditTagForm
		appG        = app.Gin{C: ctx}
	)
	httpCode, errCode := app.BindAndValid(ctx, &editTagFrom)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	tagService := tagservice.Tag{
		Name:       editTagFrom.Name,
		State:      editTagFrom.State,
		ID:         editTagFrom.ID,
		ModifiedBy: editTagFrom.ModifiedBy,
	}
	Exist, err := tagService.ExistById()
	if !Exist {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
