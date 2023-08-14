package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockConverterService é uma implementação de um serviço de conversão simulado
type mockConverterService struct {
	mock.Mock
}

//Mocks the function ConvertAmount from the service
func (m *mockConverterService) ConvertAmount(c *gin.Context, fromCurrency string, toCurrency string, amount string, rate string) string {
	args := m.Called(c, fromCurrency, toCurrency, amount, rate)
	return args.String(0)
}

//Test the controller ApplyConversion
func TestApplyConversion(t *testing.T) {
	mockConvertService := &mockConverterService{}
	mockConvertService.On("ConvertAmount", mock.Anything, "USD", "EUR", "10.25", "120").Return("123.00")

	// Create a gin-gonic router
	router := gin.Default()

	//Route to serve the handler /exchange/:amount/:from/:to/:rate
	router.GET("/exchange/:amount/:from/:to/:rate", func(c *gin.Context) {
		controller := NewConversionController(mockConvertService)
		controller.ApplyConversion(c)
	})

	//The request that will get a 200 OK Response
	req := httptest.NewRequest(http.MethodGet, "/exchange/10.25/USD/EUR/120", nil)

	// Create a recorder to get the response
	rec := httptest.NewRecorder()

	// Make the http request
	router.ServeHTTP(rec, req)

	// Verify if the response is ok
	expectedJSON := `{"valorConvertido":"123.00","simboloMoeda":"EUR"}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expectedJSON, rec.Body.String())

	mockConvertService.AssertExpectations(t)
}
