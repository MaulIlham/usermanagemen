package repository

import (
	"gorm.io/gorm"
	"log"
	"time"
	"usermanagement/models"
)

type UserController struct {
	db *gorm.DB
}

func UserNewController(db *gorm.DB) *UserController {
	return &UserController{db}
}

func (c UserController) InsertUser(newUser *models.User) error {
	newUser.CreateAt = time.Now()
	if err := c.db.Table("user").Save(newUser).Error; err!= nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (c UserController) ReadAllUser() ([]*models.User,error){
	users := []*models.User{}

	if err := c.db.Table("user").Find(&users).Error; err!= nil {
		log.Fatal(err)
		return nil,err
	}

	return users, nil
}

func (c UserController) ReadUserById(id int) (*models.User,error){
	user := models.User{}

	if err := c.db.Table("user").Where("id = ?",id).Find(user).Error; err!= nil {
		log.Fatal(err)
		return nil,err
	}

	return &user, nil
}

func (c UserController) UpdateUser(user *models.User) error {
	user.UpdateAt = time.Now()
	if err := c.db.Table("user").Updates(user).Error; err!= nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (c UserController) DeleteUser(id int) error {
	if err := c.db.Table("user").Delete(&models.User{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}


