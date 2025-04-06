package handlers

import (
	"time"
	"wxcloudrun-golang/app/handlers/request"
	"wxcloudrun-golang/app/handlers/response"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func SemiMonthSubmit(c *gin.Context) {
	glog.Info("################## SemiMonth Quotation Submit ##################")
	var reqQuotation request.ReqQuotation
	if err := c.ShouldBind(&reqQuotation); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}
	glog.Info("semimonth quotation submit user:")
	glog.Info(reqQuotation.UserID)
	err := service.AddSemiMonth(&reqQuotation)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("semimonth quotation submit successful")
	response.MakeSuccess(c, http.StatusOK, "successfully submit semimonth quotation!")
	return
}

func MonthSubmit(c *gin.Context) {
	glog.Info("################## Month Quotation Submit ##################")
	var reqQuotation request.ReqQuotation
	if err := c.ShouldBind(&reqQuotation); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}
	glog.Info("month quotation submit user:")
	glog.Info(reqQuotation.UserID)
	err := service.AddMonth(&reqQuotation)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("month quotation submit successful")
	response.MakeSuccess(c, http.StatusOK, "successfully submit month quotation!")
	return
}

func YearSubmit(c *gin.Context) {
	glog.Info("################## Year Quotation Submit ##################")
	var reqQuotation request.ReqQuotation
	if err := c.ShouldBind(&reqQuotation); err != nil {
		glog.Errorln(err.Error())
		response.MakeFail(c, http.StatusNotAcceptable, err.Error())
		return
	}
	glog.Info("year quotation submit user:")
	glog.Info(reqQuotation.UserID)
	err := service.AddYear(&reqQuotation)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("year quotation submit successful")
	response.MakeSuccess(c, http.StatusOK, "successfully submit year quotation!")
	return
}

func ApproveQuotation(c *gin.Context) {
	glog.Info("################## Approve Quotation ##################")
	var reqApproveQuotation request.ReqApproveQuotation
	if err := c.ShouldBind(&reqApproveQuotation); err != nil {
		glog.Errorln("approve quotation error")
		response.MakeFail(c, http.StatusNotAcceptable, "approve quotation failure!")
		return
	}
	var modelType interface{}
	switch reqApproveQuotation.Type {
	case "semimonth":
		modelType = &model.SemiMonthQuotation{}
	case "month":
		modelType = &model.MonthQuotation{}
	case "year":
		modelType = &model.YearQuotation{}
	default:
		response.MakeFail(c, http.StatusBadRequest, "quotation type error!")
		return
	}
	err := service.ApproveQuotation(reqApproveQuotation.QID, modelType)
	if err != nil {
		glog.Errorln("approve quotation error!")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	response.MakeSuccess(c, http.StatusOK, "approve quotation successfully")

}
func SemiMonthPublish(c *gin.Context) {
	glog.Info("################## Publish SemiMonth Quotation ##################")
	timestr := c.Query("time")
	t, err := time.Parse("2006-01-02 15:04:05", timestr)
	res, err := service.GetApprovedSemimonthQuotations(t)
	if err != nil {
		glog.Errorln("publish semimonth quotation error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("publish semimonth quotation successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
func MonthPublish(c *gin.Context) {
	glog.Info("################## Publish Month Quotation ##################")
	timestr := c.Query("time")
	t, err := time.Parse("2006-01-02 15:04:05", timestr)
	res, err := service.GetApprovedMonthQuotations(t)
	if err != nil {
		glog.Errorln("publish month quotation error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("publish month quotation successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
func YearPublish(c *gin.Context) {
	glog.Info("################## Publish Year Quotation ##################")
	timestr := c.Query("time")
	t, err := time.Parse("2006-01-02 15:04:05", timestr)
	res, err := service.GetApprovedYearQuotations(t)
	if err != nil {
		glog.Errorln("publish year quotation error")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Info("publish year quotation successful")
	response.MakeSuccess(c, http.StatusOK, res)
	return
}
