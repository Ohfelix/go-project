package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST Api v1 Opening Endpoint",
	})
}
