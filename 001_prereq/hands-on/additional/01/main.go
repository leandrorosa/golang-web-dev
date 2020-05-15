package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(s)
	fmt.Println("Printing just the indexes")
	for i, _ := range s {
		fmt.Println(i)
	}
	fmt.Println("Printing indexes and values")
	for i, v := range s {
		fmt.Println(i, v)
	}
}
