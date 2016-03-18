package triangle

const testVersion = 2

// Kind tells you whether a triangle with the given sides exists if so,
// it tells you whether the triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {

}

// Notice it returns this type.  Pick something suitable.
type Kind

// Pick values for the following identifiers used by the test program.
NaT // not a triangle
Equ // equilateral
Iso // isosceles
Sca // scalene
