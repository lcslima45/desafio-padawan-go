package controllers

import "github.com/gin-gonic/gin"

type ConversionController interface {
	ApplyConversion(c *gin.Context)
}
