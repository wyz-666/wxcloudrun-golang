package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
)

const JWTSECRET = "CPIFJWT666"

type CPIFUserClaims struct {
	UserID   string
	UserType int
	jwt.StandardClaims
}

// CreateJwtToken Create JWT Token
func CreateJwtToken(userID string, userType int) (string, error) {
	// 2 hours
	var TokenExpireDuration = time.Second * time.Duration(3600) * 2

	glog.Info("duration:", TokenExpireDuration)
	//jwt secert
	jwtSecret := []byte(JWTSECRET)

	c := CPIFUserClaims{
		userID,
		userType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "CPIF-SYSTEM",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(jwtSecret)
}
