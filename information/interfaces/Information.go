package interfaces

type InformationResponse struct {
	ID      string `json:"id"`
	NAME    string `json:"name"`
	URL     string `json:"url"`
	VERSION string `json:"version"`
}
