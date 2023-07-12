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
	"errors"
	"strconv"
	"testing"

	"github.com/progxeno/validate/pkg/validate"
)

func TestMinInt(t *testing.T) {
	t.Parallel()

	type in struct {
		n   int
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
			name: "valid input",
			in: in{
				n:   5,
				min: 2,
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				n:   1,
				min: 2,
			},
			want: want{
				result: false,
				err:    errors.New("value is less than minimum"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.NumericMinInt(tt.in.n, tt.in.min)
			if got != tt.want.result {
				t.Errorf("NumericMinInt(%+v) = %v, want %v", tt.in, got, tt.want.result)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	t.Parallel()

	type in struct {
		n   int
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
			name: "valid input",
			in: in{
				n:   5,
				max: 10,
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				n:   15,
				max: 10,
			},
			want: want{
				result: false,
				err:    errors.New("value is greater than maximum"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.NumericMaxInt(tt.in.n, tt.in.max)
			if got != tt.want.result {
				t.Errorf("NumericMaxInt(%+v) = %v, want %v", tt.in, got, tt.want.result)
			}
		})
	}
}

func TestMinFloat(t *testing.T) {
	t.Parallel()

	type in struct {
		f   float64
		min float64
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
			name: "valid input",
			in: in{
				f:   1.5,
				min: 1.0,
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				f:   0.5,
				min: 1.0,
			},
			want: want{
				result: false,
				err:    errors.New("value is less than minimum"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.NumericMinFloat(tt.in.f, tt.in.min)
			if got != tt.want.result {
				t.Errorf("NumericMinFloat(%+v) = %v, want %v", tt.in, got, tt.want.result)
			}
		})
	}
}

func TestMaxFloat(t *testing.T) {
	t.Parallel()

	type in struct {
		f   float64
		max float64
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
			name: "valid input",
			in: in{
				f:   1.5,
				max: 2.0,
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				f:   2.5,
				max: 2.0,
			},
			want: want{
				result: false,
				err:    errors.New("value is greater than maximum"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.NumericMaxFloat(tt.in.f, tt.in.max)
			if got != tt.want.result {
				t.Errorf("NumericMaxFloat(%+v) = %v, want %v", tt.in, got, tt.want.result)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	t.Parallel()

	type in struct {
		n string
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
			name: "valid input",
			in: in{
				n: "123",
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				n: "abc",
			},
			want: want{
				result: false,
				err:    strconv.ErrSyntax,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.NumericIsInt(tt.in.n)
			if got != tt.want.result {
				t.Errorf("NumericIsInt(%+v) = %v, want %v", tt.in, got, tt.want.result)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
	t.Parallel()

	type in struct {
		n string
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
			name: "valid input",
			in: in{
				n: "1.23",
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				n: "abc",
			},
			want: want{
				result: false,
				err:    strconv.ErrSyntax,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.NumericIsFloat(tt.in.n)
			if got != tt.want.result {
				t.Errorf("NumericIsFloat(%+v) = %v, want %v", tt.in, got, tt.want.result)
			}
		})
	}
}
