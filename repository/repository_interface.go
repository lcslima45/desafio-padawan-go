package repository

import "github.com/gin-gonic/gin"

type Repository interface {
	Registry(c *gin.Context, from string, to string, amount int64, rate int64, convertedValue int64) bool
}
