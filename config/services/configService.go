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
	return interfaces.Config{
		Image: interfaces.ImageConfig{
			CacheExpiresInMillis: 10000,
			MaxImagesInCache:     100,
			ImageQuality:         512,
		},
		Application: interfaces.ApplicationConfig{
			CacheExpiresInMillis: 10000,
		},
		Network: interfaces.NetworkConfig{
			Address: "0.0.0.0:3000",
		},
		Security: interfaces.SecurityConfig{
			ExpirationInDays: 99,
			Secret:           "authenticationCode",
		},
		Information: interfaces.Information{
			Id:   "rmtly",
			Name: "rmtly",
		},
		ScriptConfig: interfaces.ScriptConfig{
			Path:    filepath.Join(configDir, "scripts"),
			Enabled: true,
		},
	}
}
