package services

import (
	"encoding/json"
	"io/ioutil"
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

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		createDefaultConfig()
	}

	readConfig()
}

func GetConfig() interfaces.Config {
	if currentConfig == nil {
		log.Fatal("try to init config first")
	}

	return *currentConfig
}

func createDefaultConfig() {
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

func readConfig() {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := new(interfaces.Config)

	if !json.Valid(data) {
		log.Fatal("invalid config file format: use valid json")
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal(err)
	}

	currentConfig = config
}

func getDefaultConfig() interfaces.Config {
	config := new(interfaces.Config)

	// image config
	config.Image.CacheExpiresInMillis = 10000
	config.Image.MaxImagesInCache = 100
	config.Image.ImageQuality = 512

	// application config
	config.Application.CacheExpiresInMillis = 10000

	// network config
	config.Network.Address = "0.0.0.0:3000"

	// security config
	config.Security.ExpirationInDays = 99
	config.Security.Secret = "authenticationCode"

	return *config
}
