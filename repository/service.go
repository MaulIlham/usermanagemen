package repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"usermanagement/models"
)

type ServiceController struct {
	db *gorm.DB
}

func ServiceNewController(db *gorm.DB) *ServiceController {
	return &ServiceController{db}
}

func (c ServiceController) InsertService(newService *models.Service) error {
	newService.CreateAt = time.Now()
	newService.UpdateAt = time.Now()
	if err := c.db.Table("service").Save(newService).Error; err!= nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c ServiceController) ReadAllService() ([]*models.Service,error){
	services := []*models.Service{}

	if err := c.db.Table("service").Find(&services).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return services, nil
}

func (c ServiceController) ReadServiceById(id int) (*models.Service,error){
	service := models.Service{}

	if err := c.db.Table("service").Where("id = ?",id).Find(service).Error; err!= nil {
		log.Println(err)
		return nil,err
	}

	return &service, nil
}

func (c ServiceController) UpdateService(service *models.Service) error {
	service.UpdateAt = time.Now()
	if err := c.db.Table("service").Updates(service).Error; err!= nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c ServiceController) DeleteService(id int) error {
	if err := c.db.Table("service").Delete(&models.Service{}, id).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

