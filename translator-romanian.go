package validity

// author cristian-sima

import "strings"

// RomanianTranslator shows the messages in Romanian language
type RomanianTranslator struct {
	translator
}

//
// Translate transforms the error messages from result set to an array of
// human messages
func (translator RomanianTranslator) Translate(results *Results) map[string][]string {
	return translator.work(translator, results)
}

// TranslateRule translates a method into a english human message
func (translator RomanianTranslator) TranslateRule(method string, options string) string {

	message := "Validarea pentru regula [" + method + "] nu a reuşit!"

	switch method {
	// Float
	case "float":
		message = "Câmpul trebuie să fie un număr real"
	case "value":
		betweenMessage := translator.getMessageBetween(options)
		message = "Valoarea câmpului trebuie să fie între " + betweenMessage + " (inclusiv intervalele)"
	case "value_strict":
		betweenMessage := translator.getMessageBetween(options)
		message = "Valoarea câmpului trebuie să fie între " + betweenMessage + " (intervalele nu sunt acceptate)"

		// Int
	case "int":
		message = "Câmpul trebuie să fie un număr întreg"

		// String
	case "between":
		betweenMessage := translator.getMessageBetween(options)
		message = "Trebuie să conțină între " + betweenMessage + " de caractere (inclusiv intervale)"
	case "between_strict":
		betweenMessage := translator.getMessageBetween(options)
		message = "Trebuie să conțină între " + betweenMessage + " de caractere (fară intervalele)"

		// Shared

	case "digits":
		message = "Trebuie să aibă " + options + " cifre"
	case "digits_between":
		betweenMessage := translator.getMessageBetween(options)
		message = "Trebuie să fie de tip numeric şi trebuie să fie între " + betweenMessage + " cifre (inclusiv intervale)"
	case "digits_between_strict":
		betweenMessage := translator.getMessageBetween(options)
		message = "Trebuie să fie de tip numeric şi trebuie să fie între " + betweenMessage + " cifre (fără intervalele)"
	case "len":
		message = "Trebuie să aibă exact " + options + " caractere"
	case "min":
		message = "Lungimea minimă permisă este de " + options
	case "max":
		message = "Lungimea maximă permisă este de " + options

		// Special
	case "url", "email":
		message = "Câmpul trebuie sa fie un " + strings.ToUpper(method) + " valid."
	case "regex":
		message = "Trebuie să se potrivească acestei expresii regulate: " + options
	case "date":
		message = "Trebuie să fie o dată calendaristică lungă. De exemplu: 02.01.2006T15:04:05"
	case "short_date":
		message = "Trebuie să fie o dată calendaristică scurtă. De exemplu: 02.01.2006"
	case "cnp":
		message = "Trebuie să fie un cod numeric personal (CNP) valid"
	case "cif":
		message = "Trebuie să fie un cod de identificare fiscală valid (CIF)"
	case "iban":
		message = "Trebuie să fie un cont bancar valid (IBAN)"
	}
	return message
}

// NewRomanianTranslator creates a new romanian translator
func NewRomanianTranslator() Translater {
	return RomanianTranslator{
		translator: translator{
			and: "și",
		},
	}
}
