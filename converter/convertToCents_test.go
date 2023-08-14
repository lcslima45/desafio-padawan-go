package converter

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Registry(c *gin.Context, from string, to string, amount int64, rate int64, convertedValue int64) bool {
	args := m.Called(c, from, to, amount, rate, convertedValue)
	return args.Bool(0)
}

//TestConvertToCents -> Test the cases for conversions to cents
func TestConvertToCents(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"100", 10000},
		{"100.00", 10000},
		{"100,00", 10000},
		{"100.0", 10000},
		{"100.000", 10000},
		{"100,0", 10000},
		{"100,000", 10000},
		{"100.23", 10023},
		{"100,23", 10023},
		{"100.2", 10020},
		{"100,2", 10020},
		{"100.234", 10023}, // Should round to 100.23
		{"100,234", 10023}, // Should round to 100.23
	}

	converterService := NewConverterService(nil) // Passing nil as the repository for the test

	for _, test := range tests {
		result, err := converterService.ConvertFullMoneyToCents(test.input)
		if err != nil {
			t.Errorf("Error converting %s: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("For input %s, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}

//TestConvertCentsToFullMoney -> Test of the conversion to money string with some cases
func TestConvertCentsToFullMoney(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{10000, "100.00"},
		{10023, "100.23"},
		{12323, "123.23"},
	}

	converterService := NewConverterService(nil) // Passing nil as the repository for the test

	for _, test := range tests {
		result := converterService.ConvertCentsToFullMoney(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}

//TestConvertAmount -> Test the main functionality of the Converter
func TestConvertAmount(t *testing.T) {
	mockRepo := &mockRepository{}
	service := NewConverterService(mockRepo)

	c := &gin.Context{}

	mockRepo.On("Registry", c, "USD", "EUR", int64(1025), int64(12000), int64(123000)).Return(true)

	result := service.ConvertAmount(c, "USD", "EUR", "10.25", "120")

	assert.Equal(t, "1230.00", result)

	mockRepo.AssertExpectations(t)
}
