package datastore

import (
	"database/sql"
)

type store struct {
	d *sql.DB
}

func NewDatastore(d *sql.DB) Store {
	return &store{d}
}

type Store interface {
	Get(ID int) (int, error)
}

func (d *store) Get(ID int) (int, error) {
	return 0, nil
}
