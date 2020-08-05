package interfaces

type SignUpRequest struct {
	QrCode   string `json:"qrCode"`
	DeviceId string `json:"deviceId"`
}

type SignUpResponse struct {
	Token string `json:"token"`
}
