package repository

import (
	"github.com/jinzhu/gorm"
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
	newUser.CreatedAt = time.Now()
	newUser.UpdateAt = time.Now()
	if err := c.db.Table("user").Save(newUser).Error; err!= nil {
		log.Println(err)
		return err
	}

	err := c.InsertUserHasRole(newUser.Role, newUser.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c UserController) ReadAllUser() ([]*models.User,error){
	users := []*models.User{}

	if err := c.db.Table("user").Find(&users).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return users, nil
}

func (c UserController) ReadUserById(id int) (*models.User,error){
	user := models.User{}
	listRole := []*models.Role{}
	newRole := RoleNewController(c.db)

	if err := c.db.Table("user").Where("id = ?",id).Find(&user).Error; err!= nil {
		log.Println(err)

		return nil,err
	}

	list, err := c.ReadAllUserHasRoleById(id)
	if err != nil {
		return nil, err
	}

	for _, data := range list {
		role, _ := newRole.ReadRoleById(data.RoleId)
		log.Println(role)
		listRole = append(listRole, role)
	}

	user.Role = listRole


	return &user, nil
}

func (c UserController) UpdateUser(user *models.User) error {
	user.UpdateAt = time.Now()

	err := c.DeleteUserHasRole(user.Id)
	if err != nil {
		return err
	}

	if err := c.db.Table("user").Updates(&user).Error; err!= nil {
		log.Println(err)
		return err
	}

	err = c.InsertUserHasRole(user.Role, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c UserController) DeleteUser(id int) error {

	err := c.DeleteUserHasRole(id)
	if err != nil {
		return err
	}

	if err := c.db.Table("user").Delete(&models.User{}, id).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c UserController) ReadAllUserHasRoleById(id int) ([]*models.RoleHasUser,error){
	list := []*models.RoleHasUser{}
	log.Println("asdadsasdasdse")

	if err := c.db.Debug().Table("role_has_user").Where("user_id = ?",id).Find(&list).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return list, nil
}

func (c UserController) InsertUserHasRole(listRole []*models.Role, id int) error {

	for _, role := range listRole {

		new := models.RoleHasUser{
			role.Id,
			id,
		}

		if err := c.db.Table("role_has_user").Save(&new).Error; err!= nil {
			log.Println(err)
			return err
		}

	}

	return nil
}

func (c UserController) DeleteUserHasRole(id int) error {
	list, err := c.ReadAllUserHasRoleById(id)
	if err != nil {
		return err
	}

	for _, data := range list {
		if err := c.db.Debug().Table("role_has_user").Where("user_id = ?",data.UserId).Delete(&models.RoleHasUser{}).Error; err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}


