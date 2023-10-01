package model

type User struct {
	Id       int64  `gorm:"column:id;type:bigint;not null;primaryKey" json:"id"`
	Username string `gorm:"column:username;type:varchar(20);not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(20);not null" json:"password"`
}

// 指定表名
func (User) TableName() string {
	return "user"
}
