package userModels

import "time"

type User struct {
	ID           uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
	CreateTime   time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime   time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (User) TableName() string {
	return "user"
}
