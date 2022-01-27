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

func (h Handler) SaveUser(c *gin.Context) {
	user := models.User{}

	newController := repository.UserNewController(h.DB.Conn)

	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, user); err != nil {
		log.Println("could not parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body: %s", err.Error())})
		return
	}

	err := newController.InsertUser(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Insert Data Success",
		Data: user,
	})
}

func (h Handler) ReadAllUser(c *gin.Context) {
	newController := repository.UserNewController(h.DB.Conn)

	listUser, err := newController.ReadAllUser()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Read All Data Success",
		Data: listUser,
	})
}

func (h Handler) ReadDataUserById(c *gin.Context) {
	var param string

	newController := repository.UserNewController(h.DB.Conn)

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

	user, err := newController.ReadUserById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: fmt.Sprintf("Read Data With ID : %d Success",id),
		Data: user,
	})
}

func (h Handler) DeleteDataUser(c *gin.Context) {
	var param string

	newController := repository.UserNewController(h.DB.Conn)

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

	err = newController.DeleteUser(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: fmt.Sprintf("Delete Data With ID : %d Success",id),
	})
}

func (h Handler) UpdateUser(c *gin.Context) {
	user := models.User{}

	newController := repository.UserNewController(h.DB.Conn)

	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, user); err != nil {
		log.Println("could not parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body: %s", err.Error())})
		return
	}

	err := newController.UpdateUser(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Update Data Success",
		Data: user,
	})
}
