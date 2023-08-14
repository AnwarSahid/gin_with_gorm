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
		"message":    "data collected successfully",
		"data":       users,
	})
}

func GetUserById(ctx *gin.Context) {
	db := database.GetDB()

	var user models.User
	res := db.Find(&user, ctx.Param("userID"))

	if res == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message ":   "data not found",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "data collected succesfully",
			"datas":   user,
		})
	}
}

func CreateUser(ctx *gin.Context) {
	db := database.GetDB()
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := db.Create(&user)

	if res.Error != nil {
		ctx.JSON(500, gin.H{
			"message": res.Error,
		})
	}
	ctx.JSON(200, &user)
}

func UpdateUser(ctx *gin.Context) {
	db := database.GetDB()
	var user = models.User{}
	err := ctx.ShouldBind(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return // Menghentikan eksekusi jika terjadi kesalahan
	}
	param := ctx.Param("userID")
	updateErr := db.Model(&user).Where("id = ?", param).Updates(models.User{
		Name: user.Name,
	}).Error

	if updateErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": updateErr.Error(),
		})
		return // Menghentikan eksekusi jika terjadi kesalahan dalam meng-update
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}
func DeleteUser(ctx *gin.Context) {
	db := database.GetDB()
	user := models.User{}
	param := ctx.Param("userID")
	err := db.Where("id = ?", param).Delete(&user).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}
