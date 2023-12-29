package tokenManager

import (
	"PartyRoom.API/internal/config"
	"PartyRoom.API/internal/domain"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRefreshToken(userId string) string {
	length := 120
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() + "_" + userId

	return str
}

func GenerateAccessToken(user domain.User) (string, error) {
	cfg, err := config.New(".")
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()
	expirationTime := now.Add(22 * time.Hour)
	var roleNames []string
	for i := 0; i < len(user.Roles); i++ {
		roleNames = append(roleNames, user.Roles[i].Name)
	}
	_ = roleNames
	tokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"exp":   expirationTime.Unix(),
		"iat":   now.Unix(),
		"roles": roleNames,
		"email": user.Email,
	})

	tokenString, err := tokenByte.SignedString([]byte(cfg.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
