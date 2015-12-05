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
func (translator RomanianTranslator) Translate(results *ValidationResults) map[string][]string {
	return translator.work(translator, results)
}

// TranslateRule translates a method into a english human message
func (translator RomanianTranslator) TranslateRule(method string, options string) string {

	message := "Validarea pentru regula [" + method + "] nu a reuşit!"

	switch method {
	case "accepted":
		message = "Acest câmp trebuie să fie 'yes', 'on', true, sau 1. Sunt permise valori numerice sau text."
	case "url", "email", "ipv4", "ipv6", "date", "ip":
		message = "Câmpul trebuie sa fie un " + strings.ToUpper(method) + " valid."
	case "regex":
		message = "Trebuie să se potrivească acestei expresii regulate: " + options
	case "between":
		betweenMessage := translator.getMessageBetween(options)
		message = "Lungimea trebuie să fie între " + betweenMessage + " de caractere (fară intervale)"
	case "between_inclusive":
		betweenMessage := translator.getMessageBetween(options)
		message = "Lungimea trebuie să fie între " + betweenMessage + " de caractere (inclusiv intervalele)"
	case "digits_between":
		betweenMessage := translator.getMessageBetween(options)
		message = "Câmpul trebuie să fie de tip numeric şi trebuie să fie între " + betweenMessage + " cifre (fără intervale)"
	case "digits_between_inclusive":
		betweenMessage := translator.getMessageBetween(options)
		message = "Câmpul trebuie să fie de tip numeric şi trebuie să fie între " + betweenMessage + " cifre (inclusiv intervalele)"
	case "min":
		message = "Lungimea minimă permisă este de " + options
	case "max":
		message = "Lungimea maximă permisă este de " + options
	case "len":
		message = "Trebuie să aibă exact " + options + " cifre"
	case "alpha":
		message = "Acesta trebuie să fie în întregime caractere alfabetice"
	case "alpha_dash":
		message = "Aceasta poate avea caractere alfanumerice, precum și liniuțe de subliniere și cratime."
	case "alpha_num":
		message = "Acesta trebuie să aibă în întregime caractere alfanumerice"
	case "digits":
		message = "Trebuie să fie un număr şi să aibă un number de " + options + " de cifre"
	case "required":
		message = "Câmpul în curs de validare trebuie sa fie prezent. Sunt acceptate toate tipurile. Notă: opţiunile nu funcționează atunci când se încearcă să se valideze structs, deoarece nu este posibil să se știe dacă valorile nule sunt zero, deoarece acestea nu sunt stabilite, sau pentru că ar trebui să fie de fapt de zero"
	case "full_name":
		message = "Câmpul trebuie să conțină caractere alfanumerice, puncte sau spații"
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
