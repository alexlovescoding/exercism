package raindrops

import "strconv"

const testVersion = 2

func Convert(num int) (r string) {
	factors := [3]int{3, 5, 7}
	outputs := [3]string{"Pling", "Plang", "Plong"}

	for i := 0; i < len(factors); i++ {
		if num%factors[i] == 0 {
			r += outputs[i]
		}
	}

	if len(r) > 0 { 
		return r
	}

	return strconv.Itoa(num)
}
