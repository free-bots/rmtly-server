package interfaces

import GoppilcationEntry "github.com/free-bots/GopplicationEntry"

type IconResponse struct {
	ApplicationId string `json:"applicationId"`
	IconBase64    string `json:"iconBase64"`
}

type SortedApplicationResponse struct {
	SortedBy string        `json:"sortedBy"`
	Values   []SortedValue `json:"values"`
}

type SortedValue struct {
	SortedValue        string                                `json:"sortedValue"`
	ApplicationEntries []*GoppilcationEntry.ApplicationEntry `json:"applicationEntries"`
}

type ExecuteResponse struct {
	Application *GoppilcationEntry.ApplicationEntry `json:"application"`
}
