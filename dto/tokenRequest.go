package dto

type TokenRequest struct {
	Issuer  string `json:"issuer"`
	Subject string `json:"subject"`
}
