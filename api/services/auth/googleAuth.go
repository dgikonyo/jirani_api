package auth

import (
	"errors"
	"fmt"
	"jirani-api/api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Validate google JWT
func ValidateGoogleJwt(tokenString string) (models.GoogleClaims, error) {
	claimsStruct := models.GoogleClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := models.GetGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return models.GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*models.GoogleClaims)

	if !ok {
		return models.GoogleClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return models.GoogleClaims{}, errors.New("iss is invalid")
	}

	if claims.Audience != "YOUR_CLIENT_ID_HERE" {
		return models.GoogleClaims{}, errors.New("aud is invalid")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return models.GoogleClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}
