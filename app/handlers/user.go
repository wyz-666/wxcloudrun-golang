package handlers

import (
	"net/http"
	"wxcloudrun-golang/app/handlers/request"
	"wxcloudrun-golang/app/handlers/response"
	"wxcloudrun-golang/service"
	"wxcloudrun-golang/utils/crypto"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func Register(c *gin.Context) {

	glog.Info("################## User Register ##################")
	var reqUser request.ReqUser
	if err := c.ShouldBind(&reqUser); err != nil || !checkRegisterParams(&reqUser) {
		glog.Errorln("User registration error")
		response.MakeFail(c, http.StatusNotAcceptable, "user registration failure!")
		return
	}
	glog.Info("request user parameters:")
	glog.Info(reqUser)
	err := service.AddUser(&reqUser)
	if err != nil {
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	glog.Info("add new user successful!")
	response.MakeSuccess(c, http.StatusOK, "successfully register the user!")
	return
}

func checkRegisterParams(reqUser *request.ReqUser) bool {
	if reqUser.Name == "" || reqUser.Account == "" || reqUser.Type == 0 {
		glog.Errorln("Missing user registration parameters")
		return false
	}
	ps := reqUser.Password
	if len(ps) < 9 {
		glog.Errorln("password len is < 9")
		return false
	}
	return true
}

func Login(c *gin.Context) {
	glog.Info("################## User Login ##################")
	var reqLogin request.ReqLogin
	if err := c.ShouldBind(&reqLogin); err != nil || !checkLoginParams(&reqLogin) {
		response.MakeFail(c, http.StatusBadRequest, "login parameters error!")
		return
	}
	reqPwdHash := crypto.CalculateSHA256(reqLogin.Password, "FDUCPIF")
	userID, pwdHash, userType, approved, err := service.QueryUserInfo(reqLogin.Account)
	if err != nil {
		glog.Errorln("query fsims password hash error!")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}

	if userType == 2 && !approved {
		glog.Errorln("该用户资格审批尚未通过!")
		response.MakeFail(c, http.StatusBadRequest, "该用户资格审批尚未通过!")
		return
	}

	// check password
	if reqPwdHash != pwdHash {
		glog.Errorln("password incorrect!")
		response.MakeFail(c, http.StatusBadRequest, "password incorrecrt!")
		return
	}

	// create jwt token with uuid
	token, err := service.CreateJwtToken(userID, userType)
	if err != nil {
		glog.Errorln("create jwt token error!")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	glog.Infoln(reqLogin.Account, " login successful!")
	glog.Infoln("token:", token)

	resLogin := response.ResLogin{
		UserID: userID,
		Token:  token,
	}
	response.MakeSuccess(c, http.StatusOK, resLogin)
}

func ApproveUser(c *gin.Context) {
	glog.Info("################## Approve VIP User ##################")
	var reqApprove request.ReqApproveUser
	if err := c.ShouldBind(&reqApprove); err != nil {
		glog.Errorln("approve user error")
		response.MakeFail(c, http.StatusNotAcceptable, "approve user failure!")
		return
	}
	err := service.ApproveUser(reqApprove.UserID)
	if err != nil {
		glog.Errorln("approve vip user error!")
		response.MakeFail(c, http.StatusBadRequest, err.Error())
		return
	}
	response.MakeSuccess(c, http.StatusOK, "approve vip user successfully")

}

func checkLoginParams(reqLogin *request.ReqLogin) bool {
	if reqLogin.Account == "" || reqLogin.Password == "" {
		glog.Errorln("Missing login parameters")
		return false
	}
	return true
}
