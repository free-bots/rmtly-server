package interfaces

type Config struct {
	Application ApplicationConfig `json:"application"`
	Image       ImageConfig       `json:"image"`
	Security    SecurityConfig    `json:"security"`
	Network     NetworkConfig     `json:"network"`
}

type SecurityConfig struct {
	ExpirationInDays int    `json:"expirationInDays"`
	Secret           string `json:"secret"`
	KeyFile          string `json:"keyFile"` // todo remove secret and use key file
}

type NetworkConfig struct {
	Address string `json:"address"`
}

type ApplicationConfig struct {
	CacheExpiresInMillis int `json:"cacheExpiresInMillis"`
}

type ImageConfig struct {
	CacheExpiresInMillis int `json:"cacheExpiresInMillis"`
	MaxImagesInCache     int `json:"maxImagesInCache"`
	ImageQuality         int `json:"imageQuality"`
}
