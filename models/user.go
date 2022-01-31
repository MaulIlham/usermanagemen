package models

import "time"

type User struct {
	Id        int       `json:"id" gorm:"id"`
	Username  string    `json:"username" gorm:"username"`
	Password  string    `json:"password" gorm:"password"`
	Email     string    `json:"email" gorm:"email"`
	Status    bool      `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdateAt  time.Time `json:"update_at" gorm:"update_at"`
	CreatedBy string    `json:"created_by" gorm:"created_by"`
	UpdateBy  string    `json:"update_by" gorm:"update_by"`
}

func (l User) TableName() string {
	return "user"
}
