package validity

import (
	"strings"
	"reflect"
)

// ValidityChecker is the base interface from which type validators must implement.
type ValidityChecker interface {
	// GetErrors is the method which is actually called to run validation. It should return a slice of validation rules
	// which are invalid, or nil if there are no invalid rules.
	GetItem()   interface{}
	GetKey()    string
	GetRules()  []string
	GetErrors() []string
}

// GetErrors runs the validation! What it does is, for each validation rule in the format "rule:arg1,arg2". Surrounding
// spaces will be trimmed out. It attempts to call a function defined like:
//
//		func ValidateRule(arg1 string, arg2 string) bool { ... }
//
// It must return a boolean value (true if validation passed, false if it did not) and take string arguments.
func GetCheckerErrors(rules []string, instance ValidityChecker) []string {
	errors := []string{}

	for _, rule := range rules {
		// First we want to split the rule into its method and arguments parts,
		// so we have a []string{"rule", "arg1,arg2"}.
		parts  := strings.SplitN(rule, ":", 2)
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
			msg := GetHumanMessage(rule)
			errors = append(errors, msg)
		}
	}

	// And finally return any errors which occured.
	return errors
}

func GetHumanMessage(rule string) string {

		parts  := strings.SplitN(rule, ":", 2)
		method 	 := strings.ToLower(parts[0])

		// In case a new validation rule is added
		message := "Validation for rule [" + rule + "] failed!"

		switch method {
		case "accepted":
				message = "It must be 'yes', 'on', true, or 1. Permits numeric and text"
		case "url", "email", "ipv4", "ipv6", "date", "ip":
				message = "It must be a valid " + strings.ToUpper(method) + " type"
		case "regex":
				message = "It must match this patern: " + parts[1]
		case "between":
				betweenMessage := getBetweenMessage(parts[1])
				message = "The length should be between " + betweenMessage + " characters (the boundaries are not allowed)"
		case "between_inclusive":
				betweenMessage := getBetweenMessage(parts[1])
				message = "The length should be between " + betweenMessage + " characters (the boundaries are allowed)"
		case "digits_between":
				betweenMessage := getBetweenMessage(parts[1])
				message = "The field must be a numeric type and has between " + betweenMessage + " digits"
		case "min":
				message = "The minimum length allowed is " + parts[1]
		case "max":
				message = "The maximum length allowed is " + parts[1]
		case "len":
				message = "It must have exactly " + parts[1] + " digits"
		case "alpha":
				message = "It must be entirely alphabetic characters"
		case "alpha_dash":
				message = "It may have alpha-numeric characters, as well as dashes and underscores"
		case "alpha_num":
					message = "It must be entirely alpha-numeric characters"
		case "digits":
					message = "It must be a number and it must have exactly " + parts[1] + " of digits"
		case "required":
					message = "The field under validation must be present. Accepts any type. Note optionality does not function when trying to validate structs, as it isn't possible to know if their zero values are zero because they aren't set, or because they should actually be zero"
		case "full_name":
					message = "The field must contain alpha-numeric characters or dots or spaces"
		}
		return message;
}

func getBetweenMessage (old string) string {
		newString  := strings.Replace(old, ",", " and ", -1)
		return newString
}
