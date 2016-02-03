# Validity

[![Build Status](https://travis-ci.org/Assembli/validity.svg)](https://travis-ci.org/Assembli/validity)

Package to parse and validate unsafe input. Includes a nice range of build-in validators, and is extensible for custom validation rules. This was inspired by Laravel's validators, and was written partially as an experiment while learning Go.

### Example

```go
import "validity"

// ...

// Data may be a map[string]interface{}, such as could be returned by JSON decoding.
// Alternately you can validate a struct using the function ValidateStruct
rules := Rules{
    "username": []string{"String", "required", "between:4,30"},
    "email":    []string{"String", "email"}
}

results := ValidateMap(data, rules)

if !results.IsValid {
    fmt.Printf("Error validating data! The following have failed:")

    // The following could print output like:
    //
    //  The validator between for username has failed!
    //  The valiator email for email has failed!

    for key, failures := range results.Errors {
        for _, method := range failures {
            fmt.Printf("The validator '%s' for '%s' has failed!", method, key);
        }
    }
} else {
    fmt.Print("Data is valid!")
}
```

### Usage

#### Basic Maps/Structs

Rules, passed into ValidateMap or ValidateStruct, is a map of strings to slices of "things". The keys of the map should be the field names to validate, in the struct or map of input given. The values should be slices of validators to run. For example:

```go
rules := Rules{"username": []string{"String", "required", "between: 4, 30"}}
```

#### Tagged Structs

You may also declare your rules as structure tags, in a field `validators`. Each rule should be seperated by ` and `, like so:

```go
type TestStructTags struct {
	Foo string 	`validators:"between:2,10 and email"`
	Bar int		`validators:"digits:3"`
	Baz float32
}

results := ValidateStructTags(TestStructTags{})
```

#### Built-In Rules

... would ensure the "username" is present and between four and 30 characters long. The first element of the map MUST be a value of the type to convert to. Any numeric or string type is valid. If the value cannot be converted to the given type, then it fails validation. The available types are: Int, String, Float.

Possible rules include:
 * `accepted`: The field under validation must be "yes", "on", true, or 1. Permits numeric and string types.
 * `alpha`: The field under validation must be entirely alphabetic characters. Permits string types.
 * `alpha_dash`: The field under validation may have alpha-numeric characters, as well as dashes and underscores. Permits string types.
 * `alpha_num`: The field under validation must be entirely alpha-numeric characters. Permits string types.
 * `between:,a,b`: The field under validation must be between "a" and "b" characters long, or between the values a and b (if numeric). Permits string and numeric types.
 * `between:,a,b`	The field under validation must be between "a" and "b" *(including the boundaries)* characters long, or between the values a and b inclusive a and b (if numeric). Permits string and numeric types. (is is the same as min, max combined)
 * `date`: The field under validation must parse to a date. Accepts string types.
 * `digits:num`: The field under validation must have exactly `num` of digits. Accepts numeric types.
 * `digits_between:a,b`: The field under validation must have between (a, b) digits. Accepts numeric types.
 * `digits_between:a,b`: The field under validation must have between [a, b] digits. Accepts numeric types.
 * `email`: The field under validation must be an email.
 * `ip`: The field under validation must be an IP, either ipv4 or ipv6. Accepts string types.
 * `ipv4`: The field under validation must be in IPv4 format. Accepts string types.
 * `ipv6`: The field under validation must be in IPv6 format. Accepts string types.
 * `len:num`: The field under validation must be be `num` characters long. Accepts string types.
 * `full_name`: The field must contain alpha-numeric characters or dots or spaces
 * `max`: The field under validation must be equal to or shorter than "a" (if a string), or equal to or smaller than "a" (if numeric). Accepts string and numeric types.
 * `min`: The field under validation must be equal to or longer than "a" (if a string), or equal to or greater than "a" (if numeric). Accepts string and numeric types.
 * `regex:pattern`: The field under validation must match the given pattern. Accepts string types.
 * `required`: The field under validation must be present. Accepts any type. Note optionality does not function when trying to validate structs, as it isn't possible to know if their zero values are zero because they aren't set, or because they should actually be zero.
 * `url`: The field under validation must be a URL. Accepts string types.

The return from the validation functions is a struct Results:

```go
type Results struct {
	// Indicates whether the data under validation has passed the set of rules.
	IsValid bool
	// This is a map of strings to slices of strings. Its keys will be any validation fields which had an error, and
	// the values will be the rules which failed.
	Errors  map[string][]string
	// The results is a map of everything after validation. This will be the same data, excluding extraneous values, and
	// values which did not passed validation. They will also be converted to the correct types. Integers will be of
	// type int64, Floats of float64, and Strings of string.
	//
	// The reason that values which did not pass validation are not returned, is because it is not possible to know
	// their types without reflecting them - validation can fail if a value is not able to be converted to a type.
	// This can lead to pitfalls - assuming a value is a of type - not to mention extra work on behalf of the
	// programmer.
	Data map[string]interface{}
}
```

#### Translators


A translator converts the errors from a validity result to human messages for a certain language.

You can use the translator in 2 ways:

* translate the entire map of errors, by returning a map of messages similar to errors (this allows validity to be used by the older versions)
  In order to translate you call this: `result.TranslateTo(language)`

  In case the language is not supported, go will panic.

  The `TranslateTo` method of ValidityResults is a factory method which creates the translator. It takes care of the translation process.

* translate one particular rule by: creating a translator and calling `TranslateRule(method, advance)`. For instance, `TranslateRule('between', '7,90')`

There are 2 languages supported:

* English
* Romanian
Note:
* For previous versions, you can use the function (ExtractMethod) to trim the full method which is now returned.




### Custom Validators

There are currently three "types" of validators: `IntValidityChecker`, `FloatValidityChecker`, and `StringValidityChecker`. These contain methods like `ValidateRule() bool`, and can therefore be extended easily. Let's make a silly validator:

```go
import "validity"

// ...

func (v validity.StringValidityChecker) ValidateSomethingSilly(suffix string) bool {
    return v.Item == "silly" + suffix
}
```

You can now use the validator like so:

```go
rules := Rules{"someString": []string{"String", "required", "something_silly:String"}}
```

The validator will now pass only if "someString" is given and is equal to `sillyString`. You will notice that:

 * "arguments" of the validators get passed in as strings to the function.
 * You do not need a type assertion on `v.Item`
 * Rules which are in snake\_case are converted to StudlyCase automatically.
