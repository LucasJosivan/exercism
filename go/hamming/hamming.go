package hamming

import "errors"

func Distance(a, b string) (diff int, err error) {
	ar := []rune(a)
	br := []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("a must be the same quantity of caracters than b")
	}
	for i := 0; i < len(ar); i++ {
		if ar[i] != br[i] {
			diff++
		}
	}
	return
}
