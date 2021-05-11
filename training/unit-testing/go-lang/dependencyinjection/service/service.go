package service

import (
	"fmt"

	"github.com/ONSdigital/dp/training/unit-testing/go-lang/3-dependency-injection/datastore"
)

// Service contains the datastore store
type Service struct {
	Store datastore.Store
}

// GetNumber gets the ID
func (s *Service) GetNumber(ID int) error {
	// Use the `Get` method of the dependency to retreive the value of the database entry
	result, err := s.Store.Get(ID)
	if err != nil {
		return err
	}
	// Perform some validation, and output an error if it is too high
	if result > 10 {
		return fmt.Errorf("result too high: %d", result)
	}
	// Return nil, if the result is valid
	return nil
}

// NewGetNumber gets the
func NewGetNumber(store datastore.Store) func(int) error {
	return func(ID int) error {
		// Use the `Get` method of the dependency to retreive the value of the database entry
		result, err := store.Get(ID)
		if err != nil {
			return err
		}
		// Perform some validation, and output an error if it is too high
		if result > 10 {
			return fmt.Errorf("result too high: %d", result)
		}
		// Return nil, if the result is valid
		return nil
	}
}
