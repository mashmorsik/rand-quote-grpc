package auth

import "github.com/mashmorsik/rand-quote-grpc/logger"

func (d *DB) userAuth(u *User) (string, error) {
	id, err := d.getUser(u)
	if err != nil {
		logger.Nlog.Error().Msgf("can't get userId: %v", u)
		return "", err
	}

	u.Id = id

	token, err := generateToken(u)
	if err != nil {
		logger.Nlog.Error().Msgf("can't generate token for id: %v", u.Id)
	}

	return token, nil
}
