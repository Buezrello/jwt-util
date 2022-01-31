package service

import (
	"time"

	"gindin.com/jwt-util/domain"
)

type Maker interface {
	CreateToken(issuer string, subject string, duration time.Duration) (string, error)
	DecodeToken(token string) (*domain.Payload, error)
}
