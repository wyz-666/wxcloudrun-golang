package handlers

import (
	"log"
	"net/http"
	"time"
	"wxcloudrun-golang/app/handlers/response"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func GetMonthlyCCERStats(c *gin.Context) {
	log.Println("################## GetMonthlyCCERStats ##################")
	result, err := service.GetMonthlyCCERStats()

	if err != nil {
		glog.Errorln("get monthly ccer stats error!")
		response.MakeFail(c, http.StatusBadRequest, "get  monthly ccer stats error")
		return
	}
	log.Println("get monthly ccer stats successfully")
	response.MakeSuccess(c, http.StatusOK, result)
	return
}

func GetMonthlyCEAStats(c *gin.Context) {
	log.Println("################## GetMonthlyCEAStats ##################")
	result, err := service.GetMonthlyCEAStats()

	if err != nil {
		glog.Errorln("get monthly cea stats error!")
		response.MakeFail(c, http.StatusBadRequest, "get  monthly cea stats error")
		return
	}
	log.Println("get monthly cea stats successfully")
	response.MakeSuccess(c, http.StatusOK, result)
	return
}

func GetGECMonthlyStatsByType(c *gin.Context) {
	log.Println("################## GetMonthlyCCERStats ##################")
	result, err := service.GetGECMonthlyStatsByType()

	if err != nil {
		glog.Errorln("get monthly gec stats error!")
		response.MakeFail(c, http.StatusBadRequest, "get  monthly gec stats error")
		return
	}
	log.Println("get monthly gec stats successfully")
	response.MakeSuccess(c, http.StatusOK, result)
	return
}

func GetCEAGroupAVG(c *gin.Context) {
	log.Println("################## Get CEA Group AVG ##################")
	nowTimeStr := c.Query("nowTime")
	t, err := time.Parse("2006-01-02 15:04:05", nowTimeStr)
	if err != nil {
		c.JSON(400, gin.H{"msg": "nowTime 格式错误", "error": err.Error()})
		return
	}
	result, err := service.GetCEAStatsNextMonth(t)
	if err != nil {
		log.Printf("Get CEA Group AVG error!")
		response.MakeFail(c, http.StatusBadRequest, "Get CEA Group AVG error")
		return
	}
	log.Println("Get CEA Group AVG successfully")
	response.MakeSuccess(c, http.StatusOK, result)
	return
}
