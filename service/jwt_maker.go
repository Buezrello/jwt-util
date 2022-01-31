package service

import (
	"errors"
	"log"
	"time"

	"gindin.com/jwt-util/domain"
	"github.com/dgrijalva/jwt-go"
)

type JWTMaker struct {
	SecretKey string
}

func NewJWTMaker(secretKey string) Maker {
	return &JWTMaker{SecretKey: secretKey}
}

func (maker *JWTMaker) CreateToken(issuer string, subject string, duration time.Duration) (string, error) {
	payload := domain.NewPayload(issuer, subject, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.SecretKey))
}

func (maker *JWTMaker) DecodeToken(token string) (*domain.Payload, error) {
	var claim *domain.Payload
	var ok bool
	if jwtToken, err := jwtTokenFromString(token, maker.SecretKey); err != nil {
		return nil, err
	} else {
		if jwtToken.Valid {
			claim, ok = jwtToken.Claims.(*domain.Payload)
			if !ok {
				return nil, errors.New("invalid tolen")
			}
		}
	}

	return claim, nil
}

func jwtTokenFromString(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Payload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Println("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}
