package validity

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {

	// value

	Convey("Given the validation rule \"value:0,100\"", t, func() {

		rule := Rules{"Bar": []string{"Int", "value:0,100"}}
		Convey("Given the value -1", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: -1}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 0", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 0}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 1", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 1}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 100", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 100}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 101", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 101}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// value_strict

	Convey("Given the validation rule \"value_strict:0,100\"", t, func() {

		rule := Rules{"Bar": []string{"Int", "value_strict:0,100"}}
		Convey("Given the value -1", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: -1}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 0", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 0}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 1", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 1}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 100", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 100}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 101", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 101}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// digits

	Convey("Given the validation rule \"digits:3\"", t, func() {

		rule := Rules{"Bar": []string{"Int", "digits:3"}}

		Convey("Given the value 10", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 10}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 100", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 100}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 1000", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 1000}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// digits_between

	Convey("Given the validation rule \"digits_between:2,5\"", t, func() {

		rule := Rules{"Bar": []string{"Int", "digits_between:2,5"}}

		Convey("Given the value 1", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 1}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 12", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 12}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 123", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 123}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 12345", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 12345}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 123456", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 123456}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// digits_between_strict

	Convey("Given the validation rule \"digits_between_strict:2,5\"", t, func() {

		rule := Rules{"Bar": []string{"Int", "digits_between_strict:2,5"}}

		Convey("Given the value 1", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 1}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 12", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 12}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 123", func() {

			Convey("The result should be valid", func() {
				data := TestStruct{Bar: 123}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 12345", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 12345}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 123456", func() {

			Convey("The result should not be valid", func() {
				data := TestStruct{Bar: 123456}
				result := ValidateStruct(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

}
