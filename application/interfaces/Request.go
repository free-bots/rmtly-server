package interfaces

type ExecuteRequest struct {
	Id           int `json:"id"`
	ExecuteDelay int `json:"executeDelay"`
}
