package interfaces

type IconTheme struct {
	Name              string          `json:"name"`
	Comment           string          `json:"comment"`
	Inherits          []string        `json:"inherits"`
	Directories       []string        `json:"directories"`
	ScaledDirectories []string        `json:"scaled_directories"`
	Hidden            bool            `json:"hidden"`
	Example           string          `json:"example"`
	DirectoriesData   []IconDirectory `json:"directories"`
}

type IconDirectory struct {
	Size      int    `json:"size"`
	Scale     int    `json:"scale"`
	Context   string `json:"context"`
	Type      string `json:"type"`
	MaxSize   int    `json:"max_size"`
	MinSize   int    `json:"min_size"`
	Threshold int    `json:"threshold"`
}

func (theme IconTheme) HasThemesParents() bool {
	return len(theme.Inherits) > 0
}

func (theme IconTheme) GetParents() []IconTheme {
	return make([]IconTheme, 0)
}

// Context as enum
// Actions
// Devices
// FileSystems
// MimeTypes
