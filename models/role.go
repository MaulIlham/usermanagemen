package models

import "time"

type Role struct {
	Id       int       `json:"id" gorm:"id"`
	Name     string    `json:"name" gorm:"name"`
	CreateAt time.Time `json:"createAt" gorm:"createAt"`
	UpdateAt time.Time `json:"updateAt" gorm:"updateAt"`
	CreateBy string    `json:"createBy" gorm:"createBy"`
	UpdateBy string    `json:"updateBy" gorm:"updateBy"`
}

func (l Role) TableName() string {
	return "role"
}
