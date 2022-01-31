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

func (h Handler) SaveRole(c *gin.Context) {
	role := models.Role{}

	newRole := repository.RoleNewController(h.DB.Conn)

	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, role); err != nil {
		log.Fatal("could not parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body: %s", err.Error())})
		return
	}

	err := newRole.InsertRole(&role)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusCreated, models.Logger{
		Status: "Ok",
		Message: "Insert Data Success",
		Data: role,
	})
}

func (h Handler) ReadAllRole(c *gin.Context) {
	newRole := repository.RoleNewController(h.DB.Conn)

	role, err := newRole.ReadAllRole()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusOK, models.Logger{
		Status: "Ok",
		Message: "Read All Data Success",
		Data: role,
	})
}

func (h Handler) ReadDataRoleById(c *gin.Context) {
	var param string

	newRole := repository.RoleNewController(h.DB.Conn)

	if param, _ = c.GetQuery("q"); param == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Id Param Request"})
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	role, err := newRole.ReadRoleById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusOK, models.Logger{
		Status: "Ok",
		Message: fmt.Sprintf("Read Data With ID : %d Success",id),
		Data: role,
	})
}

func (h Handler) DeleteDataRole(c *gin.Context) {
	var param string

	newRole := repository.RoleNewController(h.DB.Conn)

	if param, _ = c.GetQuery("q"); param == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Id Param Request"})
		return
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	err = newRole.DeleteRole(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusOK, models.Logger{
		Status: "Ok",
		Message: fmt.Sprintf("Delete Data With ID : %d Success",id),
	})
}

func (h Handler) UpdateRole(c *gin.Context) {
	role := models.Role{}

	newController := repository.RoleNewController(h.DB.Conn)

	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, role); err != nil {
		log.Println("could not parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body: %s", err.Error())})
		return
	}

	err := newController.UpdateRole(&role)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusOK, models.Logger{
		Status: "Ok",
		Message: "Update Data Success",
		Data: role,
	})
}