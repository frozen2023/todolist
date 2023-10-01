package route

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todolist/controller"
)

/*
*route层, 负责路由转发
 */
func GetRouter() *gin.Engine {
	r := gin.Default()
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 注册
	r.POST("/user", controller.Register)
	// 登录
	r.POST("/user/login", controller.Login)
	return r
}
