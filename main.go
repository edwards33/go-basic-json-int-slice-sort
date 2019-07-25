package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{5, 2, 12, 14, 56, 24, 11, 76, 7}

	fmt.Println(num)
	sort.Sort(sort.IntSlice(num))
	fmt.Println(num)

}