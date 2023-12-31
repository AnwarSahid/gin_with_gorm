package routers

import (
	"gin_with_gorm/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/getuser", controllers.GetUser)
	router.GET("/getuser/:userID", controllers.GetUserById)
	router.POST("/createuser", controllers.CreateUser)
	router.PUT("/updateuser/:userID", controllers.UpdateUser)
	router.DELETE("/deleteuser/:userID", controllers.DeleteUser)

	return router
}
