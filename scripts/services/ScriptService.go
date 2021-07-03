package services

import (
	"path/filepath"
	interfaces3 "rmtly-server/application/interfaces"
	"rmtly-server/application/services"
	interfaces2 "rmtly-server/config/interfaces"
	configService "rmtly-server/config/services"
	"rmtly-server/scripts/interfaces"
	"rmtly-server/scripts/repositories"
	"rmtly-server/scripts/utils"
	"strings"
)

func GetAllScripts() []*interfaces.ScriptInformation {
	if !getConfig().Enabled {
		return make([]*interfaces.ScriptInformation, 0)
	}

	return repositories.FindAll()
}

func Exec(fileName string, request interfaces3.ExecuteRequest) {
	path := getConfig().Path

	utils.OpenFolderAndConsumeFiles(path, utils.Consumer{OnError: func() {
	}, OnFileName: func(consumerFileName string) {
		if strings.Contains(consumerFileName, ".sh") && strings.ReplaceAll(consumerFileName, ".sh", "") == fileName {
			services.ExecuteScript(filepath.Join(path, consumerFileName), request)
		}
	}})
}

func getConfig() interfaces2.ScriptConfig {
	return configService.GetConfig().ScriptConfig
}
