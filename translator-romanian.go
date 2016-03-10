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

	generalMessage := "There is no translation rule for [" + method + ":" + options + "]"

	getFloatMessage := func(rule string) string {
		switch rule {
		case "value":
			betweenMessage := translator.getMessageBetween(options)
			return "Trebuie să fie un număr real între " + betweenMessage + " (inclusiv intervalele)"
		case "value_strict":
			betweenMessage := translator.getMessageBetween(options)
			return "Trebuie să fie un număr real între " + betweenMessage + " (intervalele nu sunt acceptate)"
		case "min":
			return "Trebuie să fie un număr real mai mare sau egal cu " + options
		case "max":
			return "Trebuie să fie un număr real mai mic sau egal cu " + options
		case "digits":
			return "Trebuie să fie un număr real și să aibă exact " + options + " cifre"
		}
		return generalMessage
	}
	getIntMessage := func(rule string) string {
		switch rule {
		case "value":
			betweenMessage := translator.getMessageBetween(options)
			return "Trebuie să fie un număr întreg între " + betweenMessage + " (inclusiv intervalele)"
		case "value_strict":
			betweenMessage := translator.getMessageBetween(options)
			return "Trebuie să fie un număr întreg între " + betweenMessage + " (intervalele nu sunt acceptate)"
		case "digits_between":
			betweenMessage := translator.getMessageBetween(options)
			return "Trebuie să fie un număr întreg care să aibă între " + betweenMessage + " cifre (inclusiv intervalele)"
		case "digits_between_strict":
			betweenMessage := translator.getMessageBetween(options)
			return "Trebuie să fie un număr întreg care să aibă între " + betweenMessage + " cifre (intervalele nu sunt acceptate)"
		case "min":
			return "Trebuie să fie un număr întreg mai mare sau egal cu " + options
		case "max":
			return "Trebuie să fie un număr întreg mai mic sau egal cu " + options
		case "digits":
			return "Trebuie să fie un număr întreg și să aibă exact " + options + " cifre"
		}
		return generalMessage
	}
	getStringMessage := func(rule string) string {
		switch rule {
		case "regex":
			return "Trebuie să se potrivească acestei expresii regulate: " + options
		case "between":
			betweenMessage := translator.getMessageBetween(options)
			return "Trebuie să conțină între " + betweenMessage + " caractere (inclusiv intervalele)"
		case "between_strict":
			betweenMessage := translator.getMessageBetween(options)
			return "Trebuie să conțină între " + betweenMessage + " caractere (intervalele nu sunt acceptate)"
		case "len_min":
			return "Trebuie să conțină cel puțin " + options + " caractere "
		case "len_max":
			return "Trebuie să conțină cel puțin " + options + " caractere "
		case "len":
			return "Trebuie să conțină exact " + options + " caractere"
		}
		return generalMessage
	}
	getSpecialMessage := func(rule string) string {
		switch rule {
		case "email":
			return "Trebuie să fie o adresă de e-mail validă"
		case "longDate":
			return "Trebuie să fie o dată calendaristică lungă. De exemplu: 02.01.2006T15:04:05"
		case "shortDate":
			return "Trebuie să fie o dată calendaristică scurtă. De exemplu: 02.01.2006"
		case "cnp":
			return "Trebuie să fie un cod numeric personal (CNP) valid"
		case "cif":
			return "Trebuie să fie un cod de identificare fiscală valid (CIF)"
		case "iban":
			return "Trebuie să fie un cont bancar valid (IBAN)"
		}
		return generalMessage
	}

	parts := strings.SplitN(method, "#", 2)

	if len(parts) == 1 {
		switch method {
		case "FLOAT":
			return "Trebuie să fie un număr real (ex. 1,45)"
		case "INT":
			return "Trebuie să fie un număr întreg (ex. 563)"
		}
	}
	rule := parts[1]
	switch parts[0] {
	case "FLOAT":
		return getFloatMessage(rule)
	case "INT":
		return getIntMessage(rule)
	case "STRING":
		return getStringMessage(rule)
	case "SPECIAL":
		return getSpecialMessage(rule)
	}
	return generalMessage
}

// NewRomanianTranslator creates a new romanian translator
func NewRomanianTranslator() Translater {
	return RomanianTranslator{
		translator: translator{
			and: "și",
		},
	}
}
