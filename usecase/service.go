package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"usermanagement/models"
	"usermanagement/repository"
)

func (h Handler) SaveService(c *gin.Context) {
	service := models.Service{}

	newController := repository.ServiceNewController(h.DB.Conn)

	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, service); err != nil {
		log.Fatal("could not parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body: %s", err.Error())})
		return
	}

	err := newController.InsertService(&service)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Insert Data Success",
		Data: service,
	})
}

func (h Handler) UpdateService(c *gin.Context) {
	service := models.Service{}

	newController := repository.ServiceNewController(h.DB.Conn)

	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, service); err != nil {
		log.Fatal("could not parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body: %s", err.Error())})
		return
	}

	err := newController.UpdateService(&service)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Update Data Success",
		Data: service,
	})
}

func (h Handler) ReadAllService(c *gin.Context) {
	newController := repository.ServiceNewController(h.DB.Conn)

	listUser, err := newController.ReadAllService()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Read All Data Success",
		Data: listUser,
	})
}

func (h Handler) ReadDataServiceById(c *gin.Context) {
	var param string

	newController := repository.ServiceNewController(h.DB.Conn)

	if param, _ = c.GetQuery("q"); param == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Id Param Request"})
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	service, err := newController.ReadServiceById(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: fmt.Sprintf("Read Data With ID : %d Success",id),
		Data: service,
	})
}

func (h Handler) DeleteDataService(c *gin.Context) {
	var param string

	newController := repository.ServiceNewController(h.DB.Conn)

	if param, _ = c.GetQuery("q"); param == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Id Param Request"})
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	err = newController.DeleteService(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: fmt.Sprintf("Delete Data With ID : %d Success",id),
	})
}
