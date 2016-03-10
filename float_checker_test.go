package validity

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFloat(t *testing.T) {

	// value

	Convey("Given the rule \"value:0,100\"", t, func() {

		rules := Rules{"Baz": Field{
			Name:  "Baz",
			Type:  "Float",
			Rules: []string{"value:0,100"},
		}}

		Convey("Given the value 50 (between intervals)", func() {

			Convey("The value should be valid", func() {
				data := map[string]interface{}{"Baz": 20}
				result := Validate(data, rules)
				log.Println(result.Errors)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 0 (lower interval)", func() {

			Convey("The value should be valid", func() {
				data := map[string]interface{}{"Baz": 0}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 100 (upper interval)", func() {

			Convey("The value should be valid", func() {
				data := map[string]interface{}{"Baz": 100}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value -1 (lower interval)", func() {

			Convey("The value should not be valid", func() {
				data := map[string]interface{}{"Baz": -1}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 101 (upper interval)", func() {

			Convey("The value should not be valid", func() {
				data := map[string]interface{}{"Baz": 101}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// value_strict

	Convey("Given the rule \"value_strict:500.2,2000.8\"", t, func() {
		rules := Rules{"Baz": Field{
			Type:  "Float",
			Name:  "Foo",
			Rules: []string{"value_strict:500.2,2000.8"},
		},
		}

		Convey("Given the value 1000.78 (between intervals)", func() {

			Convey("The value should be valid", func() {
				data := map[string]interface{}{"Baz": 1000.78}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 500.3 (lower interval)", func() {

			Convey("The value should be valid", func() {
				data := map[string]interface{}{"Baz": 500.3}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 2000.7 (upper interval)", func() {

			Convey("The value should be valid", func() {
				data := map[string]interface{}{"Baz": 2000.7}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 500.2 (lower interval)", func() {

			Convey("The value should not be valid", func() {
				data := map[string]interface{}{"Baz": 500.2}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 2000.8 (upper interval)", func() {

			Convey("The value should not be valid", func() {
				data := map[string]interface{}{"Baz": 2000.8}
				result := Validate(data, rules)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// digits

	Convey("Given a validation which has the rule \"digits:4\"", t, func() {

		rule := Rules{"Baz": Field{
			Type:  "Float",
			Name:  "Foo",
			Rules: []string{"digits:4"},
		},
		}

		Convey("Given the value 5000", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Baz": 5000}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 500", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Baz": 500}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 50000", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Baz": 50000}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// max

	Convey("Given a validation which has the rule \"max:100\"", t, func() {

		rule := Rules{"Baz": Field{
			Type:  "Float",
			Name:  "Foo",
			Rules: []string{"max:100"},
		},
		}

		Convey("Given the value 50", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Baz": 50}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 100", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Baz": 100}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 101", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Baz": 101}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// min

	Convey("Given a validation which has the rule \"min:0\"", t, func() {

		rule := Rules{"Baz": Field{
			Type:  "Float",
			Name:  "Foo",
			Rules: []string{"min:0"},
		},
		}

		Convey("Given the value -1", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Baz": -1}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 0", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Baz": 0}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 1", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Baz": 1}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

	})

}
