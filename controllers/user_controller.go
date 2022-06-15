package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pgRepository "github.com/luccasalves/apigo-sample/repositories/postgres"
	pgService "github.com/luccasalves/apigo-sample/services/postgres"
)

// func GetAll(ctx *gin.Context) {
// 	db := database.GetDb()

// 	var users []models.User

// 	err := db.Find(&users).Error

// 	if err != nil {
// 		// ctx.JSON(200, gin.H{"oi": "oi"})
// 		ctx.JSON(400, gin.H{"error": "cannot list users: " + err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, users)
// }

// func GetUser(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	newid, err := strconv.Atoi(id)
// 	if err != nil {
// 		ctx.JSON(400, gin.H{
// 			"error": "ID has to be integer",
// 		})
// 		return
// 	}
// 	db := database.GetDb()

// 	var user models.User
// 	err = db.First(&user, newid).Error
// 	if err != nil {
// 		ctx.JSON(400, gin.H{"error": "cannot find user: " + err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, user)
// }

func Create(c *gin.Context) {
	var dto pgService.UserDTO
	repo := &pgRepository.UserRepository{}
	service := pgService.NewUserService(repo)

	result, err := service.CreateUser(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, result)

}

func Show(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	repo := &pgRepository.UserRepository{}
	service := pgService.NewUserService(repo)

	result, err := service.ShowUser(uint(newid))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func ShowAll(c *gin.Context) {
	repo := &pgRepository.UserRepository{}
	service := pgService.NewUserService(repo)

	result, err := service.ShowAllUsers()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}
