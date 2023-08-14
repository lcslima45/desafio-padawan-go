package models

import "time"

type ConversionModel struct {
	ID               uint      `gorm:"primaryKey"`
	FromSimbol       string    `gorm:"not null"`
	ToSimbol         string    `gorm:"not null"`
	Amount           int64     `gorm:"not null"`
	Rate             int64     `gorm:"not null"`
	AmountConverted  int64     `gorm:"not null"`
	HourOfConversion time.Time `gorm:"column:hourOfConversion;not null"`
}

func NewConversionModel(fromSimbol string, toSimbol string, amount int64, rate int64, amountConverted int64) *ConversionModel {
	return &ConversionModel{
		FromSimbol:       fromSimbol,
		ToSimbol:         toSimbol,
		Amount:           amount,
		Rate:             rate,
		AmountConverted:  amountConverted,
		HourOfConversion: time.Now(),
	}
}
