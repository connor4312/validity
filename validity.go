package validity

import (
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

// ValidationRules is a map of strings to slices of things. The keys of the map should be the field names to validate,
// in the struct or map of input given. The values should be slices of validators to run. For example:
//
//		rules := ValidationRules{"username": []string{"String", "required", "between: 4, 30"}}
//
// ... would ensure the "username" is present and between four and 30 characters long. The first element of the map
// MUST be a value of the type to convert to. Any numeric or string type is valid. If the value cannot be
// converted to the given type, then it fails validation. The available types are: Int, String, Float.
//
// Possible rules include:
//
//		accepted	   		The field under validation must be "yes", "on", true, or 1.
// 							 	Permits numeric and string types.
// 		alpha      			The field under validation must be entirely alphabetic characters. Permits string types.
//		alpha_dash 			The field under validation may have alpha-numeric characters, as
// 								well as dashes and underscores. Permits string types.
//		alpha_num  			The field under validation must be entirely alpha-numeric characters. Permits string types.
//      between:,a,b  		The field under validation must be between "a" and "b" characters long, or between
// 								the values a and b (if numeric). Permits string and numeric types.
//      between_inclusive:,a,b  		The field under validation must be between "a" and "b" (inclusive) characters long, or between
// 								the values a and b inclusive a and b (if numeric). Permits string and numeric types.
//todo: same:key,v   	  	The field under validation must be equal to another field. Accepts any comparable types.
//		date            	The field under validation must parse to a date. Accepts string types.
//todo: different:key   	The field under validation must not equal the other given
// 							 	field. Accepts any comparable types.
//		digits:num			The field under validation must have exactly `num` of digits. Accepts numeric types.
// 		digits_between:a,b	The field under validation must have between a and b digits. Accepts numeric types.
//		email				The field under validation must be an email.
//		ip					The field under validation must be an IP, either ipv4 or ipv6. Accepts string types.
//		ipv4				The field under validation must be in IPv4 format. Accepts string types.
//		ipv6				The field under validation must be in IPv6 format. Accepts string types.
//		len:num				The field under validation must be be `num` characters long. Accepts string types.
//		max				    The field under validation must be equal to or shorter than "a" (if a string), or
// 								 equal to or smaller than "a" (if numeric). Accepts string and numeric types.
//		min				    The field under validation must be equal to or longer than "a" (if a string), or
// 								equal to or greater than "a" (if numeric). Accepts string and numeric types.
//		regex:pattern		The field under validation must match the given pattern. Accepts string types.
//		required			The field under validation must be present. Accepts any type. Note optionality does not
//								function when trying to validate structs, as it isn't possible to know if their zero
//								values are zero because they aren't set, or because they should actually be zero.
//		required_if,key,v	The field under validation is required only if another field
// 								equals the given value. Accepts any type.
//		required_w,key...	The field under validation must be present if any of the other fields are present. Accepts
//								any type.
//		required_wo,key...	The field under validation must be present if any of the other fields is not present.
//								 Accepts any type.
//		url              	The field under validation must be a URL. Accepts string types.
//
type ValidationRules map[string][]string

//
// // Returns a map of the invalid value keys, to sentences.
// func (v *ValidationResults) Translate(t Translator) map[string]string {
// 	t.Translate(v)
// }

// This struct is returned from validation functions.
type ValidationResults struct {
	// Indicates whether the data under validation has passed the set of rules.
	IsValid bool
	// This is a map of strings to slices of strings. Its keys will be any validation fields which had an error, and
	// the values will be the rules which failed.
	Errors map[string][]string
	// The results is a map of everything after validation. This will be the same data, excluding extraneous values, and
	// values which did not passed validation. They will also be converted to the correct types. Integers will be of
	// type int64, Floats of float64, and Strings of string.
	//
	// The reason that values which did not pass validation are not returned, is because it is not possible to know
	// their types without reflecting them - validation can fail if a value is not able to be converted to a type.
	// This can lead to pitfalls - assuming a value is a of type - not to mention extra work on behalf of the
	// programmer.
	Data map[string]interface{}
}

// TranslateTo translates the errors into a language and returns a map[string]string
func (v *ValidationResults) TranslateTo(language string) map[string][]string {
	var translator Translater
	switch language {
	case "english":
		translator = NewEnglishTranslator()
		break
	case "romanian":
		translator = NewRomanianTranslator()
		break
	default:
		panic("This language " + language + " is not supported.")
	}
	return translator.Translate(v)
}

// ExtractMethod returns the method for complex
// For instance: if you have between_inclusive:8,9 ---> it returns between_inclusive
// It can be used for older versions which have returned between_inclusive rather than between_inclusive:8,9
func ExtractMethod(element string) string {
	parts := strings.SplitN(element, ":", 2)
	return parts[0]
}

func inferValidationType(t interface{}) string {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		return "Int"
	case reflect.Float32, reflect.Float64:
		return "Float"
	default:
		return "String"
	}
}

// This function converts the struct into a map, then runs ValidateMap() on it. See ValidateMap's documentation for
// usage details.
func ValidateStruct(s interface{}, rules ValidationRules) *ValidationResults {
	return ValidateMap(structs.Map(s), rules)
}

func ValidateStructTags(s interface{}) *ValidationResults {
	input := structs.New(s)
	rules := ValidationRules{}
	data := map[string]interface{}{}

	fields := input.Fields()

	for _, field := range fields {
		name := field.Name()
		val := field.Value()

		tag := field.Tag("validators")
		rules[name] = []string{inferValidationType(val)}

		if tag != "" {
			rules[name] = append(rules[name], strings.Split(tag, " and ")...)
		}

		data[name] = val
	}

	return ValidateMap(data, rules)
}

// Validates a map against a set of rules. "Data" is obviously a map of string keys to mixed type values, while rules
// is an instance of the rules to validate the data against. Returns a pointer to ValidationResults
func ValidateMap(data map[string]interface{}, rules ValidationRules) *ValidationResults {
	results := new(ValidationResults)

	ValidityQueue{Data: data, Rules: rules, Results: results}.Run()

	return results
}
