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
	newMenu.CreatedAt = time.Now()
	newMenu.UpdateAt = time.Now()
	if err := c.db.Table("menu").Save(newMenu).Error; err!= nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c MenuController) ReadAllMenu() ([]*models.Menu,error){
	menu := []*models.Menu{}

	if err := c.db.Table("menu").Find(&menu).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return menu, nil
}

func (c MenuController) ReadMenuById(id int) (*models.Menu,error){
	menu := models.Menu{}

	log.Println(id)
	if err := c.db.Table("menu").Where("id = ?",id).Find(&menu).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return &menu, nil
}

func (c MenuController) UpdateMenu(menu *models.Menu) error {
	menu.UpdateAt = time.Now()
	if err := c.db.Table("menu").Updates(menu).Error; err!= nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c MenuController) DeleteMenu(id int) error {

	err := c.DeleteRoleHasmenu(id)
	if err != nil {
		return err
	}

	if err := c.db.Table("menu").Delete(&models.Menu{}, id).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c MenuController) DeleteRoleHasmenu(id int) error {
	list, err := c.ReadAllRoleHasMenuById(id)
	if err != nil {
		return err
	}

	for _, data := range list {
		if err := c.db.Debug().Table("role_has_menu").Where("menu_id = ?",data.MenuId).Delete(&models.RoleHasMenu{}).Error; err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (c MenuController) ReadAllRoleHasMenuById(id int) ([]*models.RoleHasMenu,error){
	list := []*models.RoleHasMenu{}

	if err := c.db.Debug().Table("role_has_menu").Where("menu_id = ?",id).Find(&list).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return list, nil
}
