// MIT License
//
// Copyright (c) 2023 progxeno
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/progxeno/validate/pkg/validate"
)

func main() {

	// String validation
	valid := validate.StringIsEmpty("") // true
	fmt.Println("StringIsEmpty:", valid)
	valid = validate.StringMinLength("hello", 5) // true
	fmt.Println("StringMinLength:", valid)
	valid = validate.StringMaxLength("hello", 5) // true
	fmt.Println("StringMaxLength:", valid)
	valid = validate.StringMatchesRegex("hello123", "^[a-z]+$") // false
	fmt.Println("StringMatchesRegex:", valid)
	valid = validate.StringContainsChars("hello", "aeiou") // false
	fmt.Println("StringContainsChars:", valid)

	// Numeric validation
	valid = validate.NumericMinInt(5, 10) // false
	fmt.Println("NumericMinInt:", valid)
	valid = validate.NumericMaxInt(5, 10) // true
	fmt.Println("NumericMaxInt:", valid)
	valid = validate.NumericMinFloat(1.5, 2.0) // false
	fmt.Println("NumericMinFloat:", valid)
	valid = validate.NumericMaxFloat(1.5, 2.0) // true
	fmt.Println("NumericMaxFloat:", valid)
	valid = validate.NumericIsInt("123") // true
	fmt.Println("NumericIsInt:", valid)
	valid = validate.NumericIsFloat("1.23") // true
	fmt.Println("NumericIsFloat:", valid)

	// Email validation
	valid = validate.EmailIsValid("user@example.com") // true
	fmt.Println("EmailIsValid:", valid)
	valid = validate.EmailIsValid("userexample.com") // false
	fmt.Println("EmailIsValid:", valid)

	// URL validation
	valid = validate.URLIsValid("https://example.com") // true
	fmt.Println("URLIsValid:", valid)
	valid = validate.URLIsValid("https//example.com") // false
	fmt.Println("URLIsValid:", valid)

	// Date and time validation
	valid = validate.DateTimeIsValid("2022-01-01", "2006-01-02") // true
	fmt.Println("DateTimeIsValid:", valid)
	valid = validate.DateTimeIsFuture(time.Now().Add(time.Hour)) // true
	fmt.Println("DateTimeIsFuture:", valid)
	valid = validate.DateTimeIsPast(time.Now().Add(-time.Hour)) // true
	fmt.Println("DateTimeIsPast:", valid)

	// Password strength validation
	valid = validate.PasswordMatchesPolicy("P@ssw0rd2", 8, 2, 1) // true
	fmt.Println("PasswordMatchesPolicy:", valid)

	// File extension validation
	valid = validate.FileIsValidExtension("example.jpg", []string{".jpg", ".png"}) // true
	fmt.Println("FileIsValidExtension:", valid)

	// Custom validation
	v := validate.NewValidate()
	v.AddRule(func(value interface{}) error {
		if value.(int) < 0 {
			return errors.New("value must be positive")
		}
		return nil
	})
	err := v.Validate(5) // nil
	fmt.Println("Custom validation:", err)
	err = v.Validate(-5) // error: value must be positive
	fmt.Println("Custom validation:", err)
}
