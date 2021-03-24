package helloworld

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloWorldX(t *testing.T) {
	a := 100
	b := 20
	Convey("Given two integers", t, func() {
		//Lets check the first and last one first
		Convey("Check a is larger than b", func() {
			aTestA(t)
			aTestB(t)
			So(a, ShouldBeGreaterThan, b)
		})
	})
}

func aTestA(t *testing.T) {
	Print("within testA")
	Convey("Scenario: one", func() {
		
		Convey("Given ", func() {

			Convey("When ", func() {

				Convey("Then ", func() {
				})
			})
		})
	})
}

func aTestB(t *testing.T) {
	Print("within testB")
	Convey("Scenario: two ", func() {
		Convey("Given ", func() {
			Convey("When ", func() {
				Convey("Then ", func() {
				})
			})
		})
	})
}
