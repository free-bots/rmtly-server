package services

import (
	configService "rmtly-server/config/services"
	"rmtly-server/contants"
	"rmtly-server/information/interfaces"
)

func GetInformation() *interfaces.InformationResponse {
	information := configService.GetConfig().Information
	return &interfaces.InformationResponse{
		NAME:    information.Name,
		URL:     "",
		ID:      information.Id,
		VERSION: contants.VERSION,
	}
}
