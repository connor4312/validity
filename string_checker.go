package validity

import (
	"net"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

type StringValidityChecker struct {
	Key   string
	Rules []string
	Item  string
}

func (v StringValidityChecker) GetKey() string {
	return v.Key
}

func (v StringValidityChecker) GetItem() interface{} {
	return v.Item
}

func (v StringValidityChecker) GetRules() []string {
	return v.Rules
}

func (v StringValidityChecker) GetErrors() []string {
	return GetCheckerErrors(v.Rules[1:], &v)
}

func (v StringValidityChecker) toInt(s string) int {
	out, _ := strconv.ParseInt(s, 10, 64)

	return int(out)
}

func (v StringValidityChecker) checkRegexp(r string) bool {
	expression, _ := regexp.Compile(r)

	return expression.MatchString(v.Item)
}

func (v StringValidityChecker) parseIP() net.IP {
	return net.ParseIP(v.Item)
}

//----------------------------------------------------------------------------------------------------------------------
// For explanation involving validation rules, checkout the first huge comment in validity.go.
//----------------------------------------------------------------------------------------------------------------------

func (v StringValidityChecker) ValidateCNP(rawCNP string) bool {

	pattern := "^\\d{1}\\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\\d|3[01])(0[1-9]|[1-4]\\d| 5[0-2]|99)\\d{4}$"

	expression, errCompile := regexp.Compile(pattern)

	if errCompile == nil && expression.MatchString(rawCNP) {

		var (
			bigSum    int
			ctrlDigit int
			control   = []int{2, 7, 9, 1, 4, 6, 3, 5, 8, 2, 7, 9}
		)

		for i := 0; i < 12; i++ {
			current, errCurrent := strconv.Atoi(string(rawCNP[i]))
			if errCurrent != nil {
				return false
			}
			bigSum += current * control[i]
		}
		ctrlDigit = bigSum % 11

		if ctrlDigit == 10 {
			ctrlDigit = 1
		}

		return strconv.Itoa(ctrlDigit) == string(rawCNP[12])
	}
	return false

}

func (v StringValidityChecker) ValidateAccepted() bool {
	return v.Item == "yes" || v.Item == "on" || v.Item == "1"
}

func (v StringValidityChecker) ValidateAlpha() bool {
	return v.checkRegexp("^[A-Za-z]*$")
}

func (v StringValidityChecker) ValidateAlphaDash() bool {
	return v.checkRegexp("^[A-Za-z0-9\\-_]*$")
}

func (v StringValidityChecker) ValidateAlphaNum() bool {
	return v.checkRegexp("^[A-Za-z0-9]*$")
}

func (v StringValidityChecker) ValidateBetween(min string, max string) bool {
	length := len([]rune((v.Item)))

	return length > v.toInt(min) && length < v.toInt(max)
}

func (v StringValidityChecker) ValidateBetweenInclusive(min string, max string) bool {
	length := len([]rune(v.Item))
	return length >= v.toInt(min) && length <= v.toInt(max)
}

func (v StringValidityChecker) ValidateDate() bool {
	_, err := time.Parse("Jan 2, 2006 at 3:04pm (MST)", v.Item)

	return err == nil
}

func (v StringValidityChecker) ValidateEmail() bool {
	return v.checkRegexp("^.+\\@.+\\..+$")
}

func (v StringValidityChecker) ValidateIpv4() bool {
	parsed := v.parseIP()

	return parsed != nil && parsed.To4() != nil
}

func (v StringValidityChecker) ValidateIpv6() bool {
	parsed := v.parseIP()

	return parsed != nil && parsed.To16() != nil
}

func (v StringValidityChecker) ValidateIp() bool {
	return v.parseIP() != nil
}

func (v StringValidityChecker) ValidateLen(length string) bool {
	return len([]rune(v.Item)) != v.toInt(length)
}

func (v StringValidityChecker) ValidateFullName() bool {
	return v.checkRegexp(`^[A-Za-z0-9\s\.]*$`)
}

func (v StringValidityChecker) ValidateMax(length string) bool {
	return len([]rune(v.Item)) <= v.toInt(length)
}

func (v StringValidityChecker) ValidateMin(length string) bool {
	return len([]rune(v.Item)) >= v.toInt(length)
}

func (v StringValidityChecker) ValidateRegexp(r string) bool {
	return v.checkRegexp(r)
}

func (v StringValidityChecker) ValidateUrl() bool {
	_, err := url.ParseRequestURI(v.Item)

	return err == nil
}
