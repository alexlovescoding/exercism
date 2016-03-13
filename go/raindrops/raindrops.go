package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 2

func Convert(num int) string {
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		return strconv.Itoa(num)
	}

	var sb bytes.Buffer

	if num%3 == 0 {
		sb.WriteString("Pling")
	}
	if num%5 == 0 {
		sb.WriteString("Plang")
	}
	if num%7 == 0 {
		sb.WriteString("Plong")
	}
	return sb.String()
}
