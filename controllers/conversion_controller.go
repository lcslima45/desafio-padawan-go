package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lcslima45/desafio-padawan-go/converter"
)

//conversionController -> Implements the ConversionController interface
type conversionController struct {
	convert converter.ConverterService
}

//NewConversionController -> Builds the conversionController
func NewConversionController(convert converter.ConverterService) *conversionController {
	return &conversionController{
		convert: convert,
	}
}

//isValidSimbol -> Testes if the http request is with the right cases for conversion
func (controller *conversionController) isValidSimbol(simbol string) bool {
	supportedCurrencies := map[string]bool{
		"BTC": true,
		"USD": true,
		"EUR": true,
		"BRL": true,
	}

	return supportedCurrencies[simbol]
}

//ApplyConversion - Handles the request for conversion and returns the JSON with conversion if everything is ok
func (controller *conversionController) ApplyConversion(c *gin.Context) {
	log.Println("[conversionController] Using ControllerConversion on ApplyConversion func")
	amount := c.Param("amount")
	from := strings.ToUpper(c.Param("from")) //The simbol is saved in uppercase
	log.Println("FROM", from)
	if !controller.isValidSimbol(from) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Wrong simbol %s in from query", from),
		})
		return
	}
	to := strings.ToUpper(c.Param("to"))
	if !controller.isValidSimbol(to) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Wrong simbol %s in to query", to),
		})
		return
	}
	rate := c.Param("rate")

	conversion := controller.convert.ConvertAmount(c, from, to, amount, rate)
	simbol := to
	if conversion != "" {
		c.JSON(http.StatusOK, gin.H{
			"valorConvertido": conversion,
			"simboloMoeda":    simbol,
		})
	}
}
