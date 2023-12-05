package models

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.StandardClaims
}

func getEnvironmentVars(key string) string {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetGooglePublicKey(keyID string) (string, error) {
	google_url := getEnvironmentVars("GOOGLE_URL")

	resp, err := http.Get(google_url)
	if err != nil {
		log.Fatal(err)
		return "Problem fetching public key", err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "Problem loading public key", err
	}

	key_map := map[string]string{}
	err = json.Unmarshal(data, &key_map)

	if err != nil {
		log.Fatal(err)
		return "Problem getting json data", err
	}

	key, ok := key_map[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}
