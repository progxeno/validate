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

func TestFileHasValidExtension(t *testing.T) {
	t.Parallel()

	type in struct {
		filename   string
		extensions []string
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
			name: "valid extension",
			in: in{
				filename:   "file.txt",
				extensions: []string{".txt", ".doc"},
			},
			want: want{
				result: true,
				err:    nil,
			},
		},
		{
			name: "invalid extension",
			in: in{
				filename:   "file.txt",
				extensions: []string{".doc", ".pdf"},
			},
			want: want{
				result: false,
				err:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate.FileIsValidExtension(tt.in.filename, tt.in.extensions)
			if got != tt.want.result {
				t.Errorf("FileIsValidExtension(%q, %v) = %v, want %v", tt.in.filename, tt.in.extensions, got, tt.want.result)
			}
		})
	}
}
