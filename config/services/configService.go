package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"rmtly-server/config/interfaces"
)

var currentConfig *interfaces.Config

var configPath string
var configDir string
var configFile string

func InitConfig() {

	err := error(nil)
	configPath, err = os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	configDir = filepath.Join(configPath, "rmtly-server")
	configFile = filepath.Join(configPath, "rmtly-server", "config.json")

	_ = os.Mkdir(configDir, os.ModePerm)

	file, err := os.Open(configFile)
	if err != nil {
		defaultConfig, err := os.Create(configFile)
		if err != nil {
			log.Fatal(err)
		}

		data, err := json.Marshal(getDefaultConfig())
		if err != nil {
			log.Fatal(err)
		}

		_, err = defaultConfig.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		if defaultConfig != nil {
			_ = defaultConfig.Close()
		}
	}

	if file != nil {
		_ = file.Close()
	}
}

func GetConfig() interfaces.Config {
	if currentConfig == nil {
		file, err := os.Open(configFile)
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(file)

		configDate := make([]byte, 1024)

		_, err = reader.Read(configDate)

		fmt.Println(err)
		var config *interfaces.Config

		_ = json.Unmarshal(configDate, config)

		currentConfig = config
	}

	return *currentConfig
}

func getDefaultConfig() interfaces.Config {
	config := new(interfaces.Config)

	config.ImageQuality = 512

	return *config
}
