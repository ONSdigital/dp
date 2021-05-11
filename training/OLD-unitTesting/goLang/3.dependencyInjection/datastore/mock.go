package datastore

import (
	"github.com/stretchr/testify/mock"
)

//MockStore is a struct
type MockStore struct {
	mock.Mock
}

//Get will get the int based on the ID
func (m *MockStore) Get(ID int) (int, error) {
	returnVals := m.Called(ID)
	return returnVals.Get(0).(int), returnVals.Error(1)
}
