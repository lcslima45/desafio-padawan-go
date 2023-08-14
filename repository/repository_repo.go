package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lcslima45/desafio-padawan-go/models"
	gorm "gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) Registry(c *gin.Context, from string, to string, amount int64, rate int64, convertedValue int64) bool {
	conversion := models.NewConversionModel(from, to, amount, rate, convertedValue)
	//log.Println(conversion.ConTime)
	result := repo.db.Select("from_simbol", "to_simbol", "amount", "rate", "amount_converted", "hourOfConversion").Create(conversion)

	if result.Error != nil {
		log.Println("[Registry] Error:", result.Error)
		return false
	}

	return true
}
