package main

import "fmt"

func jumpingClouds(c [7]int32) int32 {
	size := int32(len(c))
	var step int32
	stepCount := 0

	for step < (size - 1) {

		if step+2 < size && c[step+2] == 0 {
			step = step + 2
			stepCount++
			continue
		}

		step++
		stepCount++

	}
	return int32(stepCount)
}

func main() {
	clouds := [...]int32{0, 1, 0, 0, 0, 1, 0}
	//clouds := [...]int{0, 0, 1, 0, 0, 1, 0}

	fmt.Println("Steps count: ", jumpingClouds(clouds))
}
