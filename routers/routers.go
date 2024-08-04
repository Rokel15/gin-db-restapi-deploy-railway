package routers

import (
	"golang-gin-db-restapi-lokal/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/chickens", controllers.GetAllAfkeur)
	router.POST("/input-chick", controllers.InserAfkeur)
	router.PUT("/update-chick/:id", controllers.UpdateAfkeur)
	router.DELETE("/delete-chick/:id", controllers.DeleteAfkeur)

	return router
}
