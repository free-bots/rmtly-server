package repositories

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	interfaces2 "rmtly-server/config/interfaces"
	configService "rmtly-server/config/services"
	"rmtly-server/scripts/interfaces"
	"rmtly-server/scripts/utils"
	"strings"
)

func FindAll() []*interfaces.ScriptInformation {
	scripts := createEmptyArray()

	path := getConfig().Path

	utils.OpenFolderAndConsumeFiles(path, utils.Consumer{OnError: func() {
	}, OnFileName: func(fileName string) {
		if strings.Contains(fileName, ".json") {
			scriptInformation := createScriptInformation(path, fileName)
			if scriptInformation != nil {
				scripts = append(scripts, scriptInformation)
			}
		}
	}})

	return scripts
}

func getConfig() interfaces2.ScriptConfig {
	return configService.GetConfig().ScriptConfig
}

func createEmptyArray() []*interfaces.ScriptInformation {
	return make([]*interfaces.ScriptInformation, 0)
}

func createScriptInformation(path string, name string) *interfaces.ScriptInformation {
	file, err := os.Open(filepath.Join(path, name))

	defer func() {
		_ = file.Close()
	}()

	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	scriptInformation := new(interfaces.ScriptInformation)

	if !json.Valid(data) {
		return nil
	}

	err = json.Unmarshal(data, scriptInformation)
	if err != nil {
		return nil
	}

	return scriptInformation
}
