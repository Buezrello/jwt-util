package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"
const TOKEN_DURATION = time.Hour

type Payload struct {
	Issuer   string    `json:"issuer"`
	Subject  string    `json:"subject"`
	IssuedAt time.Time `json:"issued_at"`
	jwt.StandardClaims
}

func NewPayload(issuer string, subject string, duration time.Duration) *Payload {
	payload := &Payload{
		Issuer:   issuer,
		Subject:  subject,
		IssuedAt: time.Now(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	return payload
}
