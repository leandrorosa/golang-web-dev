package main

import "fmt"

func main() {
	m := map[string]int{
		"Leandro": 34,
		"Gandalf": 10000,
		"Heisenberg": 55,
	}
	fmt.Println(m)

	fmt.Println("Printing keys")
	for k,_ := range m {
		fmt.Println(k)
	}
	fmt.Println("Printing keys and values")
	for k,v := range m {
		fmt.Println(k, v)
	}

}
