package auth

import (
	"github.com/mashmorsik/rand-quote-grpc/infrastructure/auth/data"
	"github.com/mashmorsik/rand-quote-grpc/pkg/log"
)

type UserData struct {
	db data.Data
}

func NewUserData(db data.Data) *UserData {
	return &UserData{db: db}
}

type Userer interface {
	isUser(u *data.User) (bool, error)
	getUser(u *data.User) (int, error)
	createUser(u *data.User) (int, error)
	UserAuthentication(u *data.User) (string, error)
}

func (d *UserData) isUser(u *data.User) (bool, error) {
	isUser, err := d.db.IsUser(u)
	if err != nil {
		log.Nlog.Error().Msgf("can't check if user exists %s", err)
	}
	return isUser, nil
}

func (d *UserData) getUser(u *data.User) (int, error) {
	id, err := d.db.GetUser(u)
	if err != nil {
		log.Nlog.Error().Msgf("can't create user: %v", u)
	}

	return id, nil
}

func (d *UserData) createUser(u *data.User) (int, error) {
	id, err := d.db.CreateUser(u)
	if err != nil {
		log.Nlog.Error().Msgf("can't create user: %v", u)
	}

	return id, nil
}
