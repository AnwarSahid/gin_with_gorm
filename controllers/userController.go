package controllers

import (
	"gin_with_gorm/database"
	"gin_with_gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	db := database.GetDB()
	if db == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database connection is nil",
		})
		return
	}

	var users []models.User
	res := db.Find(&users)
	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": res.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       users,
	})
}
