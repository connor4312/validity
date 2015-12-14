package validity

import (
	"math"
	"strconv"
)

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

func (v FloatValidityChecker) GetKey() string {
	return v.Key
}

func (v FloatValidityChecker) GetItem() interface{} {
	return v.Item
}

func (v FloatValidityChecker) GetRules() []string {
	return v.Rules
}

func (v FloatValidityChecker) GetErrors() []string {
	return GetCheckerErrors(v.Rules[1:], &v)
}

//----------------------------------------------------------------------------------------------------------------------
// For explanation involving validation rules, checkout the first huge comment in validity.go.
//----------------------------------------------------------------------------------------------------------------------

func (v FloatValidityChecker) ValidateAccepted() bool {
	return v.Item > 0
}

func (v FloatValidityChecker) ValidateBetween(min string, max string) bool {
	return v.Item > v.toFloat(min) && v.Item < v.toFloat(max)
}

func (v FloatValidityChecker) ValidateBetweenInclusive(min string, max string) bool {
	return v.Item >= v.toFloat(min) && v.Item <= v.toFloat(max)
}

func (v FloatValidityChecker) ValidateDigits(num string) bool {
	return v.getDigits() == v.toInt(num)
}

func (v FloatValidityChecker) ValidateDigitsBetween(min string, max string) bool {
	digits := v.getDigits()

	return digits > v.toInt(min) && digits < v.toInt(max)
}

func (v FloatValidityChecker) ValidateMax(max string) bool {
	return v.Item < v.toFloat(max)
}

func (v FloatValidityChecker) ValidateMin(min string) bool {
	return v.Item > v.toFloat(min)
}
