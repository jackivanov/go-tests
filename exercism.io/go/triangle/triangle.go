// Package triangle can help you to get a type of a triangle
package triangle

import "math"

// Kind int type of triangles
type Kind int

const (
	// NaT Not a triangle
	NaT = iota
	// Equ equilateral triangle
	Equ
	// Iso isosceles triangle
	Iso
	// Sca scalene triangle
	Sca
)

// KindFromSides determines whether the given triangle is equilateral, isosceles, scalene or not a triangle at all
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch true {
	case a <= 0 || b <= 0 || c <= 0:
		k = NaT
	case a+b < c || a+c < b || b+c < a:
		k = NaT
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		k = NaT
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		k = NaT
	case a == b && b == c:
		k = Equ
	case a+b == a*2 || b+c == b*2 || a+c == c*2:
		k = Iso
	case a != b && a != c:
		k = Sca
	}
	return k
}
