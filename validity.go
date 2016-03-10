package validity

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

// ValidateMap validates a map against a set of rules. "Data" is obviously a map of string keys to mixed type values, while rules
// is an instance of the rules to validate the data against. Returns a pointer to Results
func ValidateMap(data map[string]interface{}, rules Rules) *Results {
	return Validate(data, rules)
}

// Result is returned by a guard
type Result struct {
	IsValid bool
	Errors  []string
	Data    interface{}
}

// Guard ensures that the value is ok
type Guard interface {
	Check() Result
}

// Validate validates the things
func Validate(mapData map[string]interface{}, rulesMap Rules) *Results {

	results := Results{
		IsValid: true,
		Errors:  map[string][]string{},
		Data:    map[string]interface{}{},
	}

	for index, field := range rulesMap {

		var (
			result Result
			guard  Guard
		)

		value := mapData[index]
		key := index

		switch field.Type {
		case "String":
			guard = StringGuard{
				Value: value.(string),
				Rules: field.Rules,
			}
			break
		case "Float":
			guard = FloatGuard{
				Raw:   value.(string),
				Rules: field.Rules,
			}
			break
		case "Int":
			guard = IntGuard{
				Raw:   value.(string),
				Rules: field.Rules,
			}
			break
		case "Special":
			guard = SpecialGuard{
				Value: value.(string),
				Rules: field.Rules,
			}
			break
		}
		result = guard.Check()

		if !result.IsValid {
			results.IsValid = false
		}

		results.Data[key] = result.Data
		results.Errors[key] = result.Errors
	}
	return &results
}
