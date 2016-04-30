package triangle

import (
	"math"
	"sort"
)

const testVersion = 2

// KindFromSides tells you whether a triangle with the given sides exists if so,
// it tells you whether the triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
  sides := []float64{a,b,c}
  sort.Sort(sort.Float64Slice(sides))
  switch {
    case math.IsNaN(sides[0]) || sides[0] <= 0 || sides[2] >= math.Inf(1) ||
          sides[0]+sides[1] < sides[2]:
      return NaT
    case sides[0] == sides[2]:
      return Equ
    case sides[0] == sides[1] || sides[1] == sides[2]:
      return Iso
    default:
      return Sca
  }
}

// Kind: The type that gets returned
type Kind string
const (
  NaT = "NaT" // not a triangle
  Equ = "Equ" // equilateral
  Iso = "Iso" // isosceles
  Sca = "Sca" // scalene
)
