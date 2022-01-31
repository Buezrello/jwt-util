package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gindin.com/jwt-util/domain"
	"gindin.com/jwt-util/dto"
	"gindin.com/jwt-util/service"
)

/*
Sample URL
http://localhost:8181/createtoken

Sample body
{
	"issuer": "Igor Gindin",
	"subject": "My First Token"
}
*/
func CreateToken(w http.ResponseWriter, r *http.Request) {
	var tokenRequest dto.TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&tokenRequest); err != nil {
		log.Println("Error while decoding create token request:" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		jwtMaker := service.JWTMaker{
			SecretKey: domain.HMAC_SAMPLE_SECRET,
		}
		token, err := jwtMaker.CreateToken(tokenRequest.Issuer, tokenRequest.Subject, domain.TOKEN_DURATION)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			accessToken := &dto.TokenResponse{
				AccessToken: token,
			}
			writeResponse(w, http.StatusOK, *accessToken)
		}
	}
}

/*
Sample URL string
http://localhost:8181/decodetoken?token=somevalidtokenstring
*/
func DecodeToken(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	token := keys.Get("token")

	if token != "" {
		jwtMaker := service.JWTMaker{
			SecretKey: domain.HMAC_SAMPLE_SECRET,
		}
		if claim, err := jwtMaker.DecodeToken(token); err != nil {
			writeResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			writeResponse(w, http.StatusOK, claim)
		}
	} else {
		writeResponse(w, http.StatusForbidden, errors.New("missing token"))
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
