package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/iamzhiyudong/go-gin-example/models"
	"github.com/iamzhiyudong/go-gin-example/pkg/app"
	"github.com/iamzhiyudong/go-gin-example/pkg/e"
	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
	"github.com/iamzhiyudong/go-gin-example/pkg/util"
	"github.com/unknwon/com"
)

// @Summary 获取多个文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{ list: [], total: 0},"msg":"ok"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) { // c *gin.Context是Gin很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的 JSON 和呈现 JSON 响应
	name := c.Query("name") // c.Query可用于获取?name=test&state=1这类 URL 参数，而c.DefaultQuery则支持设置一个默认值

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query string false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	appG := app.Gin{C: c}

	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{} // 用于表单验证
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	data := make(map[string]string)

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, data)
		return
	}

	if models.ExistTagByName(name) {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, data)
		return
	} else {
		models.AddTag(name, state, createdBy)
		appG.Response(http.StatusOK, e.SUCCESS, data)
		return
	}
}

// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if !models.ExistTagByID(id) {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	} else {
		data := make(map[string]interface{})
		data["modified_by"] = modifiedBy
		if name != "" {
			data["name"] = name
		}
		if state != -1 {
			data["state"] = state
		}

		models.EditTag(id, data)
		appG.Response(http.StatusOK, e.SUCCESS, make(map[string]string))
	}
}

// @Summary 删除文章标签
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt() // param 和 query

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
