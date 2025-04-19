package handlers

import (
	"log"
	"net/http"
	"wxcloudrun-golang/app/handlers/request"
	"wxcloudrun-golang/app/handlers/response"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func MarketSubmit(c *gin.Context) {
	log.Println("################## CEA Market Submit ##################")
	var reqMarket request.ReqMarket
	if err := c.ShouldBind(&reqMarket); err != nil {
		log.Printf("[ERROR] : %v", err)
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}

	err := service.MarketSubmit(&reqMarket, reqMarket.Product)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("market submit successful")
	response.MakeSuccess(c, http.StatusOK, "successfully market submit!")
	return
}

// func CCERMarketSubmit(c *gin.Context) {
// 	log.Println("################## CCER Market Submit ##################")
// 	var reqMarket request.ReqMarket
// 	if err := c.ShouldBind(&reqMarket); err != nil {
// 		log.Printf("[ERROR] : %v", err)
// 		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
// 		return
// 	}
// 	err := service.MarketSubmit(&reqMarket, "CCER")
// 	if err != nil {
// 		response.MakeFail(c, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	log.Println("CCER market submit successful")
// 	response.MakeSuccess(c, http.StatusOK, "successfully CCER market submit!")
// 	return
// }

func GetCCERMarket(c *gin.Context) {
	var ms []model.CCERMarket
	cli := db.Get()
	err := cli.Find(&ms).Error
	if err != nil {
		log.Printf("[ERROR] get all ccer market error!")
		response.MakeFail(c, http.StatusBadRequest, "get all ccer market error")
		return
	}
	log.Println("get ccer market successful")
	response.MakeSuccess(c, http.StatusOK, ms)
	return
}

func GetCEAMarket(c *gin.Context) {
	var ms []model.CEAMarket
	cli := db.Get()
	err := cli.Find(&ms).Error
	if err != nil {
		log.Printf("[ERROR] get all ccer market error!")
		response.MakeFail(c, http.StatusBadRequest, "get all ccer market error")
		return
	}
	log.Println("get ccer market successful")
	response.MakeSuccess(c, http.StatusOK, ms)
	return
}
