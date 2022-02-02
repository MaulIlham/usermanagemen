package repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"usermanagement/models"
)

type RoleController struct {
	db *gorm.DB
}

func RoleNewController(db *gorm.DB) *RoleController {
	return &RoleController{db}
}

func (c RoleController) InsertRole(newRole *models.Role) error {
	newRole.CreatedAt = time.Now()
	newRole.UpdateAt = time.Now()
	if err := c.db.Table("role").Save(newRole).Error; err!= nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c RoleController) ReadAllRole() ([]*models.Role,error){
	role := []*models.Role{}

	if err := c.db.Table("role").Find(&role).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return role, nil
}

func (c RoleController) ReadRoleById(id int) (*models.Role,error){
	role := models.Role{}

	if err := c.db.Table("role").Where("id = ?",id).Find(&role).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return &role, nil
}

func (c RoleController) UpdateRole(role *models.Role) error {
	role.UpdateAt = time.Now()
	if err := c.db.Table("role").Updates(role).Error; err!= nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c RoleController) DeleteRole(id int) error {
	if err := c.db.Table("role").Delete(&models.Role{}, id).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c RoleController) InsertRoleService(services []*models.Service, idRole int) error {

	for _, service := range services {

		new := models.RoleHasService{
			idRole,
			service.Id,
		}

		if err := c.db.Table("role_has_service").Save(&new).Error; err!= nil {
			log.Println(err)
			return err
		}

	}

	return nil
}

func (c RoleController) InsertRoleMenu(listMenu []*models.Menu, idRole int) error {

	for _, menu := range listMenu {

		new := models.RoleHasMenu{
			idRole,
			menu.Id,
		}

		if err := c.db.Table("role_has_menu").Save(&new).Error; err!= nil {
			log.Println(err)
			return err
		}

	}

	return nil
}

func (c RoleController) InsertRoleUser(listUser []*models.User, idRole int) error {

	for _, user := range listUser {

		new := models.RoleHasUser{
			idRole,
			user.Id,
		}

		if err := c.db.Table("role_has_user").Save(&new).Error; err!= nil {
			log.Println(err)
			return err
		}

	}

	return nil
}