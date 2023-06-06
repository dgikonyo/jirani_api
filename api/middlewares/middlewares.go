package middlewares

import (
	"errors"
	"net/http"

	"jirani-api/api/auth"
	"jirani-api/api/responses"
)

func setMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next(res, req)
	}
}

func setMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		err := auth.validateToken(req)
		if err != nil {
			responses.ERROR(res, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(res, req)
	}
}