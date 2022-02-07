package models

import "time"

type Role struct {
	Id        int       `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Menu      []*Menu    `json:"menu"`
	Service   []*Service `json:"service"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdateAt  time.Time `json:"update_at" gorm:"update_at"`
	CreatedBy string    `json:"created_by" gorm:"created_by"`
	UpdateBy  string    `json:"update_by" gorm:"update_by"`
}

type RoleHasMenu struct {
	RoleId int `json:"role_id"`
	MenuId int `json:"menu_id"`
}

type RoleHasService struct {
	RoleId    int `json:"role_id"`
	ServiceId int `json:"service_id"`
}

func (l Role) TableName() string {
	return "role"
}
