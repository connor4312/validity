package validity

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFloat(t *testing.T) {

	Convey("Given the rule \"value:0,100\"", t, func() {

		rules := Rules{"Baz": []string{"Float", "value:0,100"}}

		Convey("Given the value 50 (between intervals)", func() {

			Convey("The value should be valid", func() {
				data := TestStruct{Baz: 20}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 0 (lower interval)", func() {

			Convey("The value should be valid", func() {
				data := TestStruct{Baz: 0}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 100 (upper interval)", func() {

			Convey("The value should be valid", func() {
				data := TestStruct{Baz: 100}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value -1 (lower interval)", func() {

			Convey("The value should not be valid", func() {
				data := TestStruct{Baz: -1}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 101 (upper interval)", func() {

			Convey("The value should not be valid", func() {
				data := TestStruct{Baz: 101}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	Convey("Given the rule \"value_strict:500.2,2000.8\"", t, func() {
		rules := Rules{"Baz": []string{"Float", "value_strict:500.2,2000.8"}}

		Convey("Given the value 1000.78 (between intervals)", func() {

			Convey("The value should be valid", func() {
				data := TestStruct{Baz: 1000.78}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 500.3 (lower interval)", func() {

			Convey("The value should be valid", func() {
				data := TestStruct{Baz: 500.3}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 2000.7 (upper interval)", func() {

			Convey("The value should be valid", func() {
				data := TestStruct{Baz: 2000.7}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 500.2 (lower interval)", func() {

			Convey("The value should not be valid", func() {
				data := TestStruct{Baz: 500.2}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 2000.8 (upper interval)", func() {

			Convey("The value should not be valid", func() {
				data := TestStruct{Baz: 2000.8}
				result := ValidateStruct(data, rules)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	Convey("Given a validation which has the rule \"digits:4\"", t, func() {

		rule := Rules{"Baz": []string{"Float", "digits:4"}}

		Convey("Given the value 5000", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Baz: 5000}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 500", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Baz: 500}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 50000", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Baz: 50000}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	Convey("Given a validation which has the rule \"max:100\"", t, func() {

		rule := Rules{"Baz": []string{"Float", "max:100"}}

		Convey("Given the value 50", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Baz: 50}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 100", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Baz: 100}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 101", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Baz: 101}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

}
