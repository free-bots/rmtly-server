package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	configService "rmtly-server/config/services"
	"rmtly-server/security/interfaces"
	"time"
)

var secret = []byte("sdfsf") // todo from keyfile or config

func CreateJwtToken(deviceId string) (token string, err error) {
	expirationInDays := configService.GetConfig().Security.ExpirationInDays
	expirationDate := time.Now().Add(24 * time.Hour * time.Duration(expirationInDays))

	claims := &interfaces.Claims{
		DeviceId: deviceId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDate.Unix(),
		},
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := tokenWithClaims.SignedString(secret)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return signedToken, nil
}

func VerifyJwtToken(signedToken string) (deviceId string, err error) {
	claims := &interfaces.Claims{}

	_, err = jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})

	if err != nil {
		return "", err
	}

	return claims.DeviceId, nil
}
