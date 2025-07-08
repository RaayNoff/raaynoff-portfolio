package blogModels

import "time"

type Blog struct {
	ID               uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"shortDescription" gorm:"column:shortDescription"`
	Content          string    `json:"content"`
	CreateTime       time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime       time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (Blog) TableName() string {
	return "blog"
}
