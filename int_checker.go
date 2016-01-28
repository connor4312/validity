package validity

import (
	"math"
	"strconv"
)

// IntValidityChecker is the validtor for int
type IntValidityChecker struct {
	Key   string
	Rules []string
	Item  int64
}

// Converts a string to an integer. That's all there is!
func (v IntValidityChecker) toInt(s string) int64 {
	out, _ := strconv.ParseInt(s, 10, 64)

	return out
}

// Gets the number of digits from the item.
func (v IntValidityChecker) getDigits() int64 {
	log := math.Log10(float64(v.Item))

	return int64(math.Ceil(log))
}

// GetKey returns the key
func (v IntValidityChecker) GetKey() string {
	return v.Key
}

// GetItem returns the item
func (v IntValidityChecker) GetItem() interface{} {
	return v.Item
}

// GetRules returns the rules
func (v IntValidityChecker) GetRules() []string {
	return v.Rules
}

// GetErrors returns the errors
func (v IntValidityChecker) GetErrors() []string {
	return GetCheckerErrors(v.Rules[1:], &v)
}

//----------------------------------------------------------------------------------------------------------------------
// For explanation involving validation rules, checkout the first huge comment in validity.go.
//----------------------------------------------------------------------------------------------------------------------

// ValidateBetweenStrict checks if the number: min < len(number) > max
func (v IntValidityChecker) ValidateBetweenStrict(min string, max string) bool {
	return v.Item > v.toInt(min) && v.Item < v.toInt(max)
}

// ValidateBetween checks if the number: min <= len(number) => max
func (v IntValidityChecker) ValidateBetween(min string, max string) bool {
	return v.Item >= v.toInt(min) && v.Item <= v.toInt(max)
}

// ValidateDigits checks if the number of digits is that number
func (v IntValidityChecker) ValidateDigits(num string) bool {
	return v.getDigits() == v.toInt(num)
}

// ValidateDigitsBetweenStrict checks if the number of digits: min < digits > max
func (v IntValidityChecker) ValidateDigitsBetweenStrict(min string, max string) bool {
	digits := v.getDigits()

	return digits > v.toInt(min) && digits < v.toInt(max)
}

// ValidateDigitsBetween checks if the number of digits: min <= digits => max
func (v IntValidityChecker) ValidateDigitsBetween(min string, max string) bool {
	digits := v.getDigits()

	return digits >= v.toInt(min) && digits <= v.toInt(max)
}

// ValidateMax checks if the number: len(number) <= max
func (v IntValidityChecker) ValidateMax(max string) bool {
	return v.Item <= v.toInt(max)
}

// ValidateMin checks if the number: min <= len(number)
func (v IntValidityChecker) ValidateMin(min string) bool {
	return v.Item >= v.toInt(min)
}
