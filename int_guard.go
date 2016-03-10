package validity

import (
	"errors"
	"strconv"
	"strings"
)

// IntGuard is a validator for float types
type IntGuard struct {
	Raw   string
	Value int
	Rules []string
}

// Check ensures that the value is ok
func (guard IntGuard) Check() Result {
	intValue, errInt := strconv.Atoi(guard.Raw)
	if errInt != nil {
		return Result{
			Errors:  []string{"INT"},
			IsValid: false,
			Data:    guard.Value,
		}
	}
	guard.Value = intValue
	return guard.checkRules()
}

func (guard IntGuard) checkRules() Result {
	result := Result{
		Errors: []string{},
	}
	for _, rule := range guard.Rules {
		isValid, err := guard.checkRule(rule)
		if err != nil {
			panic("The guardian INT does not have the rule [" + rule + "]")
		}
		if !isValid {
			result.Errors = append(result.Errors, "INT#"+rule)
			result.IsValid = false
			result.Data = guard.Value
		}
	}
	return result
}

func (guard IntGuard) checkRule(fullRule string) (bool, error) {
	rule := fullRule
	parts := strings.SplitN(fullRule, ":", 2)

	if len(parts) != 1 {
		rule = parts[0]
	}

	getArguments := func(part string) []string {
		return strings.SplitN(part, ",", 2)
	}

	switch rule {
	case "value":
		args := getArguments(parts[1])
		if len(args) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [value:a,b]")
		}
		min := args[0]
		max := args[1]
		return guard.validateValue(min, max), nil
	case "value_strict":
		args := getArguments(parts[1])
		if len(args) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [value_strict:a,b]")
		}
		min := args[0]
		max := args[1]
		return guard.validateValueStrict(min, max), nil

	case "digits_between":
		args := getArguments(parts[1])
		if len(args) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [digits_between:a,b]")
		}
		min := args[0]
		max := args[1]
		return guard.validateValue(min, max), nil

	case "digits_between_strict":
		args := getArguments(parts[1])
		if len(args) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [digits_between_strict:a,b]")
		}
		min := args[0]
		max := args[1]
		return guard.validateValueStrict(min, max), nil

	case "max":
		if len(parts) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [max:value]")
		}
		max := parts[1]
		return guard.validateMax(max), nil
	case "min":
		if len(parts) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [min:value]")
		}
		min := parts[1]
		return guard.validateMin(min), nil
	case "digits":
		if len(parts) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [digits:value]")
		}
		digits := parts[1]
		return guard.validateDigits(digits), nil
	}
	return false, errors.New("No rule such that")
}

// Converts a string to an integer. That's all there is!
func (guard IntGuard) toInt(s string) int {
	out, _ := strconv.Atoi(s)
	return out
}

// Gets the number of digits from the item.
func (guard IntGuard) getDigits() int {

	number := guard.Value
	digits := 0

	for ; number > 0; digits++ {
		number = number / 10
	}

	return digits
}

//----------------------------------------------------------------------------------------------------------------------
// For explanation involving validation rules, checkout the first huge comment in validity.go.
//----------------------------------------------------------------------------------------------------------------------

// validateValue checks if the number: min <= len(number) => max
func (guard IntGuard) validateValue(min string, max string) bool {
	return guard.Value >= guard.toInt(min) && guard.Value <= guard.toInt(max)
}

// validateValueStrict checks if the number: min < len(number) > max
func (guard IntGuard) validateValueStrict(min string, max string) bool {
	return guard.Value > guard.toInt(min) && guard.Value < guard.toInt(max)
}

// validateDigits checks if the number of digits is that number
func (guard IntGuard) validateDigits(num string) bool {
	return guard.getDigits() == guard.toInt(num)
}

// validateDigitsBetweenStrict checks if the number of digits: min < digits > max
func (guard IntGuard) validateDigitsBetweenStrict(min string, max string) bool {
	digits := guard.getDigits()

	return digits > guard.toInt(min) && digits < guard.toInt(max)
}

// validateDigitsBetween checks if the number of digits: min <= digits => max
func (guard IntGuard) validateDigitsBetween(min string, max string) bool {
	digits := guard.getDigits()

	return digits >= guard.toInt(min) && digits <= guard.toInt(max)
}

// validateMax checks if the number: len(number) <= max
func (guard IntGuard) validateMax(max string) bool {
	return guard.Value <= guard.toInt(max)
}

// validateMin checks if the number: min <= len(number)
func (guard IntGuard) validateMin(min string) bool {
	return guard.Value >= guard.toInt(min)
}
