package validity

import (
	"fmt"
	"strconv"
)

// Queue object is where we're going to store all our data related to validation of a particular set.
// It should probably not be used directly, unless you're really going for better control over things.
type Queue struct {
	// A list of checkers to run. Each conformer of ValidityChecker corresponds to a data type and has unique
	// validation checks which it can run. Each individual ValidityChecker is responsible for storing a data point,
	// the key of the data point, and the rules intended to be run against the data point.
	Checkers []Checker
	// Data is a map of the raw input data for the queue to parse. This is set by ValidateMap, usually.
	Data    map[string]interface{}
	Rules   Rules
	Results *Results
}

// Parsers is a map of functions to parse the given types with. Each one is responsible for converting the value
// to the correct type (if possible) and inserting an appropriate checker in the Queue, or adding an error to
// the output (if not possible to convert). By default these convert to the highest precision available. That is,
// int64, float64, and string.
//
// Apologies if this is not The Go Way™, put in a PR if you have a better solution :)
type Parsers struct{}

// Run is reponsible for resetting the results, then running parsers/checkers. The actual validation occurs in two
// stages. First, the data is run through to have its types fixed and appropriate checkers added to the queue. Inability
// to convert types will put an error in the output, and no checker will be added for the second stage.
//
// In the second stage, each checker is run. The checker simply returns a slice of any validators which did not pass,
// empty if everything is good and happy. We take that output and affix it onto the results appropriately. It is also in
// this stage that, if everything passed, the converted, safe value is put in the ValidationResult.Data map.
func (c Queue) Run() {
	c.Results.IsValid = true
	c.Results.Errors = map[string][]string{}
	c.Results.Data = map[string]interface{}{}

	c.RunParsers()
	c.RunCheckers()
}

// RunParsers runs the parsers, for the second stage. See Run() for explaination.
func (c *Queue) RunParsers() {
	for key, validator := range c.Rules {
		item, exists := c.Data[key]

		if !exists {
			if inSlice("required", validator) {
				c.AddError(key, "required")
			}
			continue
		}

		// This calls a function like "ParseInt" present on the Parsers map.
		callIn(Parsers{}, "Parse"+validator[0], c, key, item, validator)
	}
}

// RunCheckers runs the checkers, for the second stage. See Run() for explaination.
func (c *Queue) RunCheckers() {
	for _, checker := range c.Checkers {
		// Add errors from the checker. If no errors occured, the chcker returns
		// an empty slice and no errors are added.
		errors := checker.GetErrors()
		c.AddErrors(checker.GetKey(), errors)

		// If there were no errors, then save the typed item in the results data.
		if len(errors) == 0 {
			c.Results.Data[checker.GetKey()] = checker.GetItem()
		}
	}
}

// AddError inserts an error into the Results.Error, with the specified key. If there are already more than
// zero errors for the key, then the error is simply appended on the list.
func (c *Queue) AddError(key string, error string) {
	if _, exists := c.Results.Errors[key]; !exists {
		c.Results.Errors[key] = []string{}
	}

	c.Results.Errors[key] = append(c.Results.Errors[key], error)
	c.Results.IsValid = false
}

// AddErrors inserts a list of errors for the given key into the Results.Error.
// Any existing errors are not replaced, only appended to.
func (c *Queue) AddErrors(key string, errors []string) {
	for _, error := range errors {
		c.AddError(key, error)
	}
}

// ParseInt converts the given value to an integer. Parses it as a string. Not super amazing, but the most elegant solution that
// I was able to find. Before we used reflection of original type with switch statements... *shudders*
func (v Parsers) ParseInt(c *Queue, key string, value interface{}, rules []string) {
	item := fmt.Sprintf("%v", value)

	val, err := strconv.ParseInt(item, 10, 64)
	if err != nil {
		c.AddError(key, "Int")
		return
	}

	c.Checkers = append(c.Checkers, IntValidityChecker{Key: key, Item: val, Rules: rules})
}

// ParseFloat converte the given value to a float, using the same method as was used to convert to int.
func (v Parsers) ParseFloat(c *Queue, key string, value interface{}, rules []string) {
	item := fmt.Sprintf("%v", value)

	val, err := strconv.ParseFloat(item, 64)
	if err != nil {
		c.AddError(key, "Float")
		return
	}

	c.Checkers = append(c.Checkers, FloatValidityChecker{Key: key, Item: val, Rules: rules})
}

// ParseString converts the given value to a string.
func (v Parsers) ParseString(c *Queue, key string, item interface{}, rules []string) {
	c.Checkers = append(c.Checkers, StringValidityChecker{Key: key, Item: fmt.Sprintf("%s", item), Rules: rules})
}
