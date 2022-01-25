package models

import "time"

type User struct {
	Id       int       `json:"id" gorm:"id"`
	Username string    `json:"username" gorm:"username"`
	Password string    `json:"password" gorm:"password"`
	Email    string    `json:"email" gorm:"email"`
	Status   bool      `json:"status" gorm:"status"`
	CreateAt time.Time `json:"createAt" gorm:"createAt"`
	UpdateAt time.Time `json:"updateAt" gorm:"updateAt"`
	CreateBy string    `json:"createBy" gorm:"createBy"`
	UpdateBy string    `json:"updateBy" gorm:"updateBy"`
}

func (l User) TableName() string {
	return "user"
}