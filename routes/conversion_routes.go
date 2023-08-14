package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lcslima45/desafio-padawan-go/controllers"
)

const (
	PORT = ":8000"
)

func ConvertRoute(router *gin.Engine, controller controllers.ConversionController) {
	router.GET("/exchange/:amount/:from/:to/:rate", controller.ApplyConversion)
}
