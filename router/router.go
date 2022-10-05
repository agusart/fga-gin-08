package router

import (
	"fga/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/person/:id", handler.GetPerson(db))
		v1.GET("/persons", handler.GetPersons(db))
		v1.POST("/person", handler.CreatePerson(db))
		v1.PUT("/person/:id", handler.UpdatePerson(db))
		v1.DELETE("/person/:id", handler.DeletePerson(db))
	}

	return router
}
