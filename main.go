package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "todolist/docs"
	"todolist/route"
)

// @BasePath /api/v1

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object}  "成功"
// @Failure 400 {object}  "请求错误"
// @Failure 500 {object}  "内部错误"
// @Router /api/v1/tags [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

const server_port = ":9000"

func main() {
	// 通过路由获取engine
	r := route.GetRouter()
	r.Run(server_port)
}
