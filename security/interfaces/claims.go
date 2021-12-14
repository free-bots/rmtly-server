package interfaces

import "github.com/golang-jwt/jwt"

type Claims struct {
	DeviceId string
	jwt.StandardClaims
}
