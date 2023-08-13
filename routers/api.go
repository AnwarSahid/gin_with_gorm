package routers

import (
	"gin_with_gorm/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/getuser", controllers.GetUser)
	// router.GET("/createuser", controllers.GetUser)

	return router
}
