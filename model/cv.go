package model

import (
	"database/sql"
	"errors"

	"github.com/docsocsf/sponsor-portal/auth"
)

type CV struct {
	Name string `json:"name"`
	File string `json:"-"`
}

type CVReader interface {
	Get(auth.UserIdentifier) (CV, error)
}

type CVWriter interface {
	Put(auth.UserIdentifier, CV) error
}

type cvImpl struct {
	db *sql.DB
}

func NewCVReader(db *sql.DB) CVReader {
	return &cvImpl{db}
}

func NewCVWriter(db *sql.DB) CVWriter {
	return &cvImpl{db}
}

const (
	getCV    = `SELECT name, file FROM user_cv WHERE user_id = $1`
	insertCV = `
		INSERT INTO user_cv (user_id, name, file)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id)
		DO UPDATE SET name = $2, file = $3
	`
)

func (c *cvImpl) Get(id auth.UserIdentifier) (CV, error) {
	cv := CV{}
	err := c.db.QueryRow(getCV, id.User).Scan(&cv.Name, &cv.File)

	switch {
	case err == sql.ErrNoRows:
		return CV{}, DbError{NotFound: true, Err: errors.New("User does not have a CV")}
	case err != nil:
		return CV{}, DbError{Err: err}
	default:
		return cv, nil
	}
}

func (c *cvImpl) Put(id auth.UserIdentifier, cv CV) error {
	_, err := c.db.Exec(insertCV, id.User, cv.Name, cv.File)
	return err
}
