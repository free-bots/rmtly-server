package interfaces

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	DeviceId string
	jwt.StandardClaims
}
