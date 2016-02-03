package validity

import "strings"

// EnglishTranslator shows the messages in English
type EnglishTranslator struct {
	translator
}

// Translate transforms the error messages from result set to an array of
// human messages
func (translator EnglishTranslator) Translate(results *Results) map[string][]string {
	return translator.work(translator, results)
}

// TranslateRule translates a method into a english human message
func (translator EnglishTranslator) TranslateRule(method string, options string) string {
	// In case a new validation rule is added
	message := "Validation for rule [" + method + "] failed!"

	switch method {
	case "accepted":
		message = "It must be 'yes', 'on', true, or 1. Permits numeric and text"
	case "url", "email", "ipv4", "ipv6", "date", "ip":
		message = "It must be a valid " + strings.ToUpper(method) + " type"
	case "regex":
		message = "It must match this patern: " + options
	case "between":
		betweenMessage := translator.getMessageBetween(options)
		message = "The length should be between " + betweenMessage + " characters (the boundaries are not allowed)"
	case "between_inclusive":
		betweenMessage := translator.getMessageBetween(options)
		message = "The length should be between " + betweenMessage + " characters (the boundaries are allowed)"
	case "digits_between":
		betweenMessage := translator.getMessageBetween(options)
		message = "The field must be a numeric type and has between " + betweenMessage + " digits"
	case "min":
		message = "The minimum length allowed is " + options
	case "max":
		message = "The maximum length allowed is " + options
	case "len":
		message = "It must have exactly " + options + " digits"
	case "alpha":
		message = "It must be entirely alphabetic characters"
	case "alpha_dash":
		message = "It may have alpha-numeric characters, as well as dashes and underscores"
	case "alpha_num":
		message = "It must be entirely alpha-numeric characters"
	case "digits":
		message = "It must be a number and it must have exactly " + options + " of digits"
	case "required":
		message = "The field under validation must be present. Accepts any type. Note optionality does not function when trying to validate structs, as it isn't possible to know if their zero values are zero because they aren't set, or because they should actually be zero"
	case "full_name":
		message = "The field must contain alpha-numeric characters or dots or spaces"
	}
	return message
}

// NewEnglishTranslator creates a new English translator
func NewEnglishTranslator() Translater {
	return EnglishTranslator{
		translator: translator{
			and: "and",
		},
	}
}
