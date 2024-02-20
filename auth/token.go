package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/mashmorsik/rand-quote-grpc/logger"
	"time"
)

const JwtKey = "JWT_SECRET"

func generateToken(u *User) (string, error) {
	payload := jwt.MapClaims{
		"exp": time.Now().Add(60 * time.Minute),
		"id":  u.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)

	tokenStr, err := token.SignedString(JwtKey)
	if err != nil {
		logger.Nlog.Error().Msgf("can't generate token for %s", u.Id)
		return "", nil
	}

	return tokenStr, nil
}
