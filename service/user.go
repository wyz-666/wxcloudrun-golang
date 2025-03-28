package service

import (
	"errors"
	"wxcloudrun-golang/app/handlers/request"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/utils/crypto"

	"github.com/golang/glog"
)

const (
	UUID_PREFIX      = "CPIF"
	ADMIN_USER_TYPE  = 1
	VIP_USER_TYPE    = 2
	COMMON_USER_TYPE = 3
)

func QueryUserInfo(account string) (uuid, password string, userType int, err error) {
	var user model.User
	cli := db.Get()
	err = cli.Where("account = ?", account).First(&user).Error
	if err != nil {
		return "", "", 0, err
	}
	return user.UserID, user.PasswordHash, user.Type, nil

}

func AddUser(user *request.ReqUser) error {
	//添加判断唯一性的内容（uuid唯一，用户名唯一，公司唯一）
	uuid, err := generateUuid(user.Account, user.Type)
	if err != nil {
		return err
	}
	passwordHash := crypto.CalculateSHA256(user.Password, "FDUCPIF")
	approved := true
	if user.Type == 2 {
		approved = false
	}
	addUser := model.User{
		UserID:       uuid,
		UserName:     user.Name,
		CompanyName:  user.Company,
		Type:         user.Type,
		Phone:        user.Phone,
		Email:        user.Email,
		Account:      user.Account,
		PasswordHash: passwordHash,
		Approved:     approved,
	}
	cli := db.Get()
	err = cli.Create(&addUser).Error
	if err != nil {
		glog.Errorln("create new user error: %v", err)
		return err
	}
	return nil
}

func generateUuid(account string, userType int) (string, error) {
	accountHash := crypto.CalculateSHA256(account, "uuid")

	var uuid string
	switch userType {
	case ADMIN_USER_TYPE:
		uuid = UUID_PREFIX + "0" + accountHash
	case VIP_USER_TYPE:
		uuid = UUID_PREFIX + "A" + accountHash
	case COMMON_USER_TYPE:
		uuid = UUID_PREFIX + "B" + accountHash
	default:
		return "", errors.New("user type is not exist")
	}
	return uuid, nil
}
