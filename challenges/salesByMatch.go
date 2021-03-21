package challenges

import (
	"math"
)

func SalesByMatch(n int, arr *[]int) (m int) {
	s := 0

	pairs := make(map[int]int)

	for i := 0; i < n; i++ {
		pairs[(*arr)[i]]++
	}

	for _, v := range pairs {
		s += int(math.Round(float64(v / 2)))
	}

	return s
}
