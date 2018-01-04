package main

import (
	"fmt"

	binarySearch "github.com/Alexandr-Penkin/grokking-algorithms/binary-search"
	"github.com/Alexandr-Penkin/grokking-algorithms/factorial"
	sort "github.com/Alexandr-Penkin/grokking-algorithms/selection-sort"
	search "github.com/Alexandr-Penkin/grokking-algorithms/typical-search"
)

func main() {
	slice := []int{}

	for i := 0; i <= 1000000; i++ {
		slice = append(slice, i)
	}

	value := 643323

	index, ok := search.Search(slice, value)

	if ok {
		println(index)
	} else {
		println("not found")
	}

	index, ok = binarySearch.Search(slice, value)

	if ok {
		println(index)
	} else {
		println("not found")
	}

	slice = []int{5, 3, 6, 2, 10}
	fmt.Println(sort.SelectionSort(slice))

	fmt.Println(factorial.Factorial(10))
}
