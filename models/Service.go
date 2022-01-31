package models

import "time"

type Service struct {
	Id          int       `json:"id" gorm:"id"`
	Name        string    `json:"name" gorm:"name"`
	Url         string    `json:"url" gorm:"url"`
	Description string    `json:"description" gorm:"description"`
	CreateAt time.Time `json:"created_at" gorm:"created_at"`
	UpdateAt time.Time `json:"update_at" gorm:"update_at"`
	CreateBy string    `json:"created_by" gorm:"created_by"`
	UpdateBy string    `json:"update_by" gorm:"update_by"`
}

func (l Service) TableName() string {
	return "service"
}
