package validity

import (
	"testing"
)

func TestStringValidateAcceptedPass(t *testing.T) {
	data := TestStruct{Foo: "on"}
	rules := ValidationRules{"Foo": []string{"String", "Accepted"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String accepted validator does not pass.")
	}
}
func TestStringValidateAcceptedFail(t *testing.T) {
	data := TestStruct{Foo: "off"}
	rules := ValidationRules{"Foo": []string{"String", "Accepted"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String accepted validator does not fail.")
	}
}



func TestStringValidateAlphaPass(t *testing.T) {
	data := TestStruct{Foo: "ThisIsAlpha"}
	rules := ValidationRules{"Foo": []string{"String", "Alpha"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String alpha validator does not pass.")
	}
}
func TestStringValidateAlphaFail(t *testing.T) {
	data := TestStruct{Foo: "This isn't alpha!"}
	rules := ValidationRules{"Foo": []string{"String", "Alpha"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String alpha validator does not fail.")
	}
}



func TestStringValidateAlphaDashPass(t *testing.T) {
	data := TestStruct{Foo: "This_Is-Alpha_Dash"}
	rules := ValidationRules{"Foo": []string{"String", "alpha_dash"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String alpha_dash validator does not pass.")
	}
}
func TestStringValidateAlphaDashFail(t *testing.T) {
	data := TestStruct{Foo: "This isn't alpha-dash!"}
	rules := ValidationRules{"Foo": []string{"String", "alpha_dash"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String alpha_dash validator does not fail.")
	}
}



func TestStringValidateAlphaNumPass(t *testing.T) {
	data := TestStruct{Foo: "This1Is2Alpha4Num"}
	rules := ValidationRules{"Foo": []string{"String", "alpha_num"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String alpha_num validator does not pass.")
	}
}
func TestStringValidateAlphaNumFail(t *testing.T) {
	data := TestStruct{Foo: "This isn't alpha-num!!!11!"}
	rules := ValidationRules{"Foo": []string{"String", "alpha_num"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String alpha_num validator does not fail.")
	}
}



func TestStringValidateBetweenPass(t *testing.T) {
	data := TestStruct{Foo: "fooo"}
	rules := ValidationRules{"Foo": []string{"String", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String between validator does not pass.")
	}
}
func TestStringValidateBetweenFailLower(t *testing.T) {
	data := TestStruct{Foo: "fo"}
	rules := ValidationRules{"Foo": []string{"String", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String between validator does not fail on lower.")
	}
}
func TestStringValidateBetweenFailUpper(t *testing.T) {
	data := TestStruct{Foo: "fooooooooooooooo"}
	rules := ValidationRules{"Foo": []string{"String", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String between validator does not fail on upper.")
	}
}



func TestStringValidateDatePass(t *testing.T) {
	data := TestStruct{Foo: "Jan 2, 2006 at 3:04pm (MST)"}
	rules := ValidationRules{"Foo": []string{"String", "Date"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String date validator does not pass.")
	}
}
func TestStringValidateDateFail(t *testing.T) {
	data := TestStruct{Foo: "Jan 2, 20Nope not a date!06 at 3:04pm (MST)"}
	rules := ValidationRules{"Foo": []string{"String", "Date"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String date validator does not fail.")
	}
}



func TestStringValidateEmailPass(t *testing.T) {
	data := TestStruct{Foo: "connor@peet.io"}
	rules := ValidationRules{"Foo": []string{"String", "Email"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String email validator does not pass.")
	}
}
func TestStringValidateEmailFail(t *testing.T) {
	data := TestStruct{Foo: "connor@invalid"}
	rules := ValidationRules{"Foo": []string{"String", "Email"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String email validator does not fail.")
	}
}



func TestStringValidateIpv4Pass(t *testing.T) {
	data := TestStruct{Foo: "127.0.0.1"}
	rules := ValidationRules{"Foo": []string{"String", "Ipv4"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String ipv4 validator does not pass.")
	}
}
func TestStringValidateIpv4Fail(t *testing.T) {
	data := TestStruct{Foo: "127.foo.bar.1"}
	rules := ValidationRules{"Foo": []string{"String", "Ipv4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String ipv4 validator does not fail.")
	}
}



func TestStringValidateIpv6Pass(t *testing.T) {
	data := TestStruct{Foo: "::1"}
	rules := ValidationRules{"Foo": []string{"String", "Ipv6"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String ipv6 validator does not pass.")
	}
}
func TestStringValidateIpv6Fail(t *testing.T) {
	data := TestStruct{Foo: "8888888::1"}
	rules := ValidationRules{"Foo": []string{"String", "Ipv6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String ipv6 validator does not fail.")
	}
}



func TestStringValidateMaxPass(t *testing.T) {
	data := TestStruct{Foo: "foo"}
	rules := ValidationRules{"Foo": []string{"String", "Max:4"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String max validator does not pass.")
	}
}
func TestStringValidateMaxFail(t *testing.T) {
	data := TestStruct{Foo: "fooooooo"}
	rules := ValidationRules{"Foo": []string{"String", "Max:4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String max validator does not fail.")
	}
}



func TestStringValidateMinPass(t *testing.T) {
	data := TestStruct{Foo: "fooooooo"}
	rules := ValidationRules{"Foo": []string{"String", "Min:4"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String min validator does not pass.")
	}
}
func TestStringValidateMinFail(t *testing.T) {
	data := TestStruct{Foo: "foo"}
	rules := ValidationRules{"Foo": []string{"String", "Min:4"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String min validator does not fail.")
	}
}



func TestStringValidateRegexPass(t *testing.T) {
	data := TestStruct{Foo: "FooBar"}
	rules := ValidationRules{"Foo": []string{"String", "Regexp:^Fo.*ar$"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String regexp validator does not pass.")
	}
}
func TestStringValidateRegexFail(t *testing.T) {
	data := TestStruct{Foo: "BarFoo"}
	rules := ValidationRules{"Foo": []string{"String", "Regexp:^Fo.*ar$"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String regexp validator does not fail.")
	}
}



func TestStringValidateUrlPass(t *testing.T) {
	data := TestStruct{Foo: "http://peet.io"}
	rules := ValidationRules{"Foo": []string{"String", "Url"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String url validator does not pass.")
	}
}
func TestStringValidateUrlFail(t *testing.T) {
	data := TestStruct{Foo: "NotAUrl!"}
	rules := ValidationRules{"Foo": []string{"String", "Url"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String url validator does not fail.")
	}
}
