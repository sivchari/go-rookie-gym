package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

var jwtKey = struct{}{}

func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("jwt-secret")
		tokenStr := r.Header.Get("Authorization")
		trimTokenStr := strings.TrimPrefix(tokenStr, "Bearer ")
		token, err := jwt.Parse(trimTokenStr, func(token *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			log.Printf("failed to parse jwt token err = %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), jwtKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
