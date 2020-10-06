package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "brasilia")
	fmt.Println(capitais, len(capitais), cap(capitais))

	for indice, capital := range capitais {
		fmt.Printf("Capital[%d] = %s\n\r", indice, capital)
	}
}
