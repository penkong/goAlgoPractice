package algos

import (
	"fmt"
	"math/rand"
)

// func generateSlice(size int) []int {

// 	slice := make([]int, size)
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < size; i++ {
// 			slice[i] = rand.Intn(999) - rand.Intn(999)
// 	}
// 	return slice
// }


func Quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	
	left, right := 0, len(a)-1
	
	// []int{434, 235, 435, 432, 65, 3, 23, 7, 34, 765, 234, 8}
	pivot := rand.Int() % len(a)
	fmt.Println(a[left], "left ")
	fmt.Println(a[right],"right ")
	fmt.Println(a[pivot], "before pivot")
	a[pivot], a[right] = a[right], a[pivot]
	fmt.Println(a[left], "left")
	fmt.Println(a[right], "right")
	fmt.Println(a[pivot], "after pivot")
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	fmt.Println(a[left],"AFTER LOOP")
	fmt.Println(a)
	a[left], a[right] = a[right], a[left]
	fmt.Println(a[right], "AFTER")
	fmt.Println(a[left], "after")
	fmt.Println(a, "after")
	Quicksort(a[:left])
	Quicksort(a[left+1:])

	return a
}
