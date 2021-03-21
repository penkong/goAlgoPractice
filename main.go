package main

import (
	"fmt"
	"log"

	"github.com/penkong/goAlgoPractice/algos"
)


func main() {
	fmt.Println("started")
	el,ok := algos.BinarySearch([]int{1,2,3,4,5,6,7,8,9,10}, 10)
	if !ok {
		log.Fatal("not found")
	} 
	fmt.Println("index of item is", el)
}