package main

import "fmt"

//

func main() {
	clouds := [...]int{0, 1, 0, 0, 0, 1, 0}
	//clouds := [...]int{0, 0, 1, 0, 0, 1, 0}
	size := int32(len(clouds))
	var step int32
	stepCount := 0

	//steps 2 1 1 2
	//
	for step < (size - 1) {

		if step+2 < size && clouds[step+2] == 0 {
			step = step + 2
			fmt.Println("step : ", step)
			stepCount++
			continue
		}

		step++
		fmt.Println("step : ", step)
		stepCount++

	}

	fmt.Println("Steps count: ", stepCount)
}
