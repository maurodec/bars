// The MIT License (MIT)

// Copyright (c) 2014 Mauro de Carvalho

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package provides functions for generating simple Unicode graphs based on a
// series of float64 inputs.
package bars

// Holds the set of characters that will be used to render the graph. Each
// character will be used to represent a different value in the graph. Values
// go from Zero to One in one eighth increments.
type BarSet struct {
	Zero          rune
	Eighth        rune
	Quarter       rune
	ThreeEighths  rune
	Half          rune
	FiveEighths   rune
	ThreeQuarters rune
	SevenEighths  rune
	One           rune
}

// A pre-defined character set based on Unicode blocks.
var NiceBarSet = BarSet{
	' ',
	'▁',
	'▂',
	'▃',
	'▄',
	'▅',
	'▆',
	'▇',
	'█',
}

// A pre-defined character set based on Braile characters.
var BraileBarSet = BarSet{
	' ',
	'⣀',
	'⣀',
	'⣤',
	'⣤',
	'⣶',
	'⣶',
	'⣿',
	'⣿',
}

// Renders a graph, one character per input value to display the given values.
// the lowest value given will be represented with a value of Zero int he graph,
// while the highest value will be represented with a value of One.
// If only one value is given it will be represented as one half.
// The graph is rendered as an array of runes, one rune per value given.
func MakeBar(seq []float64, set BarSet) []rune {
	// If the sequence is empty just return an empty result.
	if len(seq) == 0 {
		return []rune{}
	}

	// If the sequence is only one number it is not possible to calculate the
	// result. We are returning half. The value is completely arbitrary.
	if len(seq) == 1 {
		return []rune{set.Half}
	}

	// We can calculate bars. Determine the minimum and maximum values.
	min := seq[0]
	max := seq[0]
	for i := 1; i < len(seq); i++ {
		if min > seq[i] {
			min = seq[i]
			continue
		}

		if max < seq[i] {
			max = seq[i]
		}
	}

	// for each value, calculate the percentage and add it to the solution
	var diff = max - min
	result := make([]rune, len(seq))

	// Create a lookup table so that we can fetch a char by index
	lookup := []rune{
		set.Zero,
		set.Eighth,
		set.Quarter,
		set.ThreeEighths,
		set.Half,
		set.FiveEighths,
		set.ThreeQuarters,
		set.SevenEighths,
		set.One,
	}

	for i := 0; i < len(seq); i++ {
		// Get what percentage the value is of the total between min and max
		percent := (seq[i] - min) / diff

		// Figure out what char to display it with. 0, 1/16 will use Zero,
		// 1/16, 3/16 will use 1/8, 3/16, 5/16 will use 1/4 and so on...
		index := int((percent + 1/16.) * 8)

		result[i] = lookup[index]
	}

	return result
}
