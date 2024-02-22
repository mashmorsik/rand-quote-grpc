package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/mashmorsik/rand-quote-grpc/infrastructure/auth/data"
	"github.com/mashmorsik/rand-quote-grpc/pkg/log"
	"time"
)

const JwtKey = "JWT_SECRET"

func generateToken(u *data.User) (string, error) {
	payload := jwt.MapClaims{
		"exp": time.Now().Add(60 * time.Minute).Unix(),
		"id":  u.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		log.Nlog.Error().Msgf("can't generate token for %s", u.Id)
		return "", err
	}

	return tokenStr, nil
}
