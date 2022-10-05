package main

import (
	"fga/database"
	"fga/docs"
	"fga/router"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	host     = "fga-db"
	port     = 5432
	user     = "user1"
	password = "password"
	dbname   = "fga"
)

func main() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	database.StartDB()
	db := database.GetDB()

	router := router.StartRouter(db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")

}
