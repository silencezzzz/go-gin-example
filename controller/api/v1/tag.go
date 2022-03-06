package v1

import (
	"fmt"
	"github.com/Silencezzzz/go-gin-example/kafka"
	"github.com/Silencezzzz/go-gin-example/models"
	"github.com/Silencezzzz/go-gin-example/pkg/e"
	"github.com/Silencezzzz/go-gin-example/pkg/setting"
	"github.com/Silencezzzz/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// GetTags 批量获取标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
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
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  e.GetMsg(code),
	})
}

// AddTag 批量新增标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := models.TagNormalState
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := e.InvalidParams

	if !valid.HasErrors() {
		if !models.TagIsExist(name) {
			code = e.SUCCESS
			models.AddTag(name, createdBy, state)
		} else {
			code = e.ErrorExistTag
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除标签
func DelTag(c *gin.Context) {

}

//修改标签
func EditTag(c *gin.Context) {

}

//修改标签
func SendMsg(c *gin.Context) {
	code := e.InvalidParams
	msg := c.Query("msg")
	valid := validation.Validation{}
	valid.Required(msg, "msg").Message("名称不能为空")
	if !valid.HasErrors() {
		//校验成功
		fmt.Println("接收到的数据")
		fmt.Println(msg)
		bl := kafka.SyncProducer(msg)
		if !bl {
			code = e.ErrorProductKafka
		} else {
			code = e.SUCCESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
