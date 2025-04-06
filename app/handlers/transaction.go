package handlers

import (
	"net/http"
	"wxcloudrun-golang/app/handlers/request"
	"wxcloudrun-golang/app/handlers/response"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func SellerTxSubmit(c *gin.Context) {
	glog.Info("################## Seller Transaction Submit ##################")
	var reqSeller request.ReqTransaction
	if err := c.ShouldBind(&reqSeller); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}
	glog.Info("Seller:")
	glog.Info(reqSeller.UserID)
	err := service.SellerTxSubmit(&reqSeller)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("seller transaction submit successful")
	response.MakeSuccess(c, http.StatusOK, "successfully submit seller transaction!")
	return
}

func BuyerTxSubmit(c *gin.Context) {
	glog.Info("################## Buyer Transaction Submit ##################")
	var reqBuyer request.ReqTransaction
	if err := c.ShouldBind(&reqBuyer); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}
	glog.Info("Buyer:")
	glog.Info(reqBuyer.UserID)
	err := service.BuyerTxSubmit(&reqBuyer)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("buyer transaction submit successful")
	response.MakeSuccess(c, http.StatusOK, "successfully submit buyer transaction!")
	return
}

func SellerTxPublish(c *gin.Context) {
	glog.Info("################## Publish Seller Transaction ##################")
	res, err := service.GetSellerTx()
	if err != nil {
		glog.Errorln("publish seller transaction error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("publish seller transaction successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}

func BuyerTxPublish(c *gin.Context) {
	glog.Info("################## Publish Buyer Transaction ##################")
	res, err := service.GetBuyerTx()
	if err != nil {
		glog.Errorln("publish buyer transaction error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("publish buyer transaction successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
