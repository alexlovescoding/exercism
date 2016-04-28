package hamming

import "errors"

const testVersion = 4

func Distance(s1, s2 string) (count int, err error) {
	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) != len(r2) {
		err = errors.New("The strings need to be the same length")
		return
	}

	for i, c := range r1 {
		if r2[i] != c {
			count++
		}
	}
	return
}
