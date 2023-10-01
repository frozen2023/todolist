package service

import (
	"todolist/config"
	"todolist/model"
	"todolist/util"
)

var SnowFlakeUtil = new(util.Snowflake)

func Register(user *model.User) model.CommonResult {
	user.Id = SnowFlakeUtil.NextVal()
	result := config.SqlSession.Create(user)
	if result.Error != nil {
		return model.OkResult(nil)
	}
	return model.FailResult()
}

func Login(user *model.User) model.CommonResult {
	newUser := model.User{}
	config.SqlSession.Where("username = ?", user.Username).First(&newUser)
	// 判断密码是否正确
	if newUser.Password == user.Password {
		data := map[string]any{"token": util.MakeToken(newUser.Id)}
		return model.OkResult(data)
	}
	return model.FailResult()
}
