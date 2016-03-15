package raindrops

import "strconv"

const testVersion = 2

func Convert(num int) string {
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		return strconv.Itoa(num)
	}

	result := ""

	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}
	return result
}
