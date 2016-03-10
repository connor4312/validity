package validity

import (
	"errors"
	"log"
	"math"
	"strconv"
	"strings"
)

// FloatGuard is a validator for float types
type FloatGuard struct {
	Raw   string
	Value float64
	Rules []string
}

// Check ensures that the value is ok
func (guard FloatGuard) Check() Result {
	log.Println("The raw is: " + guard.Raw)
	float32Value, errFloat := strconv.ParseFloat(guard.Raw, 32)
	if errFloat != nil {
		log.Println("The value " + guard.Raw + " IS NOT a float")
		return Result{
			Errors:  []string{"FLOAT"},
			IsValid: false,
			Data:    guard.Value,
		}
	}
	log.Println("The value " + guard.Raw + " IS float")
	guard.Value = float64(float32Value)
	return guard.checkRules()
}

func (guard FloatGuard) checkRules() Result {
	result := Result{
		Errors: []string{},
	}
	for _, rule := range guard.Rules {
		isValid, err := guard.checkRule(rule)
		if err != nil {
			panic(err)
		}
		if !isValid {
			result.Errors = append(result.Errors, "FLOAT#"+rule)
			result.IsValid = false
			result.Data = guard.Value
		}
	}
	return result
}

func (guard FloatGuard) checkRule(fullRule string) (bool, error) {
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
	return false, errors.New("The guardian FLOAT does not have the rule [" + rule + "]")
}

// Converts a string to an float. That's all there is!
func (guard FloatGuard) toFloat(s string) float64 {
	out, _ := strconv.ParseFloat(s, 64)
	return out
}

// Converts a string to an integer. That's all there is!
func (guard FloatGuard) toInt(s string) int64 {
	out, _ := strconv.ParseInt(s, 10, 64)
	return out
}

//----------------------------------------------------------------------------------------------------------------------
// For explanation involving validation rules, checkout the first huge comment in validity.go.
//----------------------------------------------------------------------------------------------------------------------

// validateValueStrict checks if the number: min < len(number) > max
func (guard FloatGuard) validateValueStrict(min string, max string) bool {
	return guard.Value > guard.toFloat(min) && guard.Value < guard.toFloat(max)
}

// validateValue checks if the number: min <= len(number) => max
func (guard FloatGuard) validateValue(min string, max string) bool {
	return guard.Value >= guard.toFloat(min) && guard.Value <= guard.toFloat(max)
}

// validateDigits checks if the number of digits is that number
func (guard FloatGuard) validateDigits(num string) bool {

	// Gets the number of non-decimal digits from the item.
	getDigits := func(value float64) int64 {
		return int64(math.Ceil(math.Log10(value)))
	}
	return getDigits(guard.Value) == guard.toInt(num)
}

// validateMax checks if the number: len(number) <= max
func (guard FloatGuard) validateMax(max string) bool {
	return guard.Value <= guard.toFloat(max)
}

// validateMin checks if the number: min <= len(number)
func (guard FloatGuard) validateMin(min string) bool {
	return guard.Value >= guard.toFloat(min)
}
