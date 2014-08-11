package validity

import (
	"testing"
)

func TestIntValidateAcceptedPass(t *testing.T) {
	data := TestStruct{Bar: 1}
	rules := ValidationRules{"Bar": []string{"Int", "Accepted"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Int accepted validator does not pass.")
	}
}
func TestIntValidateAcceptedFail(t *testing.T) {
	data := TestStruct{Bar: 0}
	rules := ValidationRules{"Bar": []string{"Int", "Accepted"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Int accepted validator does not fail.")
	}
}



func TestIntValidateBetweenPass(t *testing.T) {
	data := TestStruct{Bar: 5}
	rules := ValidationRules{"Bar": []string{"Int", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Int between validator does not pass.")
	}
}
func TestIntValidateBetweenFailLower(t *testing.T) {
	data := TestStruct{Bar: 1}
	rules := ValidationRules{"Bar": []string{"Int", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Int between validator does not fail on lower.")
	}
}
func TestIntValidateBetweenFailUpper(t *testing.T) {
	data := TestStruct{Bar: 8}
	rules := ValidationRules{"Bar": []string{"Int", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Int between validator does not fail on upper.")
	}
}



func TestIntValidateDigitsPass(t *testing.T) {
	data := TestStruct{Bar: 500}
	rules := ValidationRules{"Bar": []string{"Int", "digits:3"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Int digits validator does not pass.")
	}
}
func TestIntValidateBetweenFail(t *testing.T) {
	data := TestStruct{Bar: 500}
	rules := ValidationRules{"Bar": []string{"Int", "digits:4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Int digits validator does not fail.")
	}
}



func TestIntValidateDigitsBetweenPass(t *testing.T) {
	data := TestStruct{Bar: 500}
	rules := ValidationRules{"Bar": []string{"Int", "digits_between:2,4"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Int between validator does not pass.")
	}
}
func TestIntValidateDigitsBetweenFailLower(t *testing.T) {
	data := TestStruct{Bar: 5}
	rules := ValidationRules{"Bar": []string{"Int", "between:2,4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Int digits between validator does not fail on lower.")
	}
}
func TestIntValidateDigitsBetweenFailUpper(t *testing.T) {
	data := TestStruct{Bar: 5000000}
	rules := ValidationRules{"Bar": []string{"Int", "between:2,4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Int digits between validator does not fail on upper.")
	}
}



func TestIntValidateMaxPass(t *testing.T) {
	data := TestStruct{Bar: 4}
	rules := ValidationRules{"Bar": []string{"Int", "max:5"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Int max validator does not pass.")
	}
}
func TestIntValidateMaxFail(t *testing.T) {
	data := TestStruct{Bar: 8}
	rules := ValidationRules{"Bar": []string{"Int", "max:5"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Int max validator does not fail.")
	}
}



func TestIntValidateMinPass(t *testing.T) {
	data := TestStruct{Bar: 8}
	rules := ValidationRules{"Bar": []string{"Int", "min:5"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("Int min validator does not pass.")
	}
}
func TestIntValidateMinFail(t *testing.T) {
	data := TestStruct{Bar: 4}
	rules := ValidationRules{"Bar": []string{"Int", "min:5"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("Int min validator does not fail.")
	}
}
