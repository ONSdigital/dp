package datastore

import (
	"database/sql"
)

func NewDatastore(db *sql.DB) Store {
	return &store{db}
}

type Store interface {
	Get(ID int) (int, error)
}

type store struct {
	db *sql.DB
}

func (d *store) Get(ID int) (int, error) {
	return 0, nil
}
