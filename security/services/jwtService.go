package services

import (
	"fmt"
	configService "rmtly-server/config/services"
	"time"
)

func createJwtToken() (token string, err error) {
	expirationInDays := configService.GetConfig().Security.ExpirationInDays
	expirationDate := time.Now().Add(24 * time.Hour * time.Duration(expirationInDays))
	fmt.Println(expirationDate)

	return "", nil
}

func verifyJwtToken(token string) error {
	return nil
}
