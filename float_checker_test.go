package validity

import (
	"testing"
)

func TestFloatValidateAcceptedPass(t *testing.T) {
	data := TestStruct{Baz: 1}
	rules := ValidationRules{"Baz": []string{"Float", "Accepted"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Float accepted validator does not pass.")
	}
}
func TestFloatValidateAcceptedFail(t *testing.T) {
	data := TestStruct{Baz: 0}
	rules := ValidationRules{"Baz": []string{"Float", "Accepted"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Float accepted validator does not fail.")
	}
}



func TestFloatValidateBetweenPass(t *testing.T) {
	data := TestStruct{Baz: 5}
	rules := ValidationRules{"Baz": []string{"Float", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Float between validator does not pass.")
	}
}
func TestFloatValidateBetweenFailLower(t *testing.T) {
	data := TestStruct{Baz: 1}
	rules := ValidationRules{"Baz": []string{"Float", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Float between validator does not fail on lower.")
	}
}
func TestFloatValidateBetweenFailUpper(t *testing.T) {
	data := TestStruct{Baz: 8}
	rules := ValidationRules{"Baz": []string{"Float", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Float between validator does not fail on upper.")
	}
}



func TestFloatValidateDigitsPass(t *testing.T) {
	data := TestStruct{Baz: 500.5252}
	rules := ValidationRules{"Baz": []string{"Float", "digits:3"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Float digits validator does not pass.")
	}
}
func TestFloatValidateBetweenFail(t *testing.T) {
	data := TestStruct{Baz: 500.5}
	rules := ValidationRules{"Baz": []string{"Float", "digits:4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Float digits validator does not fail.")
	}
}



func TestFloatValidateDigitsBetweenPass(t *testing.T) {
	data := TestStruct{Baz: 500}
	rules := ValidationRules{"Baz": []string{"Float", "digits_between:2,4"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Float between validator does not pass.")
	}
}
func TestFloatValidateDigitsBetweenFailLower(t *testing.T) {
	data := TestStruct{Baz: 5}
	rules := ValidationRules{"Baz": []string{"Float", "between:2,4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Float digits between validator does not fail on lower.")
	}
}
func TestFloatValidateDigitsBetweenFailUpper(t *testing.T) {
	data := TestStruct{Baz: 5000000}
	rules := ValidationRules{"Baz": []string{"Float", "between:2,4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Float digits between validator does not fail on upper.")
	}
}



func TestFloatValidateMaxPass(t *testing.T) {
	data := TestStruct{Baz: 4}
	rules := ValidationRules{"Baz": []string{"Float", "max:5"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Float max validator does not pass.")
	}
}
func TestFloatValidateMaxFail(t *testing.T) {
	data := TestStruct{Baz: 8}
	rules := ValidationRules{"Baz": []string{"Float", "max:5"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Float max validator does not fail.")
	}
}



func TestFloatValidateMinPass(t *testing.T) {
	data := TestStruct{Baz: 8}
	rules := ValidationRules{"Baz": []string{"Float", "min:5"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Float min validator does not pass.")
	}
}
func TestFloatValidateMinFail(t *testing.T) {
	data := TestStruct{Baz: 4}
	rules := ValidationRules{"Baz": []string{"Float", "min:5"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Float min validator does not fail.")
	}
}
