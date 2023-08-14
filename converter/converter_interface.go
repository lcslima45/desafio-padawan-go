package converter

import "github.com/gin-gonic/gin"

type ConverterService interface {
	ConvertAmount(c *gin.Context, from string, to string, amount string, rate string) string
}
