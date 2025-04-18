package service

import (
	"log"
	"wxcloudrun-golang/app/handlers/request"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

func MarketSubmit(m *request.ReqMarket, s string) error {
	cli := db.Get()
	if s == "CEA" {
		ceaMarket := model.CEAMarket{
			Date:         m.Date,
			LowerPrice:   m.LowerPrice,
			HigherPrice:  m.HigherPrice,
			ClosingPrice: m.ClosingPrice,
		}
		err := cli.Create(&ceaMarket).Error
		if err != nil {
			log.Printf("[ERROR] Submit cea market error: %v", err)
			return err
		}
	} else {
		ccerMarket := model.CCERMarket{
			Date:         m.Date,
			LowerPrice:   m.LowerPrice,
			HigherPrice:  m.HigherPrice,
			ClosingPrice: m.ClosingPrice,
		}
		err := cli.Create(&ccerMarket).Error
		if err != nil {
			log.Printf("[ERROR] Submit ccer market error: %v", err)
			return err
		}
	}
	return nil
}
