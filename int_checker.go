package validity

import (
	"math"
	"strconv"
)

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

func (v IntValidityChecker) GetKey() string {
	return v.Key
}

func (v IntValidityChecker) GetItem() interface{} {
	return v.Item
}

func (v IntValidityChecker) GetRules() []string {
	return v.Rules
}

func (v IntValidityChecker) GetErrors() []string {
	return GetCheckerErrors(v.Rules[1:], &v)
}

//----------------------------------------------------------------------------------------------------------------------
// For explanation involving validation rules, checkout the first huge comment in validity.go.
//----------------------------------------------------------------------------------------------------------------------

func (v IntValidityChecker) ValidateAccepted() bool {
	return v.Item > 0
}

func (v IntValidityChecker) ValidateBetween(min string, max string) bool {
	return v.Item > v.toInt(min) && v.Item < v.toInt(max)
}

func (v IntValidityChecker) ValidateDigits(num string) bool {
	return v.getDigits() == v.toInt(num)
}

func (v IntValidityChecker) ValidateDigitsBetween(min string, max string) bool {
	digits := v.getDigits()

	return digits > v.toInt(min) && digits < v.toInt(max)
}

func (v IntValidityChecker) ValidateDigitsBetweenInclusive(min string, max string) bool {
	digits := v.getDigits()

	return digits >= v.toInt(min) && digits <= v.toInt(max)
}

func (v IntValidityChecker) ValidateMax(max string) bool {
	return v.Item <= v.toInt(max)
}

func (v IntValidityChecker) ValidateMin(min string) bool {
	return v.Item >= v.toInt(min)
}
