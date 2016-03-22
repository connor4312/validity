package validity

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// SpecialGuard is a validator for string types
type SpecialGuard struct {
	Value string
	Rules []string
}

// Check ensures that the value is ok
func (guard SpecialGuard) Check() Result {
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
			result.Errors = append(result.Errors, "SPECIAL#"+rule)
			result.IsValid = false
		}
	}
	return result
}

func (guard SpecialGuard) checkRule(rule string) (bool, error) {
	switch rule {
	case "iban":
		return guard.validateIBAN(), nil
	case "cif":
		return guard.validateCIF(), nil
	case "cnp":
		return guard.validateCNP(), nil
	case "shortDate":
		return guard.validateShortDate(), nil
	case "longDate":
		return guard.validateLongDate(), nil
	case "email":
		return guard.validateEmail(), nil
	}
	return false, errors.New("The guardian SPECIAL does not have the rule [" + rule + "]")
}

func (guard SpecialGuard) toInt(s string) int {
	out, _ := strconv.ParseInt(s, 10, 64)

	return int(out)
}

func (guard SpecialGuard) checkRegexp(r string) bool {
	expression, _ := regexp.Compile(r)

	return expression.MatchString(guard.Value)
}

// validateIBAN validates a bank account
// It must NOT have whitespaces
func (guard SpecialGuard) validateIBAN() bool {

	iban := strings.ToUpper(guard.Value)

	if len(iban) < 10 {
		// log.Println("The IBAN must have at least 10 characters")
		return false
	}

	countrycode := iban[0:2]

	// Check the country code and find the country specific format
	bbancountrypatterns := map[string]string{
		// "AL": "\\d{8}[\\dA-Z]{16}",
		// "AD": "\\d{8}[\\dA-Z]{12}",
		// "AT": "\\d{16}",
		// "AZ": "[\\dA-Z]{4}\\d{20}",
		// "BE": "\\d{12}",
		// "BH": "[A-Z]{4}[\\dA-Z]{14}",
		// "BA": "\\d{16}",
		// "BR": "\\d{23}[A-Z][\\dA-Z]",
		// "BG": "[A-Z]{4}\\d{6}[\\dA-Z]{8}",
		// "CR": "\\d{17}",
		// "HR": "\\d{17}",
		// "CY": "\\d{8}[\\dA-Z]{16}",
		// "CZ": "\\d{20}",
		// "DK": "\\d{14}",
		// "DO": "[A-Z]{4}\\d{20}",
		// "EE": "\\d{16}",
		// "FO": "\\d{14}",
		// "FI": "\\d{14}",
		// "FR": "\\d{10}[\\dA-Z]{11}\\d{2}",
		// "GE": "[\\dA-Z]{2}\\d{16}",
		// "DE": "\\d{18}",
		// "GI": "[A-Z]{4}[\\dA-Z]{15}",
		// "GR": "\\d{7}[\\dA-Z]{16}",
		// "GL": "\\d{14}",
		// "GT": "[\\dA-Z]{4}[\\dA-Z]{20}",
		// "HU": "\\d{24}",
		// "IS": "\\d{22}",
		// "IE": "[\\dA-Z]{4}\\d{14}",
		// "IL": "\\d{19}",
		// "IT": "[A-Z]\\d{10}[\\dA-Z]{12}",
		// "KZ": "\\d{3}[\\dA-Z]{13}",
		// "KW": "[A-Z]{4}[\\dA-Z]{22}",
		// "LV": "[A-Z]{4}[\\dA-Z]{13}",
		// "LB": "\\d{4}[\\dA-Z]{20}",
		// "LI": "\\d{5}[\\dA-Z]{12}",
		// "LT": "\\d{16}",
		// "LU": "\\d{3}[\\dA-Z]{13}",
		// "MK": "\\d{3}[\\dA-Z]{10}\\d{2}",
		// "MT": "[A-Z]{4}\\d{5}[\\dA-Z]{18}",
		// "MR": "\\d{23}",
		// "MU": "[A-Z]{4}\\d{19}[A-Z]{3}",
		// "MC": "\\d{10}[\\dA-Z]{11}\\d{2}",
		// "MD": "[\\dA-Z]{2}\\d{18}",
		// "ME": "\\d{18}",
		// "NL": "[A-Z]{4}\\d{10}",
		// "NO": "\\d{11}",
		// "PK": "[\\dA-Z]{4}\\d{16}",
		// "PS": "[\\dA-Z]{4}\\d{21}",
		// "PL": "\\d{24}",
		// "PT": "\\d{21}",
		// "SM": "[A-Z]\\d{10}[\\dA-Z]{12}",
		// "SA": "\\d{2}[\\dA-Z]{18}",
		// "RS": "\\d{18}",
		// "SK": "\\d{20}",
		// "SI": "\\d{15}",
		// "ES": "\\d{20}",
		// "SE": "\\d{20}",
		// "CH": "\\d{5}[\\dA-Z]{12}",
		// "TN": "\\d{20}",
		// "TR": "\\d{5}[\\dA-Z]{17}",
		// "AE": "\\d{3}\\d{16}",
		// "GB": "[A-Z]{4}\\d{14}",
		// "VG": "[\\dA-Z]{4}\\d{16}",
		"RO": "[A-Z]{4}[\\dA-Z]{16}",
	}

	_, isInTheList := bbancountrypatterns[countrycode]

	if !isInTheList {
		// log.Println("The country code is not in the list")
		return false
	}

	// // As new countries will start using IBAN in the
	// // future, we only check if the countrycode is known.
	// // This prevents false negatives, while almost all
	// // false positives introduced by this, will be caught
	// // by the checksum validation below anyway.
	// // Strict checking should return FALSE for unknown
	// // countries.
	// if (typeof bbanpattern !== "undefined") {
	// 	ibanregexp = new RegExp("^[A-Z]{2}\\d{2}" + bbanpattern + "$", "");
	// 	if (!(ibanregexp.test(iban))) {
	// 		return false; // Invalid country specific format
	// 	}
	// }

	// Now check the checksum, first convert to digits
	ibancheck := iban[4:len(iban)] + iban[0:4]
	lenCheck := len(ibancheck)
	leadingZeroes := true
	ibancheckdigits := ""

	for i := 0; i < lenCheck; i++ {
		character := ibancheck[i]
		if character != '0' {
			leadingZeroes = false
		}
		if !leadingZeroes {
			ibancheckdigits += strconv.Itoa(strings.Index("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(character)))
		}
	}

	// Calculate the result of: ibancheckdigits % 97
	cRest := 0
	lenD := len(ibancheckdigits)

	for p := 0; p < lenD; p++ {
		cChar := ibancheckdigits[p]
		cOperator, _ := strconv.Atoi("" + strconv.Itoa(cRest) + "" + string(cChar))
		cRest = cOperator % 97
	}

	return cRest == 1
}

// validateCIF checks the romanian id for company
func (guard SpecialGuard) validateCIF() bool {

	rawCif := guard.Value

	if lenght := len(rawCif); lenght > 10 || lenght < 6 {
		// log.Println("The length must be between 6 and 10 characters")
		return false
	}

	intCif, errInt := strconv.Atoi(rawCif)

	if errInt != nil {
		// log.Println("The CIF must contain only integers")
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

// validateCNP checks the romanian security id - CNP
func (guard SpecialGuard) validateCNP() bool {

	rawCNP := guard.Value

	if len(rawCNP) != 13 {
		// log.Println("The length of CNP is not 13 characters")
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
			// log.Println("The character at position " + strconv.Itoa(i) + "[" + string(rawCNP[i]) + "] is not a digit")
			return false
		}
		bigSum += control[i] * current
		digits[i] = current
	}

	// check last digit
	_, errLastDigit := strconv.Atoi(string(rawCNP[12]))
	if errLastDigit != nil {
		// log.Println("The character at position " + strconv.Itoa(12) + "[" + string(rawCNP[12]) + "] is not a digit")
		return false
	}

	// Sex -  allowed only 1 -> 9
	if digits[0] == 0 {
		// log.Println("Sex can not be 0")
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
		// log.Println("Wrong year: " + strconv.Itoa(year))
		return false
	}

	// Month - allowed only 1 -> 12
	month := digits[3]*10 + digits[4]
	if month < 1 || month > 12 {
		// log.Println("Wrong month: " + strconv.Itoa(month))
		return false
	}

	day := digits[5]*10 + digits[6]

	// check date
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	if int(t.Year()) != year || int(t.Month()) != month || t.Day() != day {
		// log.Println("The date does not exist: " + strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(day))
		return false
	}

	// County - allowed only 1 -> 52
	county := digits[7]*10 + digits[8]
	if county < 1 || county > 52 {
		// log.Println("Wrong county id: " + strconv.Itoa(county))
		return false
	}

	// Number - allowed only 001 --> 999
	number := digits[9]*100 + digits[10]*10 + digits[11]
	if number < 1 || number > 999 {
		// log.Println("Wrong number: " + strconv.Itoa(number))
		return false
	}

	// Check control
	ctrlDigit = bigSum % 11
	if ctrlDigit == 10 {
		ctrlDigit = 1
	}
	return strconv.Itoa(ctrlDigit) == string(rawCNP[12])
}

// validateShortDate validates a date in the format "02.01.2006"
func (guard SpecialGuard) validateShortDate() bool {
	dateFormat := "02.01.2006"
	_, err := time.Parse(dateFormat, guard.Value)
	return err == nil
}

// validateDate validates a date in the format "02.01.2006T15:04:05"
func (guard SpecialGuard) validateLongDate() bool {
	dateFormat := "02.01.2006T15:04:05"
	_, err := time.Parse(dateFormat, guard.Value)
	return err == nil
}

// validateEmail checks if the value is an email
func (guard SpecialGuard) validateEmail() bool {
	length := len([]rune(guard.Value))
	return (length >= 8 && length <= 25 && guard.checkRegexp("^.+\\@.+\\..+$"))
}
