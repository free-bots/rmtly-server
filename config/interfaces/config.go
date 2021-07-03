package interfaces

type Config struct {
	Application  ApplicationConfig `json:"application"`
	Image        ImageConfig       `json:"image"`
	Security     SecurityConfig    `json:"security"`
	Network      NetworkConfig     `json:"network"`
	Information  Information       `json:"information"`
	ScriptConfig ScriptConfig      `json:"script_config"`
}

type SecurityConfig struct {
	ExpirationInDays int    `json:"expirationInDays"`
	Secret           string `json:"secret"`
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

type Information struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ScriptConfig struct {
	Enabled bool   `json:"enabled"`
	Path    string `json:"path"`
}
