package main

import (
	"fmt"
	"math/rand"
)

func bubblesort(v []int) {

	mudou := false
	for j := 0; j < len(v)-1; j++ {

		for i := 0; i < len(v)-1-j; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				mudou = true
			}
		}
		if !mudou {
			return
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
	bubblesort(array)
	fmt.Println(array)

}
