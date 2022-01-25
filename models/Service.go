package models

import "time"

type Service struct {
	Id          int       `json:"id" gorm:"id"`
	Name        string    `json:"name" gorm:"name"`
	Url         string    `json:"url" gorm:"url"`
	Description string    `json:"description" gorm:"description"`
	CreateAt    time.Time `json:"createAt" gorm:"createAt"`
	UpdateAt    time.Time `json:"updateAt" gorm:"updateAt"`
	CreateBy    string    `json:"createBy" gorm:"createBy"`
	UpdateBy    string    `json:"updateBy" gorm:"updateBy"`
}

func (l Service) TableName() string {
	return "service"
}
