package usecase

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"usermanagement/config"
)

type Handler struct {
	DB config.Database
}

func NewHandler(db config.Database) *Handler {
	return &Handler{db}
}

func (h Handler) Requesthandler(group *gin.RouterGroup) {

	//role
	group.GET("/roles", h.ReadAllRole)
	group.GET("/role", h.ReadDataRoleById)
	group.POST("/role", h.SaveRole)
	group.PUT("/role", h.UpdateRole)
	group.DELETE("/role", h.DeleteDataRole)

	//user
	group.GET("/users", h.ReadAllUser)
	group.GET("/user", h.ReadDataUserById)
	group.POST("/user", h.SaveUser)
	group.PUT("/user", h.UpdateUser)
	group.DELETE("/user", h.DeleteDataUser)

	//menu
	group.GET("/menus", h.ReadAllMenu)
	group.GET("/menu", h.ReadDataMenuById)
	group.POST("/menu", h.SaveMenu)
	group.PUT("/menu", h.UpdateMenu)
	group.DELETE("/menu", h.DeleteDataMenu)

	//service
	group.GET("/services", h.ReadAllService)
	group.GET("/service", h.ReadDataServiceById)
	group.POST("/service", h.SaveService)
	group.PUT("/service", h.UpdateService)
	group.DELETE("/service", h.DeleteDataService)

}

func Init() {
	var err error

	db, err := config.InitDB("MYSQL_USER", "MYSQL_PASSWORD", "MYSQL_HOST", "MYSQL_PORT", "MYSQL_SCHEMA")
	if err != nil {
		log.Fatal("Connection failed")
		os.Exit(1)
	}
	log.Fatal("Database connection established")

	handler := NewHandler(db)
	router := gin.Default()
	rg := router.Group("/api/v1")
	handler.Requesthandler(rg)
	router.Run(os.Getenv("PORT"))
}
