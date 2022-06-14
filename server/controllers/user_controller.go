package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luccasalves/apigo-sample/database"
	"github.com/luccasalves/apigo-sample/models"
)

func GetAll(ctx *gin.Context) {
	db := database.GetDb()

	var users []models.User

	err := db.Find(&users).Error

	if err != nil {
		// ctx.JSON(200, gin.H{"oi": "oi"})
		ctx.JSON(400, gin.H{"error": "cannot list users: " + err.Error()})
		return
	}

	ctx.JSON(200, users)
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.GetDb()

	var user models.User
	err = db.First(&user, newid).Error
	if err != nil {
		ctx.JSON(400, gin.H{"error": "cannot find user: " + err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func CreateUser(ctx *gin.Context) {
	db := database.GetDb()

	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot find Json: " + err.Error(),
		})
		return
	}

	err = db.Create(&user).Error

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}
	ctx.JSON(201, user)
}
