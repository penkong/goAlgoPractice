package main

import (
	"fmt"

	"github.com/penkong/goAlgoPractice/challenges"
	// "github.com/penkong/goAlgoPractice/algos"
)

func main() {
	fmt.Println("started")
	// el,ok := algos.BinarySearch([]int{1,2,3,4,5,6,7,8,9,10}, 10)
	// if !ok {
	// 	log.Fatal("not found")
	// }
	// fmt.Println("index of item is", el)

	bn := make([]int, 0)
	m := append(bn, 1, 2, 1, 2, 1, 3, 3, 4, 3, 3, 4, 3)
	el := challenges.SalesByMatch(12, &m)
	fmt.Println(el)
}
