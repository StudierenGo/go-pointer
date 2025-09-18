package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(array *[5]int) {
	for index, value := range *array {
		(*array)[len(array)-1-index] = value
	}
}
