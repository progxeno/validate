# Validate

The `validate` package provides a collection of functions for validating various types of data, including strings, numbers, dates, and more. It also offers the flexibility to define custom validation rules for any data type.

## Installation

To use the `validate` package in your Go project, you can install it using the `go get` command:

```shell
go get github.com/progxeno/validate
```

## Usage

The `validate` package offers validation functions that can be used to validate different types of data.

Available functions:

- validate.StringIsEmpty
  > Checks if a string is empty or contains only whitespace.

- validate.StringMinLength
  > Validates that a string has a minimum length.

- validate.StringMaxLength
  > Validates that a string has a maximum length.

- validate.StringMatchesRegex
  > Validates that a string matches a given regular expression pattern.

- validate.StringContainsChars
  > Validates that a string contains any of the specified characters.

- validate.NumericMinInt
  > Validates that an integer is greater than or equal to a minimum value.

- validate.NumericMaxInt
  > Validates that an integer is less than or equal to a maximum value.

- validate.NumericMinFloat
  > Validates that a float is greater than or equal to a minimum value.

- validate.NumericMaxFloat
  > Validates that a float is less than or equal to a maximum value.

- validate.NumericIsInt
  > Checks if a string represents a valid integer.

- validate.NumericIsFloat
  > Checks if a string represents a valid float.

- validate.EmailIsValid
  > Checks if a string represents a valid email address.

- validate.URLIsValid
  > Checks if a string represents a valid URL.

- validate.DateTimeIsValid
  > Checks if a string represents a valid date in the specified format.

- validate.DateTimeIsFuture
  > Checks if a date is in the future.

- validate.DateTimeIsPast
  > Checks if a date is in the past.

- validate.PasswordMatchesPolicy
  > Checks if a password matches the specified policy.
  
- validate.FileIsValidExtension
  > Checks if a filename has a valid extension.

In addition, the `validate` package includes a `Validate` struct that allows you to define custom validation rules. You can create a new instance of the `Validate` using the `NewValidate` function, add validation rules using the `AddRule` method, and validate any value using the `Validate` method.

Example of custom validation function:

```go
package main

import (
 "fmt"
 "github.com/progxeno/validate"
)

func main() {
 // Create a new validate
    v := validate.NewValidate()

 // Add a custom validation rule
    v.AddRule(func(value interface{}) error {
        if value.(int) < 0 {
            return errors.New("value must be positive")
        }
        return nil
    })

 // Validate a value using the custom rule
    err := v.Validate(5) // nil
    fmt.Println("Custom validation:", err)
    err = v.Validate(-5) // error: value must be positive
    fmt.Println("Custom validation:", err)
}
```

## License

The `validate` package is licensed under the MIT License. See the LICENSE file for more information.
