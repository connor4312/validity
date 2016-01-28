package validity

import (
	"testing"
)

func isValidCNP(cnp string) bool {
	data := map[string]interface{}{
		"NID": cnp,
	}
	rules := map[string][]string{
		"NID": []string{"String", "cnp"},
	}
	results := ValidateMap(data, rules)
	return results.IsValid
}

func TestStringInvalidCNP(t *testing.T) {

	list := []string{

		// less than 13 characters
		"00",

		// more than 13 characters
		"00000000000000",

		// not all digits
		"             ",
		"a234567890123",
		"1x34567890123",
		"12y4567890123",
		"123(567890123",
		"1234*67890123",
		"12345-7890123",
		"123456+890123",
		"1234567`90123",
		"12345678!0123",
		"123456789@123",
		"12345678901&3",
		"123456789012=",

		// Digit 1 not 0
		"0000000000000",

		// Wrong month
		// Too small
		"1230000000000",
		// Too big
		"1231300000000",
		"1234500000000",
		"1239900000000",

		// Wrong date
		"1000100000000",
		"1000100000000",
		"1231200000000",
		"3431250000000",
		"8430100000000",
		"6431132000000",
		"5160230000000", // this one does not exist

		// Wrong County
		// 0
		"1930426000000",

		// Too big
		"1930426530000",
		"1930426990000",

		// Number - allowed only 001 --> 999
		"1930426010000",

		// Check controll all posible conbinations
		"1930426450030",
		"1930426450031",
		"1930426450032",
		"1930426450033",
		"1930426450034",
		// "1930426450035", // --> this is good
		"1930426450036",
		"1930426450037",
		"1930426450038",
		"1930426450039",
	}
	for _, cnp := range list {
		if isValidCNP(cnp) {
			t.Errorf("The CNP " + cnp + " should NOT pass the test ")
		}
	}
}

func TestStringValidCNP(t *testing.T) {

	list := []string{
		// list from http://cnp-orange-young.blogspot.ro/
		"1851021345131",
		"1920617149053",
		"1870505168646",
		"1870619152998",
		"1921204325416",
		"1931011351347",
		"1860601214764",
		"1930114152084",
		"1930729213031",
		"1931110257176",
		"1900806182888",
		"1920201181713",
		"1911202393786",
		"1860114313017",
		"1900129321797",
		"1910319357474",
		"1881110136241",
		"1910313389932",
		"1930324128027",
		"1860713267955",
		"1920110059921",
		"1920323137670",
		"1910524257786",
		"1910219353588",
		"1921113049026",
		"1910510171180",
		"1921219189118",
		"1930614296921",
		"1930408137659",
		"1860525171600",
		"1871008113318",
		"1900809164551",
		"1870418305339",
		"1920622093590",
		"1860228422041",
		"1881024199710",
		"1901119289825",
		"1870623151354",
		"1880717053385",
		"1860331281263",
		"1921127239892",
		"1860726084246",
		"1891209392109",
		"1870610429939",
		"1891101155151",
		"1900515188610",
		"1900817347850",
		"1850603113108",
		"1920219031621",
		"1880716319353",
		"1851123072213",
		"1890503074079",
		"1920929194685",
		"1930604396824",
		"1920721211943",
		"1911125017410",
		"1920328331453",
		"1891012055602",
		"1850412422567",
		"1900830012030",
		"1901109372453",
		"1901119063169",
		"1910801238953",
		"1850126336642",
		"1870619334654",
		"1880930146221",
		"1861019351849",
		"1930502103999",
		"1930425361934",
		"1871108269030",
		"1860622273209",
		"1920723362204",
		"1880919016077",
		"1860627124609",
		"1930210191367",
		"1920306379899",
		"1860621221169",
		"1890514289320",
		"1920421395061",
		"1930426272907",
		"1861224076805",
		"1850204207530",
		"1930726241832",
		"1920229184105",
		"1910318245177",
		"1871225154992",
		"1890425099446",

		"1920918187490",
		"1911017023863",
		"1900514401426",
		"1860117328230",
		"1861019339785",
		"1870119334824",
		"1880316079025",
		"1910216018631",
		"1920425405811",
		"1920904145205",
		"1881201016278",
		"1920731125084",
		"1901003395568",

		"1590512521591",
		"1660418230010",
		"1891119450065",
		"1920310430037",
		"1920504521696",
		"1930426450035",
		"2770926521590",
		"2830111410090",
		"2880330521699",
		"2900117521690",
		"2910825522143",
	}
	for _, cnp := range list {
		if !isValidCNP(cnp) {
			t.Errorf("The CNP " + cnp + " should PASS the test ")
		}
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

func TestStringValidateBetweenInclusivePass(t *testing.T) {
	data := TestStruct{Foo: "fooo"}
	rules := ValidationRules{"Foo": []string{"String", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String between validator does not pass.")
	}
}

func TestStringValidateBetweenInclusiveFailLower(t *testing.T) {
	data := TestStruct{Foo: "fo"}
	rules := ValidationRules{"Foo": []string{"String", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String between validator does not fail on lower.")
	}
}

func TestStringValidateBetweenInclusiveFailUpper(t *testing.T) {
	data := TestStruct{Foo: "fooooooooooooooo"}
	rules := ValidationRules{"Foo": []string{"String", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if results.IsValid {
		t.Errorf("String between validator does not fail on upper.")
	}
}

func TestStringValidateBetweenInclusivePassLowerB(t *testing.T) {
	data := TestStruct{Foo: "foo"}
	rules := ValidationRules{"Foo": []string{"String", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String between validator does not pass.")
	}
}

func TestStringValidateBetweenInclusivePassUpperB(t *testing.T) {
	data := TestStruct{Foo: "f23456"}
	rules := ValidationRules{"Foo": []string{"String", "between:3,6"}}

	results := ValidateStruct(data, rules)
	if !results.IsValid {
		t.Errorf("String between validator does not pass.")
	}
}

func TestStringValidateDatePass(t *testing.T) {
	data := TestStruct{Foo: "28.01.2016T14:03:15"}
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
