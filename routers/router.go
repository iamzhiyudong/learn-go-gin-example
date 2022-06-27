package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/iamzhiyudong/go-gin-example/docs"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/iamzhiyudong/go-gin-example/middleware/jwt"
	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
	"github.com/iamzhiyudong/go-gin-example/pkg/upload"
	"github.com/iamzhiyudong/go-gin-example/routers/api"
	v1 "github.com/iamzhiyudong/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	// 创建路由 Handlers
	// 后期可以绑定各类的路由规则和函数、中间件等
	// router := gin.Default()

	r := gin.New() // 源码中会调用打印启动日志的函数

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 创建不同的 HTTP 方法绑定到Handlers中
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })

	// 当访问 $HOST/upload/images 时
	// 将会读取到 $GOPATH/src/github.com/xxx/go-gin-example/runtime/upload/images 下的文件
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/upload", api.UploadImage)

	r.GET("/auth", api.GetAuth)

	// 注册 v1 版本接口
	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
