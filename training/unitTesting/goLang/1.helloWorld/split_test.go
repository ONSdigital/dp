package helloworld

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//Test_MainT is one method for splitting up tests into multiple functions.
func Test_MainController(t *testing.T) {
	Print("Within parent")
	testA(t)
	testB(t)
}

//func TestHelloWorld3(t *testing.T) {
func testA(t *testing.T) {
	Print("within testA")
	Convey("Scenario: one", t, func() {
		Convey("Given ", func() {
			Convey("When ", func() {
				Convey("Then ", func() {
				})
			})
		})
	})
}

func testB(t *testing.T) {
	Print("within testB")
	Convey("Scenario: two ", t, func() {
		Convey("Given ", func() {
			Convey("When ", func() {
				Convey("Then ", func() {
				})
			})
		})
	})
}
