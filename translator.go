package validity

import "strings"

// Translater ... defines the methods for a Language translato. The translator
// is the object which transoforms the error codes to human messages
// It may transform just a particular rule or the entire map
type Translater interface {
	TranslateRule(method string, options string) string
	Translate(results *Results) map[string][]string
}

// Translator is the basic type of a translator
// It must be inherited
type translator struct {
	and string
}

//
func (translator translator) getMessageBetween(old string) string {
	newString := strings.Replace(old, ",", " "+translator.and+" ", -1)
	return newString
}

// So, it takes a Translator and the result
//
// It iterates through the errors
// It splits the error message
//
// And then it passes to the child translator which provides the translation
//
func (translator translator) work(child Translater, results *Results) map[string][]string {
	humanMessages := map[string][]string{}
	for element, fieldErrors := range results.Errors {
		for _, fullMethod := range fieldErrors {
			parts := strings.SplitN(fullMethod, ":", 2)
			method := strings.ToLower(parts[0])
			options := ""
			if len(parts) == 2 {
				options = parts[1]
			}
			item := child.TranslateRule(method, options)
			currentMessages, exists := humanMessages[element]
			if !exists {
				humanMessages[element] = []string{}
			}
			humanMessages[element] = append(currentMessages, item)
		}
	}
	return humanMessages
}
