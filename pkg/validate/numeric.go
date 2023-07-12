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

package validate

import (
	"strconv"
)

func NumericMinInt(n int, min int) bool {
	return n >= min
}

func NumericMaxInt(n int, max int) bool {
	return n <= max
}

func NumericMinFloat(f float64, min float64) bool {
	return f >= min
}

func NumericMaxFloat(f float64, max float64) bool {
	return f <= max
}

func NumericIsInt(n string) bool {
	_, err := strconv.Atoi(n)
	return err == nil
}

func NumericIsFloat(n string) bool {
	_, err := strconv.ParseFloat(n, 64)
	return err == nil
}
