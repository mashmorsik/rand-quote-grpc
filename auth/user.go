package auth

import (
	"github.com/mashmorsik/rand-quote-grpc/auth/data"
	"github.com/mashmorsik/rand-quote-grpc/logger"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type DB struct {
	db data.Data
}

func (d *DB) isUser(u *User) (bool, error) {
	isUser, err := d.db.IsUser(u)
	if err != nil {
		logger.Nlog.Error().Msgf("can't check if user exists %s", err)
	}
	return isUser, nil
}

func (d *DB) getUser(u *User) (int, error) {
	id, err := d.db.GetUser(u)
	if err != nil {
		logger.Nlog.Error().Msgf("can't create user: %v", u)
	}

	return id, nil
}

func (d *DB) createUser(u *User) (int, error) {
	id, err := d.db.CreateUser(u)
	if err != nil {
		logger.Nlog.Error().Msgf("can't create user: %v", u)
	}

	return id, nil
}
