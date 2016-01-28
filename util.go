package validity

import (
	"reflect"
	"strings"
	"unicode"
)

// Uppercases the first letter of the given string. This reasonably
// nice solution is from kortschak on the go-nuts mailing list!
func firstToUpper(s string) string {
	if s == "" {
		return s
	}

	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])

	return string(a)
}

// Converts a snake_cased string to a camelCased one.
func snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	out := parts[0]

	for _, part := range parts[1:] {
		out += firstToUpper(part)
	}

	return out
}

// Converts a snake_cased string to a StudlyCased one.
func snakeToStudly(s string) string {
	return firstToUpper(snakeToCamel(s))
}

// Calls the function `name` on the map `m` with the given splat of arguments, and return the result (if possible).
// Returns an error in the event of a parameter mismatch.
func callIn(m interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m)

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.MethodByName(name).Call(in)
	return result, nil
}

// Checks to see if the given string appears in the slice.
func inSlice(a string, list []string) bool {
	a = strings.ToLower(a)

	for _, b := range list {
		if strings.ToLower(b) == a {
			return true
		}
	}
	return false
}
