package validity

import (
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

// Rules is a map of strings to slices of things. The keys of the map should be the field names to validate,
// in the struct or map of input given. The values should be slices of validators to run. For example:
type Rules map[string]Field

// Field represents the field to test
type Field struct {
	Name  string
	Type  string
	Rules []string
}

// Results is returned from validation functions.
type Results struct {
	IsValid bool
	Errors  map[string][]string
	Data    map[string]interface{}
}

// TranslateTo translates the errors into a language and returns a map[string]string
func (v *Results) TranslateTo(language string) map[string][]string {
	var translator Translater
	switch language {
	case "romanian":
		translator = NewRomanianTranslator()
		break
	default:
		panic("This language " + language + " is not supported.")
	}
	return translator.Translate(v)
}

// ExtractMethod returns the method for complex
// For instance: if you have between:8,9 ---> it returns between
// It can be used for older versions which have returned between rather than between:8,9
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

// ValidateStruct converts the struct into a map, then runs ValidateMap() on it. See ValidateMap's documentation for
// usage details.
func ValidateStruct(s interface{}, rules Rules) *Results {
	return ValidateMap(structs.Map(s), rules)
}

//
// // ValidateStructTags returns the validation results
// func ValidateStructTags(s interface{}) *Results {
// 	input := structs.New(s)
// 	rules := Rules{}
// 	data := map[string]interface{}{}
//
// 	fields := input.Fields()
//
// 	for _, field := range fields {
// 		name := field.Name()
// 		val := field.Value()
//
// 		tag := field.Tag("validators")
// 		rules[name] = []string{inferValidationType(val)}
//
// 		if tag != "" {
// 			rules[name] = append(rules[name], strings.Split(tag, " and ")...)
// 		}
//
// 		data[name] = val
// 	}
//
// 	return ValidateMap(data, rules)
// }

// ValidateMap validates a map against a set of rules. "Data" is obviously a map of string keys to mixed type values, while rules
// is an instance of the rules to validate the data against. Returns a pointer to Results
func ValidateMap(data map[string]interface{}, rules Rules) *Results {
	results := new(Results)

	Queue{Data: data, Rules: rules, Results: results}.Run()

	return results
}
