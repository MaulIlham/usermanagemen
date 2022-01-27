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

func (h Handler) SaveMenu(c *gin.Context) {
	menu := models.Menu{}

	newController := repository.MenuNewController(h.DB.Conn)

	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, menu); err != nil {
		log.Fatal("could not parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body: %s", err.Error())})
		return
	}

	err := newController.InsertMenu(&menu)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Insert Data Success",
		Data: menu,
	})
}

func (h Handler) ReadAllMenu(c *gin.Context) {
	newController := repository.MenuNewController(h.DB.Conn)

	listMenu, err := newController.ReadAllMenu()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Read All Data Success",
		Data: listMenu,
	})
}

func (h Handler) ReadDataMenuById(c *gin.Context) {
	var param string

	newController := repository.MenuNewController(h.DB.Conn)

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

	menu, err := newController.ReadMenuById(id)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: fmt.Sprintf("Read Data With ID : %d Success",id),
		Data: menu,
	})
}

func (h Handler) DeleteDataMenu(c *gin.Context) {
	var param string

	newController := repository.MenuNewController(h.DB.Conn)

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

	err = newController.DeleteMenu(id)
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

func (h Handler) UpdateMenu(c *gin.Context) {
	menu := models.Menu{}

	newController := repository.MenuNewController(h.DB.Conn)

	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, menu); err != nil {
		log.Fatal("could not parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request body: %s", err.Error())})
		return
	}

	err := newController.UpdateMenu(&menu)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(err.Error())})
		return
	}

	c.JSON(http.StatusBadRequest, models.Logger{
		Status: "Ok",
		Message: "Update Data Success",
		Data: menu,
	})
}