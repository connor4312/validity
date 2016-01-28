package validity

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

// StringValidityChecker is a validator for string types
type StringValidityChecker struct {
	Key   string
	Rules []string
	Item  string
}

// GetKey returnst the key
func (v StringValidityChecker) GetKey() string {
	return v.Key
}

// GetItem returns the item
func (v StringValidityChecker) GetItem() interface{} {
	return v.Item
}

// GetRules returns the rules
func (v StringValidityChecker) GetRules() []string {
	return v.Rules
}

// GetErrors returns the errors
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

//----------------------------------------------------------------------------------------------------------------------
// For explanation involving validation rules, checkout the first huge comment in validity.go.
//----------------------------------------------------------------------------------------------------------------------

// ValidateCif checks the romanian id for company
func (v StringValidityChecker) ValidateCif() bool {

	rawCif := v.Item

	if lenght := len(rawCif); lenght > 10 || lenght < 6 {
		log.Println("The length must be between 6 and 10 characters")
		return false
	}

	intCif, errInt := strconv.Atoi(rawCif)

	if errInt != nil {
		log.Println("The CIF must contain only integers")
		return false
	}

	var (
		controlNumber = 753217532
		controlDigit1 = intCif % 10
		controlDigit2 = 0
	)

	// delete last digit
	intCif = intCif / 10

	t := 0

	for intCif > 0 {
		t += (intCif % 10) * (controlNumber % 10)

		intCif = intCif / 10
		controlNumber = controlNumber / 10
	}

	controlDigit2 = t * 10 % 11

	if controlDigit2 == 10 {
		controlDigit2 = 0
	}

	return controlDigit1 == controlDigit2
}

// ValidateCnp checks the romanian security id - CNP
func (v StringValidityChecker) ValidateCnp() bool {

	rawCNP := v.Item

	if len(rawCNP) != 13 {
		log.Println("The length of CNP is not 13 characters")
		return false
	}

	var (
		bigSum    int
		ctrlDigit int
		digits    = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		year      = 0
		control   = []int{2, 7, 9, 1, 4, 6, 3, 5, 8, 2, 7, 9}
	)

	for i := 0; i < 12; i++ {
		current, errCurrent := strconv.Atoi(string(rawCNP[i]))
		if errCurrent != nil {
			log.Println("The character at position " + strconv.Itoa(i) + "[" + string(rawCNP[i]) + "] is not a digit")
			return false
		}
		bigSum += control[i] * current
		digits[i] = current
	}

	// check last digit
	_, errLastDigit := strconv.Atoi(string(rawCNP[12]))
	if errLastDigit != nil {
		log.Println("The character at position " + strconv.Itoa(12) + "[" + string(rawCNP[12]) + "] is not a digit")
		return false
	}

	// Sex -  allowed only 1 -> 9
	if digits[0] == 0 {
		log.Println("Sex can not be 0")
		return false
	}

	// year
	year = digits[1]*10 + digits[2]

	switch digits[0] {
	case 1, 2:
		year += 1900
		break
	case 3, 4:
		year += 1800
		break
	case 5, 6:
		year += 2000
		break
		// TODO to check
	case 7, 8, 9:
		year += 2000
		now := time.Now()
		if year > now.Year()-14 {
			year -= 100
		}
		break
	}

	if year < 1800 || year > 2099 {
		log.Println("Wrong year: " + strconv.Itoa(year))
		return false
	}

	// Month - allowed only 1 -> 12
	month := digits[3]*10 + digits[4]
	if month < 1 || month > 12 {
		log.Println("Wrong month: " + strconv.Itoa(month))
		return false
	}

	day := digits[5]*10 + digits[6]

	// check date
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	if int(t.Year()) != year || int(t.Month()) != month || t.Day() != day {
		log.Println("The date does not exist: " + strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(day))
		return false
	}

	// County - allowed only 1 -> 52
	county := digits[7]*10 + digits[8]
	if county < 1 || county > 52 {
		log.Println("Wrong county id: " + strconv.Itoa(county))
		return false
	}

	// Number - allowed only 001 --> 999
	number := digits[9]*100 + digits[10]*10 + digits[11]
	if number < 1 || number > 999 {
		log.Println("Wrong number: " + strconv.Itoa(number))
		return false
	}

	// Check control
	ctrlDigit = bigSum % 11
	if ctrlDigit == 10 {
		ctrlDigit = 1
	}
	return strconv.Itoa(ctrlDigit) == string(rawCNP[12])
}

// ValidateDate validates a date in the format "02.01.2006T15:04:05"
func (v StringValidityChecker) ValidateDate() bool {
	dateFormat := "02.01.2006T15:04:05"
	_, err := time.Parse(dateFormat, v.Item)
	return err == nil
}

// ValidateEmail checks if the value is an email
func (v StringValidityChecker) ValidateEmail() bool {
	return v.checkRegexp("^.+\\@.+\\..+$")
}

// ValidateLen checks if the real len of item is that number
func (v StringValidityChecker) ValidateLen(length string) bool {
	return len([]rune(v.Item)) != v.toInt(length)
}

// ValidateMax checks if the number: len(number) <= max
func (v StringValidityChecker) ValidateMax(length string) bool {
	return len([]rune(v.Item)) <= v.toInt(length)
}

// ValidateMin checks if the number: min <= len(number)
func (v StringValidityChecker) ValidateMin(length string) bool {
	return len([]rune(v.Item)) >= v.toInt(length)
}

// ValidateBetweenStrict checks if the number: min < len(number) > max
func (v StringValidityChecker) ValidateBetweenStrict(min string, max string) bool {
	length := len([]rune((v.Item)))

	return length > v.toInt(min) && length < v.toInt(max)
}

// ValidateBetween checks if the number: min <= len(number) => max
func (v StringValidityChecker) ValidateBetween(min string, max string) bool {
	length := len([]rune(v.Item))
	return length >= v.toInt(min) && length <= v.toInt(max)
}

// ValidateRegexp validates a regex
func (v StringValidityChecker) ValidateRegexp(r string) bool {
	return v.checkRegexp(r)
}
