package validity

import (
	"fmt"
	"log"
	"sync"
)

// Rules is a map of strings to slices of things. The keys of the map should be the field names to validate,
// in the struct or map of input given. The values should be slices of validators to run. For example:
type Rules map[string]Field

// Field represents the field to test
type Field struct {
	Name     string
	Type     string
	Rules    []string
	Optional bool
}

// Error represents an error message
type Error struct {
	Messages []string
	Keys     []string
	Field    Field
}

// Results is returned from validation functions.
type Results struct {
	IsValid bool
	Errors  []*Error
}

// TranslateTo translates the errors into a language and returns a map[string]string
func (v *Results) TranslateTo(language string) {
	var translator Translater
	switch language {
	case "romanian":
		translator = NewRomanianTranslator()
		break
	default:
		panic("This language " + language + " is not supported.")
	}
	translator.Translate(v)
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
}

// Guard ensures that the value is ok
type Guard interface {
	Check() Result
}

// Validate validates the things
func Validate(mapData map[string]interface{}, rulesMap Rules) *Results {

	results := Results{
		IsValid: true,
		Errors:  []*Error{},
	}

	messages := make(chan *Error, len(rulesMap))

	wg := new(sync.WaitGroup)

	wg.Add(len(rulesMap))

	for index, field := range rulesMap {
		rawValue, isPresent := mapData[index]
		go func(field Field) {

			var (
				result Result
				guard  Guard
			)

			if !isPresent {
				if !field.Optional {

					errorObject := Error{
						Keys:  []string{"REQUIRED"},
						Field: field,
					}

					log.Println("From required: " + errorObject.Field.Name)

					messages <- &errorObject
				}

			} else {
				value := fmt.Sprintf("%v", rawValue)

				switch field.Type {
				case "String":
					guard = StringGuard{
						Value: value,
						Rules: field.Rules,
					}
					break
				case "Float":
					guard = FloatGuard{
						Raw:   value,
						Rules: field.Rules,
					}
					break
				case "Int":
					guard = IntGuard{
						Raw:   value,
						Rules: field.Rules,
					}
					break
				case "Special":
					guard = SpecialGuard{
						Value: value,
						Rules: field.Rules,
					}
					break
				}
				result = guard.Check()

				if !result.IsValid {

					errorObject := Error{
						Keys:  result.Errors,
						Field: field,
					}

					log.Println("From validate: " + errorObject.Field.Name)

					messages <- &errorObject
				}
			}
			defer wg.Done()
		}(field)
	}

	wg.Wait()
	close(messages)

	log.Println("The messages are:")

	for errorObject := range messages {

		log.Println(errorObject.Field)
		results.IsValid = false
		results.Errors = append(results.Errors, errorObject)
	}
	return &results
}
