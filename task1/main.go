package main

import "fmt"

func main() {
	array := [5]int{0, 1, 2, 3, 4}
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered: ", v)
		}
	}()
	for i := 0; i < 6; i++ {
		fmt.Println(array[i])
	}
}
