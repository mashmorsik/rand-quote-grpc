package data

import (
	"database/sql"
	"errors"
	"github.com/mashmorsik/rand-quote-grpc/auth"
	"github.com/mashmorsik/rand-quote-grpc/logger"
)

type Data struct {
	db *sql.DB
}

func NewData(db *sql.DB) *Data {
	if db == nil {
		panic("db is nil")
	}
	return &Data{db: db}
}

func (d *Data) Db() *sql.DB {
	return d.db
}

func MustConnect() *sql.DB {
	connStr := "postgres://postgres:mysecretpassword@localhost:5435/rand-quote?sslmode=disable&application_name=rand-quote&connect_timeout=5"

	connection, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = connection.Ping(); err != nil {
		logger.Nlog.Panic().Err(err)
	}

	return connection
}

func (d *Data) CreateUser(u *auth.User) (int, error) {
	var id int

	sqlCreateUser := `
	insert into users(email, password)
	values($1, $2)`

	row, err := d.db.Query(sqlCreateUser, u.Email, u.Password)
	if err != nil {
		logger.Nlog.Error().Msgf("can't send query: %s", sqlCreateUser)
		return 0, err
	}

	if row.Next() {
		if err = row.Scan(&id); err != nil {
			logger.Nlog.Error().Msgf("can't scan rows: %v and get userId", row)
			return 0, err
		}
	}
	return id, nil
}

func (d *Data) IsUser(u *auth.User) (bool, error) {
	var result bool

	sqlGetUserId := `
	SELECT * FROM users
	WHERE email = ($1) and password = ($2)`

	row := d.db.QueryRow(sqlGetUserId, u.Email, u.Password)
	err := row.Scan(result)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		logger.Nlog.Error().Msgf("can't check if user exists %s", err)
		return false, nil
	}

	return true, nil
}

func (d *Data) GetUser(u *auth.User) (int, error) {
	var id int

	sqlGetUserId := `
	SELECT * FROM users
	WHERE email = ($1) and password = ($2)`

	row, err := d.db.Query(sqlGetUserId, u.Email, u.Password)
	if row.Next() {
		err = row.Scan(&id)
		if err != nil {
			logger.Nlog.Error().Msgf("can't scan rows: %v and get userId", row)
			return 0, err
		}
	}
	return id, nil
}
