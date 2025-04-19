package model

import "gorm.io/gorm"

type CEAMarket struct {
	gorm.Model
	Date         string `gorm:"column:date;" json:"date"`
	LowerPrice   string `gorm:"column:lowerPrice" json:"lowerPrice"`
	HigherPrice  string `gorm:"column:higherPrice" json:"higherPrice"`
	ClosingPrice string `gorm:"column:closingPrice" json:"closingPrice"`
}

type CCERMarket struct {
	gorm.Model
	Date         string `gorm:"column:date;" json:"date"`
	ClosingPrice string `gorm:"column:closingPrice" json:"closingPrice"`
}
