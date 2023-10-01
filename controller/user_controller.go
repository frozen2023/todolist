package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/model"
	"todolist/service"
)

func Register(c *gin.Context) {
	user := model.User{}
	// 绑定数据
	bind(user, c)
	// 插入数据
	resp(c, service.Register(&user))
}

func Login(c *gin.Context) {
	user := model.User{}
	bind(&user, c)
	resp(c, service.Login(&user))
}

func bind(entity any, c *gin.Context) {
	if err := c.ShouldBindJSON(entity); err != nil {
		c.JSON(
			http.StatusOK,
			model.FailResult(),
		)
	}
}

func resp(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}
