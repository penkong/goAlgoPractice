package algos

func SelectionSort(arr *[]int) []int {
	newArr := make([]int, 0)
	for len(*arr) != 0 {
		smallestIndex := findSmallest(arr)
		newArr = append(newArr, (*arr)[smallestIndex])
		*arr = remove(arr, smallestIndex)
	}
	return newArr
}

func findSmallest(arr *[]int) int {
	smallest := (*arr)[0]
	smallestIndex := 0
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] < smallest {
			smallest = (*arr)[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func remove(arr *[]int, i int) []int {
	copy((*arr)[i:], (*arr)[i+1:])
	return (*arr)[:len(*arr)-1]

}
