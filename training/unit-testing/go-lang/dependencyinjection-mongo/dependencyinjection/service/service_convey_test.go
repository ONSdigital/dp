package service

import (
	"errors"
	"github.com/ONSdigital/dp/training/unit-testing/go-lang/dependencyinjection/datastore"
	"testing"

	ass "github.com/smartystreets/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

// TestServiceSuccessWithConvey tests...
func TestServiceSuccessWithConvey(t *testing.T) {
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
		// and checked that they were called as specified with On and Return.
		// Calls may have occurred in any order.
		So(err, ShouldBeNil)
		// Finally, we assert that we shouldn't get any error
		if err != nil {
			t.Errorf("error should be nil, got: %v", err)
		}
	})
}

//TestServiceResultTooHighWithConvey tests when service returns too high
func TestServiceResultTooHighWithConvey(t *testing.T) {
	Convey("Given a result from Service which is too high", t, func() {
		m := new(datastore.MockStore)
		// In this case, we simulate a return value of 24, which would fail the services validation
		m.On("Get", 2).Return(24, nil)
		s := Service{m}
		err := s.GetNumber(2)
		// We expect the "result too high" error given by the service
		if ok, message := ass.So(err.Error(), ShouldEqual, "result too high: 24"); !ok {
			t.Errorf("error should be 'result too high: 24', got: %v \n%v", err, message)
		}
	})
}

//TestConveyServiceStoreError tests when service returns a Error
func TestServiceStoreErrorWithConvey(t *testing.T) {
	m := new(datastore.MockStore)
	// In this case, we simulate the case where the store returns an error, which may occur if it is unable to fetch the value
	m.On("Get", 2).Return(0, errors.New("failed"))
	s := Service{m}
	err := s.GetNumber(2)
	// We expect the error to be failed
	if ok, message := ass.So(err.Error(), ShouldEqual, "failed"); !ok {
		t.Errorf("error should be 'failed', got: %v \n%v", err, message)
	}
}
