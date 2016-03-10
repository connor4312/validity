package validity

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// StringGuard is a validator for string types
type StringGuard struct {
	Value string
	Rules []string
}

// Check ensures that the value is ok
func (guard StringGuard) Check() Result {
	result := Result{
		Errors:  []string{},
		IsValid: true,
	}
	for _, rule := range guard.Rules {
		isValid, err := guard.checkRule(rule)
		if err != nil {
			panic(err)
		}
		if !isValid {
			result.Errors = append(result.Errors, "STRING#"+rule)
			result.IsValid = false
			result.Data = guard.Value
		}
	}
	return result
}

func (guard StringGuard) checkRule(fullRule string) (bool, error) {
	rule := fullRule
	parts := strings.SplitN(fullRule, ":", 2)

	if len(parts) != 1 {
		rule = parts[0]
	}

	getArguments := func(part string) []string {
		return strings.SplitN(part, ",", 2)
	}

	switch rule {
	case "regexp":
		if len(parts) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [regexp:exp]")
		}
		exp := parts[1]
		return guard.validateRegexp(exp), nil
	case "between":
		args := getArguments(parts[1])
		if len(args) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [between:a,b]")
		}
		min := args[0]
		max := args[1]
		return guard.validateBetween(min, max), nil
	case "between_strict":
		args := getArguments(parts[1])
		if len(args) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [between_strict:a,b]")
		}
		min := args[0]
		max := args[1]
		return guard.validateBetweenStrict(min, max), nil
	case "max_len":
		if len(parts) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [max_len:value]")
		}
		max := parts[1]
		return guard.validateMaxLen(max), nil
	case "min_len":
		if len(parts) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [min_len:value]")
		}
		min := parts[1]
		return guard.validateMinLen(min), nil
	case "len":
		if len(parts) != 2 {
			return false, errors.New("This rule is incorrect [" + fullRule + "]. The good format is [len:value]")
		}
		len := parts[1]
		return guard.validateLen(len), nil
	}
	return false, errors.New("The guardian STRING does not have the rule [" + rule + "]")
}

func (guard StringGuard) toInt(s string) int {
	out, _ := strconv.ParseInt(s, 10, 64)
	return int(out)
}

// validateRegexp validates a regex
func (guard StringGuard) validateRegexp(exp string) bool {
	expression, _ := regexp.Compile(exp)
	return expression.MatchString(guard.Value)
}

// validateBetween checks if the number: min <= len(number) => max
func (guard StringGuard) validateBetween(min string, max string) bool {
	length := len([]rune(guard.Value))
	return length >= guard.toInt(min) && length <= guard.toInt(max)
}

// validateBetweenStrict checks if the number: min < len(number) > max
func (guard StringGuard) validateBetweenStrict(min string, max string) bool {
	length := len([]rune((guard.Value)))
	return length > guard.toInt(min) && length < guard.toInt(max)
}

// validateMaxLen checks if the number: len(number) <= max
func (guard StringGuard) validateMaxLen(length string) bool {
	return len([]rune(guard.Value)) <= guard.toInt(length)
}

// validateMinLen checks if the number: min <= len(number)
func (guard StringGuard) validateMinLen(length string) bool {
	return len([]rune(guard.Value)) >= guard.toInt(length)
}

// validateLen checks if the real len of item is that number
func (guard StringGuard) validateLen(length string) bool {
	return len([]rune(guard.Value)) == guard.toInt(length)
}
