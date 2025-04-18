package handlers

import (
	"log"
	"net/http"
	"wxcloudrun-golang/app/handlers/request"
	"wxcloudrun-golang/app/handlers/response"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func SellerTxSubmit(c *gin.Context) {
	log.Println("################## Seller Transaction Submit ##################")
	var reqSeller request.ReqTransaction
	if err := c.ShouldBind(&reqSeller); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}
	log.Println("Seller:")
	log.Println(reqSeller.UserID)
	err := service.SellerTxSubmit(&reqSeller)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("seller transaction submit successful")
	response.MakeSuccess(c, http.StatusOK, "successfully submit seller transaction!")
	return
}

func BuyerTxSubmit(c *gin.Context) {
	log.Println("################## Buyer Transaction Submit ##################")
	var reqBuyer request.ReqTransaction
	if err := c.ShouldBind(&reqBuyer); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}
	log.Println("Buyer:")
	log.Println(reqBuyer.UserID)
	err := service.BuyerTxSubmit(&reqBuyer)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("buyer transaction submit successful")
	response.MakeSuccess(c, http.StatusOK, "successfully submit buyer transaction!")
	return
}

func SellerTxPublish(c *gin.Context) {
	log.Println("################## Publish Seller Transaction ##################")
	res, err := service.GetSellerTx()
	if err != nil {
		glog.Errorln("publish seller transaction error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("publish seller transaction successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func BuyerTxPublish(c *gin.Context) {
	log.Println("################## Publish Buyer Transaction ##################")
	res, err := service.GetBuyerTx()
	if err != nil {
		glog.Errorln("publish buyer transaction error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("publish buyer transaction successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
