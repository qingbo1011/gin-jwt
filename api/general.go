package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GeneralApi 定义一个普通的api，作为验证接口的逻辑
func GeneralApi(c *gin.Context)  {
	// 我们在JWTAuth()中间中将'claims'写入到gin.Context的指针对象中，因此在这里可以将之解析出来
	claims, exists := c.Get("claims")
	if !exists{
		c.JSON(http.StatusOK,gin.H{
			"status":-1,
			"msg":"c.GET claims 失败",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"status":0,
		"msg":"token有效",
		"data":claims,
	})
}
