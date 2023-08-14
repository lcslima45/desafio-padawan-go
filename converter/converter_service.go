package converter

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lcslima45/desafio-padawan-go/repository"
	"github.com/shopspring/decimal"
)

//converterService -> implements the ConverterService interface
type converterService struct {
	repo repository.Repository
}

//Builder of the converterService

func NewConverterService(repo repository.Repository) *converterService {
	return &converterService{
		repo: repo,
	}
}

// ConvertToCents -> Convert money to cent format to avoid errors with float types
func (service *converterService) ConvertFullMoneyToCents(money string) (int64, error) {
	//remove , to .
	if strings.Contains(money, ",") {
		money = strings.ReplaceAll(money, ",", ".")
	}
	//transform string into money type decimal.Decimal
	moneyInCents, err := decimal.NewFromString(money)
	//rounding of the decimal part of the number
	moneyInCents = moneyInCents.Round(2)
	//transform to the cents type
	moneyInCents = moneyInCents.Mul(decimal.NewFromInt(100))
	return moneyInCents.IntPart(), err
}

// ConvertCentsToFullMoney -> Convert the quantity of money on input to a formatted string $0.00
func (service *converterService) ConvertCentsToFullMoney(moneyToCents int64) string {
	fullMoneyDecimal := decimal.NewFromInt(moneyToCents).Div(decimal.NewFromInt(100))
	fullMoneyFormattedToString := fullMoneyDecimal.StringFixed(2)
	log.Println(fullMoneyFormattedToString)
	return fullMoneyFormattedToString
}

// ConvertAmount -> Convert amount based on the input rate the result is saved on database repository
// we calculate the conversion and save the values in cents format to avoid errors with float types
func (service *converterService) ConvertAmount(c *gin.Context, from string, to string, amount string, rate string) string {
	var convertedValueInCents int64
	//Convert amount to cents
	amountToCents, err := service.ConvertFullMoneyToCents(amount)
	log.Println("AmountToCents:", amountToCents)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Invalid amount format",
			},
		)
		return ""
	}
	//Convert rate to cents
	rateToCents, err := service.ConvertFullMoneyToCents(rate)
	log.Println("RateToCents:", rateToCents)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Invalid rate format",
			},
		)
		return ""
	}
	//Calculate conversion but divide by 100 to deal with conversion to cents
	convertedValueInCents = (amountToCents * rateToCents) / 100
	//Registry on the repository
	ok := service.repo.Registry(c, from, to, amountToCents, rateToCents, convertedValueInCents)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to registry conversion in the database",
		})
		return ""
	}
	//Return string of the conversion in $0.00 style
	return service.ConvertCentsToFullMoney(convertedValueInCents)
}
