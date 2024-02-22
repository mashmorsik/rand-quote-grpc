package auth

import (
	"github.com/mashmorsik/rand-quote-grpc/infrastructure/auth/data"
	"github.com/pkg/errors"
)

func (d *UserData) UserAuthentication(u *data.User) (string, error) {
	id, err := d.getUser(u)
	if err != nil {
		return "", errors.WithMessagef(err, "can't get userId: %v\n", u)
	}

	u.Id = id

	token, err := generateToken(u)
	if err != nil {
		return "", errors.WithMessagef(err, "can't generate token for id: %v", u.Id)
	}

	return token, nil
}
