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

package validate_test

import (
	"testing"

	"github.com/progxeno/validate/pkg/validate"
)

func TestStringIsEmpty(t *testing.T) {
	t.Parallel()

	type in struct {
		s string
	}

	type want struct {
		result bool
		err    error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "empty string",
			in: in{
				s: "",
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "non-empty string",
			in: in{
				s: "hello",
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.StringIsEmpty(tt.in.s)
			if got != tt.want.result {
				t.Errorf("StringIsEmpty(%q) = %v, want %v", tt.in.s, got, tt.want.result)
			}
		})
	}
}

func TestStringMinLength(t *testing.T) {
	t.Parallel()

	type in struct {
		s   string
		min int
	}

	type want struct {
		result bool
		err    error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "string is too short",
			in: in{
				s:   "hello",
				min: 6,
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
		{
			name: "string is exactly the minimum length",
			in: in{
				s:   "hello",
				min: 5,
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "string is longer than the minimum length",
			in: in{
				s:   "hello",
				min: 4,
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.StringMinLength(tt.in.s, tt.in.min)
			if got != tt.want.result {
				t.Errorf("StringMinLength(%q, %d) = %v, want %v", tt.in.s, tt.in.min, got, tt.want.result)
			}
		})
	}
}

func TestStringMaxLength(t *testing.T) {
	t.Parallel()

	type in struct {
		s   string
		max int
	}

	type want struct {
		result bool
		err    error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "string is too long",
			in: in{
				s:   "hello",
				max: 4,
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
		{
			name: "string is exactly the maximum length",
			in: in{
				s:   "hello",
				max: 5,
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "string is shorter than the maximum length",
			in: in{
				s:   "hello",
				max: 6,
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.StringMaxLength(tt.in.s, tt.in.max)
			if got != tt.want.result {
				t.Errorf("StringMaxLength(%q, %d) = %v, want %v", tt.in.s, tt.in.max, got, tt.want.result)
			}
		})
	}
}

func TestStringMatchesRegex(t *testing.T) {
	t.Parallel()

	type in struct {
		s       string
		pattern string
	}

	type want struct {
		result bool
		err    error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "string matches pattern",
			in: in{
				s:       "hello123",
				pattern: "^[a-z]+$",
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
		{
			name: "string does not match pattern",
			in: in{
				s:       "hello123",
				pattern: "^[a-z0-9]+$",
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.StringMatchesRegex(tt.in.s, tt.in.pattern)
			if got != tt.want.result {
				t.Errorf("StringMatchesRegex(%q, %q) = %v, want %v", tt.in.s, tt.in.pattern, got, tt.want.result)
			}
		})
	}
}

func TestStringContainsChars(t *testing.T) {
	t.Parallel()

	type in struct {
		s     string
		chars string
	}

	type want struct {
		result bool
		err    error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "string contains all characters",
			in: in{
				s:     "hello",
				chars: "helo",
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "string does not contain all characters",
			in: in{
				s:     "hello",
				chars: "heloq",
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.StringContainsChars(tt.in.s, tt.in.chars)
			if got != tt.want.result {
				t.Errorf("StringContainsChars(%q, %q) = %v, want %v", tt.in.s, tt.in.chars, got, tt.want.result)
			}
		})
	}
}
