package middleware

import (
	"PartyRoom.API/internal/config"
	"context"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func ValidateJwt(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cfg, err := config.New(".")
		if err != nil {
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}
		ctx := r.Context()
		authorization := r.Header.Get("Authorization")
		var jwtString string
		if strings.HasPrefix(authorization, "Bearer") {
			jwtString = strings.TrimPrefix(authorization, "Bearer ")
		}
		if jwtString == "" {
			http.Error(w, "Не авторизован: отсутствует токен Bearer", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(cfg.JwtSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Не авторизован: токен не действителен", http.StatusUnauthorized)
			return
		}

		var userId = claims["sub"].(string)

		if err != nil {
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
