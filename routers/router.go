package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-smallproject/pkg/setting"
	"go-gin-smallproject/routers/api/v1"
)

/**
处理跨越
*/
func middleware(c *gin.Context) {
	// gin设置响应头，设置跨域
	c.Header("Access-Control-Allow-Origin", "http://localhost:8000")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware)
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// SetMode根据输入字符串设置gin模式。
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
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
		//上传表格
		apiv1.POST("/exel", v1.UpdataExel)
		//导出表格
		apiv1.GET("/downexel", v1.DownExel)

	}

	return r
}
