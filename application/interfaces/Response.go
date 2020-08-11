package interfaces

type IconResponse struct {
	ApplicationId string `json:"applicationId"`
	IconBase64    string `json:"iconBase64"`
}

type SortedApplicationResponse struct {
	SortedBy string        `json:"sortedBy"`
	Values   []SortedValue `json:"values"`
}

type SortedValue struct {
	SortedValue        string              `json:"sortedValue"`
	ApplicationEntries []*ApplicationEntry `json:"applicationEntries"`
}
