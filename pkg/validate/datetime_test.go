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
	"time"

	"github.com/progxeno/validate/pkg/validate"
)

func TestDatetimeIsValid(t *testing.T) {
	t.Parallel()

	type in struct {
		datetime string
		format   string
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
			name: "valid datetime",
			in: in{
				datetime: "2022-01-01 12:00:00",
				format:   "2006-01-02 15:04:05",
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "invalid datetime",
			in: in{
				datetime: "invalid-datetime",
				format:   "2006-01-02 15:04:05",
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.DateTimeIsValid(tt.in.datetime, tt.in.format)
			if got != tt.want.result {
				t.Errorf("DateTimeIsValid(%q, %q) = %v, want %v", tt.in.datetime, tt.in.format, got, tt.want.result)
			}
		})
	}
}

func TestIsFutureDatetime(t *testing.T) {
	t.Parallel()

	type in struct {
		datetime time.Time
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
			name: "future datetime",
			in: in{
				datetime: time.Now().AddDate(0, 0, 1),
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "past datetime",
			in: in{
				datetime: time.Now().AddDate(0, 0, -1),
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.DateTimeIsFuture(tt.in.datetime)
			if got != tt.want.result {
				t.Errorf("DateTimeIsFuture(%v) = %v, want %v", tt.in.datetime, got, tt.want.result)
			}
		})
	}
}

func TestIsPastDatetime(t *testing.T) {
	t.Parallel()

	type in struct {
		datetime time.Time
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
			name: "past datetime",
			in: in{
				datetime: time.Now().AddDate(0, 0, -1),
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "future datetime",
			in: in{
				datetime: time.Now().AddDate(0, 0, 1),
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.DateTimeIsPast(tt.in.datetime)
			if got != tt.want.result {
				t.Errorf("DateTimeIsPast(%v) = %v, want %v", tt.in.datetime, got, tt.want.result)
			}
		})
	}
}
