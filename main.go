package main

import (
	"fmt"
	"log"

	// "net/http"
	"wxcloudrun-golang/app/handlers"
	"wxcloudrun-golang/app/middlewares"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	r := gin.Default()
	r.Use(middlewares.Cors())
	user := r.Group("/user")
	{
		//注册
		user.POST("register", handlers.Register)
		//登录
		user.POST("login", handlers.Login)
		//报价提交
		// user.POST("semimonth", handlers.SemiMonthSubmit)
		// user.POST("month", handlers.MonthSubmit)
		// user.POST("year", handlers.YearSubmit)
		user.POST("submit", handlers.MultiSubmit)
		//报价公告
		user.GET("semimonthpublish", handlers.SemiMonthPublish)
		user.GET("monthpublish", handlers.MonthPublish)
		user.GET("yearpublish", handlers.YearPublish)
		//可视化数据
		user.GET("monthcea", handlers.GetMonthlyCEAExpectation)
		user.GET("monthccer", handlers.GetMonthlyCCERExpectation)
		user.GET("yearcea", handlers.GetYearlyCEAExpectation)
		user.GET("monthgec", handlers.GetGECMonthlyStatsByType)
		//场外交易提交
		user.POST("sellertxsubmit", handlers.SellerTxSubmit)
		user.POST("buyertxsubmit", handlers.BuyerTxSubmit)
		user.GET("sellertx", handlers.SellerTxPublish)
		user.GET("buyertx", handlers.BuyerTxPublish)
		//申请升级为VIP
		user.POST("applyVip", handlers.ApplyToVip)
	}
	// all := r.Group("/all",middlewares.JwtAuth())
	// {
	// 	all.POST("semimonth",handlers.)
	// 	all.POST("month",handlers.)
	// 	all.POST("year",handlers.)
	// }
	common := r.Group("/common", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		common.GET("count", service.GetCounterHandler)
	}
	admin := r.Group("/admin", middlewares.JwtAuth(), middlewares.CheckPermission())
	{
		// admin.POST("/approveuser", handlers.ApproveUser)
		admin.GET("approvequotation", handlers.ApproveQuotation)
		admin.GET("/upgradeVip", handlers.UpToVipByAdmin)
		admin.GET("/downVip", handlers.DownToCommonByAdmin)
		admin.GET("/userlist", handlers.GetAllUser)
		admin.GET("/approvinguser", handlers.GetAllApprovingUser)
		// admin.GET("approvingsemimonth", handlers.GetApprovingSemiMonthQuotations)
		// admin.GET("approvingmonth", handlers.GetApprovingMonthQuotations)
		// admin.GET("approvingyear", handlers.GetApprovingYearQuotations)
		admin.POST("uploadMarket", handlers.MarketSubmit)
		admin.POST("submitStats", handlers.StatsSubmit)
		admin.POST("submitGECStats", handlers.GECStatsSubmit)
		admin.GET("getCEA", handlers.GetCEAMarket)
		admin.GET("getCCER", handlers.GetCCERMarket)
		admin.GET("getNewMonthAvg", handlers.GetMonthGroupAVG)
		admin.GET("getNewYearAvg", handlers.GetYearGroupAVG)
		admin.GET("getScoreList", handlers.GetCEAMonthScore)
		admin.GET("getMonthQuotation", handlers.AdminGetMonthQuotation)
		admin.GET("getYearQuotation", handlers.AdminGetYearQuotation)
		admin.GET("getCEAMonthExpectation", handlers.GetCEAMonthExpectation)
		admin.GET("getCCERMonthExpectation", handlers.GetCCERMonthExpectation)
		admin.GET("getCEAYearExpectation", handlers.GetCEAYearExpectation)
		admin.GET("getGECMonthExpectation", handlers.GetGECMonthExpectation)
	}
	r.GET("/", service.IndexHandler)
	r.GET("/api/count", service.GetCounterHandler)
	r.POST("/api/count", service.PostCounterHandler)
	r.GET("/api/user", service.GetUserInfo)
	log.Fatal(r.Run(":80"))
}
