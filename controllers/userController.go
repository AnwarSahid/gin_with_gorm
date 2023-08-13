package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"datas ":     "anwar Sahid",
	})

}
