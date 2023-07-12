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
	"testing"

	"github.com/progxeno/validate/pkg/validate"
)

func TestValidate_Validate(t *testing.T) {
	t.Parallel()

	type in struct {
		value interface{}
		rules []validate.ValidationRule
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "valid input",
			in: in{
				value: "hello",
				rules: []validate.ValidationRule{
					func(v interface{}) error {
						s, ok := v.(string)
						if !ok {
							return errors.New("not a string")
						}
						if len(s) < 5 {
							return errors.New("too short")
						}
						return nil
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "invalid input",
			in: in{
				value: "hi",
				rules: []validate.ValidationRule{
					func(v interface{}) error {
						s, ok := v.(string)
						if !ok {
							return errors.New("not a string")
						}
						if len(s) < 5 {
							return errors.New("too short")
						}
						return nil
					},
				},
			},
			want: want{
				err: errors.New("too short"),
			},
		},
		{
			name: "multiple rules",
			in: in{
				value: "hello",
				rules: []validate.ValidationRule{
					func(v interface{}) error {
						s, ok := v.(string)
						if !ok {
							return errors.New("not a string")
						}
						if len(s) < 5 {
							return errors.New("too short")
						}
						return nil
					},
					func(v interface{}) error {
						s, ok := v.(string)
						if !ok {
							return errors.New("not a string")
						}
						if len(s) > 10 {
							return errors.New("too long")
						}
						return nil
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "multiple rules with errors",
			in: in{
				value: "hi",
				rules: []validate.ValidationRule{
					func(v interface{}) error {
						s, ok := v.(string)
						if !ok {
							return errors.New("not a string")
						}
						if len(s) < 5 {
							return errors.New("too short")
						}
						return nil
					},
					func(v interface{}) error {
						s, ok := v.(string)
						if !ok {
							return errors.New("not a string")
						}
						if len(s) > 10 {
							return errors.New("too long")
						}
						return nil
					},
				},
			},
			want: want{
				err: errors.New("too short"),
			},
		},
		{
			name: "empty rules",
			in: in{
				value: "hello",
				rules: []validate.ValidationRule{},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validate.NewValidate()
			for _, rule := range tt.in.rules {
				v.AddRule(rule)
			}
			err := v.Validate(tt.in.value)
			if err == nil {
				if tt.want.err != nil {
					t.Errorf("Validate() error = nil, want err %v", tt.want.err)
				}
			} else if err.Error() != tt.want.err.Error() {
				t.Errorf("Validate() error = %v, want err %v", err, tt.want.err)
			}
		})
	}
}
