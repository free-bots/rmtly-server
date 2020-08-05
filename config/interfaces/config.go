package interfaces

type Config struct {
	ImageQuality int            `json:"imageQuality"`
	Security     SecurityConfig `json:"security"`
}

type SecurityConfig struct {
	ExpirationInDays int    `json:"expirationInDays"`
	Secret           string `json:"secret"`
	KeyFile          string `json:"keyFile"` // todo remove secret and use key file
}
