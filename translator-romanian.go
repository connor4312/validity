package validity

// author cristian-sima

// RomanianTranslator shows the messages in Romanian language
type RomanianTranslator struct {
	Translator
}

// NewRomanianTranslator creates a new romanian translator
func NewRomanianTranslator() Translater {
	return RomanianTranslator{
		Translator: Translator{
			floatT: floatT{
				floatNumber: "un număr real (ex. 1,45)",
				value:       "un număr real între %s (inclusiv intervalele)",
				valueStrict: "un număr real între %s (intervalele nu sunt acceptate)",
				min:         "un număr real mai mare sau egal cu %s",
				max:         "un număr real mai mic sau egal cu %s",
				digits:      "un număr real și să aibă exact %s cifre",
			},
			intT: intT{
				intNumber:           "un număr întreg (ex. 563)",
				value:               "un număr întreg între %s (inclusiv intervalele)",
				valueStrict:         "un număr întreg între %s (intervalele nu sunt acceptate)",
				digitsBetween:       "un număr întreg care să aibă între %s cifre (inclusiv intervalele)",
				digitsBetweenStrict: "un număr întreg care să aibă între %s cifre (intervalele nu sunt acceptate)",
				min:                 "un număr întreg mai mare sau egal cu %s",
				max:                 "un număr întreg mai mic sau egal cu %s",
				digits:              "un număr întreg și să aibă exact %s cifre",
			},
			stringT: stringT{
				regexp:        "Trebuie să se potrivească acestei expresii regulate: %s",
				between:       "Trebuie să conțină între %s caractere (inclusiv intervalele)",
				betweenStrict: "Trebuie să conțină între %s caractere (intervalele nu sunt acceptate)",
				lenMin:        "Trebuie să conțină cel puțin %s caractere ",
				lenMax:        "Trebuie să conțină cel puțin %s caractere ",
				len:           "Trebuie să conțină exact %s caractere",
			},
			specialT: specialT{
				iban:      "un cont bancar valid (IBAN)",
				cif:       "un cod de identificare fiscală valid (CIF)",
				cnp:       "un cod numeric personal (CNP) valid",
				shortDate: "o dată calendaristică scurtă. De exemplu: 02.01.2006",
				longDate:  "o dată calendaristică lungă. De exemplu: 02.01.2006T15:04:05",
				email:     "o adresă de e-mail validă",
			},
			itMustBe: "Trebuie să fie",
			and:      "și",
		},
	}
}
