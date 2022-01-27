package repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"usermanagement/models"
)

type MenuController struct {
	db *gorm.DB
}

func MenuNewController(db *gorm.DB) *MenuController {
	return &MenuController{db}
}

func (c MenuController) InsertMenu(newMenu *models.Menu) error {
	newMenu.CreateAt = time.Now()
	if err := c.db.Table("menu").Save(newMenu).Error; err!= nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (c MenuController) ReadAllMenu() ([]*models.Menu,error){
	menu := []*models.Menu{}

	if err := c.db.Table("menu").Find(&menu).Error; err!= nil {
		log.Fatal(err)
		return nil,err
	}

	return menu, nil
}

func (c MenuController) ReadMenuById(id int) (*models.Menu,error){
	menu := models.Menu{}

	if err := c.db.Table("menu").Where("id = ?",id).Find(menu).Error; err!= nil {
		log.Fatal(err)
		return nil,err
	}

	return &menu, nil
}

func (c MenuController) UpdateMenu(menu *models.Menu) error {
	menu.UpdateAt = time.Now()
	if err := c.db.Table("menu").Updates(menu).Error; err!= nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (c MenuController) DeleteMenu(id int) error {
	if err := c.db.Table("menu").Delete(&models.Menu{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

