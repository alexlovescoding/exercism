package raindrops

import "strconv"

const testVersion = 2

func Convert(num int) string {
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		return strconv.Itoa(num)
	}

	r := ""

	if num%3 == 0 {
		r += "Pling"
	}
	if num%5 == 0 {
		r += "Plang"
	}
	if num%7 == 0 {
		r += "Plong"
	}
	return r
}
