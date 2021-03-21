package algos

import (
	"math"
)

func BinarySearch (arr []int, item int) (el int, res bool) {
	lower := (0)
	higher := (len(arr) -1)

	for lower <= higher {
		mid := int(math.Round(float64((lower + higher )/2)))
		guess := arr[mid]
		if guess == item {
			return mid, true
		} else if guess > item {
			higher = mid - 1
		} else {
			lower = mid + 1
		}
	}

	return 0, false
}