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
	if err := c.db.Debug().Table("role").Save(newRole).Error; err!= nil {
		log.Println(err)
		return err
	}

	err := c.InsertRoleMenu(newRole.Menu,newRole.Id)
	if err != nil {
		return err
	}

	err = c.InsertRoleService(newRole.Service, newRole.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c RoleController) ReadAllRole() ([]*models.Role,error){
	role := []*models.Role{}

	if err := c.db.Debug().Table("role").Find(&role).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return role, nil
}

func (c RoleController) ReadRoleById(id int) (*models.Role,error){
	role := models.Role{}

	if err := c.db.Debug().Table("role").Where("id = ?",id).Find(&role).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return &role, nil
}

func (c RoleController) UpdateRole(role *models.Role) error {
	role.UpdateAt = time.Now()

	//deleting helping table
	err := c.DeleteRoleHasService()
	if err != nil {
		return err
	}

	err = c.DeleteRoleHasmenu()
	if err != nil {
		return err
	}

	// update data role
	updateRole := models.Role{
		Id: role.Id,
		Name: role.Name,
		CreatedAt: role.CreatedAt,
		UpdateAt: role.UpdateAt,
		CreatedBy: role.CreatedBy,
		UpdateBy: role.UpdateBy,
	}
	if err := c.db.Debug().Table("role").Updates(&updateRole).Error; err!= nil {
		log.Println(err)
		return err
	}

	// save new helping table
	err = c.InsertRoleService(role.Service, role.Id)
	if err != nil {
		return err
	}

	err = c.InsertRoleMenu(role.Menu, role.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c RoleController) DeleteRole(id int) error {

	err := c.DeleteRoleHasmenu()
	if err != nil {
		return err
	}

	err = c.DeleteRoleHasService()
	if err != nil {
		return err
	}

	if err := c.db.Debug().Table("role").Delete(&models.Role{}, id).Error; err != nil {
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

func (c RoleController) ReadAllRoleService() ([]*models.RoleHasService,error){
	list := []*models.RoleHasService{}

	if err := c.db.Debug().Table("role_has_service").Find(&list).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return list, nil
}

func (c RoleController) ReadAllRoleMenu() ([]*models.RoleHasMenu,error){
	list := []*models.RoleHasMenu{}

	if err := c.db.Debug().Table("role_has_menu").Find(&list).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return list, nil
}

func (c RoleController) DeleteRoleHasService() error {
	list, err := c.ReadAllRoleService()
	if err != nil {
		return err
	}

	for _, data := range list {
		if err := c.db.Debug().Table("role_has_service").Where("role_id = ?",data.RoleId).Delete(&models.RoleHasService{}).Error; err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (c RoleController) DeleteRoleHasmenu() error {
	list, err := c.ReadAllRoleMenu()
	if err != nil {
		return err
	}

	log.Println(list)

	for _, data := range list {
		if err := c.db.Debug().Table("role_has_menu").Where("role_id = ?",data.RoleId).Delete(&models.RoleHasMenu{}).Error; err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}