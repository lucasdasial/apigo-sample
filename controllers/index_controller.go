package controllers

import "github.com/gin-gonic/gin"

func Greets(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "statusOk"})
}
