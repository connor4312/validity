package validity

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {

	// regex

	Convey("Given the validation rule \"regex:^Fo.*ar$\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "String",
			Name:  "Foo",
			Rules: []string{"regexp:^Fo.*ar$"},
		}}

		Convey("Given the value \"FooBar\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "FooBar"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"BarFoo\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "BarFoo"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// between

	Convey("Given the validation rule \"between:1,5\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "String",
			Name:  "Foo",
			Rules: []string{"between:1,5"},
		},
		}

		Convey("Given the value \"\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": ""}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"a\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "a"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"aa\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "aa"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"aaaaa\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "aaaaa"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"aaaaaa'", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "aaaaaa"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// between_strict

	Convey("Given the validation rule \"between_strict:1,5\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "String",
			Name:  "Foo",
			Rules: []string{"between_strict:1,5"},
		},
		}

		Convey("Given the value \"\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": ""}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"a\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "a"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"aa\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "aa"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"aaaaa\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "aaaaa"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"aaaaaa'", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "aaaaaa"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// max_len

	Convey("Given the validation rule \"max_len:2\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "String",
			Name:  "Foo",
			Rules: []string{"max_len:2"},
		},
		}

		Convey("Given the value \"a\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "a"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"ab\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "ab"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"abc\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "abc"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// min_len

	Convey("Given the validation rule \"min_len:2\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "String",
			Name:  "Foo",
			Rules: []string{"min_len:2"},
		},
		}

		Convey("Given the value \"aaa\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "aaa"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"ab\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "ab"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"a\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "a"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// len

	// min_len

	Convey("Given the validation rule \"len:2\"", t, func() {

		rule := Rules{"Foo": Field{
			Type:  "String",
			Name:  "Foo",
			Rules: []string{"len:2"},
		},
		}

		Convey("Given the value \"a\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "a"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value \"ab\"", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Foo": "ab"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value \"abc\"", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Foo": "abc"}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})
}
