package service

import (
	"errors"

	"testing"

	"github.com/ONSdigital/dp/training/unitTesting/goLang/3.dependencyInjection/datastore"
	. "github.com/smartystreets/goconvey/convey"
)

//TestServiceSuccessWithConvey tests...
func TestServiceSuccess1(t *testing.T) {
	Convey("Given an ID of 2", t, func() {
		// Create a new instance of the mock store
		m := new(datastore.MockStore)
		// In the "On" method, we assert that we want the "Get" method
		// to be called with one argument, that is 2
		// In the "Return" method, we define the return values to be 7, and nil (for the result and error values)
		m.On("Get", 2).Return(7, nil)
		// Next, we create a new instance of our service with the mock store as its "store" dependency
		s := Service{m}
		// The "GetNumber" method call is then made
		err := s.GetNumber(2)
		// The expectations that we defined for our mock store earlier are asserted here
		m.AssertExpectations(t)
		// Finally, we assert that we should'nt get any error
		if err != nil {
			t.Errorf("error should be nil, got: %v", err)
		}

	})
}

//TestServiceResultTooHighWithConvey tests
func TestServiceResultTooHigh(t *testing.T) {
	m := new(datastore.MockStore)
	// In this case, we simulate a return value of 24, which would fail the services validation
	m.On("Get", 2).Return(24, nil)
	s := Service{m}
	err := s.GetNumber(2)
	m.AssertExpectations(t)
	// We assert that we expect the "result too high" error given by the service
	if err.Error() != "result too high: 24" {
		t.Errorf("error should be 'result too high: 24', got: %v", err)
	}
}

//TestConveyServiceStoreError uses Convey to check StoreError is
func TestServiceStoreError(t *testing.T) {
	m := new(datastore.MockStore)
	// In this case, we simulate the case where the store returns an error, which may occur if it is unable to fetch the value
	m.On("Get", 2).Return(0, errors.New("failed"))
	s := Service{m}
	err := s.GetNumber(2)
	m.AssertExpectations(t)
	if err.Error() != "failed" {
		t.Errorf("error should be 'failed', got: %v", err)
	}
}

func TestServiceSuccess(t *testing.T) {
	// Create a new instance of the mock store
	m := new(datastore.MockStore)
	// In the "On" method, we assert that we want the "Get" method
	// to be called with one argument, that is 2
	// In the "Return" method, we define the return values to be 7, and nil (for the result and error values)
	m.On("Get", 2).Return(7, nil)
	// Next, we create a new instance of our service with the mock store as its "store" dependency
	s := Service{m}
	// The "GetNumber" method call is then made
	err := s.GetNumber(2)
	// The expectations that we defined for our mock store earlier are asserted here
	m.AssertExpectations(t)
	// Finally, we assert that we should'nt get any error
	if err != nil {
		t.Errorf("error should be nil, got: %v", err)
	}
}

func TestServiceResultTooHigh(t *testing.T) {
	m := new(datastore.MockStore)
	// In this case, we simulate a return value of 24, which would fail the services validation
	m.On("Get", 2).Return(24, nil)
	s := Service{m}
	err := s.GetNumber(2)
	m.AssertExpectations(t)
	// We assert that we expect the "result too high" error given by the service
	if err.Error() != "result too high: 24" {
		t.Errorf("error should be 'result too high: 24', got: %v", err)
	}
}

func TestServiceStoreError(t *testing.T) {
	m := new(datastore.MockStore)
	// In this case, we simulate the case where the store returns an error, which may occur if it is unable to fetch the value
	m.On("Get", 2).Return(0, errors.New("failed"))
	s := Service{m}
	err := s.GetNumber(2)
	m.AssertExpectations(t)
	if err.Error() != "failed" {
		t.Errorf("error should be 'failed', got: %v", err)
	}
}
