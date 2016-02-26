package validity

import (
	"math"
	"strconv"
)

// FloatValidityChecker is a validator for floats
type FloatValidityChecker struct {
	Key   string
	Rules []string
	Item  float64
}

// Converts a string to an float. That's all there is!
func (v FloatValidityChecker) toFloat(s string) float64 {
	out, _ := strconv.ParseFloat(s, 64)

	return out
}

// Converts a string to an integer. That's all there is!
func (v FloatValidityChecker) toInt(s string) int64 {
	out, _ := strconv.ParseInt(s, 10, 64)

	return out
}

// Gets the number of non-decimal digits from the item.
func (v FloatValidityChecker) getDigits() int64 {
	return int64(math.Ceil(math.Log10(v.Item)))
}

// GetKey returns the key
func (v FloatValidityChecker) GetKey() string {
	return v.Key
}

// GetItem returns the item
func (v FloatValidityChecker) GetItem() interface{} {
	return v.Item
}

// GetRules returns the rules
func (v FloatValidityChecker) GetRules() []string {
	return v.Rules
}

// GetErrors returns the errors
func (v FloatValidityChecker) GetErrors() []string {
	return GetCheckerErrors(v.Rules[1:], &v)
}

//----------------------------------------------------------------------------------------------------------------------
// For explanation involving validation rules, checkout the first huge comment in validity.go.
//----------------------------------------------------------------------------------------------------------------------

// ValidateValueStrict checks if the number: min < len(number) > max
func (v FloatValidityChecker) ValidateValueStrict(min string, max string) bool {
	return v.Item > v.toFloat(min) && v.Item < v.toFloat(max)
}

// ValidateValue checks if the number: min <= len(number) => max
func (v FloatValidityChecker) ValidateValue(min string, max string) bool {
	return v.Item >= v.toFloat(min) && v.Item <= v.toFloat(max)
}

// ValidateDigits checks if the number of digits is that number
func (v FloatValidityChecker) ValidateDigits(num string) bool {
	return v.getDigits() == v.toInt(num)
}

// ValidateMax checks if the number: len(number) <= max
func (v FloatValidityChecker) ValidateMax(max string) bool {
	return v.Item <= v.toFloat(max)
}

// ValidateMin checks if the number: min <= len(number)
func (v FloatValidityChecker) ValidateMin(min string) bool {
	return v.Item >= v.toFloat(min)
}
