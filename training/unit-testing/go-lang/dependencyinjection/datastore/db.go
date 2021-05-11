package datastore

import (
	"database/sql"
)

type store struct {
	db *sql.DB
}

func NewDatastore(db *sql.DB) Store {
	return &store{db}
}

type Store interface {
	Get(ID int) (int, error)
}

func (d *store) Get(ID int) (int, error) {
	return 0, nil
}
