package validity

// author cristian-sima

import "strings"

// RomanianTranslator shows the messages in Romanian language
type RomanianTranslator struct {
	translator
}

// Translate transforms the error messages from result set to an array of
// human messages
func (translator RomanianTranslator) Translate(results *ValidationResults) map[string][]string {
	return translator.work(translator, results)
}

// TranslateRule translates a method into a english human message
func (translator RomanianTranslator) TranslateRule(method string, options string) string {

	message := "Validarea pentru regula [" + method + "] nu a reusit!"

	switch method {
	case "accepted":
		message = "Trebuie sa fie 'yes', 'on', true, or 1. Permite valori numberice si text"
	case "url", "email", "ipv4", "ipv6", "date", "ip":
		message = "Trebuie sa fie un " + strings.ToUpper(method) + " valid"
	case "regex":
		message = "Trebuie sa se potriveasca acestei expresii regulate: " + options
	case "between":
		betweenMessage := translator.getMessageBetween(options)
		message = "Lungimea trebuie sa fie intre " + betweenMessage + " de caractere (limitele nu sunt permise)"
	case "between_inclusive":
		betweenMessage := translator.getMessageBetween(options)
		message = "Lungimea trebuie sa fie intre " + betweenMessage + " de caractere (limitele sunt permise)"
	case "digits_between":
		betweenMessage := translator.getMessageBetween(options)
		message = "Campul trebuie sa fie de tip numeric si trebuie sa fie intre " + betweenMessage + " cifre"
	case "min":
		message = "Lungimea minimă permisă este de " + options
	case "max":
		message = "Lungimea maximă permisă este de " + options
	case "len":
		message = "Trebuie sa aiba exact " + options + " cifre"
	case "alpha":
		message = "Acesta trebuie să fie în întregime caractere alfabetice"
	case "alpha_dash":
		message = "Aceasta poate avea caractere alfanumerice, precum și liniuțe de subliniere și cratime."
	case "alpha_num":
		message = "Acesta trebuie să aiba în întregime caractere alfanumerice"
	case "digits":
		message = "Trebuie sa fie un numar si sa aiba un number de " + options + " de cifre"
	case "required":
		message = "Câmpul în curs de validare trebuie sa fie prezent. Este acceptat orice tip. Notă:  optiunile nu funcționează atunci când încearcă să valideze structs, deoarece nu este posibil să se știe dacă valorile nule sunt zero, deoarece acestea nu sunt stabilite, sau pentru că ar trebui să fie de fapt de zero"
	case "full_name":
		message = "Câmpul trebuie să conțină caractere alfanumerice sau puncte sau spații"
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
