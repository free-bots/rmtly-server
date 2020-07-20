package interfaces

type ApplicationEntry struct {
	Id         string    `json:"id"`
	Version    float32   `json:"version"`
	Type       string    `json:"type"`
	Name       string    `json:"name"`
	Comment    string    `json:"comment"`
	TryExec    string    `json:"tryExec"`
	Exec       string    `json:"exec"`
	Icon       string    `json:"icon"`
	MimeType   []string  `json:"mimeType"`
	Actions    []*Action `json:"actions"`
	Categories []string  `json:"categories"`
}

type Action struct {
	Name string `json:"name"`
	Exec string `json:"exec"`
	Icon string `json:"icon"`
}
