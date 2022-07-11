package route

import (
	"gin-JWT/api"
	"gin-JWT/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	user := r.Group("/api/user")
	{
		user.POST("/login",api.Login)
	}
	general := r.Group("/api/test")
	general.Use(middleware.JWTAuth())	// 加载自定义JWTAuth()中间件，在整合general路由都生效
	{
		general.GET("/ga1",api.GeneralApi)
	}
	return r
}
