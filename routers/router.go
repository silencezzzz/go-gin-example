package routers

import (
	v1 "github.com/Silencezzzz/go-gin-example/controller/api/v1"
	"github.com/Silencezzzz/go-gin-example/pkg/e"
	"github.com/Silencezzzz/go-gin-example/pkg/setting"
	"github.com/Silencezzzz/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		apiV1.GET("/test", v1.AddTag)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.GET("/sendMsg", v1.SendMsg)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DelTag)
		apiV1.GET("/getToken", func(c *gin.Context) {
			token, _ := util.GenerateToken("name", "pwd")
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  e.GetMsg(200),
				"data": map[string]string{"token": token},
			})
		})
	}

	return r
}
