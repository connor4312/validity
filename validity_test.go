package validity

import (
	"testing"
)

type TestStruct struct {
	Foo string
	Bar int
	Baz float32
}

func TestValidatesMap(t *testing.T) {
	var v interface {}
	v = "42"

	data := make(map[string]interface{})
	data["foo"] = v

	rules := ValidationRules{"foo": []string{"Int"}}

	results := ValidateMap(data, rules)
	if !results.IsValid {
		t.Errorf("Does not validate a basic map of data!")
	}
}

func TestValidatesStruct(t *testing.T) {
	data := TestStruct{Foo: "42"}
	rules := ValidationRules{"Foo": []string{"Int"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Does not validate a basic struct of data!")
	}
}

func TestHandlesBasicTypeConversions(t *testing.T) {
	data  := TestStruct{Foo: "42", Bar: 55, Baz: 12.34}
	rules := ValidationRules{"Foo": []string{"Int"}, "Bar": []string{"String"}, "Baz": []string{"Float"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Doesn't handle basic type conversions! Errors: %s", results.Errors)
	}
}

func TestHandlesInvalidTypeConversions(t *testing.T) {
	data  := TestStruct{Foo: "Not A Number!", Bar: 1234, Baz: 12.34}
	rules := ValidationRules{"Foo": []string{"Int"}, "Bar": []string{"String"}, "Baz": []string{"Float"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Validator thinks that invalid numbers are numbers!")
	}
}


func TestReturnsProperResultsOnTypeFail(t *testing.T) {
	data  := TestStruct{Foo: "Not A Number!"}
	rules := ValidationRules{"Foo": []string{"Int"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Validator did not return a correct isValid.")
	}
	if results.Errors["Foo"][0] != "Int" {
		t.Errorf("Validator did not return the failing key. Returned instead: %s", results.Errors)
	}
	if _, exists := results.Data["Foo"]; exists {
		t.Errorf("Validator should not return data which failed.")
	}
}


func TestReturnsProperResultsOnFailedValidators(t *testing.T) {
	data  := TestStruct{Foo: "42"}
	rules := ValidationRules{"Foo": []string{"Int", "Min:50", "Digits:3", "Max: 60"}}

	results := ValidateStruct(data, rules)

	if results.Errors["Foo"][0] != "Min" || results.Errors["Foo"][1] != "Digits" {
		t.Errorf("Validator did not return the failing rules. Wanted a []string{\"Min\", \"Digits\"} Returned instead: %s", results.Errors["Foo"])
	}
	if _, exists := results.Data["Foo"]; exists {
		t.Errorf("Validator should not return data which failed.")
	}
}


func TestReturnsProperResultsOnSuccess(t *testing.T) {
	data  := TestStruct{Foo: "42"}
	rules := ValidationRules{"Foo": []string{"Int", "Min:40", "Digits:2", "Max: 60"}}

	results := ValidateStruct(data, rules)

	if len(results.Errors["Foo"]) != 0 {
		t.Errorf("Errors should have been zeron on success. Instead: Returned instead: %s", results.Errors["Foo"])
	}
	if _, exists := results.Data["Foo"]; !exists {
		t.Errorf("Validator should return data which which passed.")
	}
}
