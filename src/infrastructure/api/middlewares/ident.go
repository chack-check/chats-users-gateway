package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/chack-check/chats-users-gateway/infrastructure/api/settings"
	"github.com/golang-jwt/jwt/v5"
)

type TokenSubject struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

func GetTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Settings.APP_SECRET_KEY), nil
	})

	return token, err
}

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header["Authorization"]
		ctx := r.Context()

		if len(authorization) != 0 {
			tokenString := strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1)
			token, err := GetTokenFromString(tokenString)
			if err == nil && token.Valid {
				ctx = context.WithValue(r.Context(), "token", token)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		ctx = context.WithValue(r.Context(), "token", nil)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTokenSubject(token *jwt.Token) (TokenSubject, error) {
	tokenSubject := TokenSubject{}

	subject, err := token.Claims.GetSubject()
	if err != nil {
		return tokenSubject, err
	}

	err = json.Unmarshal([]byte(subject), &tokenSubject)
	if err != nil {
		return tokenSubject, err
	}

	return tokenSubject, nil
}
