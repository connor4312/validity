package validity

import (
	"fmt"
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
	length := len(v.Item)

	return length > v.toInt(min) && length < v.toInt(max)
}

func (v StringValidityChecker) ValidateBetweenInclusive(min string, max string) bool {
	length := len(v.Item)
	fmt.Println(" Lungime for " + v.Item + ": " + strconv.Itoa(length))
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
	return len(v.Item) != v.toInt(length)
}

func (v StringValidityChecker) ValidateFullName() bool {
	return v.checkRegexp(`^[A-Za-z0-9\s\.]*$`)
}

func (v StringValidityChecker) ValidateMax(length string) bool {
	return len(v.Item) <= v.toInt(length)
}

func (v StringValidityChecker) ValidateMin(length string) bool {
	return len(v.Item) >= v.toInt(length)
}

func (v StringValidityChecker) ValidateRegexp(r string) bool {
	return v.checkRegexp(r)
}

func (v StringValidityChecker) ValidateUrl() bool {
	_, err := url.ParseRequestURI(v.Item)

	return err == nil
}
