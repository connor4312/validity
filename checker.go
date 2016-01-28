package validity

import (
	"reflect"
	"strings"
)

// Checker is the base interface from which type validators must implement.
type Checker interface {
	// GetErrors is the method which is actually called to run validation. It should return a slice of validation rules
	// which are invalid, or nil if there are no invalid rules.
	GetItem() interface{}
	GetKey() string
	GetRules() []string
	GetErrors() []string
}

// GetCheckerErrors runs the validation! What it does is, for each validation rule in the format "rule:arg1,arg2". Surrounding
// spaces will be trimmed out. It attempts to call a function defined like:
//
//		func ValidateRule(arg1 string, arg2 string) bool { ... }
//
// It must return a boolean value (true if validation passed, false if it did not) and take string arguments.
func GetCheckerErrors(rules []string, instance Checker) []string {
	errors := []string{}

	for _, rule := range rules {
		// First we want to split the rule into its method and arguments parts,
		// so we have a []string{"rule", "arg1,arg2"}.
		parts := strings.SplitN(rule, ":", 2)
		method := snakeToStudly(strings.ToLower(parts[0]))
		// The parameters to call is a list of reflection values.
		params := []reflect.Value{}

		// If we do have arguments (some rules do not require them), the split the arguments part by commas and
		// put their Values into the params array.
		if len(parts) > 1 {
			for _, arg := range strings.Split(parts[1], ",") {
				params = append(params, reflect.ValueOf(strings.Trim(arg, " ")))
			}
		}

		// Finall, call the validator...
		valid := reflect.ValueOf(instance).MethodByName("Validate" + method).Call(params)[0].Bool()

		// and if it is not valid, then we need to store it in the errors...
		if !valid {
			errors = append(errors, rule)
		}
	}

	// And finally return any errors which occured.
	return errors
}
