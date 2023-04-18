package main

import (
	"fmt"
	"math/rand"
)

func selectSort(v []int) {

	for j := 0; j < len(v); j++ {

		for i := 0; i < len(v); i++ {
			if v[i] >= v[j] {
				v[j], v[i] = v[i], v[j]
			}
		}

	}
}

func main() {

	array := []int{1, 6, 2, 0, 5}

	for i := 0; i < 5; i++ {
		array[i] = rand.Intn(100)
		i++
	}

	fmt.Println(array)
	selectSort(array)
	fmt.Println("-------------------")
	fmt.Println(array)

}
